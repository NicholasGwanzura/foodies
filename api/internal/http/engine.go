package httpapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"foodies/api/internal/db"
	"foodies/api/internal/http/apigen"
	"foodies/api/internal/service"
)

// ListTasks lists tasks, with assignee=me for the inbox.
func (s *Server) ListTasks(ctx context.Context, req apigen.ListTasksRequestObject) (apigen.ListTasksResponseObject, error) {
	assignee := ""
	if req.Params.Assignee != nil {
		assignee = *req.Params.Assignee
	}
	return listPage[apigen.ListTasks200JSONResponse](ctx, req.Params, map[string]any{"assignee": nil},
		func(ctx context.Context, p db.ListTasksForParams) ([]db.Task, error) {
			return s.service.ListTasks(ctx, p, assignee)
		})
}

// CreateTask adds a task.
func (s *Server) CreateTask(ctx context.Context, req apigen.CreateTaskRequestObject) (apigen.CreateTaskResponseObject, error) {
	in, err := input[db.CreateTaskRowParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateTask(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Task](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateTask201JSONResponse{Body: body, Headers: apigen.CreateTask201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

func taskBody(row db.Task) (apigen.Task, *string, error) {
	body, err := output[apigen.Task](row)
	return body, etag(row.RowVersion), err
}

// ClaimTask takes a task.
func (s *Server) ClaimTask(ctx context.Context, req apigen.ClaimTaskRequestObject) (apigen.ClaimTaskResponseObject, error) {
	row, err := s.service.ClaimTask(ctx, req.TaskId)
	if err != nil {
		return nil, err
	}
	body, tag, err := taskBody(row)
	if err != nil {
		return nil, err
	}
	return apigen.ClaimTask200JSONResponse{Body: body, Headers: apigen.ClaimTask200ResponseHeaders{ETag: tag}}, nil
}

// CompleteTask finishes a task.
func (s *Server) CompleteTask(ctx context.Context, req apigen.CompleteTaskRequestObject) (apigen.CompleteTaskResponseObject, error) {
	row, err := s.service.CompleteTask(ctx, req.TaskId)
	if err != nil {
		return nil, err
	}
	body, tag, err := taskBody(row)
	if err != nil {
		return nil, err
	}
	return apigen.CompleteTask200JSONResponse{Body: body, Headers: apigen.CompleteTask200ResponseHeaders{ETag: tag}}, nil
}

// SkipTask skips a task with a reason.
func (s *Server) SkipTask(ctx context.Context, req apigen.SkipTaskRequestObject) (apigen.SkipTaskResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	row, err := s.service.SkipTask(ctx, req.TaskId, req.Body.Reason, req.Body.AlternativeControl)
	if err != nil {
		return nil, err
	}
	body, tag, err := taskBody(row)
	if err != nil {
		return nil, err
	}
	return apigen.SkipTask200JSONResponse{Body: body, Headers: apigen.SkipTask200ResponseHeaders{ETag: tag}}, nil
}

// ReassignTask gives a task to someone.
func (s *Server) ReassignTask(ctx context.Context, req apigen.ReassignTaskRequestObject) (apigen.ReassignTaskResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	row, err := s.service.ReassignTask(ctx, req.TaskId, req.Body.UserId)
	if err != nil {
		return nil, err
	}
	body, tag, err := taskBody(row)
	if err != nil {
		return nil, err
	}
	return apigen.ReassignTask200JSONResponse{Body: body, Headers: apigen.ReassignTask200ResponseHeaders{ETag: tag}}, nil
}

// ListApprovalRequests lists approvals, with assignee=me for the inbox.
func (s *Server) ListApprovalRequests(ctx context.Context, req apigen.ListApprovalRequestsRequestObject) (apigen.ListApprovalRequestsResponseObject, error) {
	if req.Params.Assignee != nil && *req.Params.Assignee == "me" {
		return listPage[apigen.ListApprovalRequests200JSONResponse](ctx, struct {
			Limit  *int    `json:"limit,omitempty"`
			Cursor *string `json:"cursor,omitempty"`
		}{req.Params.Limit, req.Params.Cursor}, nil, func(ctx context.Context, p db.ApprovalInboxParams) ([]db.ApprovalRequest, error) {
			return s.service.ApprovalInbox(ctx, p.CursorID, p.RowLimit)
		})
	}
	return listPage[apigen.ListApprovalRequests200JSONResponse](ctx, req.Params, nil, s.service.ListApprovalRequest)
}

func approvalBody(ctx context.Context, s *Server, row db.ApprovalRequest) (apigen.ApprovalRequest, *string, error) {
	kids, err := s.service.ApprovalRequestChildren(ctx, row.ID)
	if err != nil {
		return apigen.ApprovalRequest{}, nil, err
	}
	body, err := output[apigen.ApprovalRequest](row, kids)
	return body, etag(row.RowVersion), err
}

// ApproveApprovalRequest approves the open step.
func (s *Server) ApproveApprovalRequest(ctx context.Context, req apigen.ApproveApprovalRequestRequestObject) (apigen.ApproveApprovalRequestResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	row, err := s.service.ApproveApprovalRequest(ctx, req.ApprovalRequestId, req.Body.Comment, req.Body.SignatureId)
	if err != nil {
		return nil, err
	}
	body, tag, err := approvalBody(ctx, s, row)
	if err != nil {
		return nil, err
	}
	return apigen.ApproveApprovalRequest200JSONResponse{Body: body, Headers: apigen.ApproveApprovalRequest200ResponseHeaders{ETag: tag}}, nil
}

// RejectApprovalRequest rejects the request.
func (s *Server) RejectApprovalRequest(ctx context.Context, req apigen.RejectApprovalRequestRequestObject) (apigen.RejectApprovalRequestResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	row, err := s.service.RejectApprovalRequest(ctx, req.ApprovalRequestId, req.Body.Comment, req.Body.SignatureId)
	if err != nil {
		return nil, err
	}
	body, tag, err := approvalBody(ctx, s, row)
	if err != nil {
		return nil, err
	}
	return apigen.RejectApprovalRequest200JSONResponse{Body: body, Headers: apigen.RejectApprovalRequest200ResponseHeaders{ETag: tag}}, nil
}

// ReturnApprovalRequest sends the record back.
func (s *Server) ReturnApprovalRequest(ctx context.Context, req apigen.ReturnApprovalRequestRequestObject) (apigen.ReturnApprovalRequestResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	row, err := s.service.ReturnApprovalRequest(ctx, req.ApprovalRequestId, req.Body.Comment)
	if err != nil {
		return nil, err
	}
	body, tag, err := approvalBody(ctx, s, row)
	if err != nil {
		return nil, err
	}
	return apigen.ReturnApprovalRequest200JSONResponse{Body: body, Headers: apigen.ReturnApprovalRequest200ResponseHeaders{ETag: tag}}, nil
}

// ReassignApprovalRequest hands the step to someone.
func (s *Server) ReassignApprovalRequest(ctx context.Context, req apigen.ReassignApprovalRequestRequestObject) (apigen.ReassignApprovalRequestResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	row, err := s.service.ReassignApprovalRequest(ctx, req.ApprovalRequestId, req.Body.UserId, req.Body.Reason)
	if err != nil {
		return nil, err
	}
	body, tag, err := approvalBody(ctx, s, row)
	if err != nil {
		return nil, err
	}
	return apigen.ReassignApprovalRequest200JSONResponse{Body: body, Headers: apigen.ReassignApprovalRequest200ResponseHeaders{ETag: tag}}, nil
}

// CreateSignatureRequest asks someone to sign.
func (s *Server) CreateSignatureRequest(ctx context.Context, req apigen.CreateSignatureRequestRequestObject) (apigen.CreateSignatureRequestResponseObject, error) {
	in, err := input[db.CreateSignatureRequestRowParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateSignatureRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SignatureRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateSignatureRequest201JSONResponse{Body: body, Headers: apigen.CreateSignatureRequest201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SignSignatureRequest signs with a PIN.
func (s *Server) SignSignatureRequest(ctx context.Context, req apigen.SignSignatureRequestRequestObject) (apigen.SignSignatureRequestResponseObject, error) {
	row, err := s.service.SignSignatureRequest(ctx, req.SignatureRequestId, str(req.Body.Pin))
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SignatureRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.SignSignatureRequest200JSONResponse{Body: body, Headers: apigen.SignSignatureRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// DeclineSignatureRequest refuses to sign.
func (s *Server) DeclineSignatureRequest(ctx context.Context, req apigen.DeclineSignatureRequestRequestObject) (apigen.DeclineSignatureRequestResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	row, err := s.service.DeclineSignatureRequest(ctx, req.SignatureRequestId, req.Body.Reason)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SignatureRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.DeclineSignatureRequest200JSONResponse{Body: body, Headers: apigen.DeclineSignatureRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// GetAttachment returns a file record with a fresh download link.
func (s *Server) GetAttachment(ctx context.Context, req apigen.GetAttachmentRequestObject) (apigen.GetAttachmentResponseObject, error) {
	row, err := s.service.GetAttachment(ctx, req.AttachmentId)
	if err != nil {
		return nil, err
	}
	row.DownloadUrl = s.service.DownloadURL(row)
	body, err := output[apigen.Attachment](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetAttachment200JSONResponse{Body: body, Headers: apigen.GetAttachment200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListNotifications lists the caller's notifications.
func (s *Server) ListNotifications(ctx context.Context, req apigen.ListNotificationsRequestObject) (apigen.ListNotificationsResponseObject, error) {
	extra := map[string]any{"unread_only": req.Params.Unread != nil && *req.Params.Unread}
	return listPage[apigen.ListNotifications200JSONResponse](ctx, req.Params, extra, s.service.ListNotifications)
}

// GetNotification returns one notification.
func (s *Server) GetNotification(ctx context.Context, req apigen.GetNotificationRequestObject) (apigen.GetNotificationResponseObject, error) {
	row, err := s.service.GetNotification(ctx, req.NotificationId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Notification](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetNotification200JSONResponse{Body: body, Headers: apigen.GetNotification200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReadNotification marks one read.
func (s *Server) ReadNotification(ctx context.Context, req apigen.ReadNotificationRequestObject) (apigen.ReadNotificationResponseObject, error) {
	row, err := s.service.ReadNotification(ctx, req.NotificationId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Notification](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReadNotification200JSONResponse{Body: body, Headers: apigen.ReadNotification200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// MarkAllNotificationsRead clears the unread count.
func (s *Server) MarkAllNotificationsRead(ctx context.Context, _ apigen.MarkAllNotificationsReadRequestObject) (apigen.MarkAllNotificationsReadResponseObject, error) {
	if err := s.service.ReadAllNotifications(ctx); err != nil {
		return nil, err
	}
	return apigen.MarkAllNotificationsRead204Response{}, nil
}

// GetNotificationPreferences returns channel choices.
func (s *Server) GetNotificationPreferences(ctx context.Context, _ apigen.GetNotificationPreferencesRequestObject) (apigen.GetNotificationPreferencesResponseObject, error) {
	raw, err := s.service.NotificationPreferences(ctx)
	if err != nil {
		return nil, err
	}
	return output[apigen.GetNotificationPreferences200JSONResponse](map[string]any{"categories": raw})
}

// SetNotificationPreferences stores channel choices.
func (s *Server) SetNotificationPreferences(ctx context.Context, req apigen.SetNotificationPreferencesRequestObject) (apigen.SetNotificationPreferencesResponseObject, error) {
	raw, err := json.Marshal(req.Body.Categories)
	if err != nil {
		return nil, fmt.Errorf("encode preferences: %w", err)
	}
	if err := s.service.SetNotificationPreferences(ctx, raw); err != nil {
		return nil, err
	}
	return apigen.SetNotificationPreferences200JSONResponse(*req.Body), nil
}

// GetReadiness shows what stops workflows going live.
func (s *Server) GetReadiness(ctx context.Context, req apigen.GetReadinessRequestObject) (apigen.GetReadinessResponseObject, error) {
	rows, err := s.service.Readiness(ctx, req.Params.SiteId)
	if err != nil {
		return nil, err
	}
	items := make([]map[string]any, 0, len(rows))
	for _, r := range rows {
		module, _, _ := strings.Cut(r.WorkflowKey, ".")
		if req.Params.Module != nil && *req.Params.Module != module {
			continue
		}
		owners := map[string]bool{}
		for _, key := range r.MissingRuleKeys {
			if _, owner := s.service.RuleOwner(ctx, key); owner != "" {
				owners[owner] = true
			}
		}
		roles := make([]string, 0, len(owners))
		for o := range owners {
			roles = append(roles, o)
		}
		items = append(items, map[string]any{
			"module": module, "workflow_key": r.WorkflowKey, "site_id": r.SiteID, "status": r.Status,
			"missing_rule_keys": r.MissingRuleKeys, "owner_roles": roles,
		})
	}
	return output[apigen.GetReadiness200JSONResponse](map[string]any{"items": items})
}

// GetDashboard returns a dashboard.
func (s *Server) GetDashboard(ctx context.Context, req apigen.GetDashboardRequestObject) (apigen.GetDashboardResponseObject, error) {
	tiles, title, err := s.service.Dashboard(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]map[string]any, 0, len(tiles))
	for _, t := range tiles {
		out = append(out, map[string]any{"key": t.Key, "label": t.Label, "value": t.Value, "status": t.Status, "drilldown": t.Drilldown})
	}
	return output[apigen.GetDashboard200JSONResponse](map[string]any{"key": req.DashboardKey, "title": title, "as_of": time.Now().UTC(), "tiles": out})
}

// Search finds records.
func (s *Server) Search(ctx context.Context, req apigen.SearchRequestObject) (apigen.SearchResponseObject, error) {
	rows, err := s.service.Search(ctx, req.Params.Q, 25)
	if err != nil {
		return nil, err
	}
	return output[apigen.Search200JSONResponse](map[string]any{
		"data": rows, "page": map[string]any{"limit": 25, "has_more": false, "next_cursor": nil, "prev_cursor": nil},
	})
}

// GetEffectiveFxRate returns the rate in force on a date.
func (s *Server) GetEffectiveFxRate(ctx context.Context, req apigen.GetEffectiveFxRateRequestObject) (apigen.GetEffectiveFxRateResponseObject, error) {
	on := time.Now()
	if req.Params.On != nil {
		on = req.Params.On.Time
	}
	row, err := s.service.EffectiveFxRate(ctx, req.Params.FromCurrency, req.Params.ToCurrency, pgtype.Date{Time: on, Valid: true})
	if err != nil {
		return nil, err
	}
	return output[apigen.GetEffectiveFxRate200JSONResponse](row)
}

// GetSyncManifest returns what a device needs offline.
func (s *Server) GetSyncManifest(ctx context.Context, _ apigen.GetSyncManifestRequestObject) (apigen.GetSyncManifestResponseObject, error) {
	m, err := s.service.SyncManifest(ctx)
	if err != nil {
		return nil, err
	}
	return output[apigen.GetSyncManifest200JSONResponse](m)
}

// SubmitSyncBatch replays queued writes. Each one re-enters the API as the
// signed-in user, so offline work passes the same permission, rule, approval
// and audit checks as online work. client_record_id is the idempotency key.
func (s *Server) SubmitSyncBatch(ctx context.Context, req apigen.SubmitSyncBatchRequestObject) (apigen.SubmitSyncBatchResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	batch := uuid.NewString()
	items := make([]map[string]any, 0, len(req.Body.Items))
	token := s.bearerFrom(ctx)
	for _, item := range req.Body.Items {
		result := map[string]any{"client_record_id": item.ClientRecordId}
		status, replayed, body := s.replay(ctx, token, item)
		if status >= 200 && status < 300 {
			result["status"] = "applied"
			if replayed {
				result["status"] = "duplicate"
			}
			if id := recordID(body); id != "" {
				result["record_id"] = id
			}
		} else {
			raw, err := json.Marshal(item.Payload)
			if err != nil {
				return nil, fmt.Errorf("encode payload: %w", err)
			}
			conflict, err := s.service.RecordConflict(ctx, batch, req.Body.DeviceId, item.ClientRecordId, string(body), raw)
			if err != nil {
				return nil, err
			}
			result["status"], result["conflict_id"] = "conflict", conflict
		}
		items = append(items, result)
	}
	return output[apigen.SubmitSyncBatch202JSONResponse](map[string]any{"batch_id": batch, "status": "completed", "items": items})
}

func recordID(body []byte) string {
	var row struct {
		ID string `json:"id"`
	}
	if json.Unmarshal(body, &row) != nil {
		return ""
	}
	return row.ID
}

type tokenKey struct{}

func (s *Server) bearerFrom(ctx context.Context) string {
	token, _ := ctx.Value(tokenKey{}).(string)
	return token
}

func (s *Server) replay(ctx context.Context, token string, item apigen.SyncItem) (int, bool, []byte) {
	op, ok := apigen.Operations[strings.ToUpper(item.OperationId[:1])+item.OperationId[1:]]
	if !ok || op.Method != http.MethodPost && op.Method != http.MethodPatch {
		return http.StatusUnprocessableEntity, false, []byte("Unknown operation")
	}
	path := op.Path
	if item.PathParams != nil {
		for name, value := range *item.PathParams {
			path = strings.ReplaceAll(path, "{"+name+"}", fmt.Sprint(value))
		}
	}
	raw, err := json.Marshal(item.Payload)
	if err != nil {
		return http.StatusUnprocessableEntity, false, []byte("Bad payload")
	}
	r, err := http.NewRequestWithContext(context.WithoutCancel(ctx), op.Method, basePath+path, bytes.NewReader(raw))
	if err != nil {
		return http.StatusUnprocessableEntity, false, []byte("Bad request")
	}
	r.Header.Set("Authorization", "Bearer "+token)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Idempotency-Key", item.ClientRecordId.String())
	if item.BaseRowVersion != nil {
		r.Header.Set("If-Match", fmt.Sprint(*item.BaseRowVersion))
	}
	rec := httptest.NewRecorder()
	s.routes.ServeHTTP(rec, r)
	return rec.Code, rec.Header().Get("Idempotent-Replayed") == "true", rec.Body.Bytes()
}

// VerifyWhatsappWebhook answers Meta's handshake.
func (s *Server) VerifyWhatsappWebhook(_ context.Context, req apigen.VerifyWhatsappWebhookRequestObject) (apigen.VerifyWhatsappWebhookResponseObject, error) {
	if !s.service.VerifyWhatsAppToken(str(req.Params.HubVerifyToken)) {
		return nil, service.Fail(http.StatusForbidden, "forbidden", "Not allowed", "Wrong verify token.")
	}
	return apigen.VerifyWhatsappWebhook200TextResponse(str(req.Params.HubChallenge)), nil
}

type rawBodyKey struct{}

// webhookBody keeps the raw body, which the signature covers.
func webhookBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, basePath+"/integrations/inbound/") {
			ctx := context.WithValue(r.Context(), integrationKey{}, r.Header.Get("X-Integration-Key"))
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		if !strings.HasPrefix(r.URL.Path, basePath+"/webhooks/") {
			next.ServeHTTP(w, r)
			return
		}
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			requestError(w, r, err)
			return
		}
		r.Body = io.NopCloser(bytes.NewReader(raw))
		ctx := context.WithValue(r.Context(), rawBodyKey{}, raw)
		ctx = context.WithValue(ctx, signatureKey{}, r.Header.Get("X-Hub-Signature-256"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type signatureKey struct{}

func (s *Server) verifiedBody(ctx context.Context) ([]byte, error) {
	raw, _ := ctx.Value(rawBodyKey{}).([]byte)
	header, _ := ctx.Value(signatureKey{}).(string)
	if !s.service.VerifyWhatsAppSignature(raw, header) {
		return nil, service.Fail(http.StatusUnauthorized, "unauthenticated", "Bad signature", "The webhook signature does not match.")
	}
	return raw, nil
}

// ReceiveWhatsappWebhook logs inbound messages.
func (s *Server) ReceiveWhatsappWebhook(ctx context.Context, _ apigen.ReceiveWhatsappWebhookRequestObject) (apigen.ReceiveWhatsappWebhookResponseObject, error) {
	raw, err := s.verifiedBody(ctx)
	if err != nil {
		return nil, err
	}
	if err := s.service.ReceiveWhatsApp(ctx, raw); err != nil {
		return nil, err
	}
	return apigen.ReceiveWhatsappWebhook200Response{}, nil
}

// serveFile handles signed upload and download URLs.
func (s *Server) serveFile(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	q := r.URL.Query()
	if err != nil || !s.service.CheckFileURL(id, r.Method, q.Get("exp"), q.Get("sig")) {
		writeProblem(w, r, problem{status: http.StatusForbidden, code: apigen.ErrorCodeForbidden, typeName: "forbidden", title: "Link expired"})
		return
	}
	if r.Method == http.MethodPut {
		if err := s.service.StoreFile(id, r.Body); err != nil {
			s.logger.ErrorContext(r.Context(), "store upload", "error", err)
			writeInternalError(w, r)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}
	f, row, err := s.service.OpenFile(r.Context(), id)
	if err != nil {
		notFound(w, r)
		return
	}
	defer f.Close() //nolint:errcheck // read-only
	w.Header().Set("Content-Type", row.ContentType)
	http.ServeContent(w, r, row.Filename, row.UpdatedAt, f)
}
