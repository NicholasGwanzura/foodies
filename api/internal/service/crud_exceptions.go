// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package service

import (
	"context"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

var _ uuid.UUID

// ListNcr returns a page of ncrs.
func (s *Service) ListNcr(ctx context.Context, p db.ListNcrParams) ([]db.Ncr, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListNcr(ctx, p))("list ncrs")
}

// GetNcr returns one row of ncrs.
func (s *Service) GetNcr(ctx context.Context, id uuid.UUID) (db.Ncr, error) {
	row, err := s.q(ctx).GetNcr(ctx, id)
	if err != nil {
		return row, notFound(err, "ncrs")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateNcr adds one row to ncrs.
func (s *Service) CreateNcr(ctx context.Context, p db.CreateNcrParams) (db.Ncr, error) {
	var zero db.Ncr
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateNcr", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "NCR", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateNcr(ctx, p)
	if err != nil {
		return zero, wrap(err, "create ncrs")
	}
	if err := s.audit(ctx, "quality.ncr.write", "ncrs", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateNcr", row)
}

// UpdateNcr changes one row of ncrs.
func (s *Service) UpdateNcr(ctx context.Context, p db.UpdateNcrParams) (db.Ncr, error) {
	before, err := s.GetNcr(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateNcr(ctx, p)
	if err != nil {
		return before, wrap(err, "update ncrs")
	}
	if err := s.voidSignatures(ctx, "ncrs", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "quality.ncr.write", "ncrs", row.ID, &row.SiteID, before, row)
}

// ContainNcr is the record containment action on ncrs.
func (s *Service) ContainNcr(ctx context.Context, p db.ActContainNcrParams) (db.Ncr, error) {
	before, err := s.GetNcr(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "quality.ncr.contain", "ncrs", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ContainNcr", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActContainNcr(ctx, p)
	if err != nil {
		return before, stateError(err, "contain ncrs")
	}
	if err := s.audit(ctx, "quality.ncr.contain", "ncrs", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ContainNcr", row)
}

// VerifyCloseNcr is the verify effectiveness and close action on ncrs.
func (s *Service) VerifyCloseNcr(ctx context.Context, p db.ActVerifyCloseNcrParams) (db.Ncr, error) {
	before, err := s.GetNcr(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "quality.ncr.verify_close", "ncrs", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "VerifyCloseNcr", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActVerifyCloseNcr(ctx, p)
	if err != nil {
		return before, stateError(err, "verify close ncrs")
	}
	if err := s.audit(ctx, "quality.ncr.verify_close", "ncrs", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "VerifyCloseNcr", row)
}

// ListHold returns a page of holds.
func (s *Service) ListHold(ctx context.Context, p db.ListHoldParams) ([]db.Hold, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListHold(ctx, p))("list holds")
}

// GetHold returns one row of holds.
func (s *Service) GetHold(ctx context.Context, id uuid.UUID) (db.Hold, error) {
	row, err := s.q(ctx).GetHold(ctx, id)
	if err != nil {
		return row, notFound(err, "holds")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateHold adds one row to holds.
func (s *Service) CreateHold(ctx context.Context, p db.CreateHoldParams) (db.Hold, error) {
	var zero db.Hold
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateHold", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "HLD", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateHold(ctx, p)
	if err != nil {
		return zero, wrap(err, "create holds")
	}
	if err := s.audit(ctx, "quality.hold.write", "holds", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateHold", row)
}

// UpdateHold changes one row of holds.
func (s *Service) UpdateHold(ctx context.Context, p db.UpdateHoldParams) (db.Hold, error) {
	before, err := s.GetHold(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateHold(ctx, p)
	if err != nil {
		return before, wrap(err, "update holds")
	}
	if err := s.voidSignatures(ctx, "holds", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "quality.hold.write", "holds", row.ID, &row.SiteID, before, row)
}

// ReleaseHold is the authorised release with qc evidence action on holds.
func (s *Service) ReleaseHold(ctx context.Context, p db.ActReleaseHoldParams) (db.Hold, error) {
	before, err := s.GetHold(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "quality.hold.release", "holds", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReleaseHold", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReleaseHold(ctx, p)
	if err != nil {
		return before, stateError(err, "release holds")
	}
	if err := s.audit(ctx, "quality.hold.release", "holds", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReleaseHold", row)
}

// ListDisposition returns a page of dispositions.
func (s *Service) ListDisposition(ctx context.Context, p db.ListDispositionParams) ([]db.Disposition, error) {
	return wrap2(s.q(ctx).ListDisposition(ctx, p))("list dispositions")
}

// GetDisposition returns one row of dispositions.
func (s *Service) GetDisposition(ctx context.Context, id uuid.UUID) (db.Disposition, error) {
	row, err := s.q(ctx).GetDisposition(ctx, id)
	if err != nil {
		return row, notFound(err, "dispositions")
	}
	return row, nil
}

// CreateDisposition adds one row to dispositions.
func (s *Service) CreateDisposition(ctx context.Context, p db.CreateDispositionParams) (db.Disposition, error) {
	var zero db.Disposition
	if err := s.guard(ctx, "CreateDisposition", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "DIS", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateDisposition(ctx, p)
	if err != nil {
		return zero, wrap(err, "create dispositions")
	}
	if err := s.audit(ctx, "quality.disposition.write", "dispositions", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateDisposition", row)
}

// SubmitDisposition is the submit action on dispositions.
func (s *Service) SubmitDisposition(ctx context.Context, p db.ActSubmitDispositionParams) (db.Disposition, error) {
	before, err := s.GetDisposition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "quality.disposition.submit", "dispositions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitDisposition", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitDisposition(ctx, p)
	if err != nil {
		return before, stateError(err, "submit dispositions")
	}
	if err := s.audit(ctx, "quality.disposition.submit", "dispositions", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "dispositions", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitDisposition", row)
}

// ExecuteDisposition is the record execution action on dispositions.
func (s *Service) ExecuteDisposition(ctx context.Context, p db.ActExecuteDispositionParams) (db.Disposition, error) {
	before, err := s.GetDisposition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "quality.disposition.execute", "dispositions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ExecuteDisposition", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActExecuteDisposition(ctx, p)
	if err != nil {
		return before, stateError(err, "execute dispositions")
	}
	if err := s.audit(ctx, "quality.disposition.execute", "dispositions", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ExecuteDisposition", row)
}

// ListRecall returns a page of recalls.
func (s *Service) ListRecall(ctx context.Context, p db.ListRecallParams) ([]db.Recall, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListRecall(ctx, p))("list recalls")
}

// GetRecall returns one row of recalls.
func (s *Service) GetRecall(ctx context.Context, id uuid.UUID) (db.Recall, error) {
	row, err := s.q(ctx).GetRecall(ctx, id)
	if err != nil {
		return row, notFound(err, "recalls")
	}
	return row, inScope(ctx, row.SiteID)
}

// CreateRecall adds one row to recalls.
func (s *Service) CreateRecall(ctx context.Context, p db.CreateRecallParams) (db.Recall, error) {
	var zero db.Recall
	if err := canWrite(ctx, p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateRecall", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "RCL", p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateRecall(ctx, p)
	if err != nil {
		return zero, wrap(err, "create recalls")
	}
	if err := s.audit(ctx, "quality.recall.write", "recalls", row.ID, p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateRecall", row)
}

// UpdateRecall changes one row of recalls.
func (s *Service) UpdateRecall(ctx context.Context, p db.UpdateRecallParams) (db.Recall, error) {
	before, err := s.GetRecall(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateRecall(ctx, p)
	if err != nil {
		return before, wrap(err, "update recalls")
	}
	if err := s.voidSignatures(ctx, "recalls", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "quality.recall.write", "recalls", row.ID, row.SiteID, before, row)
}

// AssessRecall is the record risk assessment action on recalls.
func (s *Service) AssessRecall(ctx context.Context, p db.ActAssessRecallParams) (db.Recall, error) {
	before, err := s.GetRecall(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "quality.recall.assess", "recalls", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "AssessRecall", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActAssessRecall(ctx, p)
	if err != nil {
		return before, stateError(err, "assess recalls")
	}
	if err := s.audit(ctx, "quality.recall.assess", "recalls", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "AssessRecall", row)
}

// DecideRecall is the record the decision for approval action on recalls.
func (s *Service) DecideRecall(ctx context.Context, p db.ActDecideRecallParams) (db.Recall, error) {
	before, err := s.GetRecall(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "quality.recall.decide", "recalls", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "DecideRecall", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActDecideRecall(ctx, p)
	if err != nil {
		return before, stateError(err, "decide recalls")
	}
	if err := s.audit(ctx, "quality.recall.decide", "recalls", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "DecideRecall", row)
}

// CloseRecall is the close with effectiveness report action on recalls.
func (s *Service) CloseRecall(ctx context.Context, p db.ActCloseRecallParams) (db.Recall, error) {
	before, err := s.GetRecall(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "quality.recall.close", "recalls", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CloseRecall", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "recalls", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCloseRecall(ctx, p)
	if err != nil {
		return before, stateError(err, "close recalls")
	}
	if err := s.audit(ctx, "quality.recall.close", "recalls", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CloseRecall", row)
}

// ListRecallCommunication returns a page of recall communications.
func (s *Service) ListRecallCommunication(ctx context.Context, p db.ListRecallCommunicationParams) ([]db.RecallCommunication, error) {
	return wrap2(s.q(ctx).ListRecallCommunication(ctx, p))("list recall communications")
}

// GetRecallCommunication returns one row of recall communications.
func (s *Service) GetRecallCommunication(ctx context.Context, id uuid.UUID) (db.RecallCommunication, error) {
	row, err := s.q(ctx).GetRecallCommunication(ctx, id)
	if err != nil {
		return row, notFound(err, "recall communications")
	}
	return row, nil
}

// CreateRecallCommunication adds one row to recall communications.
func (s *Service) CreateRecallCommunication(ctx context.Context, p db.CreateRecallCommunicationParams) (db.RecallCommunication, error) {
	var zero db.RecallCommunication
	if err := s.guard(ctx, "CreateRecallCommunication", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateRecallCommunication(ctx, p)
	if err != nil {
		return zero, wrap(err, "create recall communications")
	}
	if err := s.audit(ctx, "quality.recall_communication.write", "recall_communications", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateRecallCommunication", row)
}

// SubmitRecallCommunication is the submit action on recall communications.
func (s *Service) SubmitRecallCommunication(ctx context.Context, p db.ActSubmitRecallCommunicationParams) (db.RecallCommunication, error) {
	before, err := s.GetRecallCommunication(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "quality.recall_communication.submit", "recall_communications", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitRecallCommunication", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitRecallCommunication(ctx, p)
	if err != nil {
		return before, stateError(err, "submit recall communications")
	}
	if err := s.audit(ctx, "quality.recall_communication.submit", "recall_communications", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "recall_communications", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitRecallCommunication", row)
}

// MarkSentRecallCommunication is the record as sent action on recall communications.
func (s *Service) MarkSentRecallCommunication(ctx context.Context, p db.ActMarkSentRecallCommunicationParams) (db.RecallCommunication, error) {
	before, err := s.GetRecallCommunication(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "quality.recall_communication.mark_sent", "recall_communications", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "MarkSentRecallCommunication", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActMarkSentRecallCommunication(ctx, p)
	if err != nil {
		return before, stateError(err, "mark sent recall communications")
	}
	if err := s.audit(ctx, "quality.recall_communication.mark_sent", "recall_communications", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "MarkSentRecallCommunication", row)
}

// AcknowledgeRecallCommunication is the record acknowledgement action on recall communications.
func (s *Service) AcknowledgeRecallCommunication(ctx context.Context, p db.ActAcknowledgeRecallCommunicationParams) (db.RecallCommunication, error) {
	before, err := s.GetRecallCommunication(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "quality.recall_communication.acknowledge", "recall_communications", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "AcknowledgeRecallCommunication", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActAcknowledgeRecallCommunication(ctx, p)
	if err != nil {
		return before, stateError(err, "acknowledge recall communications")
	}
	if err := s.audit(ctx, "quality.recall_communication.acknowledge", "recall_communications", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "AcknowledgeRecallCommunication", row)
}

// ListRecallRetrieval returns a page of recall retrievals.
func (s *Service) ListRecallRetrieval(ctx context.Context, p db.ListRecallRetrievalParams) ([]db.RecallRetrieval, error) {
	return wrap2(s.q(ctx).ListRecallRetrieval(ctx, p))("list recall retrievals")
}

// GetRecallRetrieval returns one row of recall retrievals.
func (s *Service) GetRecallRetrieval(ctx context.Context, id uuid.UUID) (db.RecallRetrieval, error) {
	row, err := s.q(ctx).GetRecallRetrieval(ctx, id)
	if err != nil {
		return row, notFound(err, "recall retrievals")
	}
	return row, nil
}

// CreateRecallRetrieval adds one row to recall retrievals.
func (s *Service) CreateRecallRetrieval(ctx context.Context, p db.CreateRecallRetrievalParams) (db.RecallRetrieval, error) {
	var zero db.RecallRetrieval
	if err := s.guard(ctx, "CreateRecallRetrieval", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateRecallRetrieval(ctx, p)
	if err != nil {
		return zero, wrap(err, "create recall retrievals")
	}
	if err := s.audit(ctx, "quality.recall_retrieval.write", "recall_retrievals", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateRecallRetrieval", row)
}

// UpdateRecallRetrieval changes one row of recall retrievals.
func (s *Service) UpdateRecallRetrieval(ctx context.Context, p db.UpdateRecallRetrievalParams) (db.RecallRetrieval, error) {
	before, err := s.GetRecallRetrieval(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateRecallRetrieval(ctx, p)
	if err != nil {
		return before, wrap(err, "update recall retrievals")
	}
	if err := s.voidSignatures(ctx, "recall_retrievals", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "quality.recall_retrieval.write", "recall_retrievals", row.ID, nil, before, row)
}

// ListRecallReconciliation returns a page of recall reconciliations.
func (s *Service) ListRecallReconciliation(ctx context.Context, p db.ListRecallReconciliationParams) ([]db.RecallReconciliation, error) {
	return wrap2(s.q(ctx).ListRecallReconciliation(ctx, p))("list recall reconciliations")
}

// GetRecallReconciliation returns one row of recall reconciliations.
func (s *Service) GetRecallReconciliation(ctx context.Context, id uuid.UUID) (db.RecallReconciliation, error) {
	row, err := s.q(ctx).GetRecallReconciliation(ctx, id)
	if err != nil {
		return row, notFound(err, "recall reconciliations")
	}
	return row, nil
}

// CreateRecallReconciliation adds one row to recall reconciliations.
func (s *Service) CreateRecallReconciliation(ctx context.Context, p db.CreateRecallReconciliationParams) (db.RecallReconciliation, error) {
	var zero db.RecallReconciliation
	if err := s.guard(ctx, "CreateRecallReconciliation", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateRecallReconciliation(ctx, p)
	if err != nil {
		return zero, wrap(err, "create recall reconciliations")
	}
	if err := s.audit(ctx, "quality.recall_reconciliation.write", "recall_reconciliations", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateRecallReconciliation", row)
}

// UpdateRecallReconciliation changes one row of recall reconciliations.
func (s *Service) UpdateRecallReconciliation(ctx context.Context, p db.UpdateRecallReconciliationParams) (db.RecallReconciliation, error) {
	before, err := s.GetRecallReconciliation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateRecallReconciliation(ctx, p)
	if err != nil {
		return before, wrap(err, "update recall reconciliations")
	}
	if err := s.voidSignatures(ctx, "recall_reconciliations", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "quality.recall_reconciliation.write", "recall_reconciliations", row.ID, nil, before, row)
}

// ReviewRecallReconciliation is the independent review action on recall reconciliations.
func (s *Service) ReviewRecallReconciliation(ctx context.Context, p db.ActReviewRecallReconciliationParams) (db.RecallReconciliation, error) {
	before, err := s.GetRecallReconciliation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "quality.recall_reconciliation.review", "recall_reconciliations", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReviewRecallReconciliation", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "recall_reconciliations", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActReviewRecallReconciliation(ctx, p)
	if err != nil {
		return before, stateError(err, "review recall reconciliations")
	}
	if err := s.audit(ctx, "quality.recall_reconciliation.review", "recall_reconciliations", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReviewRecallReconciliation", row)
}

// ListIncident returns a page of incidents.
func (s *Service) ListIncident(ctx context.Context, p db.ListIncidentParams) ([]db.Incident, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListIncident(ctx, p))("list incidents")
}

// GetIncident returns one row of incidents.
func (s *Service) GetIncident(ctx context.Context, id uuid.UUID) (db.Incident, error) {
	row, err := s.q(ctx).GetIncident(ctx, id)
	if err != nil {
		return row, notFound(err, "incidents")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateIncident adds one row to incidents.
func (s *Service) CreateIncident(ctx context.Context, p db.CreateIncidentParams) (db.Incident, error) {
	var zero db.Incident
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateIncident", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "INC", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateIncident(ctx, p)
	if err != nil {
		return zero, wrap(err, "create incidents")
	}
	if err := s.audit(ctx, "incidents.incident.write", "incidents", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateIncident", row)
}

// UpdateIncident changes one row of incidents.
func (s *Service) UpdateIncident(ctx context.Context, p db.UpdateIncidentParams) (db.Incident, error) {
	before, err := s.GetIncident(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateIncident(ctx, p)
	if err != nil {
		return before, wrap(err, "update incidents")
	}
	if err := s.voidSignatures(ctx, "incidents", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "incidents.incident.write", "incidents", row.ID, &row.SiteID, before, row)
}

// CloseIncident is the close with outcome action on incidents.
func (s *Service) CloseIncident(ctx context.Context, p db.ActCloseIncidentParams) (db.Incident, error) {
	before, err := s.GetIncident(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "incidents.incident.close", "incidents", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CloseIncident", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "incidents", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCloseIncident(ctx, p)
	if err != nil {
		return before, stateError(err, "close incidents")
	}
	if err := s.audit(ctx, "incidents.incident.close", "incidents", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CloseIncident", row)
}

// ListComplaint returns a page of complaints.
func (s *Service) ListComplaint(ctx context.Context, p db.ListComplaintParams) ([]db.Complaint, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListComplaint(ctx, p))("list complaints")
}

// GetComplaint returns one row of complaints.
func (s *Service) GetComplaint(ctx context.Context, id uuid.UUID) (db.Complaint, error) {
	row, err := s.q(ctx).GetComplaint(ctx, id)
	if err != nil {
		return row, notFound(err, "complaints")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateComplaint adds one row to complaints.
func (s *Service) CreateComplaint(ctx context.Context, p db.CreateComplaintParams) (db.Complaint, error) {
	var zero db.Complaint
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateComplaint", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "CMP", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateComplaint(ctx, p)
	if err != nil {
		return zero, wrap(err, "create complaints")
	}
	if err := s.audit(ctx, "customer.complaint.write", "complaints", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateComplaint", row)
}

// UpdateComplaint changes one row of complaints.
func (s *Service) UpdateComplaint(ctx context.Context, p db.UpdateComplaintParams) (db.Complaint, error) {
	before, err := s.GetComplaint(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateComplaint(ctx, p)
	if err != nil {
		return before, wrap(err, "update complaints")
	}
	if err := s.voidSignatures(ctx, "complaints", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "customer.complaint.write", "complaints", row.ID, &row.SiteID, before, row)
}

// TriageComplaint is the set severity and owner action on complaints.
func (s *Service) TriageComplaint(ctx context.Context, p db.ActTriageComplaintParams) (db.Complaint, error) {
	before, err := s.GetComplaint(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "customer.complaint.triage", "complaints", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "TriageComplaint", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActTriageComplaint(ctx, p)
	if err != nil {
		return before, stateError(err, "triage complaints")
	}
	if err := s.audit(ctx, "customer.complaint.triage", "complaints", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "TriageComplaint", row)
}

// ResolveComplaint is the record resolution action on complaints.
func (s *Service) ResolveComplaint(ctx context.Context, p db.ActResolveComplaintParams) (db.Complaint, error) {
	before, err := s.GetComplaint(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "customer.complaint.resolve", "complaints", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ResolveComplaint", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActResolveComplaint(ctx, p)
	if err != nil {
		return before, stateError(err, "resolve complaints")
	}
	if err := s.audit(ctx, "customer.complaint.resolve", "complaints", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ResolveComplaint", row)
}

// CloseComplaint is the close action on complaints.
func (s *Service) CloseComplaint(ctx context.Context, p db.ActCloseComplaintParams) (db.Complaint, error) {
	before, err := s.GetComplaint(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "customer.complaint.close", "complaints", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CloseComplaint", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "complaints", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCloseComplaint(ctx, p)
	if err != nil {
		return before, stateError(err, "close complaints")
	}
	if err := s.audit(ctx, "customer.complaint.close", "complaints", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CloseComplaint", row)
}

// ListServiceRecovery returns a page of service recoveries.
func (s *Service) ListServiceRecovery(ctx context.Context, p db.ListServiceRecoveryParams) ([]db.ServiceRecovery, error) {
	return wrap2(s.q(ctx).ListServiceRecovery(ctx, p))("list service recoveries")
}

// GetServiceRecovery returns one row of service recoveries.
func (s *Service) GetServiceRecovery(ctx context.Context, id uuid.UUID) (db.ServiceRecovery, error) {
	row, err := s.q(ctx).GetServiceRecovery(ctx, id)
	if err != nil {
		return row, notFound(err, "service recoveries")
	}
	return row, nil
}

// CreateServiceRecovery adds one row to service recoveries.
func (s *Service) CreateServiceRecovery(ctx context.Context, p db.CreateServiceRecoveryParams) (db.ServiceRecovery, error) {
	var zero db.ServiceRecovery
	if err := s.guard(ctx, "CreateServiceRecovery", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "SRV", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateServiceRecovery(ctx, p)
	if err != nil {
		return zero, wrap(err, "create service recoveries")
	}
	if err := s.audit(ctx, "customer.service_recovery.write", "service_recoveries", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateServiceRecovery", row)
}

// SubmitServiceRecovery is the submit action on service recoveries.
func (s *Service) SubmitServiceRecovery(ctx context.Context, p db.ActSubmitServiceRecoveryParams) (db.ServiceRecovery, error) {
	before, err := s.GetServiceRecovery(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "customer.service_recovery.submit", "service_recoveries", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitServiceRecovery", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitServiceRecovery(ctx, p)
	if err != nil {
		return before, stateError(err, "submit service recoveries")
	}
	if err := s.audit(ctx, "customer.service_recovery.submit", "service_recoveries", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "service_recoveries", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitServiceRecovery", row)
}

// ExecuteServiceRecovery is the record execution action on service recoveries.
func (s *Service) ExecuteServiceRecovery(ctx context.Context, p db.ActExecuteServiceRecoveryParams) (db.ServiceRecovery, error) {
	before, err := s.GetServiceRecovery(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "customer.service_recovery.execute", "service_recoveries", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ExecuteServiceRecovery", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActExecuteServiceRecovery(ctx, p)
	if err != nil {
		return before, stateError(err, "execute service recoveries")
	}
	if err := s.audit(ctx, "customer.service_recovery.execute", "service_recoveries", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ExecuteServiceRecovery", row)
}

// ListPaymentIncident returns a page of payment incidents.
func (s *Service) ListPaymentIncident(ctx context.Context, p db.ListPaymentIncidentParams) ([]db.PaymentIncident, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListPaymentIncident(ctx, p))("list payment incidents")
}

// GetPaymentIncident returns one row of payment incidents.
func (s *Service) GetPaymentIncident(ctx context.Context, id uuid.UUID) (db.PaymentIncident, error) {
	row, err := s.q(ctx).GetPaymentIncident(ctx, id)
	if err != nil {
		return row, notFound(err, "payment incidents")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreatePaymentIncident adds one row to payment incidents.
func (s *Service) CreatePaymentIncident(ctx context.Context, p db.CreatePaymentIncidentParams) (db.PaymentIncident, error) {
	var zero db.PaymentIncident
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreatePaymentIncident", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "PMI", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreatePaymentIncident(ctx, p)
	if err != nil {
		return zero, wrap(err, "create payment incidents")
	}
	if err := s.audit(ctx, "payments.payment_incident.write", "payment_incidents", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreatePaymentIncident", row)
}

// UpdatePaymentIncident changes one row of payment incidents.
func (s *Service) UpdatePaymentIncident(ctx context.Context, p db.UpdatePaymentIncidentParams) (db.PaymentIncident, error) {
	before, err := s.GetPaymentIncident(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdatePaymentIncident(ctx, p)
	if err != nil {
		return before, wrap(err, "update payment incidents")
	}
	if err := s.voidSignatures(ctx, "payment_incidents", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "payments.payment_incident.write", "payment_incidents", row.ID, &row.SiteID, before, row)
}

// VerifyStatusPaymentIncident is the record verified status action on payment incidents.
func (s *Service) VerifyStatusPaymentIncident(ctx context.Context, p db.ActVerifyStatusPaymentIncidentParams) (db.PaymentIncident, error) {
	before, err := s.GetPaymentIncident(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "payments.payment_incident.verify_status", "payment_incidents", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "VerifyStatusPaymentIncident", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActVerifyStatusPaymentIncident(ctx, p)
	if err != nil {
		return before, stateError(err, "verify status payment incidents")
	}
	if err := s.audit(ctx, "payments.payment_incident.verify_status", "payment_incidents", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "VerifyStatusPaymentIncident", row)
}

// ProposeCorrectionPaymentIncident is the propose reversal or refund action on payment incidents.
func (s *Service) ProposeCorrectionPaymentIncident(ctx context.Context, p db.ActProposeCorrectionPaymentIncidentParams) (db.PaymentIncident, error) {
	before, err := s.GetPaymentIncident(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "payments.payment_incident.propose_correction", "payment_incidents", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ProposeCorrectionPaymentIncident", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActProposeCorrectionPaymentIncident(ctx, p)
	if err != nil {
		return before, stateError(err, "propose correction payment incidents")
	}
	if err := s.audit(ctx, "payments.payment_incident.propose_correction", "payment_incidents", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ProposeCorrectionPaymentIncident", row)
}

// ReconcilePaymentIncident is the record settlement reconciliation action on payment incidents.
func (s *Service) ReconcilePaymentIncident(ctx context.Context, p db.ActReconcilePaymentIncidentParams) (db.PaymentIncident, error) {
	before, err := s.GetPaymentIncident(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "payments.payment_incident.reconcile", "payment_incidents", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReconcilePaymentIncident", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReconcilePaymentIncident(ctx, p)
	if err != nil {
		return before, stateError(err, "reconcile payment incidents")
	}
	if err := s.audit(ctx, "payments.payment_incident.reconcile", "payment_incidents", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReconcilePaymentIncident", row)
}

// ClosePaymentIncident is the close action on payment incidents.
func (s *Service) ClosePaymentIncident(ctx context.Context, p db.ActClosePaymentIncidentParams) (db.PaymentIncident, error) {
	before, err := s.GetPaymentIncident(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "payments.payment_incident.close", "payment_incidents", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ClosePaymentIncident", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "payment_incidents", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActClosePaymentIncident(ctx, p)
	if err != nil {
		return before, stateError(err, "close payment incidents")
	}
	if err := s.audit(ctx, "payments.payment_incident.close", "payment_incidents", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ClosePaymentIncident", row)
}

// PaymentIncidentChildren returns the rows of a record's child tables, by API field.
func (s *Service) PaymentIncidentChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListPaymentIncidentUpdates(ctx, id)
		if err != nil {
			return nil, wrap(err, "list payment_incident_updates")
		}
		out["updates"] = childRows(rows)
	}
	return out, nil
}

// ReplacePaymentIncidentUpdates swaps a record's payment incident updates for the given rows.
func (s *Service) ReplacePaymentIncidentUpdates(ctx context.Context, id uuid.UUID, rows []db.InsertPaymentIncidentUpdatesParams) error {
	if err := s.q(ctx).DeletePaymentIncidentUpdates(ctx, id); err != nil {
		return wrap(err, "clear payment_incident_updates")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertPaymentIncidentUpdates(ctx, row); err != nil {
			return wrap(err, "add payment_incident_updates")
		}
	}
	return nil
}

// ListPosOutage returns a page of pos outages.
func (s *Service) ListPosOutage(ctx context.Context, p db.ListPosOutageParams) ([]db.PosOutage, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListPosOutage(ctx, p))("list pos outages")
}

// GetPosOutage returns one row of pos outages.
func (s *Service) GetPosOutage(ctx context.Context, id uuid.UUID) (db.PosOutage, error) {
	row, err := s.q(ctx).GetPosOutage(ctx, id)
	if err != nil {
		return row, notFound(err, "pos outages")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreatePosOutage adds one row to pos outages.
func (s *Service) CreatePosOutage(ctx context.Context, p db.CreatePosOutageParams) (db.PosOutage, error) {
	var zero db.PosOutage
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreatePosOutage", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "OUT", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreatePosOutage(ctx, p)
	if err != nil {
		return zero, wrap(err, "create pos outages")
	}
	if err := s.audit(ctx, "payments.pos_outage.write", "pos_outages", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreatePosOutage", row)
}

// UpdatePosOutage changes one row of pos outages.
func (s *Service) UpdatePosOutage(ctx context.Context, p db.UpdatePosOutageParams) (db.PosOutage, error) {
	before, err := s.GetPosOutage(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdatePosOutage(ctx, p)
	if err != nil {
		return before, wrap(err, "update pos outages")
	}
	if err := s.voidSignatures(ctx, "pos_outages", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "payments.pos_outage.write", "pos_outages", row.ID, &row.SiteID, before, row)
}

// SubmitPosOutage is the submit action on pos outages.
func (s *Service) SubmitPosOutage(ctx context.Context, p db.ActSubmitPosOutageParams) (db.PosOutage, error) {
	before, err := s.GetPosOutage(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "payments.pos_outage.submit", "pos_outages", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitPosOutage", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitPosOutage(ctx, p)
	if err != nil {
		return before, stateError(err, "submit pos outages")
	}
	if err := s.audit(ctx, "payments.pos_outage.submit", "pos_outages", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitPosOutage", row)
}

// RestorePosOutage is the record pos restored action on pos outages.
func (s *Service) RestorePosOutage(ctx context.Context, p db.ActRestorePosOutageParams) (db.PosOutage, error) {
	before, err := s.GetPosOutage(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "payments.pos_outage.restore", "pos_outages", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "RestorePosOutage", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRestorePosOutage(ctx, p)
	if err != nil {
		return before, stateError(err, "restore pos outages")
	}
	if err := s.audit(ctx, "payments.pos_outage.restore", "pos_outages", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "RestorePosOutage", row)
}

// ClosePosOutage is the close after reconciliation action on pos outages.
func (s *Service) ClosePosOutage(ctx context.Context, p db.ActClosePosOutageParams) (db.PosOutage, error) {
	before, err := s.GetPosOutage(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "payments.pos_outage.close", "pos_outages", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ClosePosOutage", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "pos_outages", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActClosePosOutage(ctx, p)
	if err != nil {
		return before, stateError(err, "close pos outages")
	}
	if err := s.audit(ctx, "payments.pos_outage.close", "pos_outages", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ClosePosOutage", row)
}

// ListReceiptBook returns a page of receipt books.
func (s *Service) ListReceiptBook(ctx context.Context, p db.ListReceiptBookParams) ([]db.ReceiptBook, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListReceiptBook(ctx, p))("list receipt books")
}

// GetReceiptBook returns one row of receipt books.
func (s *Service) GetReceiptBook(ctx context.Context, id uuid.UUID) (db.ReceiptBook, error) {
	row, err := s.q(ctx).GetReceiptBook(ctx, id)
	if err != nil {
		return row, notFound(err, "receipt books")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateReceiptBook adds one row to receipt books.
func (s *Service) CreateReceiptBook(ctx context.Context, p db.CreateReceiptBookParams) (db.ReceiptBook, error) {
	var zero db.ReceiptBook
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateReceiptBook", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateReceiptBook(ctx, p)
	if err != nil {
		return zero, wrap(err, "create receipt books")
	}
	if err := s.audit(ctx, "payments.receipt_book.write", "receipt_books", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateReceiptBook", row)
}

// UpdateReceiptBook changes one row of receipt books.
func (s *Service) UpdateReceiptBook(ctx context.Context, p db.UpdateReceiptBookParams) (db.ReceiptBook, error) {
	before, err := s.GetReceiptBook(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateReceiptBook(ctx, p)
	if err != nil {
		return before, wrap(err, "update receipt books")
	}
	if err := s.voidSignatures(ctx, "receipt_books", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "payments.receipt_book.write", "receipt_books", row.ID, &row.SiteID, before, row)
}

// IssueReceiptBook is the issue to a cashier action on receipt books.
func (s *Service) IssueReceiptBook(ctx context.Context, p db.ActIssueReceiptBookParams) (db.ReceiptBook, error) {
	before, err := s.GetReceiptBook(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "payments.receipt_book.issue", "receipt_books", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "IssueReceiptBook", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "receipt_books", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActIssueReceiptBook(ctx, p)
	if err != nil {
		return before, stateError(err, "issue receipt books")
	}
	if err := s.audit(ctx, "payments.receipt_book.issue", "receipt_books", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "IssueReceiptBook", row)
}

// ReturnReceiptBook is the return and account for numbers action on receipt books.
func (s *Service) ReturnReceiptBook(ctx context.Context, p db.ActReturnReceiptBookParams) (db.ReceiptBook, error) {
	before, err := s.GetReceiptBook(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "payments.receipt_book.return", "receipt_books", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReturnReceiptBook", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReturnReceiptBook(ctx, p)
	if err != nil {
		return before, stateError(err, "return receipt books")
	}
	if err := s.audit(ctx, "payments.receipt_book.return", "receipt_books", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReturnReceiptBook", row)
}

// ListManualReceipt returns a page of manual receipts.
func (s *Service) ListManualReceipt(ctx context.Context, p db.ListManualReceiptParams) ([]db.ManualReceipt, error) {
	return wrap2(s.q(ctx).ListManualReceipt(ctx, p))("list manual receipts")
}

// GetManualReceipt returns one row of manual receipts.
func (s *Service) GetManualReceipt(ctx context.Context, id uuid.UUID) (db.ManualReceipt, error) {
	row, err := s.q(ctx).GetManualReceipt(ctx, id)
	if err != nil {
		return row, notFound(err, "manual receipts")
	}
	return row, nil
}

// CreateManualReceipt adds one row to manual receipts.
func (s *Service) CreateManualReceipt(ctx context.Context, p db.CreateManualReceiptParams) (db.ManualReceipt, error) {
	var zero db.ManualReceipt
	if err := s.guard(ctx, "CreateManualReceipt", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateManualReceipt(ctx, p)
	if err != nil {
		return zero, wrap(err, "create manual receipts")
	}
	if err := s.audit(ctx, "payments.manual_receipt.write", "manual_receipts", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateManualReceipt", row)
}

// CancelManualReceipt is the cancel and retain action on manual receipts.
func (s *Service) CancelManualReceipt(ctx context.Context, p db.ActCancelManualReceiptParams) (db.ManualReceipt, error) {
	before, err := s.GetManualReceipt(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "payments.manual_receipt.cancel", "manual_receipts", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CancelManualReceipt", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActCancelManualReceipt(ctx, p)
	if err != nil {
		return before, stateError(err, "cancel manual receipts")
	}
	if err := s.audit(ctx, "payments.manual_receipt.cancel", "manual_receipts", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CancelManualReceipt", row)
}

// BackCaptureManualReceipt is the link to the pos transaction action on manual receipts.
func (s *Service) BackCaptureManualReceipt(ctx context.Context, p db.ActBackCaptureManualReceiptParams) (db.ManualReceipt, error) {
	before, err := s.GetManualReceipt(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "payments.manual_receipt.back_capture", "manual_receipts", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "BackCaptureManualReceipt", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActBackCaptureManualReceipt(ctx, p)
	if err != nil {
		return before, stateError(err, "back capture manual receipts")
	}
	if err := s.audit(ctx, "payments.manual_receipt.back_capture", "manual_receipts", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "BackCaptureManualReceipt", row)
}

// ListEquipmentFault returns a page of equipment faults.
func (s *Service) ListEquipmentFault(ctx context.Context, p db.ListEquipmentFaultParams) ([]db.EquipmentFault, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListEquipmentFault(ctx, p))("list equipment faults")
}

// GetEquipmentFault returns one row of equipment faults.
func (s *Service) GetEquipmentFault(ctx context.Context, id uuid.UUID) (db.EquipmentFault, error) {
	row, err := s.q(ctx).GetEquipmentFault(ctx, id)
	if err != nil {
		return row, notFound(err, "equipment faults")
	}
	return row, inScope(ctx, row.SiteID)
}

// CreateEquipmentFault adds one row to equipment faults.
func (s *Service) CreateEquipmentFault(ctx context.Context, p db.CreateEquipmentFaultParams) (db.EquipmentFault, error) {
	var zero db.EquipmentFault
	if err := canWrite(ctx, p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateEquipmentFault", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "FLT", p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateEquipmentFault(ctx, p)
	if err != nil {
		return zero, wrap(err, "create equipment faults")
	}
	if err := s.audit(ctx, "maintenance.equipment_fault.write", "equipment_faults", row.ID, p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateEquipmentFault", row)
}

// UpdateEquipmentFault changes one row of equipment faults.
func (s *Service) UpdateEquipmentFault(ctx context.Context, p db.UpdateEquipmentFaultParams) (db.EquipmentFault, error) {
	before, err := s.GetEquipmentFault(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateEquipmentFault(ctx, p)
	if err != nil {
		return before, wrap(err, "update equipment faults")
	}
	if err := s.voidSignatures(ctx, "equipment_faults", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "maintenance.equipment_fault.write", "equipment_faults", row.ID, row.SiteID, before, row)
}

// ListWorkOrder returns a page of work orders.
func (s *Service) ListWorkOrder(ctx context.Context, p db.ListWorkOrderParams) ([]db.WorkOrder, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListWorkOrder(ctx, p))("list work orders")
}

// GetWorkOrder returns one row of work orders.
func (s *Service) GetWorkOrder(ctx context.Context, id uuid.UUID) (db.WorkOrder, error) {
	row, err := s.q(ctx).GetWorkOrder(ctx, id)
	if err != nil {
		return row, notFound(err, "work orders")
	}
	return row, inScope(ctx, row.SiteID)
}

// CreateWorkOrder adds one row to work orders.
func (s *Service) CreateWorkOrder(ctx context.Context, p db.CreateWorkOrderParams) (db.WorkOrder, error) {
	var zero db.WorkOrder
	if err := canWrite(ctx, p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateWorkOrder", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "WO", p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateWorkOrder(ctx, p)
	if err != nil {
		return zero, wrap(err, "create work orders")
	}
	if err := s.audit(ctx, "maintenance.work_order.write", "work_orders", row.ID, p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateWorkOrder", row)
}

// UpdateWorkOrder changes one row of work orders.
func (s *Service) UpdateWorkOrder(ctx context.Context, p db.UpdateWorkOrderParams) (db.WorkOrder, error) {
	before, err := s.GetWorkOrder(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateWorkOrder(ctx, p)
	if err != nil {
		return before, wrap(err, "update work orders")
	}
	if err := s.voidSignatures(ctx, "work_orders", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "maintenance.work_order.write", "work_orders", row.ID, row.SiteID, before, row)
}

// AssignWorkOrder is the assign action on work orders.
func (s *Service) AssignWorkOrder(ctx context.Context, p db.ActAssignWorkOrderParams) (db.WorkOrder, error) {
	before, err := s.GetWorkOrder(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "maintenance.work_order.assign", "work_orders", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "AssignWorkOrder", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActAssignWorkOrder(ctx, p)
	if err != nil {
		return before, stateError(err, "assign work orders")
	}
	if err := s.audit(ctx, "maintenance.work_order.assign", "work_orders", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "AssignWorkOrder", row)
}

// StartWorkOrder is the start work action on work orders.
func (s *Service) StartWorkOrder(ctx context.Context, p db.ActStartWorkOrderParams) (db.WorkOrder, error) {
	before, err := s.GetWorkOrder(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "maintenance.work_order.start", "work_orders", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "StartWorkOrder", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActStartWorkOrder(ctx, p)
	if err != nil {
		return before, stateError(err, "start work orders")
	}
	if err := s.audit(ctx, "maintenance.work_order.start", "work_orders", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "StartWorkOrder", row)
}

// CompleteWorkOrder is the submit for clearance action on work orders.
func (s *Service) CompleteWorkOrder(ctx context.Context, p db.ActCompleteWorkOrderParams) (db.WorkOrder, error) {
	before, err := s.GetWorkOrder(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "maintenance.work_order.complete", "work_orders", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CompleteWorkOrder", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "work_orders", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCompleteWorkOrder(ctx, p)
	if err != nil {
		return before, stateError(err, "complete work orders")
	}
	if err := s.audit(ctx, "maintenance.work_order.complete", "work_orders", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CompleteWorkOrder", row)
}

// ClearWorkOrder is the clear return to service action on work orders.
func (s *Service) ClearWorkOrder(ctx context.Context, p db.ActClearWorkOrderParams) (db.WorkOrder, error) {
	before, err := s.GetWorkOrder(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "maintenance.work_order.clear", "work_orders", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ClearWorkOrder", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActClearWorkOrder(ctx, p)
	if err != nil {
		return before, stateError(err, "clear work orders")
	}
	if err := s.audit(ctx, "maintenance.work_order.clear", "work_orders", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ClearWorkOrder", row)
}

// ListMaintenancePlan returns a page of maintenance plans.
func (s *Service) ListMaintenancePlan(ctx context.Context, p db.ListMaintenancePlanParams) ([]db.MaintenancePlan, error) {
	return wrap2(s.q(ctx).ListMaintenancePlan(ctx, p))("list maintenance plans")
}

// GetMaintenancePlan returns one row of maintenance plans.
func (s *Service) GetMaintenancePlan(ctx context.Context, id uuid.UUID) (db.MaintenancePlan, error) {
	row, err := s.q(ctx).GetMaintenancePlan(ctx, id)
	if err != nil {
		return row, notFound(err, "maintenance plans")
	}
	return row, nil
}

// CreateMaintenancePlan adds one row to maintenance plans.
func (s *Service) CreateMaintenancePlan(ctx context.Context, p db.CreateMaintenancePlanParams) (db.MaintenancePlan, error) {
	var zero db.MaintenancePlan
	if err := s.guard(ctx, "CreateMaintenancePlan", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateMaintenancePlan(ctx, p)
	if err != nil {
		return zero, wrap(err, "create maintenance plans")
	}
	if err := s.audit(ctx, "maintenance.maintenance_plan.write", "maintenance_plans", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateMaintenancePlan", row)
}

// UpdateMaintenancePlan changes one row of maintenance plans.
func (s *Service) UpdateMaintenancePlan(ctx context.Context, p db.UpdateMaintenancePlanParams) (db.MaintenancePlan, error) {
	before, err := s.GetMaintenancePlan(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateMaintenancePlan(ctx, p)
	if err != nil {
		return before, wrap(err, "update maintenance plans")
	}
	if err := s.voidSignatures(ctx, "maintenance_plans", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "maintenance.maintenance_plan.write", "maintenance_plans", row.ID, nil, before, row)
}
