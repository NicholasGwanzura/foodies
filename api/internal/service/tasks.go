package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
)

// ListTasks returns tasks in scope. "me" means tasks assigned to the caller, or
// to a role they hold and not yet claimed.
func (s *Service) ListTasks(ctx context.Context, p db.ListTasksForParams, assignee string) ([]db.Task, error) {
	q := s.q(ctx)
	a := auth.From(ctx)
	p.AllSites, p.SiteIds = scopeOf(ctx)
	p.RoleIds = []uuid.UUID{}
	switch assignee {
	case "":
	case "me":
		p.Assignee = &a.UserID
		roles, err := q.UserRoleIds(ctx, a.UserID)
		if err != nil {
			return nil, fmt.Errorf("read roles: %w", err)
		}
		p.RoleIds = roles
	default:
		id, err := uuid.Parse(assignee)
		if err != nil {
			return nil, errInvalid("/assignee", "Use me, or a user id.")
		}
		p.Assignee = &id
	}
	return wrap2(q.ListTasksFor(ctx, p))("list tasks")
}

// CreateTask adds a one-off task.
func (s *Service) CreateTask(ctx context.Context, p db.CreateTaskRowParams) (db.Task, error) {
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return db.Task{}, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateTaskRow(ctx, p)
	if err != nil {
		return row, wrap(err, "create task")
	}
	if p.AssigneeUserID != nil {
		if err := s.notify(ctx, []uuid.UUID{*p.AssigneeUserID}, "tasks", "New task", p.Title, "tasks", row.ID); err != nil {
			return row, err
		}
	}
	return row, s.audit(ctx, "tasks.task.write", "tasks", row.ID, &row.SiteID, nil, row)
}

func (s *Service) taskAction(ctx context.Context, id uuid.UUID, perm, what string, apply func(q db.Querier, actor *uuid.UUID) (db.Task, error)) (db.Task, error) {
	before, err := s.GetTask(ctx, id)
	if err != nil {
		return before, err
	}
	row, err := apply(s.q(ctx), actorID(ctx))
	if err != nil {
		return before, stateError(err, what)
	}
	return row, s.audit(ctx, perm, "tasks", id, &row.SiteID, before, row)
}

// ClaimTask takes a role task for the caller.
func (s *Service) ClaimTask(ctx context.Context, id uuid.UUID) (db.Task, error) {
	return s.taskAction(ctx, id, "tasks.task.claim", "claim this task", func(q db.Querier, actor *uuid.UUID) (db.Task, error) {
		return q.ClaimTaskRow(ctx, db.ClaimTaskRowParams{ID: id, AssigneeUserID: actor})
	})
}

// CompleteTask finishes a task.
func (s *Service) CompleteTask(ctx context.Context, id uuid.UUID) (db.Task, error) {
	return s.taskAction(ctx, id, "tasks.task.complete", "complete this task", func(q db.Querier, actor *uuid.UUID) (db.Task, error) {
		return q.CompleteTaskRow(ctx, db.CompleteTaskRowParams{ID: id, UpdatedBy: actor})
	})
}

// SkipTask records a skipped control, with its reason and alternative control.
func (s *Service) SkipTask(ctx context.Context, id uuid.UUID, reason string, alternative *string) (db.Task, error) {
	if strings.TrimSpace(reason) == "" {
		return db.Task{}, errInvalid("/reason", "Say why this is being skipped.")
	}
	text := reason
	if alternative != nil && *alternative != "" {
		text += " Alternative control: " + *alternative
	}
	return s.taskAction(ctx, id, "tasks.task.skip", "skip this task", func(q db.Querier, actor *uuid.UUID) (db.Task, error) {
		return q.SkipTaskRow(ctx, db.SkipTaskRowParams{ID: id, UpdatedBy: actor, SkipReason: &text})
	})
}

// ReassignTask gives a task to one person.
func (s *Service) ReassignTask(ctx context.Context, id, user uuid.UUID) (db.Task, error) {
	row, err := s.taskAction(ctx, id, "tasks.task.reassign", "reassign this task", func(q db.Querier, actor *uuid.UUID) (db.Task, error) {
		return q.ReassignTaskRow(ctx, db.ReassignTaskRowParams{ID: id, UpdatedBy: actor, AssigneeUserID: &user})
	})
	if err != nil {
		return row, err
	}
	return row, s.notify(ctx, []uuid.UUID{user}, "tasks", "Task assigned to you", row.Title, "tasks", row.ID)
}

// frequency is a schedule's rule value: {"every":"day"|"week"|"month","at":"08:00","weekday":1,"day":1}.
type frequency struct {
	Every   string `json:"every"`
	At      string `json:"at"`
	Weekday *int   `json:"weekday"`
	Day     *int   `json:"day"`
}

func (f frequency) dueOn(day time.Time) bool {
	switch f.Every {
	case "week":
		return f.Weekday == nil || int(day.Weekday()) == *f.Weekday
	case "month":
		return f.Day == nil || day.Day() == *f.Day
	default:
		return true
	}
}

// generateTasks creates today's tasks from active schedules. A schedule whose
// frequency rule has no active value generates nothing, as readiness says.
func (s *Service) generateTasks(ctx context.Context) error {
	q := s.q(ctx)
	now := s.now().In(Harare)
	schedules, err := q.ListActiveSchedules(ctx)
	if err != nil {
		return fmt.Errorf("list schedules: %w", err)
	}
	for _, sc := range schedules {
		for _, site := range sc.SiteIds {
			if sc.WorkflowKey != nil {
				if err := s.requireLive(ctx, *sc.WorkflowKey, site); err != nil {
					continue
				}
			}
			f := frequency{Every: "day", At: "08:00"}
			if sc.FrequencyRuleKey != nil {
				v, err := s.RuleValue(ctx, *sc.FrequencyRuleKey, &site)
				if err != nil {
					continue
				}
				if err := json.Unmarshal(v.Value, &f); err != nil {
					continue
				}
			}
			if !f.dueOn(now) {
				continue
			}
			due := atTime(now, f.At)
			if sc.DueOffsetMinutes != nil {
				due = due.Add(time.Duration(*sc.DueOffsetMinutes) * time.Minute)
			}
			date := pgtype.Date{Time: now, Valid: true}
			row, err := q.CreateTaskRow(ctx, db.CreateTaskRowParams{
				Title: sc.Title, ScheduleID: &sc.ID, SiteID: site, BusinessDate: date, AssigneeRoleID: &sc.AssigneeRoleID,
				DueAt: &due, Priority: ptr("normal"),
			})
			if errors.Is(err, pgx.ErrNoRows) {
				continue // Already generated today.
			}
			if err != nil {
				return wrap(err, "generate task")
			}
			if err := s.notifyRole(ctx, sc.AssigneeRoleID, &site, "tasks", "New task", sc.Title, "tasks", row.ID); err != nil {
				return err
			}
		}
	}
	return nil
}

func atTime(day time.Time, hhmm string) time.Time {
	t, err := time.Parse("15:04", hhmm)
	if err != nil {
		return day
	}
	return time.Date(day.Year(), day.Month(), day.Day(), t.Hour(), t.Minute(), 0, 0, Harare)
}

// Tick runs the time-driven work once: schedules, reminders, overdue flags and
// escalation timers. Tests call it directly.
func (s *Service) Tick(ctx context.Context) error {
	return s.InTx(ctx, func(ctx context.Context) error {
		q := s.q(ctx)
		if err := s.generateTasks(ctx); err != nil {
			return err
		}
		due, err := q.TasksNeedingReminder(ctx)
		if err != nil {
			return fmt.Errorf("find reminders: %w", err)
		}
		for _, t := range due {
			if t.AssigneeUserID != nil {
				if err := s.notify(ctx, []uuid.UUID{*t.AssigneeUserID}, "tasks", "Task due soon", t.Title, "tasks", t.ID); err != nil {
					return err
				}
			} else if t.AssigneeRoleID != nil {
				if err := s.notifyRole(ctx, *t.AssigneeRoleID, &t.SiteID, "tasks", "Task due soon", t.Title, "tasks", t.ID); err != nil {
					return err
				}
			}
			if err := q.MarkReminded(ctx, t.ID); err != nil {
				return fmt.Errorf("mark reminded: %w", err)
			}
		}
		late, err := q.MarkOverdueTasks(ctx)
		if err != nil {
			return fmt.Errorf("mark overdue: %w", err)
		}
		for _, t := range late {
			if err := s.taskOverdue(ctx, t); err != nil {
				return err
			}
		}
		return s.checkEscalationTimers(ctx)
	})
}

func (s *Service) taskOverdue(ctx context.Context, t db.Task) error {
	if t.ScheduleID == nil {
		return nil
	}
	sc, err := s.q(ctx).GetSchedule(ctx, *t.ScheduleID)
	if err != nil {
		return wrap(err, "read schedule")
	}
	if sc.EscalateToRoleID != nil {
		return s.notifyRole(ctx, *sc.EscalateToRoleID, &t.SiteID, "tasks", "Overdue task", t.Title, "tasks", t.ID)
	}
	return nil
}

// RunBackground ticks once a minute until ctx ends.
func (s *Service) RunBackground(ctx context.Context, logger *slog.Logger) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		if err := s.Tick(ctx); err != nil {
			logger.ErrorContext(ctx, "background tick failed", "error", err)
		}
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}
