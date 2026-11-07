"""State transitions for every action operation.

Most come from the rule in `infer`, because each status list in data-model.md
runs in lifecycle order. OVERRIDES holds the ones the rule gets wrong. An entry
is (from statuses or None for any non-final status, to status or None to leave
the status alone). A function with the action's name in a hand-written file of
internal/service replaces all of this for that action.
"""
FINAL = {"closed", "cancelled", "rejected", "completed", "reimbursed", "reconciled", "retired", "superseded", "revoked", "expired", "released", "done", "skipped"}

TARGETS = {
    "submit": ["pending_approval", "pending_review", "pending_signatures", "pending_signature", "pending_verification", "pending_acceptance", "submitted", "under_review", "in_review"],
    "approve": ["approved", "active"], "reject": ["rejected"], "cancel": ["cancelled"], "return": ["returned", "draft"],
    "close": ["closed", "resolved", "completed"], "complete": ["completed", "done", "closed", "complete"],
    "verify": ["verified", "closed"], "review": ["reviewed", "pending_approval", "closed"], "acknowledge": ["acknowledged"],
    "issue": ["issued", "active"], "revoke": ["revoked"], "release": ["released"], "suspend": ["suspended"], "resolve": ["resolved", "closed"],
    "receive": ["received"], "confirm": ["confirmed"], "execute": ["executed", "completed"], "decide": ["decided"], "restore": ["restored", "active", "resolved"],
    "seal": ["sealed"], "hand-over": ["handed_over"], "dispatch": ["dispatched", "in_transit"], "accept": ["accepted"], "post": ["posted"],
    "match": ["matched"], "validate": ["validated"], "reimburse": ["reimbursed"], "retire": ["retired"], "publish": ["published"],
    "activate": ["live", "active"], "reinstate": ["active"], "reopen": ["open"], "pause": ["paused"], "resume": ["active"],
    "hold": ["on_hold", "held"], "lock": ["locked"], "unlock": ["active"], "disable": ["disabled"], "claim": ["in_progress"], "skip": ["skipped"],
}

OVERRIDES = {
    "completeAction": (["open", "in_progress"], "pending_verification"),
    "verifyAction": (["pending_verification"], "closed"),
    "unlockUser": (["locked"], "active"),
    "lockUser": (["active"], "locked"),
    "disableUser": (["invited", "active", "locked"], "disabled"),
    "resetCredentialsUser": (["invited", "active", "locked"], None),
}


# Actions named for a step the verb alone cannot place: (table, action).
STEPS = {
    ("cash_variances", "record-recovery"): (["open", "under_investigation"], "recovered"),
    ("cash_variances", "refer-hr"): (["open", "under_investigation"], "referred"),
    ("cit_transits", "verify-identity"): (["sealed"], "identity_verified"),
    ("cit_transits", "hand-over"): (["identity_verified"], "handed_over"),
    ("cit_transits", "receive-vault"): (["handed_over", "in_transit"], "received"),
    ("cit_transits", "raise-discrepancy"): (["handed_over", "in_transit", "received"], "discrepancy"),
    ("bank_deposits", "submit-proof"): (["banked"], "proof_submitted"),
    ("cit_reconciliations", "complete"): (["open"], "reconciled"),
    ("petty_cash_vouchers", "verify"): (["pending_approval"], None),
    ("petty_cash_vouchers", "attach-receipt"): (["issued"], "receipted"),
    ("petty_cash_reimbursements", "record-payment"): (["approved"], "paid"),
    ("goods_receipts", "submit"): (["draft"], "inspected"),
    ("supplier_invoices", "approve-exception"): (["on_hold"], "approved_exception"),
    ("payment_batches", "record-execution"): (["approved"], "executed"),
    ("stocktakes", "request-skip"): (["planned"], "skipped"),
    ("disposals", "record-disposal"): (["approved"], "disposed"),
    ("kitchen_orders", "acknowledge-late"): (["late_pending"], "confirmed"),
    ("batches", "release"): (["completed"], "closed"),
    ("dispatch_notes", "security-clear"): (["verified"], "security_cleared"),
    ("dispatch_notes", "confirm-delivery"): (["security_cleared", "in_transit"], "delivered"),
    ("dispatch_notes", "raise-discrepancy"): (["security_cleared", "in_transit", "delivered"], "discrepancy"),
    ("ncrs", "contain"): (["open"], "contained"),
    ("ncrs", "verify-close"): (["pending_verification"], "closed"),
    ("recalls", "assess"): (["held"], "assessed"),
    ("recall_communications", "mark-sent"): (["draft", "pending_approval"], "sent"),
    ("complaints", "triage"): (["new"], "triaged"),
    ("payment_incidents", "verify-status"): (["open"], "verified"),
    ("payment_incidents", "propose-correction"): (["verified"], "correction_pending"),
    ("payment_incidents", "reconcile"): (["correction_approved"], "reconciling"),
    ("pos_outages", "submit"): (["declared"], "pending_authorisation"),
    ("manual_receipts", "back-capture"): (["issued"], "back_captured"),
    ("work_orders", "assign"): (["open"], "assigned"),
    ("work_orders", "start"): (["assigned"], "in_progress"),
    ("work_orders", "clear"): (["in_progress"], "pending_clearance"),
    ("asset_disposals", "confirm-removal"): (["approved"], "removed"),
    ("visitor_entries", "sign-out"): (["on_site"], "departed"),
    ("lost_property_items", "dispose"): (["in_custody"], "disposed"),
    ("document_revisions", "withdraw"): (["issued", "approved"], "withdrawn"),
    ("training_sessions", "close"): (["planned"], "held"),
    ("store_openings", "record-opened"): (["go"], "opened"),
    ("footage_requests", "record-export"): (["approved", "preserved"], "exported"),
}


def infer(op, verb, statuses, table=None):
    if op in OVERRIDES:
        return OVERRIDES[op]
    if (table, verb) in STEPS:
        return STEPS[(table, verb)]
    if not statuses:
        return None, None
    for target in TARGETS.get(verb, []):
        if target in statuses:
            i = statuses.index(target)
            if verb in ("reject", "cancel", "suspend", "return", "revoke", "hold", "reopen"):
                source = [s for s in statuses if s not in FINAL and s != target]
                if verb == "reopen":
                    source = [s for s in statuses if s in ("pending_verification", "closed", "resolved", "completed")]
            else:
                source = [s for s in statuses[:i] if s not in FINAL]
                if verb == "submit":
                    source = [s for s in ("draft", "returned", "counted", "open", "requested", "planned") if s in statuses] or source
            return source or None, target
    return None, None
