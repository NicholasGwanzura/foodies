#!/usr/bin/env python3
"""The slices, in build order. `python3 plan.py` applies every slice listed.

Each entry is (migration number, migration name, data-model sections). The
extra SQL is infrastructure the data model leaves out on purpose.
"""
import os, sys
sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import scaffold as sc  # noqa: E402

INFRA_UP = """
CREATE TABLE user_credentials (
  user_id uuid PRIMARY KEY,
  password_hash text,
  pin_hash text,
  totp_secret text,
  totp_confirmed_at timestamptz,
  failed_logins integer NOT NULL DEFAULT 0,
  failed_pins integer NOT NULL DEFAULT 0,
  pin_locked_at timestamptz,
  updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE auth_sessions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  user_id uuid NOT NULL,
  channel text NOT NULL,
  device_id uuid,
  ip text,
  created_at timestamptz NOT NULL DEFAULT now(),
  last_seen_at timestamptz NOT NULL DEFAULT now(),
  revoked_at timestamptz
);
CREATE INDEX ON auth_sessions (user_id);

CREATE TABLE refresh_tokens (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  session_id uuid NOT NULL REFERENCES auth_sessions (id) ON DELETE CASCADE,
  token_hash text NOT NULL UNIQUE,
  expires_at timestamptz NOT NULL,
  used_at timestamptz
);

CREATE TABLE mfa_challenges (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  user_id uuid NOT NULL,
  device_id uuid,
  channel text NOT NULL,
  expires_at timestamptz NOT NULL,
  failed_attempts integer NOT NULL DEFAULT 0,
  consumed_at timestamptz
);

CREATE TABLE idempotency_keys (
  actor_id uuid NOT NULL,
  key text NOT NULL,
  request_hash text NOT NULL,
  status integer,
  headers jsonb,
  body bytea,
  created_at timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY (actor_id, key)
);

CREATE TABLE reference_sequences (
  prefix text NOT NULL,
  site_code text NOT NULL,
  year integer NOT NULL,
  last_value bigint NOT NULL DEFAULT 0,
  PRIMARY KEY (prefix, site_code, year)
);

-- +goose StatementBegin
CREATE FUNCTION forbid_change() RETURNS trigger LANGUAGE plpgsql AS $$
BEGIN
  RAISE EXCEPTION '% is append-only', TG_TABLE_NAME;
END;
$$;
-- +goose StatementEnd
"""
INFRA_DOWN = """
DROP FUNCTION IF EXISTS forbid_change CASCADE;
DROP TABLE IF EXISTS reference_sequences, idempotency_keys, mfa_challenges, refresh_tokens, auth_sessions, user_credentials CASCADE;
"""
AUDIT_UP = """
ALTER TABLE audit_events ADD COLUMN seq bigint GENERATED ALWAYS AS IDENTITY;
CREATE UNIQUE INDEX ON audit_events (seq);
CREATE TRIGGER audit_events_append_only BEFORE UPDATE OR DELETE ON audit_events
  FOR EACH ROW EXECUTE FUNCTION forbid_change();
"""


WORKFLOW_UP = """
-- Which rules must have an active value before a workflow can go live. The data
-- model keeps this on templates (required_rule_keys); dedicated workflows use this.
CREATE TABLE workflow_requirements (
  workflow_key text NOT NULL,
  rule_key text NOT NULL,
  PRIMARY KEY (workflow_key, rule_key)
);
"""
WORK_UP = """
CREATE TABLE notification_preferences (
  user_id uuid PRIMARY KEY,
  categories jsonb NOT NULL DEFAULT '{}',
  updated_at timestamptz NOT NULL DEFAULT now()
);
"""
WORK_POST = """
CREATE UNIQUE INDEX tasks_once_per_day ON tasks (schedule_id, site_id, business_date) WHERE schedule_id IS NOT NULL;
CREATE INDEX ON tasks (assignee_user_id);
CREATE INDEX ON notifications (recipient_user_id, read_at);
ALTER TABLE tasks ADD COLUMN reminded_at timestamptz;
"""


def tables(*sections):
    return [t for t in sc.ORDER if sc.T[t]["section"] in sections]


SLICES = [
    (2, "organisation_people", "org", ["Organisation", "People & Access", "Audit"], INFRA_UP, INFRA_DOWN, AUDIT_UP),
    (3, "rules_approvals_evidence", "core", ["Rules & Readiness", "Approvals", "Signatures", "Evidence"], WORKFLOW_UP, "DROP TABLE IF EXISTS workflow_requirements;", ""),
    (4, "tasks_forms_escalations", "work", ["Tasks & Schedules", "Forms", "Escalations & Actions", "Notifications"], WORK_UP, "DROP TABLE IF EXISTS notification_preferences;", WORK_POST),
    (5, "integrations_reporting", "ops", ["Sync", "WhatsApp", "Integrations", "Reporting", "Retention"], "", "", ""),
]


def main():
    done = set()
    for number, name, key, sections, up, down, post in SLICES:
        ts = tables(*sections)
        done = sc.write_migration(number, name, ts, done, up, down, post)
    sc.write_search([t for s in SLICES for t in tables(*s[3])])
    for number, name, key, sections, *_ in SLICES:
        sc.scaffold_section(key, tables(*sections))
    sc.write_registry([t for s in SLICES for t in tables(*s[3])])
    sc.subprocess.call("gofmt -w internal/service", shell=True, cwd=sc.API)
    import ui
    built = [t for *_x, sections, _u, _d, _p in [(s[0], s[1], s[2], s[3], s[4], s[5], s[6]) for s in SLICES] for t in tables(*sections)]
    ui.emit(built, os.path.join(sc.ROOT, "web", "lib", "resources.gen.ts"))


if __name__ == "__main__":
    main()
