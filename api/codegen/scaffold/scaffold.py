#!/usr/bin/env python3
"""Scaffolds the mechanical CRUD layer of a slice from the data model and the API.

This is not part of `make gen`. Run it by hand when a slice adds tables:

    python3 api/codegen/scaffold/scaffold.py --section "Organisation" ...

It writes a goose migration (once), sqlc queries, and `crud_*.go` files in the
service and http packages. `crud_*.go` may be rewritten at any time, so never
edit one. To give an operation real behaviour, write the function in any other
file of that package: the scaffolder sees it there and stops emitting its own.
Needs Python 3 and PyYAML.
"""
import argparse, collections, glob, os, re, subprocess, sys
import yaml

HERE = os.path.dirname(os.path.abspath(__file__))
sys.path.insert(0, HERE)
from datamodel import parse  # noqa: E402
import workflows  # noqa: E402

ROOT = os.path.abspath(os.path.join(HERE, "..", "..", ".."))
API = os.path.join(ROOT, "api")
spec = yaml.safe_load(open(os.path.join(API, "openapi.yaml")))
SCH = spec["components"]["schemas"]
T, ORDER = parse(os.path.join(ROOT, "docs", "data-model.md"))

RESERVED = {"key", "type", "user", "order", "position", "values", "value", "from", "to", "default", "check",
            "limit", "offset", "end", "group", "table", "column", "desc", "asc", "when", "case", "primary",
            "references", "grant", "window", "action", "scope", "role", "date", "time", "timestamp", "level", "name",
            "method", "reason", "result", "source", "status", "comment", "version", "year", "month", "day", "zone",
            "like", "in", "is", "on", "or", "and", "not", "null", "true", "false", "all", "any", "as", "by"}
ALIASES = {"ZReport": "z_reports", "CitPerson": "cit_personnel"}
AUDIT_COLS = {"id", "created_at", "updated_at", "created_by", "updated_by", "row_version", "reference", "status"}


def q(c):
    return f'"{c}"' if c in RESERVED else c


def snake(n):
    n = re.sub(r"(?<=[a-z0-9])([A-Z])", r"_\1", n)
    return n.lower()


def plural(n):
    if n.endswith("y") and n[-2] not in "aeiou":
        return n[:-1] + "ies"
    if n.endswith(("s", "x", "ch", "sh")):
        return n + "es"
    return n + "s"


def table_for(schema):
    if schema in ALIASES:
        return ALIASES[schema]
    s = snake(schema)
    for c in (plural(s), s):
        if c in T:
            return c
    return None


def ref(x):
    return x["$ref"].rsplit("/", 1)[1] if isinstance(x, dict) and "$ref" in x else None


def props(schema):
    """Merged properties and required list of a schema, following allOf."""
    s = SCH[schema] if isinstance(schema, str) else schema
    out, req = {}, []
    for part in s.get("allOf", []):
        r = ref(part)
        p, rq = props(r) if r else props(part)
        out.update(p)
        req += rq
    out.update(s.get("properties", {}))
    req += s.get("required", [])
    return out, req


# Operations ------------------------------------------------------------------

OPS = []
for path, item in spec["paths"].items():
    shared = item.get("parameters", [])
    for method, o in item.items():
        if method not in ("get", "post", "put", "patch", "delete"):
            continue
        o = dict(o, method=method, path=path)
        o["params"] = shared + o.get("parameters", [])
        o["resp"], o["etag"], o["code"] = None, False, None
        for code, r in o["responses"].items():
            if code.startswith("2"):
                o["code"] = code
                s = (r.get("content", {}).get("application/json", {}) or {}).get("schema")
                if s:
                    o["resp"] = ref(s)
                o["etag"] = "ETag" in r.get("headers", {})
        body = o.get("requestBody", {}).get("content", {}).get("application/json", {}).get("schema")
        o["req"] = ref(body) if body else None
        OPS.append(o)
BY_ID = {o["operationId"]: o for o in OPS}


def go_name(op):
    return op[0].upper() + op[1:]


def resource_ops():
    g = collections.defaultdict(dict)
    for o in OPS:
        s = o["resp"]
        if not s:
            continue
        base = s[:-4] if s.endswith("List") and s[:-4] in SCH else s
        t = table_for(base)
        if not t:
            continue
        i = o["operationId"]
        kind = None
        if o["method"] == "get" and s.endswith("List"):
            kind = "list"
        elif o["method"] == "get" and i.startswith("get"):
            kind = "get"
        elif o["method"] == "post" and i.startswith("create"):
            kind = "create"
        elif o["method"] == "patch" and i.startswith("update"):
            kind = "update"
        if kind:
            g[t].setdefault(kind, o)
            g[t]["schema"] = base
    for o in OPS:
        if o["method"] == "delete":
            m = re.match(r"delete(.+)", o["operationId"])
            t = table_for(m.group(1)) if m else None
            if t:
                g[t]["delete"] = o
    return g


RES = resource_ops()

# Columns ---------------------------------------------------------------------


def enum_values(notes):
    head = notes.split("Set by the server")[0]
    return re.findall(r"`(\w+)`", head)


def fk_target(col):
    m = re.search(r"FK `(\w+)(?:\.(\w+))?`", col["notes"])
    if m:
        return m.group(1), (m.group(2) or ("code" if col["type"] == "text" else "id"))
    return None


def server_set(col):
    return "Set by the server" in col["notes"]


def is_parent(col):
    return "Parent record" in col["notes"]


def sql_columns(t):
    """(name, ddl_type, required, default, extra) for each physical column."""
    tb = T[t]
    out = []
    for c in tb["cols"]:
        n, ty, req = c["name"], c["type"], c["req"] and not server_set(c)
        if ty == "money":
            out.append((n + "_amount", "numeric(18,4)", req, None, c))
            out.append((n + "_currency", "text", req, None, c))
            continue
        ddl = {"text": "text", "uuid": "uuid", "date": "date", "timestamptz": "timestamptz",
               "decimal": "numeric(18,4)", "integer": "integer", "boolean": "boolean", "json": "jsonb",
               "enum": "text", "uuid[]": "uuid[]", "text[]": "text[]", "time": "text"}[ty]
        d = None
        if ty == "boolean":
            d, req = ("true" if n == "active" else "false"), True
        out.append((n, ddl, req, d, c))
    return out


def has_col(t, name):
    return any(c[0] == name for c in sql_columns(t))


def ddl_table(t, existing):
    tb = T[t]
    L = ["id uuid PRIMARY KEY DEFAULT uuidv7()"]
    if tb["prefix"]:
        L.append("reference text NOT NULL UNIQUE")
    for n, ddl, req, d, c in sql_columns(t):
        s = f"{q(n)} {ddl}"
        if req:
            s += " NOT NULL"
        if d:
            s += f" DEFAULT {d}"
        if c["type"] == "enum":
            vals = ", ".join(f"'{v}'" for v in enum_values(c["notes"]))
            if vals:
                s += f" CHECK ({q(n)} IN ({vals}))"
        L.append(s)
    if tb["status"]:
        vals = ", ".join(f"'{v}'" for v in tb["status"])
        L.append(f"status text NOT NULL DEFAULT '{tb['status'][0]}' CHECK (status IN ({vals}))")
    L += ["created_at timestamptz NOT NULL DEFAULT now()", "updated_at timestamptz NOT NULL DEFAULT now()",
          "created_by uuid", "updated_by uuid", "row_version integer NOT NULL DEFAULT 1"]
    if has_col(t, "code") and t not in ("currencies",):
        L.append("UNIQUE (site_id, code)" if has_col(t, "site_id") else "UNIQUE (code)")
    if t == "currencies":
        L.append("UNIQUE (code)")
    if t == "employees":
        L.append("UNIQUE (employee_number)")
    if t == "users":
        L.append("UNIQUE (employee_id)")
    if t in ("roles", "sod_rules", "rule_definitions", "form_templates", "kpi_definitions", "retention_classes"):
        L.append('UNIQUE ("key")')
    if t == "workflow_activations":
        L.append("UNIQUE (workflow_key, site_id)")
    if t == "approval_request_steps":
        L.append("assigned_user_id uuid")
    body = ",\n  ".join(L)
    return f"CREATE TABLE {t} (\n  {body}\n);\n"


def fk_stmt(t, c):
    tgt, col = fk_target(c)
    extra = " ON DELETE CASCADE" if T[t]["parent"] else ""
    return f"ALTER TABLE {t} ADD FOREIGN KEY ({q(c['name'])}) REFERENCES {tgt} ({col}){extra};"


def parent_fk(t):
    tb = T[t]
    for c in tb["cols"]:
        if is_parent(c):
            return c
    return None


def indexes(t):
    out = []
    for c in ("site_id", "status", "employee_id", "business_date"):
        if c == "status" and T[t]["status"] or (c != "status" and has_col(t, c)):
            out.append(f"CREATE INDEX ON {t} ({c});")
    p = parent_fk(t)
    if p:
        out.append(f"CREATE INDEX ON {t} ({p['name']});")
    if has_col(t, "subject_id"):
        out.append(f"CREATE INDEX ON {t} (subject_type, subject_id);")
    return out


def money_money_fk(t):
    out = []
    for c in T[t]["cols"]:
        if c["type"] == "money":
            out.append(f"ALTER TABLE {t} ADD FOREIGN KEY ({c['name']}_currency) REFERENCES currencies (code);")
    return out


def write_migration(number, name, tables, done, extra_up="", extra_down="", post_up=""):
    path = os.path.join(API, "db", "migrations", f"{number:05d}_{name}.sql")
    done_now = set(done) | set(tables)
    fks = []
    for t in done_now:
        for c in T[t]["cols"]:
            if fk_target(c) and fk_target(c)[0] in done_now and (t in tables or fk_target(c)[0] in tables):
                fks.append(fk_stmt(t, c))
        if t in tables or "currencies" in tables:
            if "currencies" in done_now:
                fks += money_money_fk(t)
    if not os.path.exists(path):
        up = "\n".join(ddl_table(t, done_now) for t in tables)
        idx = "\n".join(i for t in tables for i in indexes(t))
        down = "\n".join(f"DROP TABLE IF EXISTS {t} CASCADE;" for t in reversed(tables))
        src = f"-- +goose Up\n{extra_up}\n{up}\n{idx}\n\n" + "\n".join(sorted(set(fks))) + f"\n{post_up}\n-- +goose Down\n{extra_down}\n{down}\n"
        open(path, "w").write(src)
    return done_now


# Queries ---------------------------------------------------------------------


def api_columns(t, schema):
    """Physical columns reachable from an API request schema."""
    if not schema or schema not in SCH:
        return []
    p, _ = props(schema)
    want = []
    cols = {c[0]: c for c in sql_columns(t)}
    for name in p:
        if name in cols:
            want.append(name)
        elif name + "_amount" in cols:
            want += [name + "_amount", name + "_currency"]
    return want


def list_filters(t, o):
    fs, date_col = [], None
    cols = {c[0]: c for c in sql_columns(t)}
    for prm in o["params"]:
        name = prm.get("name")
        if "$ref" in prm or prm.get("in") != "query":
            continue
        if name == "status" and T[t]["status"]:
            fs.append(("status", "text[]", "status = ANY(sqlc.narg(status)::text[])"))
        elif name in ("from", "to"):
            continue
        elif name in cols:
            ddl = cols[name][1].split("(")[0]
            if ddl.endswith("[]"):
                continue
            fs.append((name, ddl, f"{q(name)} = sqlc.narg({name})::{ddl}"))
    if any(p.get("name") == "from" for p in o["params"]):
        for c in sql_columns(t):
            if c[1] == "date":
                date_col = c[0]
                break
        col, cast = (q(date_col), "date") if date_col else ("created_at", "timestamptz")
        guard = f"sqlc.narg(from_date)::date"
        expr = f"{col}::date"
        fs.append(("from_date", "date", f"{expr} >= {guard}"))
        fs.append(("to_date", "date", f"{expr} <= sqlc.narg(to_date)::date"))
    return fs


def search_cols(t):
    pref = ["reference", "code", "name", "title", "employee_number", "last_name", "first_name", "label", "description"]
    cols = {c[0]: c for c in sql_columns(t)}
    out = [c for c in pref if (c in cols and cols[c][1] == "text") or (c == "reference" and T[t]["prefix"])]
    return out[:4]


def site_col(t):
    if t == "sites":
        return "id"
    return "site_id" if has_col(t, "site_id") else None


def child_tables(t):
    return [x for x in ORDER if T[x]["parent"] == t]


def query_name(t):
    s = RES[t]["schema"] if t in RES and "schema" in RES[t] else None
    if s:
        return s
    return "".join(w.capitalize() for w in t.split("_"))


def arg(name, cols):
    c = cols.get(name)
    return f"sqlc.arg({name})" if c and c[2] else f"sqlc.narg({name})"


def queries_for(t):
    r, name, out = RES.get(t, {}), query_name(t), []
    cols = {c[0]: c for c in sql_columns(t)}
    if "get" in r or "update" in r or "list" in r or r:
        out.append(f"-- name: Get{name} :one\nSELECT * FROM {t} WHERE id = $1;\n")
    if "list" in r:
        where = ["(sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))",
                 "(sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))"]
        sc = search_cols(t)
        if sc:
            cond = " OR ".join(f"{q(c)} ILIKE '%' || sqlc.narg(q)::text || '%'" for c in sc)
            where.append(f"(sqlc.narg(q)::text IS NULL OR {cond})")
        for fname, ddl, expr in list_filters(t, r["list"]):
            where.append(f"(sqlc.narg({fname})::{ddl} IS NULL OR {expr})")
        sc = site_col(t)
        if sc:
            where.append(f"(sqlc.arg(all_sites)::boolean OR {sc} IS NULL OR {sc} = ANY(sqlc.arg(site_ids)::uuid[]))")
        p = parent_fk(t)
        jw = "\n  AND ".join(where)
        out.append(f"-- name: List{name} :many\nSELECT * FROM {t}\nWHERE {jw}\nORDER BY id DESC\nLIMIT sqlc.arg(row_limit);\n")
    if "create" in r:
        want = set(api_columns(t, r["create"].get("req")))
        for c in sql_columns(t):
            if (c[2] and not c[3]) or c[0] == "site_id":
                want.add(c[0])
        ordered = [c[0] for c in sql_columns(t) if c[0] in want]
        names = list(ordered)
        vals = []
        for n in ordered:
            c = cols[n]
            vals.append(f"COALESCE(sqlc.narg({n})::boolean, {c[3]})" if c[3] else arg(n, cols))
        if T[t]["prefix"]:
            names.insert(0, "reference")
            vals.insert(0, "sqlc.arg(reference)")
        names.append("created_by"), vals.append("sqlc.narg(created_by)")
        names.append("updated_by"), vals.append("sqlc.narg(created_by)")
        out.append(f"-- name: Create{name} :one\nINSERT INTO {t} ({', '.join(q(n) for n in names)})\nVALUES ({', '.join(vals)})\nRETURNING *;\n")
    if "update" in r:
        want = [c for c in api_columns(t, r["update"].get("req")) if c not in ("status",)]
        sets = [f"{q(c)} = COALESCE(sqlc.narg({c}), {q(c)})" for c in want]
        sets += ["updated_at = now()", "updated_by = sqlc.narg(updated_by)", "row_version = row_version + 1"]
        out.append(f"-- name: Update{name} :one\nUPDATE {t}\nSET {', '.join(sets)}\nWHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)\nRETURNING *;\n")
    if T[t]["status"]:
        out.append(f"-- name: Set{name}Status :one\nUPDATE {t}\nSET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1\nWHERE id = sqlc.arg(id)\nRETURNING *;\n")
    if "delete" in r:
        out.append(f"-- name: Delete{name} :exec\nDELETE FROM {t} WHERE id = $1;\n")
    for ch in child_tables(t):
        out += child_queries(ch)
    out += action_queries(t)
    return "\n".join(out)


def child_name(ch):
    return "".join(w.capitalize() for w in ch.split("_"))


def child_queries(ch):
    p = parent_fk(ch)["name"]
    n = child_name(ch)
    cols = {c[0]: c for c in sql_columns(ch)}
    names = list(cols) + ["created_by"]
    vals = []
    for c in cols.values():
        vals.append(f"COALESCE(sqlc.narg({c[0]})::boolean, {c[3]})" if c[3] else arg(c[0], cols))
    vals.append("sqlc.narg(created_by)")
    return [f"-- name: List{n} :many\nSELECT * FROM {ch} WHERE {p} = $1 ORDER BY id;\n",
            f"-- name: Insert{n} :one\nINSERT INTO {ch} ({', '.join(q(x) for x in names)})\nVALUES ({', '.join(vals)})\nRETURNING *;\n",
            f"-- name: Delete{n} :exec\nDELETE FROM {ch} WHERE {p} = $1;\n"]


# Actions ---------------------------------------------------------------------

PAST = {"submit": "submitted", "approve": "approved", "close": "closed", "complete": "completed", "verify": "verified",
        "review": "reviewed", "acknowledge": "acknowledged", "issue": "issued", "return": "returned", "revoke": "revoked",
        "release": "released", "cancel": "cancelled", "suspend": "suspended", "reject": "rejected", "resolve": "resolved",
        "receive": "received", "confirm": "confirmed", "execute": "executed", "decide": "decided", "restore": "restored",
        "seal": "sealed", "dispatch": "dispatched", "accept": "accepted", "post": "posted", "publish": "published",
        "activate": "activated", "reimburse": "reimbursed", "retire": "retired", "lock": "locked", "disable": "disabled"}


def actions_of(t):
    own = RES.get(t, {}).get("get", {}).get("path")
    if not own:
        return []
    out = []
    for o in OPS:
        if o["method"] == "post" and o["path"].startswith(own + "/") and o["path"].count("/") == own.count("/") + 1 \
                and not o["path"].endswith("}") and not o["operationId"].startswith("create"):
            out.append(o)
    return out


def action_columns(t, o):
    """Columns a body sets directly, and server stamps for this verb."""
    cols = {c[0]: c for c in sql_columns(t)}
    body = [c for c in api_columns(t, o.get("req")) if c != "status"] if o.get("req") else []
    verb = o["path"].rsplit("/", 1)[1]
    past = PAST.get(verb, verb.replace("-", "_") + "d")
    stamps = {}
    for suffix, expr in ((f"{past}_at", "now()"), (f"{past}_by", "sqlc.narg(updated_by)"),
                         (f"{past}_by_user_id", "sqlc.narg(updated_by)")):
        if suffix in cols and suffix not in body:
            stamps[suffix] = expr
    return body, stamps


def action_queries(t):
    out = []
    for o in actions_of(t):
        verb = o["path"].rsplit("/", 1)[1]
        source, target = workflows.infer(o["operationId"], verb, T[t]["status"], t)
        body, stamps = action_columns(t, o)
        sets = [f"{q(c)} = COALESCE(sqlc.narg({c}), {q(c)})" for c in body]
        sets += [f"{q(c)} = {e}" for c, e in stamps.items()]
        if target:
            sets.append(f"status = '{target}'")
        sets += ["updated_at = now()", "updated_by = sqlc.narg(updated_by)", "row_version = row_version + 1"]
        where = "id = sqlc.arg(id)"
        if source:
            where += " AND status IN (" + ", ".join(f"'{x}'" for x in source) + ")"
        out.append(f"-- name: Act{go_name(o['operationId'])} :one\nUPDATE {t}\nSET {', '.join(sets)}\nWHERE {where}\nRETURNING *;\n")
    return out


def service_actions(t, qr, hw):
    B = []
    n = query_name(t)
    sc = site_col(t)
    rt = row_type(qr, f"Get{n}")
    for o in actions_of(t):
        G = go_name(o["operationId"])
        if G in hw:
            continue
        verb = o["path"].rsplit("/", 1)[1]
        source, target = workflows.infer(o["operationId"], verb, T[t]["status"], t)
        perm = o.get("x-permission") or o["operationId"]
        sptr = "nil"
        if sc == "id":
            sptr = "&row.ID"
        elif sc:
            sptr = "&row.SiteID" if col_required(t, "site_id") else "row.SiteID"
        B.append(f"// {G} is the {o.get('summary', verb).lower()} action on {t.replace('_', ' ')}.")
        B.append(f"func (s *Service) {G}(ctx context.Context, p db.Act{G}Params) (db.{rt}, error) {{")
        B.append(f"\tbefore, err := s.Get{n}(ctx, p.ID)")
        B.append("\tif err != nil {\n\t\treturn before, err\n\t}")
        B.append(f'\tif err := s.checkSod(ctx, "{perm}", "{t}", p.ID); err != nil {{\n\t\treturn before, err\n\t}}')
        B.append(f'\tif err := s.guard(ctx, "{G}", before); err != nil {{\n\t\treturn before, err\n\t}}')
        B.append("\tp.UpdatedBy = actorID(ctx)")
        if verb in ("close", "complete", "approve", "verify", "accept", "post", "issue", "review"):
            B.append(f'\tif err := s.requireSignatures(ctx, "{t}", p.ID); err != nil {{\n\t\treturn before, err\n\t}}')
        B.append(f"\trow, err := s.q(ctx).Act{G}(ctx, p)")
        B.append(f'\tif err != nil {{\n\t\treturn before, stateError(err, "{verb.replace("-", " ")} {t.replace("_", " ")}")\n\t}}')
        B.append(f'\tif err := s.audit(ctx, "{perm}", "{t}", row.ID, {sptr}, before, row); err != nil {{\n\t\treturn row, err\n\t}}')
        if target == "pending_approval":
            B.append(f'\tif err := s.startApproval(ctx, "{t}", row.ID, {sptr}, row); err != nil {{\n\t\treturn row, err\n\t}}')
        B.append(f'\treturn row, s.effect(ctx, "{G}", row)')
        B.append("}\n")
    return B


def http_actions(t, hw):
    B = []
    n = query_name(t)
    schema = RES[t]["schema"]
    for o in actions_of(t):
        G = go_name(o["operationId"])
        if G in hw:
            continue
        ids = path_ids(o)
        own = f"req.{go_param(ids[-1])}"
        B.append(f"// {G} implements the {o['operationId']} operation.")
        B.append(f"func (s *Server) {G}(ctx context.Context, req apigen.{G}RequestObject) (apigen.{G}ResponseObject, error) {{")
        body = "req.Body" if o.get("req") else "struct{}{}"
        B.append(f'\tin, err := input[db.Act{G}Params]({body}, map[string]any{{"id": {own}}})')
        B.append("\tif err != nil {\n\t\treturn nil, err\n\t}")
        if o.get("req") and "pin" in props(o["req"])[0]:
            B.append("\tif err := s.service.RequirePin(ctx, str(req.Body.Pin)); err != nil {\n\t\treturn nil, err\n\t}")
        if o.get("req") and "signature_id" in props(o["req"])[0]:
            sig = "req.Body.SignatureId"
            required = "signature_id" in props(o["req"])[1]
            B.append(f'\tif err := s.service.RequireSigned(ctx, "{t}", {own}, {"&" if required else ""}{sig}); err != nil {{\n\t\treturn nil, err\n\t}}')
        B.append(f"\trow, err := s.service.{G}(ctx, in)")
        B.append("\tif err != nil {\n\t\treturn nil, err\n\t}")
        if o["resp"] == schema:
            B.append(f"\tbody, err := output[apigen.{schema}](row)")
            B.append("\tif err != nil {\n\t\treturn nil, err\n\t}")
            c = o["code"]
            if o["etag"]:
                B.append(f"\treturn apigen.{G}{c}JSONResponse{{Body: body, Headers: apigen.{G}{c}ResponseHeaders{{ETag: etag(row.RowVersion)}}}}, nil")
            else:
                B.append(f"\treturn apigen.{G}{c}JSONResponse(body), nil")
        else:
            B.append("\t_ = row")
            B.append("\treturn nil, apigen.ErrNotImplemented")
        B.append("}\n")
    return B


# Go --------------------------------------------------------------------------


def handwritten(pkg):
    names = set()
    for f in glob.glob(os.path.join(API, "internal", pkg, "*.go")):
        b = os.path.basename(f)
        if b.startswith("crud_") or b.endswith("_test.go"):
            continue
        names |= set(re.findall(r"^func \(\w+ \*\w+\) (\w+)\(", open(f).read(), re.M))
    return names


def querier():
    src = open(os.path.join(API, "internal", "db", "querier.go")).read()
    return {m.group(1): m.group(2) for m in re.finditer(r"^\s+(\w+)\(ctx context\.Context.*?\) \(?(.*?)\)?$", src, re.M)}


def row_type(qr, name):
    return qr[name].split(",")[0].strip().lstrip("[]")


def attr(col):
    parts = col.split("_")
    return "".join("ID" if w == "id" else w.capitalize() for w in parts)


def children_in_api(t, schema):
    out = []
    p, _ = props(schema)
    for ch in child_tables(t):
        for name, v in p.items():
            items = v.get("items", {}) if isinstance(v, dict) else {}
            if v.get("type") == "array" and ref(items) and table_for(ref(items)) == ch:
                out.append((name, ch))
                break
            if v.get("type") == "array" and ref(items) and (name == ch or ch.endswith("_" + name)):
                out.append((name, ch))
                break
    return out


def service_go(section, tables, qr, hw):
    B = ["// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in", "// another file of this package and this copy disappears on the next run.", "", "package service", "", "import (", '\t"context"', "", '\t"github.com/google/uuid"', "", '\t"foodies/api/internal/db"', ")", "", "var _ uuid.UUID", ""]
    used = False
    for t in tables:
        r = RES.get(t)
        if not r:
            continue
        n, sc, prefix = query_name(t), site_col(t), T[t]["prefix"]
        rt = row_type(qr, f"Get{n}")
        human = t.replace("_", " ")
        if "list" in r and f"List{n}" not in hw:
            B.append(f"// List{n} returns a page of {human}.")
            B.append(f"func (s *Service) List{n}(ctx context.Context, p db.List{n}Params) ([]db.{rt}, error) {{")
            if sc:
                B.append("\tp.AllSites, p.SiteIds = scopeOf(ctx)")
            B.append(f'\treturn wrap2(s.q(ctx).List{n}(ctx, p))("list {human}")')
            B.append("}\n")
            used = True
        if f"Get{n}" not in hw:
            B.append(f"// Get{n} returns one row of {human}.")
            B.append(f"func (s *Service) Get{n}(ctx context.Context, id uuid.UUID) (db.{rt}, error) {{")
            B.append(f"\trow, err := s.q(ctx).Get{n}(ctx, id)")
            B.append('\tif err != nil {')
            B.append(f'\t\treturn row, notFound(err, "{human}")')
            B.append("\t}")
            if sc:
                expr = "row.ID" if sc == "id" else ("row.SiteID")
                ptr = (sc == "id") or not col_required(t, "site_id")
                B.append(f"\treturn row, inScope(ctx, {'&row.ID' if sc == 'id' else ('row.SiteID' if ptr else '&row.SiteID')})")
            else:
                B.append("\treturn row, nil")
            B.append("}\n")
        if "create" in r and f"Create{n}" not in hw:
            B.append(f"// Create{n} adds one row to {human}.")
            B.append(f"func (s *Service) Create{n}(ctx context.Context, p db.Create{n}Params) (db.{rt}, error) {{")
            B.append("\tvar zero db." + rt)
            sptr = None
            if sc == "site_id":
                sptr = "&p.SiteID" if col_required(t, "site_id") else "p.SiteID"
                B.append(f"\tif err := canWrite(ctx, {sptr}); err != nil {{\n\t\treturn zero, err\n\t}}")
            B.append(f'\tif err := s.guard(ctx, "Create{n}", &p); err != nil {{\n\t\treturn zero, err\n\t}}')
            if prefix:
                B.append(f'\tref, err := s.nextReference(ctx, "{prefix}", {sptr or "nil"})')
                B.append("\tif err != nil {\n\t\treturn zero, err\n\t}\n\tp.Reference = ref")
            B.append("\tp.CreatedBy = actorID(ctx)")
            B.append(f"\trow, err := s.q(ctx).Create{n}(ctx, p)")
            B.append(f'\tif err != nil {{\n\t\treturn zero, wrap(err, "create {human}")\n\t}}')
            B.append(f'\tif err := s.audit(ctx, "{r["create"].get("x-permission") or t + ".create"}", "{t}", row.ID, {sptr or "nil"}, nil, row); err != nil {{\n\t\treturn row, err\n\t}}')
            B.append(f'\treturn row, s.effect(ctx, "Create{n}", row)')
            B.append("}\n")
        if "update" in r and f"Update{n}" not in hw:
            B.append(f"// Update{n} changes one row of {human}.")
            B.append(f"func (s *Service) Update{n}(ctx context.Context, p db.Update{n}Params) (db.{rt}, error) {{")
            B.append(f"\tbefore, err := s.Get{n}(ctx, p.ID)")
            B.append("\tif err != nil {\n\t\treturn before, err\n\t}")
            B.append("\tif err := checkVersion(before.RowVersion, p.RowVersion); err != nil {\n\t\treturn before, err\n\t}")
            B.append("\tp.UpdatedBy = actorID(ctx)")
            B.append(f"\trow, err := s.q(ctx).Update{n}(ctx, p)")
            B.append(f'\tif err != nil {{\n\t\treturn before, wrap(err, "update {human}")\n\t}}')
            B.append(f'\tif err := s.voidSignatures(ctx, "{t}", row.ID); err != nil {{\n\t\treturn row, err\n\t}}')
            sptr = "nil"
            if sc == "id":
                sptr = "&row.ID"
            elif sc:
                sptr = "&row.SiteID" if col_required(t, "site_id") else "row.SiteID"
            B.append(f'\treturn row, s.audit(ctx, "{r["update"].get("x-permission") or t + ".update"}", "{t}", row.ID, {sptr}, before, row)')
            B.append("}\n")
        if "delete" in r and f"Delete{n}" not in hw:
            B.append(f"// Delete{n} removes one row of {human}.")
            B.append(f"func (s *Service) Delete{n}(ctx context.Context, id uuid.UUID) error {{")
            B.append(f"\tbefore, err := s.Get{n}(ctx, id)")
            B.append("\tif err != nil {\n\t\treturn err\n\t}")
            B.append(f'\tif err := s.q(ctx).Delete{n}(ctx, id); err != nil {{\n\t\treturn wrap(err, "delete {human}")\n\t}}')
            B.append(f'\treturn s.audit(ctx, "{t}.delete", "{t}", id, nil, before, nil)')
            B.append("}\n")
        B += service_actions(t, qr, hw)
        kids = children_in_api(t, r["schema"])
        if kids and f"{n}Children" not in hw:
            B.append(f"// {n}Children returns the rows of a record's child tables, by API field.")
            B.append(f"func (s *Service) {n}Children(ctx context.Context, id uuid.UUID) (map[string]any, error) {{")
            B.append("\tout := map[string]any{}")
            for name, ch in kids:
                B.append(f"\t{{\n\t\trows, err := s.q(ctx).List{child_name(ch)}(ctx, id)\n\t\tif err != nil {{\n\t\t\treturn nil, wrap(err, \"list {ch}\")\n\t\t}}\n\t\tout[\"{name}\"] = childRows(rows)\n\t}}")
            B.append("\treturn out, nil\n}\n")
            for name, ch in kids:
                c = child_name(ch)
                if f"Replace{c}" in hw:
                    continue
                B.append(f"// Replace{c} swaps a record's {ch.replace('_', ' ')} for the given rows.")
                B.append(f"func (s *Service) Replace{c}(ctx context.Context, id uuid.UUID, rows []db.Insert{c}Params) error {{")
                B.append(f"\tif err := s.q(ctx).Delete{c}(ctx, id); err != nil {{\n\t\treturn wrap(err, \"clear {ch}\")\n\t}}")
                B.append("\tfor _, row := range rows {\n\t\trow.CreatedBy = actorID(ctx)")
                B.append(f"\t\tif _, err := s.q(ctx).Insert{c}(ctx, row); err != nil {{\n\t\t\treturn wrap(err, \"add {ch}\")\n\t\t}}\n\t}}\n\treturn nil\n}}\n")
    return "\n".join(B)


def write_registry(tables):
    qr = querier()
    B = ["// Code scaffolded by api/codegen/scaffold. Do not edit.", "", "package service", "", "import (", '\t"context"', "", '\t"github.com/google/uuid"', "", '\t"foodies/api/internal/db"', ")", "",
         "// subjects lets shared services reach any record by table name.", "var subjects = map[string]subject{"]
    for t in tables:
        n = query_name(t)
        if f"Get{n}" not in qr:
            continue
        B.append(f'\t"{t}": {{')
        B.append(f"\t\tget: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {{ return q.Get{n}(ctx, id) }},")
        if T[t]["status"]:
            vals = ", ".join(f'"{x}"' for x in T[t]["status"])
            B.append(f"\t\tstatuses: []string{{{vals}}},")
            B.append(f"\t\tsetStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {{")
            B.append(f"\t\t\t_, err := q.Set{n}Status(ctx, db.Set{n}StatusParams{{ID: id, Status: status, UpdatedBy: actor}})")
            B.append("\t\t\treturn err\n\t\t},")
        B.append("\t},")
    B.append("}\n")
    open(os.path.join(API, "internal", "service", "crud_registry.go"), "w").write("\n".join(B))


def write_search(tables):
    parts = []
    for t in tables:
        if not T[t]["prefix"] or "list" not in RES.get(t, {}):
            continue
        title = next((c for c in ("title", "name", "description", "subject") if has_col(t, c)), "reference")
        site = "site_id" if has_col(t, "site_id") else "NULL::uuid"
        status = "status" if T[t]["status"] else "NULL::text"
        parts.append(f"SELECT '{t}'::text AS type, id, reference, {q(title)}::text AS title, {site} AS site_id, {status} AS status FROM {t} WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR {q(title)}::text ILIKE '%' || sqlc.arg(term)::text || '%'")
    sql = "-- name: SearchRecords :many\nSELECT * FROM (\n" + "\nUNION ALL\n".join(parts) + "\n) hits\nWHERE sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[])\nORDER BY id DESC\nLIMIT sqlc.arg(row_limit);\n"
    open(os.path.join(API, "db", "queries", "crud_search.sql"), "w").write(sql)


def col_required(t, name):
    return any(c[0] == name and c[2] for c in sql_columns(t))


def path_ids(o):
    """Path parameters besides the record's own id, as (param, column-ish)."""
    return re.findall(r"\{(\w+)\}", o["path"])


def go_param(p):
    return "".join("Id" if w == "id" else w.capitalize() for w in p.split("_"))


def http_go(section, tables, qr, hw):
    B = ["// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in", "// another file of this package and this copy disappears on the next run.", "", "package httpapi", "", "import (", '\t"context"', "", '\t"foodies/api/internal/db"', '\t"foodies/api/internal/http/apigen"', ")", "", "var _ db.Querier", ""]
    for t in tables:
        r = RES.get(t)
        if not r:
            continue
        n, schema = query_name(t), r["schema"]
        kids = children_in_api(t, schema)
        for kind in ("list", "get", "create", "update", "delete"):
            o = r.get(kind)
            if not o:
                continue
            G = go_name(o["operationId"])
            if G in hw:
                continue
            ids = path_ids(o)
            cols = {c[0] for c in sql_columns(t)}
            extra = [(p, f"req.{go_param(p)}") for p in ids if p in cols]
            sig = f"func (s *Server) {G}(ctx context.Context, req apigen.{G}RequestObject) (apigen.{G}ResponseObject, error) {{"
            ex = ("map[string]any{" + ", ".join(f'"{a}": {b}' for a, b in extra) + "}") if extra else "nil"
            own = f"req.{go_param(ids[-1])}" if ids else None
            B.append(f"// {G} implements the {o['operationId']} operation.")
            B.append(sig)
            if kind == "list":
                params = "req.Params" if any(True for _ in o["params"]) else "struct{}{}"
                m = {"from": "from_date", "to": "to_date"}
                B.append(f"\treturn listPage[apigen.{G}200JSONResponse](ctx, {params}, {ex}, s.service.List{n})")
            elif kind == "get":
                B.append(f"\trow, err := s.service.Get{n}(ctx, {own})")
                B.append("\tif err != nil {\n\t\treturn nil, err\n\t}")
                kv = ""
                if kids:
                    B.append(f"\tkids, err := s.service.{n}Children(ctx, row.ID)")
                    B.append("\tif err != nil {\n\t\treturn nil, err\n\t}")
                    kv = ", kids"
                B.append(f"\tbody, err := output[apigen.{schema}](row{kv})")
                B.append("\tif err != nil {\n\t\treturn nil, err\n\t}")
                if o["etag"]:
                    B.append(f"\treturn apigen.{G}200JSONResponse{{Body: body, Headers: apigen.{G}200ResponseHeaders{{ETag: etag(row.RowVersion)}}}}, nil")
                else:
                    B.append(f"\treturn apigen.{G}200JSONResponse(body), nil")
            elif kind in ("create", "update"):
                ptype = f"db.{'Create' if kind == 'create' else 'Update'}{n}Params"
                B.append(f"\tin, err := input[{ptype}](req.Body, {ex})")
                B.append("\tif err != nil {\n\t\treturn nil, err\n\t}")
                if kind == "update":
                    B.append(f"\tin.ID = {own}")
                    B.append("\tif in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {\n\t\treturn nil, err\n\t}")
                B.append(f"\trow, err := s.service.{'Create' if kind == 'create' else 'Update'}{n}(ctx, in)")
                B.append("\tif err != nil {\n\t\treturn nil, err\n\t}")
                kv = ""
                if kids:
                    B.append("\traw, err := toMap(req.Body)\n\tif err != nil {\n\t\treturn nil, err\n\t}")
                    pk = None
                    for name, ch in kids:
                        pk = parent_fk(ch)["name"]
                        B.append(f'\tif items, ok := raw["{name}"]; ok {{')
                        B.append(f'\t\trows, err := childInputs[db.Insert{child_name(ch)}Params](items, "{pk}", row.ID)')
                        B.append("\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}")
                        B.append(f"\t\tif err := s.service.Replace{child_name(ch)}(ctx, row.ID, rows); err != nil {{\n\t\t\treturn nil, err\n\t\t}}")
                        B.append("\t}")
                    B.append(f"\tkids, err := s.service.{n}Children(ctx, row.ID)")
                    B.append("\tif err != nil {\n\t\treturn nil, err\n\t}")
                    kv = ", kids"
                B.append(f"\tbody, err := output[apigen.{schema}](row{kv})")
                B.append("\tif err != nil {\n\t\treturn nil, err\n\t}")
                c = o["code"]
                if o["etag"]:
                    B.append(f"\treturn apigen.{G}{c}JSONResponse{{Body: body, Headers: apigen.{G}{c}ResponseHeaders{{ETag: etag(row.RowVersion)}}}}, nil")
                else:
                    B.append(f"\treturn apigen.{G}{c}JSONResponse(body), nil")
            elif kind == "delete":
                B.append(f"\tif err := s.service.Delete{n}(ctx, {own}); err != nil {{\n\t\treturn nil, err\n\t}}")
                B.append(f"\treturn apigen.{G}204Response{{}}, nil")
            B.append("}\n")
        B += http_actions(t, hw)
    return "\n".join(B)


def sqlc_generate():
    subprocess.check_call("go tool -modfile=tools/go.mod sqlc generate", shell=True, cwd=API)


def scaffold_section(key, tables):
    qdir = os.path.join(API, "db", "queries")
    for t in tables:
        s = queries_for(t) if not T[t]["parent"] else ""
        if s:
            open(os.path.join(qdir, f"crud_{t}.sql"), "w").write(s)
    sqlc_generate()
    qr = querier()
    svc = service_go(key, [t for t in tables if not T[t]["parent"]], qr, handwritten("service"))
    open(os.path.join(API, "internal", "service", f"crud_{key}.go"), "w").write(svc)
    h = http_go(key, [t for t in tables if not T[t]["parent"]], qr, handwritten("http"))
    open(os.path.join(API, "internal", "http", f"crud_{key}.go"), "w").write(h)
    subprocess.call("gofmt -w internal/service internal/http/crud_*.go", shell=True, cwd=API)
