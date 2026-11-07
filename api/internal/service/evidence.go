package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

const uploadURLTTL = 15 * time.Minute

func (s *Service) fileSignature(id uuid.UUID, method string, expires int64) string {
	mac := hmac.New(sha256.New, []byte(s.config.SigningSecret))
	mac.Write(fmt.Appendf(nil, "%s|%s|%d", method, id, expires))
	return hex.EncodeToString(mac.Sum(nil))
}

func (s *Service) fileURL(id uuid.UUID, method string) string {
	expires := s.now().Add(uploadURLTTL).Unix()
	return fmt.Sprintf("%s/files/%s?exp=%d&sig=%s", strings.TrimRight(s.config.PublicBaseURL, "/"), id, expires, s.fileSignature(id, method, expires))
}

// CheckFileURL verifies a signed file URL.
func (s *Service) CheckFileURL(id uuid.UUID, method, exp, sig string) bool {
	expires, err := strconv.ParseInt(exp, 10, 64)
	if err != nil || s.now().Unix() > expires {
		return false
	}
	return hmac.Equal([]byte(sig), []byte(s.fileSignature(id, method, expires)))
}

func (s *Service) filePath(id uuid.UUID) string {
	return filepath.Join(s.config.UploadDir, id.String())
}

// StoreFile saves an upload behind a signed URL.
func (s *Service) StoreFile(id uuid.UUID, body io.Reader) error {
	if err := os.MkdirAll(s.config.UploadDir, 0o750); err != nil {
		return fmt.Errorf("create upload dir: %w", err)
	}
	f, err := os.Create(s.filePath(id))
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer f.Close() //nolint:errcheck // written data is synced by Copy's success path
	if _, err := io.Copy(f, body); err != nil {
		return fmt.Errorf("write file: %w", err)
	}
	return nil
}

// OpenFile opens a stored upload for download.
func (s *Service) OpenFile(ctx context.Context, id uuid.UUID) (*os.File, db.Attachment, error) {
	row, err := s.queries.GetAttachment(ctx, id)
	if err != nil {
		return nil, row, notFound(err, "file")
	}
	f, err := os.Open(s.filePath(id))
	if err != nil {
		return nil, row, errNotFound("file")
	}
	return f, row, nil
}

// CreateAttachment records a file and returns where to upload it.
func (s *Service) CreateAttachment(ctx context.Context, p db.CreateAttachmentParams) (db.Attachment, error) {
	if p.SizeBytes <= 0 {
		return db.Attachment{}, errInvalid("/size_bytes", "The file is empty.")
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateAttachment(ctx, p)
	if err != nil {
		return row, wrap(err, "create attachment")
	}
	url := s.fileURL(row.ID, http.MethodPut)
	row.UploadUrl = &url
	return row, s.audit(ctx, "evidence.attachment.write", "attachments", row.ID, nil, nil, row)
}

// CompleteAttachment checks the uploaded bytes against the declared checksum.
func (s *Service) CompleteAttachment(ctx context.Context, p db.ActCompleteAttachmentParams) (db.Attachment, error) {
	before, err := s.GetAttachment(ctx, p.ID)
	if err != nil {
		return before, err
	}
	f, err := os.Open(s.filePath(p.ID))
	if err != nil {
		return before, errState("The file has not been uploaded yet.")
	}
	defer f.Close() //nolint:errcheck // read-only
	sum := sha256.New()
	size, err := io.Copy(sum, f)
	if err != nil {
		return before, fmt.Errorf("hash file: %w", err)
	}
	if hex.EncodeToString(sum.Sum(nil)) != strings.ToLower(before.Sha256) || size != int64(before.SizeBytes) {
		if _, err := s.q(ctx).SetAttachmentStatus(ctx, db.SetAttachmentStatusParams{ID: p.ID, Status: "quarantined", UpdatedBy: actorID(ctx)}); err != nil {
			return before, wrap(err, "quarantine file")
		}
		return before, errInvalid("/sha256", "The upload does not match its checksum.")
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActCompleteAttachment(ctx, p)
	if err != nil {
		return before, stateError(err, "complete this upload")
	}
	return row, s.audit(ctx, "evidence.attachment.complete", "attachments", row.ID, nil, before, row)
}

// DownloadURL returns a signed link for an available file.
func (s *Service) DownloadURL(row db.Attachment) *string {
	if row.Status != "available" {
		return nil
	}
	return ptr(s.fileURL(row.ID, http.MethodGet))
}
