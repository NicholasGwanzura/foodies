# Build status

Where the build stands against [spec.md](spec.md). `make check` covers everything under "Built" apart from the items marked as tested by hand.

## Built

| Area | State |
|---|---|
| Contract | All 871 operations have a handler. `exchangeWhatsappFlowData` verifies the caller and then answers 501, because it needs Meta's Flow keys. |
| Data | All 200 tables, in 12 migrations that apply and roll back. |
| Sign-in | Employee number and password, TOTP, refresh rotation with reuse detection, PIN, device unlock. |
| Access | Roles, scopes, permissions checked on every operation, 404 outside scope, segregation of duties from the audit trail. |
| Safe writes | Idempotency keys, `If-Match`, one transaction per write, hash-chained append-only audit. |
| Platform | Rules and readiness, approval engine, PIN signatures, evidence uploads, schedules and tasks, form engine with consequences, escalations with timers, notifications, search. |
| Offline | Installable web app, checklists queue on the device and replay through `/sync/batches`, conflicts reach a supervisor. Tested by hand with the API stopped. |
| Money trail | Float, flush, cash-up, variances, CIT, banking, surprise counts, petty cash, with the SOP rules in `internal/service/cash.go`. |
| Other modules | State machines for every action, and rules in `internal/service/rules_*.go`: three-way match, payment blocks, HOLD, stop-sale, yield, D-Note signatures, recall reconciliation, tracing with gaps, opening gate, key duplication, lost property. |
| Web | Bespoke My Day, tasks, approvals, checklist runner, escalations, readiness, notifications and search. List, detail and form screens for all 170-odd resources come from generated metadata. |

## Not finished

| Area | What is missing |
|---|---|
| WhatsApp | Webhooks verify and log inbound messages. The menus, Flows and outbound sending (WA-01 to WA-12) are not built, and need a Meta business account to build and test. |
| Integrations | Runs, imports and postings are records only. Nothing talks to GAAP POS, GAAP Unity, SAP or Belina, and import files are not parsed. Spec N-02 is still open. |
| Background jobs | Export jobs and disposition batches record their state, but no worker produces a file or disposes of anything. KPI values return an empty series. |
| Depth of rules | The checklist has 792 items. Dedicated workflows follow spec section 10. Rules not listed above run as plain state transitions. Modules outside Cash were tested by API sweep and targeted tests, not click by click. |
| Screens | Dense desk screens such as the invoice workbench, payment runs and budget consolidation use the generic list and form, not a purpose-built layout. Offline covers checklists only, not the other offline screens in screens.md. |
| List sorting | `sort` is accepted and ignored. Lists run newest first. `prev_cursor` is always null. |
| Security | Offline data is in `localStorage`, unencrypted, against spec 9.2. PIN unlock works, but offline signatures do not exist yet. Rate limiting (429) is not built. Field encryption and the penetration test are outstanding. |
| PATCH | A field cannot be cleared back to empty, because an omitted field means no change. |

## Decisions needed

1. **Temporary credentials.** `POST /users/{id}/reset-credentials` returns a `User`, so there is no way to hand over a new password, and a reset user cannot sign in until WhatsApp first use exists. Proposed spec change, not applied: return `temporary_password`, shown once, on that response.
2. **Rule values.** Every Appendix B value is a draft until its owner approves it, and the demo's pilot branch uses them unapproved. The Authorised Limit Matrix is still needed for real approval routes.
3. **Password lockout.** Five failures locks an account, copied from the PIN rule. `auth.lockout` has no value in the sources.
