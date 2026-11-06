# Foodies product specification

**Status:** Draft for review, 6 November 2026.
**Companion documents:** [data-model.md](data-model.md), [screens.md](screens.md), [api/openapi.yaml](../api/openapi.yaml).

## 1. Overview

Foodies is a staff-only operations and compliance platform for the Foodies fast-food group in Zimbabwe. It turns the group's SOPs into daily work: tasks, checklists, registers, approvals, evidence and escalations, across every branch, the Central Kitchen and Head Office.

It is a control layer. GAAP POS, GAAP Unity, SAP Business One and Belina stay the systems of record for sales, stock, finance and payroll. Foodies holds what they do not: who did what, who approved it, what proof exists, and what went wrong.

Staff use it in two ways, with the same rules and audit trail:

- **Web app.** An installable web app for phones, tablets and desktops. Branch forms keep working offline.
- **WhatsApp.** A full working channel for all role work, so a WhatsApp-only employee can do what an app user does. Admin configuration and large desk screens stay on the web.

The scope covers all 11 checklist modules and 792 checklist items, plus the SOP library and the site hierarchy, delivered in seven phases.

### Size of the design

| Measure | Count |
|---|---:|
| API operations | 871 |
| API resources | 173 |
| Data tables | 200 |
| Screens | 151 |
| WhatsApp flows | 12 |

## 2. Goals and non-goals

### Goals

1. **Every money movement is traceable** from till to bank, and from requisition to payment, with named people and proof.
2. **Every stock movement is traceable** from supplier lot through batch and dispatch to branch.
3. **Nothing goes quiet.** Exceptions get an owner and a timer, and they escalate until someone acts.
4. **No deadline is missed unnoticed.** Reports, tax dates, maintenance and expiries are tracked.
5. **Paper registers become searchable records**, except for the paper originals kept on purpose.
6. **Supervisors adopt it.** The notes in the build-first workbook say branch supervisors and Complex Managers decide whether the app gets used, so their mobile and WhatsApp experience comes first.

### Non-goals

- Customer ordering, menus, loyalty or customer accounts.
- Replacing GAAP POS, GAAP Unity, SAP Business One, Belina or bank portals.
- Deciding disciplinary outcomes. Foodies records and refers; people decide.
- Storing CCTV video.
- Multi-company tenancy. Foodies is one organisation. Nash Factory, Nash Transport and Nash Boards are out of scope.
- Inventing any threshold, approver, frequency or retention period (see D-10).

## 3. Sources

| Source | What it gave |
|---|---|
| `FOODIES APPLICATION DEVELOPMENT CHECKLIST.xlsx` | 792 items, 124 clarifications and 107 cross-module requirements across 11 modules. This is the main source. |
| `Foodies_Build_First_Departments.xlsx` | Six build-first groups, department involvement, adoption notes. |
| `Foodies Dashboard.html` | Module tiles, branding, and the empty SOPs and Complexes modules. |
| Product interview, 6 November 2026 | Decisions D-01 to D-16. |

The workbook says "Xactco" throughout. Xactco is a benchmark system the requirements were derived from, not the product. Read each mention as "the Foodies app".

The SOPs themselves were not supplied. Where this spec quotes a limit or clause, it comes from the checklist's summary of that SOP. Rule owners should check values against the controlled SOP before approving them.

## 4. Product decisions

| ID | Topic | Decision |
|---|---|---|
| D-01 | Product | Internal operations platform. Staff only. |
| D-02 | Source | The development checklist workbook is the requirements baseline. |
| D-03 | Name | The product is Foodies. Xactco is a benchmark. |
| D-04 | Scope | All 11 modules, with every feature tagged to a release phase. |
| D-05 | SOPs and Complexes | SOPs is the searchable controlled-document library. Complexes is the site hierarchy. Both are core. |
| D-06 | System boundary | Control layer. Existing systems stay the record, and Foodies reads their data by integration or import. |
| D-07 | Clients | One responsive installable web app with offline capture. |
| D-08 | Sign-in | Employee number and password. PIN unlock on enrolled devices. Mandatory two-factor for finance and admin roles. |
| D-09 | Record model | Hybrid. Dedicated entities for high-logic workflows, and a configurable form engine for the rest. |
| D-10 | Unconfirmed rules | A workflow that depends on an unset rule cannot go live at that site. A readiness screen shows what is missing. |
| D-11 | WhatsApp | A separate full channel. WhatsApp-only users can do all role work. Admin and large desk screens are excluded. Interaction is structured menus and forms, generated from the same templates. |
| D-12 | Signatures and paper | Each signer confirms with a personal PIN, in the app or from their own verified WhatsApp number. The record stores signer, time, channel and a content hash. Listed paper originals stay, and are scanned and referenced. |
| D-13 | Organisation | Foodies only. Head Office, then Complex, then Site. A site is a branch, Central Kitchen, store, bakery or butchery. A complex may hold one site. |
| D-14 | Accounts | Every employee has an account, usually WhatsApp-only for general staff, with self-service for leave, roster, acknowledgements, training and profile updates. |
| D-15 | Locale | USD and ZiG, time zone Africa/Harare. A currencies table holds exchange rates that admins can adjust manually. USD is the reporting base. |
| D-16 | Phasing | Foundation first, then the build-first groups. WhatsApp ships with each phase. |

## 5. Users, roles and access

### 5.1 Who uses it

| Group | Examples | Main channel |
|---|---|---|
| Branch front of house | Cashiers, customer service | WhatsApp, shared tablet |
| Branch management | Shift Supervisors, Branch Managers, Complex Managers | Phone (web or WhatsApp) |
| Stores and production | Storekeepers, chefs, bakers, butchery, drivers | WhatsApp, shared tablet |
| Security | Security Officers, Supervisors, Security Manager | WhatsApp |
| Quality and safety | QC, Food Safety, SHEQ | Phone and desktop |
| Head Office | Accounts, Creditors, Procurement, Finance, HR, Training, QMS, ICT | Desktop |
| Executives | GM, MD, Chairman, Board | WhatsApp and dashboards |
| General staff | All other employees | WhatsApp |
| External parties | CIT crew, drivers, contractors, visitors, claimants | No account. An employee captures them. |

### 5.2 Access model

- **Role + scope.** A role grants permissions at a scope: one site, one complex, or all. A filter can narrow a user's view, never widen it.
- **Seeded roles.** Roles come from those named in the checklist. Synonyms are merged as below. This merge is an assumption to confirm (N-06).

| Role | Also called in the workbook |
|---|---|
| Shift Supervisor | Supervisor, Shift Manager, Manager on Duty, AM Supervisor |
| Branch Manager | Shop Manager, Store Manager, Resident Manager |
| Complex Manager | Complex/Branch Manager, where the scope is the complex |
| Finance Manager | Finance and Administration Manager |
| Operations Manager | OM |
| ICT | IT, System Administrator (permission-limited) |
| QA | QC Lead, Food Safety Officer, depending on site |

- **Two-factor.** Mandatory on the web for finance and admin roles. On WhatsApp, the verified number plus PIN serves as the two factors.
- **Sensitive data.** Medical records, disciplinary files, recipes, cost sheets, CCTV registers and key registers are restricted, and every read is logged. Dashboards show compliance status, never the medical detail.
- **Own record.** Self-service users see only their own records.

### 5.3 Segregation of duties

These are enforced on the server. The default is to block. Each is a row in `sod_rules`.

| Rule | Source |
|---|---|
| A cashier cannot approve their own cash-up or payment correction. | Sales, Store Ops |
| No one approves their own cash count, key exception or footage access. | Store Ops |
| Requester, buyer, approver, receiver, counter, adjuster and payment authoriser are separate. | Inventory |
| Cash receipt, banking, posting, reconciliation, supplier master changes and payment approval are separate. | Finance |
| A supplier change is verified by someone other than the requester. | Finance |
| A disposal needs at least two people. | Waste & Yield |
| Recall reconciliation has an independent reviewer. | Recall |
| The approver's actions are checked across delegations, so delegating cannot get around a conflict. | Design |

## 6. System boundary

| System | Holds | Foodies reads | Foodies sends |
|---|---|---|---|
| GAAP POS | Sales, tenders, receipts, Z-reports | Sales by till, shift, tender. X-reads and Z-reports. | Stop-sale instruction for a HOLD (manual until confirmed) |
| GAAP Unity | Stock balances and movements | Balances, transactions | Approved transfers, adjustments, issues |
| SAP Business One | AP, GL, bank, assets, depreciation | Supplier codes, posted status, actuals | Approved suppliers, POs, invoices, journals |
| Belina Payroll | Payroll | Employee list for reconciliation | Approved employee changes |
| Bank portals | Statements, payments | Statements by import | Nothing. Payments are made in the portal and recorded in Foodies. |
| CCTV | Video | Camera health, only if the vendor allows it | Nothing |
| WhatsApp Business Platform | Message delivery | Inbound messages and Flow data | Messages and Flows |

### 6.1 Postings

When a Foodies decision has to appear in a system of record, Foodies creates a posting. A posting stays visible as `pending`, `sent`, `posted`, `failed` or `manual_required` until it is confirmed. The record shows its posting status, so nobody assumes SAP or GAAP Unity was updated.

### 6.2 Imports as the fallback

Integration feasibility is unconfirmed for GAAP, SAP and Belina (N-02). So each inbound feed also has a file import, and each outbound posting can be done by hand and recorded. Workflows must keep running on either path.

### 6.3 What a HOLD can and cannot do

A HOLD blocks everything that passes through Foodies: transfers, dispatches, issues and disposals. Foodies does not control the till, so it cannot stop a sale directly. It creates a stop-sale posting and a task for the POS Administrator, with an escalation until someone confirms. This is a gap against the checklist's "block sale" wording, listed as N-03.

## 7. Organisation model

- **Head Office** is a site of type `head_office`, outside any complex.
- **Complex** groups sites under a Complex Manager.
- **Site** is a branch, Central Kitchen, central store, bakery or butchery. A site with `stockholding = true` is a stock node, with its own stock on hand.
- **Trading pattern** on a site (`closes_2200` or `open_24h`) selects report deadlines.
- **Business date** for a shop that trades past midnight stays on the opening date.

Site-specific checklists, patrol points and escalation contacts are configured per site.

## 8. Platform capabilities

These serve every module, and are built in Phase 0. The 107 cross-module requirements in the workbook map to them.

### 8.1 How a workflow works

A workflow is a piece of work that has to happen in a set way, by set people, and leave a record. Foodies moves the work from person to person, applies the rules and keeps the history. The rest of this section describes the shared parts, and section 10 gives the specifics of each workflow.

**The parts every workflow draws on**

| Part | What it does | Details |
|---|---|---|
| Status | Shows where the work is now. Status only changes through a defined step, never by editing the record. | Section 10, data-model.md |
| Rules | Hold the limits, deadlines and frequencies. Managers own them, not developers. | 8.2 |
| Approvals | Route the record to the right people, in order. | 8.3 |
| Signatures | Record a named person confirming something with their PIN. | 8.4 |
| Evidence | Attaches photos, scans and documents to the record. | 8.5 |
| Tasks | Send work to a person or role, on demand or on a schedule. | 8.7 |
| Escalation | Acts when something is late or goes wrong. | 8.8 |

**Example: a petty cash voucher**

| Step | Status | What happens |
|---|---|---|
| 1. Request | `draft` | The request is raised on the web app or WhatsApp, with the amount, reason, category and recipient. |
| 2. Rule check | `pending_approval` | Foodies checks the request against the rules. A prohibited category is rejected and routed to procurement. |
| 3. Approval | `approved` or `rejected` | The Complex Manager verifies and the MD approves, each with their PIN. Late approvals remind, then escalate. This step needs a connection. |
| 4. Issue | `issued` | Only the approved amount is issued. No cash can be issued before approval. The register updates at once. |
| 5. Receipt | `receipted` | The receipt is attached. Every transaction needs one. |
| 6. Reimbursement | `reimbursed` | The voucher is claimed from Finance. A request can also end as `cancelled`. |

The record keeps who did what and when, and the version of each rule it ran under. The example shows the mechanics only. The final approval chain is still being confirmed (C-01.04), and section 10.1 has the full requirements.

**Two ways workflows are built (D-09)**

- **Dedicated:** high-stakes work with a lot of logic, such as cash-up, purchasing and payments, has its own screens and data.
- **Form-based:** simpler, repetitive work, such as cleaning and kitchen checks, runs on published form templates (8.6). A failed answer can create a task, action, NCR, HOLD or incident.

Many areas mix the two. Appendix A shows the approach for each area: 40 are dedicated, 25 hybrid, 25 form-based, 8 met by the platform, and 3 done in a system of record with evidence in Foodies.

**Going live (D-10)**

Building a workflow and switching it on are separate. A workflow is switched on per site, so branches can differ, and it moves through these states:

| State | Meaning |
|---|---|
| `not_ready` | At least one blocking rule has no active value. Staff cannot use the workflow at this site. |
| `ready` | Nothing is missing, and it can be activated. |
| `live` | Activated. Staff can use it at this site. |
| `suspended` | Switched off after going live. |

Petty cash shows the two ways a rule can be unready. The US$1,000 limit is stated in the checklist, but it is only a proposal until its owner approves it (Appendix B). The reimbursement day has no value at all yet (C-01.07). While a blocking rule has no active value, the workflow stays `not_ready`, and anyone who tries to use it gets a `workflow_not_live` error naming what is missing. This is why a missing decision holds up release, not development.

### 8.2 Rules and readiness

Every threshold, deadline, frequency, limit and authority is a versioned rule with an owner, a source clause and an approval. Nothing is hard-coded.

- A rule has a value per scope (group, complex, site, item or asset) and effective dates.
- A blocking rule with no active value stops the workflow going live at that site. API calls return 409 `workflow_not_live` with the missing rule keys.
- **Readiness** (RUL-01) shows, by module, workflow and site, what is missing and who owns it.
- Values stated in the checklist are seeded as proposals (Appendix B). They do nothing until the owner approves them.
- Records keep the rule version they ran under.

### 8.3 Approval engine

- Policies are versioned routing per workflow: ordered steps, role, any or all, band conditions, deadline and escalation.
- Monetary bands come from the Authorised Limit Matrix, which has not been supplied. Workflows that need it stay `not_ready` until it is loaded.
- An approver acts from the web app or WhatsApp, with a PIN.
- Delegation is time-boxed, approved, and cannot exceed the delegator's authority.
- Retrospective approval that regularises a bypass is not allowed.

### 8.4 Signatures

A signature is a named person confirming a specific meaning (counted, witnessed, verified, approved, received, handed over, countersigned, acknowledged) on a specific version of a record.

- **Employees** enter a 6-digit PIN. The PIN is separate from the password. Five wrong attempts lock signing until it is reset.
- **Witnesses on a shared device** enter their own PIN on that device. **On WhatsApp**, each signer gets a request on their own phone.
- **PINs on WhatsApp** are only entered inside a WhatsApp Flow, sent encrypted to the Foodies endpoint. A PIN is never typed as a chat message.
- **External parties** (CIT crew, drivers, claimants, contractors) have no account. The employee records the name and ID reference and attests to it, attaching a photo of the paper slip when one exists.
- **Editing** a signed record voids its signatures, and they must be collected again.

A record that needs signatures cannot close until they are all in. Missing signers show up as blockers.

### 8.5 Evidence and paper

Photos, scans and documents attach to any record. Images are compressed on the device and uploaded with a checksum.

Paper originals stay where the workbook points to them, unless the owner approves otherwise (C-04.13): supplier invoices, stamped transfer invoices, pre-numbered stock sheets, Bin Cards, signed manual receipts. For these the attachment records `paper_original_location`.

### 8.6 Form engine

Work that does not need dedicated logic runs on form templates rather than bespoke screens: 25 areas in full, and the form parts of the 25 hybrid areas (Appendix A). A template defines:

- Sections and fields: text, number, reading, date, choice, pass/fail/N-A, OK/N-OK, photo, table.
- Validation, such as a temperature limit taken from a rule.
- Required evidence, and required sign-offs.
- Consequences: a failed answer can create a task, action, NCR, HOLD or incident, and can block completion.
- Channels and offline use.
- Retention class and source SOP clause.

Templates are versioned and approved. A submission keeps the version it used. Registers are simply templates of kind `register`. When a template is published with WhatsApp on, the platform generates the matching Flow. A template that exceeds Flow limits must be split into steps, and the builder warns about this.

### 8.7 Tasks, schedules and deadlines

- Schedules create tasks at the frequency held in a rule.
- Tasks go to a person or a role at a site.
- Clocks use hours or working days. Working days use the public-holiday calendar.
- Reminders fire before the due time. Overdue items escalate.
- A skipped control needs approval and a stated alternative control.

### 8.8 Exceptions, escalation and actions

Severity has five levels. The levels are the same numbers in every domain, and the labels are configured per domain, because QC and Customer Service use different words (N-05).

| Level | QC wording | Customer Service wording | Route |
|---|---|---|---|
| 1 | Routine | Routine | Front line, Supervisor |
| 2 | Significant | Operational | Supervisor, Complex Manager |
| 3 | Serious | Serious | Complex Manager, GM or OM, to MD |
| 4 | Critical | Executive | MD immediately, Chairman informed |
| 5 | Executive critical | Critical | MD, then Chairman |

If a reporter is unsure between two levels, the higher level applies.

An escalation notifies each person in the chain and starts an acknowledgement timer. No acknowledgement moves it up a level. Critical alerts do not wait for a form to be finished: the alert goes first, and the record follows. Recipients and timers are rules, unset until supplied (C-06.09, C-07.10, C-11.05).

Actions (CAPA) carry an owner, due date, evidence and an effectiveness check. Overdue and repeat items surface to management.

### 8.9 Notifications

Channels are in-app, web push, email and WhatsApp. A person chooses channels per category, except for critical escalations, which always go out on every channel available to them. Delivery failure moves to the next channel.

### 8.10 Audit trail

Every create, change, approval, signature, export and restricted read is written to an append-only, hash-chained log with user, role, time, channel and device. Registers such as the Occurrence Book are append-only too.

### 8.11 Retention

Retention is by record class, never a blanket period. The workbook contains conflicts, such as a 2-year cash-up record against a 6-year finance rule, and Z-reports at 5 years. A class with no approved period stays `unset`. Its records are kept until a period is approved. Legal holds stop disposition, and disposition batches need approval.

### 8.12 Dashboards, search and exports

Dashboards run at branch, complex and group level, with drill-down to the source record. Each KPI has a formula and a target rule. Search works across records by reference, branch, date, product, batch, supplier, employee and customer case. PDF and Excel exports run in the background, and restricted exports are logged.

## 9. Channels

### 9.1 Web app

A single installable web app, responsive from phone to desktop. It keeps the existing Foodies look: orange header, module tiles, numbered modules. Head Office work is table-heavy and desktop-first, while branch work is one-handed and phone-first.

Shared branch devices are enrolled. Users switch with employee number and PIN, and sessions time out quickly on shared devices.

### 9.2 Offline capture

Branches lose power and connectivity, and the work has to continue.

- **Works offline:** templates with `offline_allowed`, and the screens marked offline in screens.md, including counts, readings, checklists, flushes, cash-ups, incidents, GRVs and stocktakes.
- **Needs a connection:** approvals, executive decisions, and anything that must see current data, such as a HOLD release or a CIT limit check against live balances.
- **How it works:** the device holds the data from `GET /sync/manifest`, including a snapshot of rules. Writes queue with a client id and capture time, then replay through `POST /sync/batches`. Each item is idempotent.
- **Offline signatures:** a signer enters a PIN, checked against a device-bound verifier. The server re-validates the proof on sync, and the signature shows `offline_web`.
- **Conflicts:** records are append-only by nature, so true conflicts are rare. A write the server cannot apply goes to a supervisor (SYN-02) rather than being dropped.
- **Time:** both device capture time and server receive time are stored. A large gap is flagged.
- **Security:** offline data is encrypted on the device, and revoking a device wipes it at next contact.

### 9.3 WhatsApp channel

WhatsApp is a working channel, not just alerts. Both channels call the same services, so rules, approvals, segregation of duties and audit are identical.

| Topic | Design |
|---|---|
| Scope | All role work: tasks, checklists, registers, approvals, incidents, photo evidence, sign-offs, alerts, summary KPIs and self-service. |
| Excluded | Admin configuration and large desk screens, such as template builder, invoice workbench, payment runs, budget consolidation and full dashboards. |
| Interaction | Buttons, lists and WhatsApp Flows forms. No free-text commands. |
| Identity | HR registers a number against an employee. The user confirms a one-time code and sets a PIN in a Flow. One number maps to one employee. |
| Changing numbers | ICT or HR rebinds the number. The old number stops at once. |
| Leaving | Separation unbinds the number immediately. |
| Sign-off | PIN in a Flow, on the signer's own number. |
| Evidence | Photos and documents sent in the chat are attached to the record in progress. |
| Sessions | Conversation state is held on the server, so a user can stop and resume. Stale drafts expire with a reminder. |
| Sensitive content | Message bodies carry a reference and a summary, not personal data. Restricted detail opens only in the web app, or in a Flow after PIN. |
| Message rules | Outside the 24-hour service window, Foodies can only send approved templates, so every proactive alert has a utility template. |
| Failure | WhatsApp cannot work offline. A branch with no data at all falls back to the web app's offline mode on a device, or to paper. |

Three things the channel depends on are listed as risks (section 14): Meta's template approval, per-message cost across all staff, and staff having suitable phones and data.

Groups are not used for evidence. Records that SOPs say go to an official WhatsApp group now live in Foodies, so those SOPs need revising (N-01).

## 10. Module requirements

The tables below give the behaviour of the dedicated workflows. Fields are in data-model.md. Appendix A places all 101 workbook areas. Items that run on the form engine follow their checklist row directly, since each row already gives the function, fields, responsible person, frequency, verification and exception handling.

Limits shown below are rule keys, with stated values in Appendix B.

### 10.1 Cash

| Workflow | States | Requirements |
|---|---|---|
| Float | Certificate: draft to active | Amounts cannot exceed `cash.change_float_max` or `cash.petty_cash_max`. GM and Finance Manager set them jointly. |
| Float use | issued, acknowledged, returned | Float is for change only. Exceptional use needs prior MD approval and full reimbursement next day. Any other till withdrawal is flagged. |
| Till flush | draft, pending_signatures, closed, exception | Triggered at `cash.till_flush_threshold`. Cashier and supervisor count together, and the supervisor secures the cash in the safe with a security witness. A POS feed raises the alert where available, otherwise the cashier declares it. The threshold cannot be silently overridden. No close without exact time and all signatures. |
| Cash-up | draft to approved | Count per currency and tender against POS, X-read and Z-report. USD and ZiG are recorded separately and never netted. The rate used is stored. An independent reviewer approves it. CIT preparation cannot start until the cash-up and its variances are dealt with. A Z-report that fails to print stops the cash-up until instructed. |
| Variance | open to closed | Every variance is logged, however small. Above `cash.variance_escalation` it escalates to the Complex Manager at once. Temporary loans are not allowed. The system counts occurrences per cashier (3 in a week or 3 in a month triggers referral) and shows the sanction band, but makes no decision. Fake notes are not treated as ordinary variances. |
| CIT transit | draft to reconciled | Blocked above `cash.cit_transit_limit` per transit. Only Shift or Complex Managers prepare slips. Two slip copies are tracked. Seal numbers are unique. CIT identity is matched to the approved list, and a failed match blocks the handover. All parties countersign. Formal receipt and armoured vehicle confirmed. Vault checks seals and amount. |
| CIT discrepancy | | Any mismatch starts a 2-hour clock from discovery, with breach flagged. |
| Banking | draft to verified | Takings are banked intact. Expense net-offs and partial banking are flagged. Banking is daily or next working day, with proof within 24 hours. Direct banking needs a Head Office Finance instruction and a stamped slip. |
| Surprise count | planned to completed | Head Office only. The branch cannot see or pre-schedule it. |
| Petty cash | draft to reimbursed | No cash before MD approval. Complex Manager verifies. Only the approved amount is issued. Register updates at once. Prohibited categories (creditors, salaries, wages, major and personal expenses) are rejected and routed to procurement. Every transaction needs a receipt. Joint count at each change of custodian. End-of-day reconciliation, with any discrepancy reported the same day. |

### 10.2 Procurement and accounts payable

| Workflow | States | Requirements |
|---|---|---|
| Supplier | draft to active | Mandatory documents block activation. SAP code series LOSUP or FOSUP. Finance Manager approves independently. Bank detail changes need independent verification. |
| PR | draft to closed | No PR, no procurement. WhatsApp or verbal messages do not replace a formal PR, though a PR can be raised through the WhatsApp channel because that creates the same record. Procurement validates completeness. Routine local requests reach Procurement at least 3 working days ahead. |
| Branch purchasing | | Branches buy only minor consumables within petty cash limits. Raw materials and packaging are bought centrally. |
| Quotation | | Count comes from `procurement.quotation_count` (C-04.03). Like-for-like comparison. Reason needed if the lowest quote is not chosen. Splitting to avoid thresholds is flagged. |
| PO | draft to closed | No PO, no purchase. No backdating. GM or Operations Manager review, then MD sign-off, under the approved policy. Only Procurement commits to suppliers. |
| GRV | draft to posted | Linked to the PO, delivery note and invoice. Chilled and frozen limits apply from `stock.chilled_max_c` and `stock.frozen_max_c`. Non-conforming stock is rejected or held. |
| Invoice | received to paid | Forwarded to Finance within 3 days. Three-way match of PO, GRV and invoice. Duplicate check on supplier, number, amount and date. Local invoice QR verification for VAT. Unmatched invoices are held. |
| Payment | draft to paid | Blocked without documents, approvals and match. Approval stages come from the policy (N-04). No PO, no payment, apart from documented statutory exceptions. Payment is executed in the bank portal and proof recorded in Foodies. |
| WHT | | Suppliers without valid tax clearance have tax withheld at the configured rate. Remittance by the 10th of the next month is tracked in the compliance calendar. |
| Emergency | draft to reviewed | Genuine unforeseen need only. Written justification, MD retrospective approval within 48 hours, and a separate review. |
| Unauthorised purchase | reported to closed | Recorded, reviewed by Operations and Finance. No retrospective regularising or backdating. |
| Budget | draft to approved | Annual cycle to Board approval. Variances above the threshold rule need an explanation, and escalate to MD within 5 working days of month-end close. |

### 10.3 Stock

GAAP Unity remains the authoritative balance. Foodies adds workflow, holds and alerts.

| Workflow | States | Requirements |
|---|---|---|
| Planning | | Critical items, days of cover, the 14-day supply rule and 7-day alert (scope C-04.05). |
| Issue | draft to acknowledged | No issue without an authorised requisition. Recipient acknowledges. Expired stock is never issued. FIFO, or FEFO for shelf-life stock. |
| Transfer | requested to closed | System document travels with the stock. Security compares physical stock and stamps it at dispatch and at receipt. Receiver counts before acceptance. Discrepancies are reported and investigated before confirmation. HOLD stock cannot move. TBA follow-up and credit note within 24 hours, and discrepancy closure within 48 hours (clocks per C-04.06). |
| Stocktake | planned to posted | Types: shift open and close, close of business, cycle, month-end, year-end. Reconciles to DRS, GAAP Unity and Bin Card. Year-end uses pre-numbered sheets and two count teams. Obsolete stock is counted separately. A skipped count needs approval and an agreed alternative control. |
| Variance | open to closed | Opened for every variance. Separates timing or capture issues from real loss. Recovery follows review. |
| Adjustment | draft to posted | Mandatory reason tag: wastage, damage, off-cuts, donation, expired inventory. Approval comes before posting. |
| Waste | recorded to closed | Recorded where it happens, by category. 100% recording is the target. |
| Disposal | draft to disposed | Never one person alone. Approval comes before disposal or write-off. |

### 10.4 Central Kitchen and production

| Workflow | States | Requirements |
|---|---|---|
| Branch order | draft to received | Morning cycle closes 10:00 and evening 19:00. Late orders capture reason, expected time and acknowledgement, and are not guaranteed same-cycle dispatch. Unusual quantities are flagged against par levels and history. |
| Plan | draft to completed | Consolidates on-time orders. Material shortages go to Procurement. Late orders are logged separately. |
| Batch | in_progress to closed | Unique batch number. Ingredient supplier lots are recorded, and missing labels mean no use. Yield = usable output ÷ standard or input quantity × 100. Variance above `production.yield_variance_percent` needs review. |
| Release | | Statuses: released, hold, rework, rejected, disposed. No clear status, no release. |
| D-Note | draft to closed | Nothing leaves without a signed D-Note. Pick by FIFO or FEFO, with raw and ready-to-eat separated. Security matches D-Note and items. Loading and delivery temperatures are captured. The receiver signs after inspection, and a discrepancy is recorded before the driver leaves where possible. Rejected goods do not enter branch stock. |
| Recipes | draft to active | Versioned, restricted, approved. Changes follow the approved revision process. |

### 10.5 Quality, incidents and recall

| Workflow | States | Requirements |
|---|---|---|
| NCR | open to closed | Records deviation, containment and stop decision. Cannot close without effectiveness verification. |
| HOLD | on_hold to resolved | Blocks use, transfer, dispatch and disposal. Stop-sale per section 6.3. No use, sale, donation, rework or disposal without an authorised disposition. |
| Disposition | draft to executed | Separate approval records for rework, recovery, rejection, disposal and return. |
| Linked records | | One linked chain: a complaint, incident, NCR or HOLD can launch a recall, while each keeps its own approval and closure rules (C-06.05). |
| Recall | initiated to closed | Any employee can raise one. Immediate stop-sale and HOLD. Risk assessment, then trace, then decision (significant actions authorised by MD). External communication must be approved. Every location tracked to acknowledgement. Quantities reconcile (produced = stock + distributed + sold + returned + disposed) with independent review. Closure needs reconciliation and started CAPA. |
| Trace | | Forward and back from a lot or batch, with gaps shown. A gap triggers HOLD and escalation. POS links by product, site and date unless GAAP POS can supply batch data. Target for bakery recalls: trace within 4 hours. |
| Mock recall | | Scheduled, using the same workflow with `kind = mock`. |
| Incident | reported to closed | One register with typed detail. Time of discovery and notification are stamped. SHE internal report within 24 hours, with NSSA notification where required. Lost or stolen keys: report within 24 hours. |

### 10.6 Customer service and payment exceptions

| Workflow | States | Requirements |
|---|---|---|
| Complaint | new to closed | Unique case reference linking order, sale, payment, refund and follow-up. Open cases appear at shift handover, and a serious case needs a named owner. Food-safety allegations and foreign objects capture evidence and do not admit causation. Illness, injury or exposure escalates at once. |
| Service recovery | draft to executed | Authority by role and value from the approved matrix. Remedies above the user's authority are blocked. |
| Payment incident | open to closed | Structured error types. Status is verified with the provider, and a customer debit alert alone is not proof. A failed-looking payment is never simply repeated. Corrections go through Finance. No refund date is promised before the provider confirms. Closes after settlement reconciliation. |
| POS outage | declared to closed | Manual receipting starts only after authorisation. Foodies holds that approval. |
| Manual receipts | | Pre-numbered books are issued and returned. Cancelled receipts are kept. Missing numbers are investigated. Back-capture is duplicate-checked. |

### 10.7 Maintenance

Faults on any asset can become work orders, with downtime tracked. Critical equipment stays out of service until cleared. Critical faults alert the Complex Manager and Maintenance, and the GM when trading may be affected. Preventive plans create tasks ahead of due dates. Only authorised operators may run generators and ovens. Generator and pizza oven SOPs are the only ones supplied; other equipment needs its own schedules (C-02.12).

### 10.8 Deadlines

| Workflow | Requirements |
|---|---|
| Reporting calendar | Schedules in Appendix B. A designated deputy can submit if the head is away. Lateness is tracked and repeated or unexplained lateness escalates. Reports feed the action tracker, and actions carry forward until closed. Wednesday management meetings use both. |
| Compliance calendar | Tax, statutory, certificate, contract, licence, insurance and calibration dates, each with an owner, reminders and evidence. No rates or dates are inferred. |
| Internal audit | Findings become actions. Non-compliance is corrected within 30 days. |

### 10.9 Registers

| Register | Requirements |
|---|---|
| Fixed assets | One ID across the register, maintenance, faults and custodian. Lifecycle from CAPEX through PO, GRV and capitalisation to verification and disposal. Annual verification. Depreciation runs in SAP. Tagging method is C-02.13. |
| Keys | Register and custodian matrix. Sign-out and sign-in. Duplication is blocked without written Chairman approval and a memo signed by Complex Manager, Operations Manager, GM, MD and Chairman. Monthly audit and annual reconciliation. Lost keys create an incident. Alarm codes are never stored. Exit clearance requires key return. |
| Visitors and contractors | Digital book, with induction status and host. |
| Lost property | Bank cards: 30 days, then surrendered to the bank or securely destroyed under manager supervision. Personal items: 90 days. Claimant verified and signature captured. Monthly review. Full card numbers and PINs are never stored. |
| Documents (SOPs tile) | Master Document Register with numbering `[Company]/[Department]/[Type]/[Sequence]`. Draft to review, QMS check, approval and issue, on a five-working-day clock. Only current versions reach the point of use, and old versions are withdrawn. Temporary documents have an expiry. Revisions can require acknowledgement and training. Duplicate numbers in the current set are flagged (C-05.01, C-07.01). |
| Training | Courses linked to SOP versions. Role matrix. Sessions scheduled to avoid coverage gaps. Attendance, assessments, competency sign-off before independent work, certificates with expiry alerts. SOP revisions, incidents and failed assessments trigger refreshers. HR and Training share one employee master. |
| Leave | Self-service request. Approver sees balance and branch coverage. Entitlements come from the approved policy. Sick notes are restricted. |
| Exit clearance | Stores, Finance, ICT, Security, uniform custodian, keys and department each sign. Access is revoked at the effective date, or immediately where required. HR notifies Payroll, and Finance calculates. |

### 10.10 Remaining modules

| Area | Requirements |
|---|---|
| Store opening | Checklists feed a traffic-light gate: RED is no go, AMBER is a management decision, GREEN is go. A recorded decision is needed before doors open. A cold-chain deviation is a critical non-conformance. A till that fails to fiscalise stops trading until instructed. |
| Store closing | Runs on forms, plus the Phase 1 cash-up. Only a Supervisor can do the final POS close. |
| Security | Electronic Occurrence Book, patrols with checkpoints, searches, declarations, vehicles, gates, and cash attendance. Officers report 15 minutes before shift. Exemptions from searches are not created. |
| CCTV | Camera inventory, daily checks, outages with compensating measures. Footage requests need the approved set of approvers, with purpose limits, chain of custody and retention holds. A requester cannot approve their own request. |
| HR | Rosters, attendance, recruitment, performance, grievance and uniforms. The wider HR lifecycle is recommended design, not an existing procedure, since only two HR documents were supplied (C-09.01). Disciplinary tracking follows the approved process, with no automatic sanctions. |
| Product knowledge | Product cards, allergens, FAQs and assessments. Details come from approved sources only. Allergens must come from verified ingredients. Changes link to retraining. |
| NPD | Stage gates G0 to G5, with confidential recipe and cost data. |
| Costing | Standard cost, margin and review triggers (thresholds unset, C-01.13). |
| External visits | Request, screening, approval, disclaimer or NDA, register, close-out. |

### 10.11 Reports and KPIs

KPIs come from the checklist's dashboard rows, and each target is a rule. Examples:

| Area | KPI |
|---|---|
| Petty cash | 100% with voucher and receipt, zero unexplained end-of-day variances |
| AP | Invoice processing under 2 days, payment cycle within 48 hours, zero duplicates, 100% WHT compliance |
| Budget | 100% submissions, compliance at or above 95% |
| Bank | 100% accounts reconciled within 5 working days after month-end |
| Waste | Below 3% or approved target, 100% recording |
| Hygiene | At least 95% audit pass, 100% daily register |
| Keys | 100% monthly audit, zero unauthorised duplication, zero missing keys |
| QMS | Percentage of documents completed in five working days |

## 11. Release phases

Phases follow the workbook's build-first groups, with the cross-cutting group moved to the front because everything depends on it. WhatsApp ships with each phase.

| Phase | Name | Workbook areas | Checklist items |
|---|---|---:|---:|
| 0 | Foundation | 3 | 13 |
| 1 | Money trail | 9 | 116 |
| 2 | Stock and production | 19 | 95 |
| 3 | Exceptions and incidents | 10 | 116 |
| 4 | Deadline-driven compliance | 4 | 59 |
| 5 | Registers | 24 | 145 |
| 6 | Remaining modules | 32 | 248 |
| | **Total** | **101** | **792** |

Items are counted once, at the phase holding most of their area. Some areas span phases, as Appendix A shows.

| Build | Checklist items |
|---|---:|
| Hybrid | 329 |
| Dedicated | 282 |
| Form engine | 121 |
| Platform | 36 |
| External + form | 24 |

| Phase | Delivers | Ready when |
|---|---|---|
| 0 Foundation | Sign-in, organisation, employees, roles, rules and readiness, approvals, signatures, evidence, audit, tasks, forms, escalations, notifications, offline, WhatsApp core, integration framework, import, dashboards, retention | One pilot branch can run a published checklist end to end on the web app and on WhatsApp, offline included, and an auditor can see the result. |
| 1 Money trail | Cash, CIT, banking, float, petty cash, PR to payment, budgets | A pilot branch's daily cash reconciles from till to vault. The Authorised Limit Matrix is loaded. |
| 2 Stock and production | Receiving controls, transfers, stocktakes, waste, orders, batches, dispatch | A batch can be traced from supplier lot to branch. |
| 3 Exceptions | NCR, HOLD, recall, incidents, complaints, payment incidents, POS outage, faults | A mock recall meets its targets, and escalation reaches executives in a test. |
| 4 Deadlines | Reporting calendar, compliance calendar, maintenance plans | One full month of reports runs through the calendar. |
| 5 Registers | Assets, keys, visitors, lost property, documents, training, leave, exit | Paper registers retire branch by branch, as approved (C-11.09). |
| 6 Remaining | Store opening gate, security, CCTV, HR lifecycle, product knowledge, NPD, costing, and the remaining form-based areas | All templates published, or explicitly deferred. |

**Rollout inside a phase:** pilot one branch, then one complex, then all sites. A workflow goes live per site through readiness, so branches can differ. The Store Opening gate sits in Phase 6 by the build-first ranking, but it is daily branch work and could move up if pilot feedback supports it.

## 12. API conventions

The contract is [api/openapi.yaml](../api/openapi.yaml), OpenAPI 3.1, with 871 operations.

### 12.1 Basics

- REST over HTTPS, JSON, base path `/api/v1`. Breaking changes get a new major version.
- Field names use `snake_case`.
- `id` values are UUIDv7. People use `reference`.
- Timestamps are UTC RFC 3339. Business dates are `YYYY-MM-DD` in Africa/Harare.
- Decimals are JSON strings, so no precision is lost.
- Money is `{ "amount": "1000.00", "currency": "USD" }`.
- `status` is read-only. State changes use `POST /{resource}/{id}/{action}`, so each transition has a permission and an audit event.

### 12.2 Authentication

| Client | Scheme |
|---|---|
| Web app | `POST /auth/login` with employee number and password. Returns a 15-minute JWT and a rotating refresh token. Finance and admin roles complete TOTP through `/auth/mfa/verify`. |
| Enrolled device | `POST /auth/pin/unlock` switches user with employee number and PIN. |
| WhatsApp | The verified phone number identifies the user. The channel service acts on their behalf with a short-lived token. Sensitive steps need the PIN. Meta's calls are verified with `X-Hub-Signature-256`. |
| Integrations | `X-Integration-Key`, per system, rotated, and restricted by source address. Inbound endpoints only. |

Refresh token reuse revokes the whole session family. Revoking a user ends their tokens within the access-token lifetime, and WhatsApp stops at once.

### 12.3 Permissions

Permissions are named `module.resource.action`, such as `cash.cash_up.approve`, and appear on each operation as `x-permission`. The server checks four things on every call: permission, scope, segregation of duties and workflow state.

### 12.4 Errors

Every non-2xx response is `application/problem+json` (RFC 9457).

```json
{
  "type": "/problems/workflow-not-live",
  "title": "Workflow not live at this site",
  "status": 409,
  "detail": "Petty cash cannot start at this branch until 1 rule is approved.",
  "instance": "/api/v1/cash/petty-cash/vouchers",
  "code": "workflow_not_live",
  "trace_id": "01JC8Z3T9Q",
  "blockers": [
    { "kind": "rule_missing", "reference": "cash.petty_cash_reimbursement_day", "message": "Reimbursement day not set" }
  ]
}
```

Validation errors add field detail:

```json
{
  "type": "/problems/validation-failed",
  "title": "Validation failed",
  "status": 422,
  "code": "validation_failed",
  "trace_id": "01JC8Z4ABC",
  "errors": [
    { "pointer": "/tenders/0/actual", "code": "required", "message": "Enter the counted amount" }
  ]
}
```

| Status | Used for | Codes |
|---|---|---|
| 400 | Malformed request | `validation_failed` |
| 401 | Not signed in | `unauthenticated`, `token_expired`, `mfa_required`, `account_locked` |
| 403 | Not allowed | `forbidden`, `out_of_scope`, `sod_violation`, `channel_not_allowed` |
| 404 | Missing or outside scope | `not_found` |
| 409 | Blocked by state | `conflict`, `invalid_state_transition`, `duplicate_record`, `idempotency_key_reused`, `rule_not_configured`, `workflow_not_live`, `hold_active`, `approval_required`, `sync_conflict`, `offline_not_allowed` |
| 412 | Changed since read | `stale_version` |
| 422 | Valid request that breaks a rule | `validation_failed`, `limit_exceeded`, `signature_required`, `evidence_required`, `pin_invalid`, `pin_locked` |
| 429 | Too many requests | `rate_limited` |
| 500, 502 | Server or external failure | `internal_error`, `integration_unavailable` |

Clients branch on `code`, never on `title` or `detail`. A record out of the user's scope returns 404 rather than 403, so its existence is not revealed.

### 12.5 Pagination

All record lists use cursor pagination. `GET /permissions` and `GET /me/sessions` are short, and return `data` without a `page`.

```
GET /api/v1/tasks?assignee=me&limit=25&sort=-due_at&cursor=eyJ2Ijo...
```

```json
{
  "data": [ { "id": "...", "title": "..." } ],
  "page": { "limit": 25, "has_more": true, "next_cursor": "eyJ2Ijo...", "prev_cursor": null }
}
```

- `limit` runs from 1 to 100, default 25.
- `cursor` is opaque. Clients must not parse it, and must keep other parameters the same between pages.
- `sort` takes comma-separated fields, with `-` for descending. Default is `-created_at`.
- Lists do not return totals. Dashboards supply counts.
- `updated_since` supports incremental reads. `q` searches key fields.

### 12.6 Safe writes

- Every `POST` that creates or acts needs `Idempotency-Key`. A repeat within 72 hours returns the first result. Reusing a key with a different body returns `idempotency_key_reused`. Account and sign-in calls, webhooks and `POST /notifications/read-all` are exempt, and sync items use `client_record_id` instead.
- `PATCH` and `DELETE` need `If-Match` with the `ETag` from the last read. Ending your own session is the one exception.
- `PATCH` is a partial update. Omitted fields do not change.
- Locked records, such as approved or signed ones, reject edits with `invalid_state_transition`.

### 12.7 File uploads

1. `POST /attachments` with name, size, type and checksum. Returns a short-lived `upload_url`.
2. The client uploads to that URL.
3. `POST /attachments/{id}/complete`. The server checks the file before it becomes `available`.

### 12.8 Webhooks

`/webhooks/whatsapp` handles Meta's verification and inbound events, and `/webhooks/whatsapp/flows` serves Flow data. Both are verified, then queued and handled asynchronously, and duplicates are ignored by provider message id.

## 13. Non-functional requirements

These targets are proposals from this spec, not from the source documents. Please confirm them.

| Area | Proposed target |
|---|---|
| Performance | 95% of reads under 500 ms and writes under 1 s on Head Office connections. Branch screens usable on low-end Android phones on 3G. |
| WhatsApp latency | Reply within 3 seconds for 95% of messages. |
| Alert speed | Critical alerts sent within 60 seconds of classification. |
| Availability | 99.5% monthly, excluding planned windows. Offline capture covers branch outages. |
| Backups | Daily, as the Finance Policy requires. Restore tested each quarter. Proposed recovery point 1 hour and recovery time 4 hours. |
| Security | TLS in transit, encryption at rest. Argon2id for passwords and PINs. Field encryption for personal data. Annual penetration test. |
| Privacy | Minimum necessary personal data. Masking in lists. Confirm obligations under Zimbabwe's Cyber and Data Protection Act before go-live (N-08). |
| Scale | Designed for all sites and all employees. The current number of sites and staff is not in the source files (N-09). |
| Accessibility | WCAG 2.2 AA for the web app. Light and dark themes, as the current dashboard has. |
| Language | English for now. Confirm whether Shona or Ndebele are needed for the WhatsApp channel (N-10). |
| Hosting | Not decided. Confirm provider and data location (N-07). |
| Observability | Every request has a trace id. Integration and posting failures alert ICT. |

## 14. Risks and items needing attention

### Risks

| Risk | Effect | Response |
|---|---|---|
| Authorised Limit Matrix and other values are missing | D-10 means many workflows cannot go live | Readiness screen, with named owners. Start collecting now, and treat this as the biggest schedule risk. |
| GAAP, SAP and Belina integration may not be possible | Manual effort, delays | File import and manual posting on every path. Check feasibility in Phase 0. |
| WhatsApp as a full channel | Cost for all staff, Flow size limits, template approval, dependence on Meta | Estimate message volumes before committing. Keep the web app as a complete alternative for all role work. |
| Staff phones and data | WhatsApp-only staff may lack a smartphone or data | Survey staff. Shared branch tablets as fallback. |
| Scale of the design | 871 operations is a large build | Phasing, the form engine, and an honest review of scope after Phase 1. |
| Double entry | Staff key the same thing into POS, GAAP and Foodies | Prefill from feeds. Measure entry time in the pilot. |
| Parallel paper and digital | Two versions of the truth during rollout | Agree per register whether it is replaced or runs in parallel (C-11.09). |
| SOP conflicts | Digitising a conflict locks it in | Open decisions in Appendix C must close first. The pizza oven cleaning contradiction (C-02.09, C-02.10) is a safety example. |

### New items found while writing this spec

These are not in the workbook's clarifications.

| ID | Item | Owner |
|---|---|---|
| N-01 | SOPs that require WhatsApp group posts (petty cash, reports) need revising, since D-11 moves those records into Foodies. Until they are revised, the system does not match the SOP wording. | QMS Representative, MD |
| N-02 | Confirm that GAAP POS, GAAP Unity, SAP Business One and Belina offer APIs or scheduled exports, and who owns each. | ICT |
| N-03 | Foodies cannot block a POS sale itself. Agree the stop-sale process, or a POS-side control. | Operations, ICT |
| N-04 | Payment approval differs between sources: Finance Manager and MD, against the five-stage hierarchy including Chairman. Choose one. | Finance Manager, MD |
| N-05 | QC and Customer Service use different labels for severity levels 4 and 5. Confirm these should stay different. | QA, Operations |
| N-06 | Confirm the role merges in section 5.2. COO appears only on the CCTV form. Confirm whether that role exists. | HR, GM |
| N-07 | Choose hosting and data location. | ICT, MD |
| N-08 | Get legal review of data protection, employee monitoring, staff search records, and electronic signature validity. | Legal, HR |
| N-09 | Supply branch and employee counts for capacity planning and WhatsApp cost estimates. | Operations, HR |
| N-10 | Decide the languages for WhatsApp and the web app. | Operations |
| N-11 | The workbook names a Nash Paints HQ master key location, though Nash sites are out of scope. It is held as free text. Confirm this is right. | Security Manager |
| N-12 | Retention conflict: cash-up records are 2 years under the Store Closing SOP, and 6 years under Revenue and Finance rules. Confirm which class applies. | Finance Manager |
| N-13 | Agree how much WhatsApp use on personal phones is acceptable for restricted records and work outside hours, and who pays for data. | HR, ICT |

### Assumptions to confirm

- Foodies becomes the Central Kitchen ordering channel (C-04.12).
- PINs are 6 digits, and five wrong attempts lock signing or sign-in. Access tokens last 15 minutes. Idempotency keys last 72 hours.
- Group posting on WhatsApp is not used.
- The web app also covers all role work, so WhatsApp is never the only way in.

## Appendix A. Coverage of the checklist

Each of the 101 workbook areas, with its release phase and build approach.

- **Dedicated:** purpose-built entities and API.
- **Form engine:** runs on published templates.
- **Hybrid:** dedicated core, forms for the rest.
- **Platform:** met by shared capabilities.
- **External + form:** done in a system of record, with evidence in Foodies.

**1. CASH & FINANCE**

| Area | Items | Phase | Build | Note |
|---|---:|---|---|---|
| Revenue, POS & Cash Control | 13 | 1 | Hybrid | Flush, cash-up and variance are dedicated. Cash handling, safe and fake-note checks are forms. |
| CIT, Banking & Deposit Control | 14 | 1 | Dedicated |  |
| Change Float & Petty Cash | 16 | 1 | Dedicated |  |
| Accounts Payable & Supplier Management | 21 | 1 | Hybrid | Workflow here. GRNI, accruals, FX allocation and statements stay in SAP, with evidence in forms. |
| Budgeting & Expenditure Control | 10 | 1 | Dedicated |  |
| Bank Reconciliation & Treasury | 10 | 4 | External + form | Done in SAP. Review, ageing and evidence in Foodies. |
| Recipe Standard Costing & Profitability | 12 | 6 | Dedicated | Cost sheets. |
| Finance Governance: Payroll, Assets, Tax, Journals & Audit | 18 | 4 | Hybrid | Compliance calendar, forms and approvals. Calculations stay in SAP and Belina. |

**2. EQUIPMENT & ASSET MANAGEMENT**

| Area | Items | Phase | Build | Note |
|---|---:|---|---|---|
| Asset Register & Classification | 6 | 5 | Dedicated |  |
| CAPEX & Asset Acquisition | 8 | 5 | Dedicated |  |
| Depreciation & Impairment | 10 | 5 | External + form | Rates in asset classes. Runs in SAP. |
| Asset Verification, Transfers & Disposal | 11 | 5 | Dedicated |  |
| Generator Operation & Maintenance | 12 | 3, 4 | Hybrid | Faults and work orders are dedicated. Daily logs are forms. Plans arrive in Phase 4. |
| Pizza Oven Operation & Maintenance | 11 | 3, 4 | Hybrid | As generator. |
| Reporting, Records & System Governance | 6 | 4 | Platform |  |

**3. GENERAL CONTROLS**

| Area | Items | Phase | Build | Note |
|---|---:|---|---|---|
| Management Reporting & Submission | 25 | 4 | Hybrid | Calendar is dedicated. Report content is forms. |
| QMS Documentation & Document Control | 18 | 5 | Dedicated |  |
| Cleaning & Sanitation | 13 | 6 | Form engine |  |
| Staff Uniform Issuance & Replacement | 10 | 6 | Hybrid |  |

**4. INVENTORY & STORES**

| Area | Items | Phase | Build | Note |
|---|---:|---|---|---|
| Stock Planning & Master Data | 7 | 2 | Dedicated |  |
| Receiving & Inspection | 8 | 1, 2 | Dedicated | GRV in Phase 1. Temperature, batch and HOLD links in Phase 2. |
| Storage, Issue & Transfers | 11 | 2 | Hybrid | Issue and transfers are dedicated. Storage checks are forms. |
| Stocktake, Variance & Adjustments | 8 | 2 | Dedicated |  |
| Procurement & Payment | 15 | 1 | Dedicated |  |
| Central Kitchen Ordering | 11 | 2 | Dedicated |  |
| Waste & Yield Control | 9 | 2 | Dedicated |  |
| System, Audit & Reporting | 7 | 2 | Platform |  |

**5. PRODUCTION & KITCHEN OPERATIONS**

| Area | Items | Phase | Build | Note |
|---|---:|---|---|---|
| Production planning | 2 | 2 | Hybrid |  |
| Recipe and formulation control | 1 | 2 | Dedicated |  |
| Production records | 1 | 2 | Dedicated |  |
| Ingredient control | 1 | 2 | Dedicated |  |
| Kitchen preparation | 2 | 2 | Form engine |  |
| Pizza production | 3 | 2 | Form engine |  |
| Chicken processing | 5 | 2 | Form engine | Release blocks are template rules. |
| Potato processing | 3 | 2 | Form engine |  |
| Production reconciliation | 2 | 2 | Dedicated |  |
| Waste and yield | 3 | 2 | Dedicated |  |
| Pizza oven | 3 | 3 | Form engine |  |
| Finished goods dispatch | 5 | 2 | Dedicated |  |
| Bakery operations | 6 | 2 | Form engine |  |

**6. QUALITY, FOOD SAFETY & COMPLIANCE**

| Area | Items | Phase | Build | Note |
|---|---:|---|---|---|
| Quality Control | 19 | 3 | Hybrid | NCR, HOLD and disposition are dedicated. Inspections are forms. |
| SHE Management | 23 | 3, 6 | Hybrid | Incidents in Phase 3. Registers and inspections are forms in Phase 6. |
| Food Safety | 18 | 3, 6 | Hybrid | Unsafe food uses HOLD in Phase 3. Daily checks are forms in Phase 6. |
| New Product Development | 18 | 6 | Hybrid |  |
| Product Recall | 13 | 3 | Dedicated |  |

**7. SALES & CUSTOMER SERVICE**

| Area | Items | Phase | Build | Note |
|---|---:|---|---|---|
| Revenue Management | 12 | 1 | Hybrid | Shares cash-up. Price list and POS access are forms. |
| Customer Service | 16 | 3 | Hybrid | Cases are dedicated. Service checklists are forms. |
| Electronic Payments | 7 | 3 | Dedicated |  |
| Manual Receipting | 8 | 3 | Dedicated |  |
| Lost Property | 6 | 5 | Dedicated |  |
| External Visits | 9 | 6 | Dedicated |  |

**8. STORE OPERATIONS**

| Area | Items | Phase | Build | Note |
|---|---:|---|---|---|
| Store Opening | 17 | 6 | Hybrid | GO / NO-GO gate is dedicated. Control checks are forms. |
| Store Closing | 10 | 1, 6 | Form engine | Witnessed cash-up lands in Phase 1. The other checks are forms. |
| CCTV Management | 10 | 6 | Hybrid |  |
| Customer Sitting Area Furniture | 7 | 6 | Form engine |  |
| Key Handling | 10 | 5 | Dedicated |  |

**9. HUMAN RESOURCES**

| Area | Items | Phase | Build | Note |
|---|---:|---|---|---|
| Employee records | 6 | 0 | Dedicated |  |
| Recruitment | 5 | 6 | Dedicated |  |
| Onboarding & induction | 5 | 5 | Form engine |  |
| Attendance & timekeeping | 4 | 6 | Dedicated |  |
| Leave management | 4 | 5 | Dedicated |  |
| Payroll coordination | 4 | 6 | External + form | Approved changes post to Belina. |
| Performance management | 3 | 6 | Dedicated |  |
| Training & development | 3 | 5 | Dedicated | Uses Training. |
| Disciplinary management | 3 | 6 | Dedicated |  |
| Grievance management | 2 | 6 | Dedicated |  |
| Employee relations | 1 | 6 | Form engine |  |
| Uniform management | 10 | 6 | Hybrid |  |
| Communication & employee awareness | 8 | 6 | Form engine |  |
| Employee exit & clearance | 6 | 5 | Dedicated |  |
| HR reporting & compliance | 4 | 5 | Platform |  |
| HR access, audit & governance | 4 | 0 | Platform |  |

**10. TRAINING & PRODUCT KNOWLEDGE**

| Area | Items | Phase | Build | Note |
|---|---:|---|---|---|
| Training governance | 9 | 5 | Dedicated |  |
| Induction and onboarding | 4 | 5 | Hybrid |  |
| Food safety and hygiene | 9 | 5 | Form engine | Course content and records. |
| Health, safety and security | 4 | 5 | Form engine | Course content and records. |
| Operational skills | 4 | 5 | Form engine | Course content and records. |
| Customer service and sales | 2 | 5 | Form engine | Course content and records. |
| Supervisory and management | 1 | 5 | Form engine | Course content and records. |
| Assessment and competency | 9 | 5 | Dedicated |  |
| Product knowledge governance | 20 | 2, 6 | Hybrid | Product master in Phase 2. Knowledge assessment in Phase 6. |
| Training and product knowledge reporting | 4 | 5 | Platform |  |
| Training and product knowledge integration | 3 | 5 | Platform |  |
| Training and product knowledge controls | 3 | 0 | Platform |  |

**11. SECURITY**

| Area | Items | Phase | Build | Note |
|---|---:|---|---|---|
| Shift management & handover | 7 | 6 | Hybrid |  |
| Gate & access control | 6 | 6 | Form engine |  |
| Visitor management | 2 | 5 | Dedicated |  |
| Staff searches & declarations | 7 | 6 | Form engine |  |
| Vehicle control | 4 | 6 | Form engine |  |
| Receiving & dispatch control | 8 | 2 | Hybrid | Security steps on transfers and D-Notes. Other checks are forms. |
| Cash & cash-up security | 7 | 1 | Hybrid | Witness steps on flush and cash-up. Registers are forms. |
| Staff purchases & product testing | 5 | 6 | Form engine |  |
| Patrols & inspections | 5 | 6 | Hybrid |  |
| CCTV management | 12 | 6 | Hybrid |  |
| Keys & premises security | 3 | 5 | Dedicated |  |
| Emergency & fire safety | 4 | 6 | Form engine |  |
| Closing procedure | 5 | 6 | Form engine |  |
| Incident management & reporting | 4 | 3 | Dedicated |  |
| Security registers & records | 2 | 6 | Form engine |  |
| Security compliance & conduct | 2 | 6 | Form engine |  |
| Security reporting & dashboards | 5 | 6 | Platform |  |

## Appendix B. Rules

### Values stated in the checklist

These load as proposals. Each needs its owner's approval, checked against the SOP.

| Rule key | Value | Source |
|---|---|---|
| `cash.till_flush_threshold` | US$1,000 | Receipts, Cash Control and CIT SOP §5 |
| `cash.cit_transit_limit` | US$10,000 per transit | Receipts, Cash Control and CIT SOP §7 |
| `cash.change_float_max` | US$1,500 | Change Float and Petty Cash SOP §3–4 |
| `cash.petty_cash_max` | US$1,000 | Change Float and Petty Cash SOP §3–4 (see C-01.03) |
| `cash.variance_escalation` | Over US$5 or ZiG equivalent at RBZ rate | Store Closing |
| `cash.variance_repeat_trigger` | 3 in a week or 3 in a month, same cashier | Receipts, Cash Control and CIT SOP §6 |
| `cash.banking_deadline` | Daily, or next working day | Finance Policy §9 |
| `cash.banking_proof_hours` | 24 | Finance Policy §9 |
| `cash.cit_incident_report_hours` | 2 | CIT SOP |
| `cash.float_exception_reimbursement` | Next day | Change Float SOP |
| `procurement.invoice_forward_days` | 3 | Accounts Payables SOP §4 |
| `procurement.emergency_approval_hours` | 48 | Finance Policy §7 |
| `procurement.local_lead_time_working_days` | 3 | Procurement SOP §8 |
| `procurement.po_review_hours` | 24 (see C-04.04) | Procurement & Payment Workflow |
| `tax.wht_remittance_day` | 10th of following month | Accounts Payables SOP §5.5 |
| `finance.audit_correction_days` | 30 | Finance Policy §14 |
| `budget.variance_threshold_percent` | 10 (see C-01.11) | Finance Policy, Budgeting SOP |
| `budget.escalation_working_days` | 5 | Budgeting SOP |
| `bank.reconciliation_working_days` | 5 | Bank Reconciliation SOP |
| `stock.chilled_max_c` | ≤ 4°C | Stores SOP, Store Opening |
| `stock.frozen_max_c` | ≤ -18°C | Stores SOP, Store Opening |
| `stock.min_supply_days` | 14 (see C-04.05) | Procurement & Payment Workflow |
| `stock.pr_trigger_days_cover` | 7 (see C-04.05) | Procurement & Payment Workflow |
| `stock.tba_followup_hours` | 24 (see C-04.06) | Stores SOP |
| `stock.transfer_discrepancy_close_hours` | 48 (see C-04.06) | Stores SOP |
| `ck.order_deadline_morning` | 10:00 | Central Kitchen SOP |
| `ck.order_deadline_evening` | 19:00 | Central Kitchen SOP |
| `production.yield_variance_percent` | 3 (see C-04.09) | Waste Yield SOP |
| `production.waste_target_percent` | Below 3, or approved target | Waste Yield SOP |
| `production.dough_max_hours` | 24 | Pizza SOP |
| `production.thaw_max_hours` | 48 whole birds, 24 portions, at 5°C or lower | Butchery SOP |
| `production.chicken_pack_count` | 36 | Butchery SOP |
| `production.freezer_hold_min_hours` | 3 | Butchery SOP |
| `oven.set_point` | 500°F / 260°C | Pizza Oven SOP |
| `oven.conveyor_time` | 4 min 15 s | Pizza Oven SOP |
| `she.incident_report_hours` | 24 | SHE SOP |
| `she.drill_report_days` | 14 | SHE SOP |
| `she.oven_fryer_log_hours` | 2 | SHE SOP |
| `she.unsafe_food_hours` | Discard above 2 hours at unsafe temperature | SHE SOP |
| `quality.complaint_ack_hours` | 24 (bakery) | Bakery Policy |
| `quality.complaint_investigation_days` | 3 business days (bakery) | Bakery Policy |
| `quality.complaint_resolution_days` | 5 business days (bakery) | Bakery Policy |
| `quality.recall_trace_hours` | 4 (bakery) | Bakery Policy |
| `documents.turnaround_working_days` | 5 | QMS procedure |
| `lost_property.card_days` | 30 | Bank Cards and Personal Items SOP |
| `lost_property.personal_item_days` | 90 | Bank Cards and Personal Items SOP |
| `keys.incident_report_hours` | 24 | Key Handling SOP |
| `security.prereport_minutes` | 15 | Security SOP |
| `hr.uniform_replacement_months` | 6 | Uniform SOP |
| `assets.disposal_report_days` | 7 | Non-Current Asset Management §9 |
| `assets.depreciation_rates` | Land and buildings 2%. Kitchen, vehicles, furniture 10%. Computers 20%. | Asset SOP |
| `hygiene.audit_pass_percent` | 95 | Cleaning SOP |

### Report deadlines

| Rule key | Report | Due |
|---|---|---|
| `reports.daily_cash_books_due` | Cash books and bank balances | 11:00 |
| `reports.daily_payment_plan_due` | Payment plan | 12:00 |
| `reports.daily_turnover_2200_due` | Turnover, shops closing 22:00 | 23:00 |
| `reports.daily_turnover_24h_due` | Turnover, 24-hour shops | 08:00 |
| `reports.daily_sales_control_due` | Sales control | 10:00 |
| `reports.daily_production_due` | Production | 18:00 |
| `reports.daily_quality_due` | Quality and food safety | 18:00 |
| `reports.weekly_due` | Weekly departmental | Monday 15:00 |
| `reports.biweekly_due` | Bi-weekly executive | Every two weeks, Friday 17:00 |
| `reports.monthly_due` | Monthly departmental | On or before the 3rd |

### Known unset rules

These have no value in the sources, so their workflows wait for them. The list is not exhaustive. The readiness screen is the working list.

| Rule | Decision |
|---|---|
| Authorised Limit Matrix bands | C-01.08, C-02.02, C-04.02, C-06.04, C-07.05 |
| `cash.sanction_bands` | C-01.02 |
| `procurement.quotation_count` | C-04.03 |
| `cash.petty_cash_reimbursement_day` | C-01.07 |
| `cash.cit_collection_schedule` | C-01.06 |
| ZIMRA rates and dates | C-01.09 |
| Recipe costing review triggers | C-01.13 |
| `assets.capitalisation_threshold` | C-02.01 |
| Generator test and reading frequency | C-02.06, C-02.07 |
| Restroom and sanitiser frequency, dilution, contact time | C-03.07 to C-03.10 |
| Cold-chain monitoring intervals | C-04.10 |
| Temperature limits for hot and cold holding | C-06.01 |
| QC sampling plans | C-06.02 |
| `security.patrol_frequency` | C-11.06 |
| CCTV retention period | C-08.05, C-11.03 |
| Severity recipients and timers | C-06.09, C-07.10, C-11.05 |
| Complaint SLAs by severity | C-07.04 |
| Leave types and accruals | C-09.03 |
| Pass marks and refresher intervals | C-10.01, C-10.06, C-10.07 |
| Retention schedule | C-04.07 and others |
| `auth.lockout` duration | Not in the sources. ICT to set. |

## Appendix C. Decision register

All 124 workbook clarifications, numbered by module and row. Status now: 12 resolved, 1 assumed, 11 partly resolved, 100 open. An open item needs an answer from its owner, and most block a rule in Appendix B.

The owner column is a suggested owner by module.

**1. CASH & FINANCE**

| ID | Decision | Owner | Status | Note |
|---|---|---|---|---|
| C-01.01 | Cash-up variance recovery wording | Finance Manager | Open | Workflow records recovery. Payroll handling needs approval. |
| C-01.02 | Cash discrepancy sanction table | Finance Manager | Partly | Foodies records and refers only. The table itself is still needed. |
| C-01.03 | Change float vs petty cash | Finance Manager | Open |  |
| C-01.04 | Petty cash approvals across procedures | Finance Manager | Partly | D-11: approval moves to Foodies, on the web app or WhatsApp. SOP wording must change. |
| C-01.05 | CIT transit above insurance limit | Finance Manager | Open |  |
| C-01.06 | CIT collection schedule | Finance Manager | Open |  |
| C-01.07 | Weekly petty cash reimbursement day | Finance Manager | Open |  |
| C-01.08 | Supplier payment authority matrix | Finance Manager | Open |  |
| C-01.09 | Statutory tax rates and dates | Finance Manager | Open |  |
| C-01.10 | Budget internal submission deadline | Finance Manager | Open |  |
| C-01.11 | Budget variance threshold | Finance Manager | Open |  |
| C-01.12 | Bank reconciliation frequency | Finance Manager | Open |  |
| C-01.13 | Recipe costing review triggers | Finance Manager | Open |  |
| C-01.14 | Recipe costing record retention | Finance Manager | Open |  |
| C-01.15 | Finance Policy scope and Xactco integration | Finance Manager | Partly | D-06 sets the principle. The function list still needs confirming. |
| C-01.16 | CIT incident reporting | Finance Manager | Open |  |

**2. EQUIPMENT & ASSET MANAGEMENT**

| ID | Decision | Owner | Status | Note |
|---|---|---|---|---|
| C-02.01 | Minor asset capitalisation threshold | Finance Manager, Maintenance | Open |  |
| C-02.02 | Authorised Limit Matrix | Finance Manager, Maintenance | Open |  |
| C-02.03 | CAPEX budget and approval routing | Finance Manager, Maintenance | Open |  |
| C-02.04 | Depreciation start convention | Finance Manager, Maintenance | Open |  |
| C-02.05 | Residual values and impairment frequency | Finance Manager, Maintenance | Open |  |
| C-02.06 | Generator load-bank/full-load test frequency | Finance Manager, Maintenance | Open |  |
| C-02.07 | Generator operating readings frequency | Finance Manager, Maintenance | Open |  |
| C-02.08 | Generator maintenance schedule | Finance Manager, Maintenance | Open |  |
| C-02.09 | Pizza oven weekly cleaning instruction | Finance Manager, Maintenance | Open |  |
| C-02.10 | Pizza oven shutdown wording | Finance Manager, Maintenance | Open |  |
| C-02.11 | Pizza oven cleaning liquid and chemical controls | Finance Manager, Maintenance | Open |  |
| C-02.12 | Asset maintenance scope | Finance Manager, Maintenance | Open |  |
| C-02.13 | Asset tagging standard | Finance Manager, Maintenance | Open |  |
| C-02.14 | Internal transfer approval and handover | Finance Manager, Maintenance | Open |  |
| C-02.15 | Retention and repository | Finance Manager, Maintenance | Partly | Evidence stays in Foodies. The retention schedule is still needed. |

**3. GENERAL CONTROLS**

| ID | Decision | Owner | Status | Note |
|---|---|---|---|---|
| C-03.01 | Reports Submission Policy – monthly deadline | General Manager, QMS Representative | Open |  |
| C-03.02 | Reports Submission Policy – immediate escalation | General Manager, QMS Representative | Open |  |
| C-03.03 | Reports Submission Policy – reporting format | General Manager, QMS Representative | Open |  |
| C-03.04 | Reports Submission Policy – time zone and business calendar | General Manager, QMS Representative | Partly | D-15 sets Africa/Harare. The working-day calendar still needs loading. |
| C-03.05 | QMS documentation – authority matrix | General Manager, QMS Representative | Open |  |
| C-03.06 | QMS documentation – retention periods | General Manager, QMS Representative | Open |  |
| C-03.07 | Cleaning procedure – restroom frequency | General Manager, QMS Representative | Open |  |
| C-03.08 | Cleaning procedure – toilet cleaner contact time | General Manager, QMS Representative | Open |  |
| C-03.09 | Cleaning procedure – sanitizer dilution and contact time | General Manager, QMS Representative | Open |  |
| C-03.10 | Cleaning procedure – task schedule by shift | General Manager, QMS Representative | Open |  |
| C-03.11 | Uniform procedure – document number | General Manager, QMS Representative | Open |  |
| C-03.12 | Uniform procedure – uniform stock reorder levels | General Manager, QMS Representative | Open |  |
| C-03.13 | Uniform procedure – negligence/loss deductions | General Manager, QMS Representative | Open |  |
| C-03.14 | Cleaning procedure – numbering | General Manager, QMS Representative | Open |  |

**4. INVENTORY & STORES**

| ID | Decision | Owner | Status | Note |
|---|---|---|---|---|
| C-04.01 | Stock system integration | Operations Manager, Stores Controller | Partly | D-06: GAAP Unity holds the authoritative balance. Method and DRS update still open. |
| C-04.02 | Procurement approval limits | Operations Manager, Stores Controller | Open |  |
| C-04.03 | Quotation requirement | Operations Manager, Stores Controller | Open |  |
| C-04.04 | PR / PO approval chain | Operations Manager, Stores Controller | Open |  |
| C-04.05 | 14-day and 7-day stock cover | Operations Manager, Stores Controller | Open |  |
| C-04.06 | Transfer follow-up | Operations Manager, Stores Controller | Open |  |
| C-04.07 | Record retention conflicts | Operations Manager, Stores Controller | Open |  |
| C-04.08 | Stock variance targets | Operations Manager, Stores Controller | Open |  |
| C-04.09 | Waste and yield thresholds | Operations Manager, Stores Controller | Open |  |
| C-04.10 | Cold-chain monitoring intervals | Operations Manager, Stores Controller | Open |  |
| C-04.11 | Stocktake schedule | Operations Manager, Stores Controller | Open |  |
| C-04.12 | Central Kitchen ordering channel | Operations Manager, Stores Controller | Assumed | Foodies becomes the ordering channel. Please confirm. |
| C-04.13 | CCTV-style/technical evidence not assumed | Operations Manager, Stores Controller | Partly | D-12 keeps listed paper originals. The final list is still needed. |

**5. PRODUCTION & KITCHEN OPERATIONS**

| ID | Decision | Owner | Status | Note |
|---|---|---|---|---|
| C-05.01 | Duplicate document number | Production Manager | Open |  |
| C-05.02 | Overlapping chicken procedures | Production Manager | Open |  |
| C-05.03 | Policy content shared with other application folders | Production Manager | Resolved | Shared controls sit in one module and are linked. |
| C-05.04 | Parameters must remain configurable and controlled | Production Manager | Resolved | D-10: versioned rules. |
| C-05.05 | Retention periods differ by record type | Production Manager | Resolved | Retention by record class. |

**6. QUALITY, FOOD SAFETY & COMPLIANCE**

| ID | Decision | Owner | Status | Note |
|---|---|---|---|---|
| C-06.01 | Food safety temperature limits are not numerically stated in Food Safety Procedure | SHEQ, QA | Open |  |
| C-06.02 | QC sampling frequency and acceptance limits | SHEQ, QA | Open |  |
| C-06.03 | SHE incident reporting and notification details | SHEQ, QA | Open |  |
| C-06.04 | Recall and QC delegated approval limits | SHEQ, QA | Open |  |
| C-06.05 | Overlapping quality, food safety and recall workflows | SHEQ, QA | Resolved | One linked HOLD, NCR and incident chain. See section 10. |
| C-06.06 | Sensitive employee health and medical records | SHEQ, QA | Resolved | Dashboards show status only. See section 5. |
| C-06.07 | NPD document retention | SHEQ, QA | Open |  |
| C-06.08 | SHE record retention differs by category | SHEQ, QA | Open |  |
| C-06.09 | Critical failure escalation route | SHEQ, QA | Open |  |
| C-06.10 | Product recall external communication | SHEQ, QA | Open |  |

**7. SALES & CUSTOMER SERVICE**

| ID | Decision | Owner | Status | Note |
|---|---|---|---|---|
| C-07.01 | Revenue Management and Lost Property share document number FOODIES/FIN/SOP08 | Operations Manager, Finance | Open |  |
| C-07.02 | Revenue Management tax and pricing configuration | Operations Manager, Finance | Open |  |
| C-07.03 | Revenue reports and reconciliations | Operations Manager, Finance | Open |  |
| C-07.04 | Customer service response and resolution SLAs | Operations Manager, Finance | Open |  |
| C-07.05 | Refund, voucher and compensation authority | Operations Manager, Finance | Open |  |
| C-07.06 | Electronic payment provider processes | Operations Manager, Finance | Open |  |
| C-07.07 | Manual receipt retention | Operations Manager, Finance | Open |  |
| C-07.08 | External visit records retention | Operations Manager, Finance | Open |  |
| C-07.09 | Lost property personal data | Operations Manager, Finance | Resolved | Minimum details only. No full card numbers or PINs. |
| C-07.10 | Complaint severity routing | Operations Manager, Finance | Open |  |

**8. STORE OPERATIONS**

| ID | Decision | Owner | Status | Note |
|---|---|---|---|---|
| C-08.01 | Opening authority and approval chain | Operations Manager | Open |  |
| C-08.02 | Opening schedule | Operations Manager | Open |  |
| C-08.03 | Manual docket contingency | Operations Manager | Partly | Digital approval is available. MD or Chairman to confirm it is acceptable. |
| C-08.04 | CCTV authorization wording | Operations Manager | Open |  |
| C-08.05 | CCTV footage retention | Operations Manager | Open |  |
| C-08.06 | Furniture cleaning frequencies | Operations Manager | Open |  |
| C-08.07 | Store closing records | Operations Manager | Open |  |
| C-08.08 | Key register confidentiality | Operations Manager | Partly | Alarm codes are never stored. Branch visibility is still to confirm. |

**9. HUMAN RESOURCES**

| ID | Decision | Owner | Status | Note |
|---|---|---|---|---|
| C-09.01 | Only two HR-related source documents were supplied for this section | HR | Open |  |
| C-09.02 | Recruitment, onboarding, employee records, leave, attendance, discipline, grievances, performance and exit procedures | HR | Open |  |
| C-09.03 | Leave types, accruals and entitlements | HR | Open |  |
| C-09.04 | Disciplinary process timelines and sanctions | HR | Partly | Foodies tracks the process and never decides. Timelines are still needed. |
| C-09.05 | Payroll changes and final pay | HR | Resolved | Payroll stays with Finance and Belina. HR sends approved changes. |
| C-09.06 | HR record retention periods | HR | Open |  |
| C-09.07 | Uniform replacement responsibility | HR | Resolved | Roles configured as stated. |
| C-09.08 | Uniform schedule detail | HR | Open |  |
| C-09.09 | Communication Policy acknowledgements | HR | Resolved | Acknowledgement workflow in Documents. |
| C-09.10 | Official WhatsApp groups | HR | Open |  |
| C-09.11 | Employee exits and communication access | HR | Resolved | Exit clearance and immediate WhatsApp unbinding. |

**10. TRAINING & PRODUCT KNOWLEDGE**

| ID | Decision | Owner | Status | Note |
|---|---|---|---|---|
| C-10.01 | Approved training framework | Training Manager | Open |  |
| C-10.02 | Product master and recipes | Training Manager | Open |  |
| C-10.03 | Allergen information | Training Manager | Open |  |
| C-10.04 | Food safety parameters | Training Manager | Open |  |
| C-10.05 | Training authority | Training Manager | Open |  |
| C-10.06 | Assessment rules | Training Manager | Open |  |
| C-10.07 | Training frequency | Training Manager | Open |  |
| C-10.08 | Product launch governance | Training Manager | Open |  |
| C-10.09 | Existing HR training module | Training Manager | Resolved | One Training module. HR reuses it. |
| C-10.10 | Examples of operational/product categories | Training Manager | Open |  |

**11. SECURITY**

| ID | Decision | Owner | Status | Note |
|---|---|---|---|---|
| C-11.01 | Foodies-specific post orders | Security Manager | Open |  |
| C-11.02 | Search procedure | Security Manager | Open |  |
| C-11.03 | CCTV retention period | Security Manager | Open |  |
| C-11.04 | CCTV authorised access | Security Manager | Open |  |
| C-11.05 | Incident escalation | Security Manager | Open |  |
| C-11.06 | Patrol frequency | Security Manager | Open |  |
| C-11.07 | Cash and safe controls | Security Manager | Open |  |
| C-11.08 | Visitor and contractor data | Security Manager | Open |  |
| C-11.09 | Security registers | Security Manager | Open |  |
| C-11.10 | Site-specific rules | Security Manager | Resolved | D-13: Nash sites are out of scope. |
| C-11.11 | CCTV system integration | Security Manager | Open |  |
| C-11.12 | Approval of digital evidence | Security Manager | Partly | D-12 sets PIN sign-off. Legal acceptance still to confirm. |

## Appendix D. Glossary

| Term | Meaning |
|---|---|
| Blocking rule | A rule a workflow cannot run without |
| BOE | Bill of Entry for imports |
| CAPA | Corrective and preventive action |
| CCP | Critical control point |
| CIT | Cash in transit |
| D-Note | Delivery note for Central Kitchen dispatch |
| DRS | Stock record named in the Stocktake SOP, alongside GAAP |
| FAR | Fixed Asset Register |
| FEFO, FIFO | First expired first out, first in first out |
| FOH, BOH | Front of house, back of house |
| Go live | Switch a workflow on for staff at one site |
| GRV, GRN | Goods received voucher or note |
| GRNI | Goods received not invoiced |
| HOD | Head of Department |
| HOLD | Status that stops product being used or moved |
| HOTO | Handover and takeover |
| IMR | Integrated Management Representative, in QMS |
| NCR | Non-conformance report |
| NPD | New product development |
| NSSA | National Social Security Authority |
| OB | Occurrence Book |
| POP | Proof of payment |
| PR, PO | Purchase requisition, purchase order |
| QMS | Quality management system |
| RBZ | Reserve Bank of Zimbabwe |
| Readiness | The view of what each workflow is missing, by site |
| Rule | A versioned, owned value, such as a limit, deadline or frequency |
| SHE, SHEQ | Safety, health and environment, plus quality |
| TBA | Outstanding transfer item in the Stores SOP |
| WHT | Withholding tax |
| Workflow | Work that follows set steps and people and leaves a record (section 8.1) |
| ZiG | Zimbabwe Gold currency |
| ZIMRA | Zimbabwe Revenue Authority |
