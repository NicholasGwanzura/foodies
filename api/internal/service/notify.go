package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
)

// notify writes an in-app notification for each user. Other channels attach to
// these rows when their providers are connected.
func (s *Service) notify(ctx context.Context, users []uuid.UUID, category, title, body, subjectType string, subject uuid.UUID) error {
	channel := "in_app"
	for _, u := range users {
		_, err := s.q(ctx).CreateNotification(ctx, db.CreateNotificationParams{
			RecipientUserID: &u, Category: &category, Title: &title, Body: &body, Channel: &channel,
			SubjectType: &subjectType, SubjectID: &subject,
		})
		if err != nil {
			return wrap(err, "create notification")
		}
	}
	return nil
}

func (s *Service) notifyRole(ctx context.Context, role uuid.UUID, site *uuid.UUID, category, title, body, subjectType string, subject uuid.UUID) error {
	users, err := s.q(ctx).UsersWithRoleAtSite(ctx, db.UsersWithRoleAtSiteParams{RoleID: role, SiteID: site})
	if err != nil {
		return fmt.Errorf("find role holders: %w", err)
	}
	return s.notify(ctx, users, category, title, body, subjectType, subject)
}

// ListNotifications returns the signed-in user's notifications.
func (s *Service) ListNotifications(ctx context.Context, p db.ListMyNotificationsParams) ([]db.Notification, error) {
	p.UserID = ptr(auth.From(ctx).UserID)
	return wrap2(s.q(ctx).ListMyNotifications(ctx, p))("list notifications")
}

// GetNotification returns one of the signed-in user's notifications.
func (s *Service) GetNotification(ctx context.Context, id uuid.UUID) (db.Notification, error) {
	row, err := s.q(ctx).GetMyNotification(ctx, db.GetMyNotificationParams{ID: id, RecipientUserID: ptr(auth.From(ctx).UserID)})
	if err != nil {
		return row, notFound(err, "notification")
	}
	return row, nil
}

// ReadNotification marks one notification read.
func (s *Service) ReadNotification(ctx context.Context, id uuid.UUID) (db.Notification, error) {
	row, err := s.q(ctx).ReadNotification(ctx, db.ReadNotificationParams{ID: id, RecipientUserID: ptr(auth.From(ctx).UserID)})
	if err != nil {
		return row, notFound(err, "notification")
	}
	return row, nil
}

// ReadAllNotifications marks everything read.
func (s *Service) ReadAllNotifications(ctx context.Context) error {
	_, err := s.q(ctx).ReadAllNotifications(ctx, ptr(auth.From(ctx).UserID))
	return wrap(err, "mark notifications read")
}

// NotificationPreferences returns channel choices by category.
func (s *Service) NotificationPreferences(ctx context.Context) (json.RawMessage, error) {
	raw, err := s.q(ctx).GetNotificationPreferences(ctx, auth.From(ctx).UserID)
	if errors.Is(err, pgx.ErrNoRows) {
		return json.RawMessage(`{}`), nil
	}
	return raw, wrap(err, "read preferences")
}

// SetNotificationPreferences stores channel choices.
func (s *Service) SetNotificationPreferences(ctx context.Context, categories json.RawMessage) error {
	return wrap(s.q(ctx).SetNotificationPreferences(ctx, db.SetNotificationPreferencesParams{UserID: auth.From(ctx).UserID, Categories: categories}), "store preferences")
}
