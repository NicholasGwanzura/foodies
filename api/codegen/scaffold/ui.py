"""Writes web/lib/resources.gen.ts: what the generic screens need to know."""
import json, os, re
import scaffold as sc

ACRONYMS = {"fx": "FX", "pos": "POS", "cit": "CIT", "sap": "SAP", "ict": "ICT", "qa": "QA", "ncr": "NCR", "id": "ID", "mfa": "MFA",
            "kpi": "KPI", "grv": "GRV", "capa": "CAPA", "sop": "SOP", "hr": "HR", "vat": "VAT", "po": "PO", "pr": "PR", "npd": "NPD",
            "cctv": "CCTV", "url": "URL", "sla": "SLA", "ppe": "PPE", "wht": "WHT", "zimra": "ZIMRA", "ip": "IP", "gl": "GL", "rbz": "RBZ"}

TILES = {
    "Cash": "cash-finance", "Procurement & AP": "cash-finance", "Budgets": "cash-finance",
    "Assets": "equipment-assets", "Maintenance": "equipment-assets",
    "Management Reporting": "general-controls", "Compliance Calendar": "general-controls", "Forms": "general-controls",
    "Stock": "inventory-stores", "Central Kitchen & Production": "production-kitchen",
    "Quality": "quality-compliance", "Incidents": "quality-compliance", "External Visits": "quality-compliance",
    "Products & NPD": "quality-compliance", "Customer Service": "sales-customer", "Payment Exceptions": "sales-customer",
    "Lost Property": "sales-customer", "Store Operations": "store-operations", "Keys": "store-operations", "CCTV": "store-operations",
    "HR": "human-resources", "People & Access": "human-resources", "Training": "training",
    "Security": "security", "Visitors & Contractors": "security", "Documents": "sops", "Organisation": "complexes",
}
ADMIN_TAGS = {"Rules & Readiness", "Approvals", "Escalations & Actions", "Integrations", "WhatsApp", "Retention", "Audit",
              "Tasks & Schedules", "Notifications", "Reporting", "Sync", "Signatures", "Evidence", "Auth", "System"}


def words(name):
    out = []
    for i, w in enumerate(name.split("_")):
        w = ACRONYMS.get(w, w)
        out.append(w.capitalize() if i == 0 and w.islower() else w)
    return " ".join(out)


def label(name):
    n = re.sub(r"_employee_id$|_user_id$|_id$", "", name)
    n = re.sub(r"_employee_ids$|_ids$", "s", n)
    n = re.sub(r"_at$", "", n) if name.endswith("_at") else n
    return words(n or name)


def singular(t):
    w = t.split("_")
    last = w[-1]
    if last.endswith("ies"):
        last = last[:-3] + "y"
    elif last.endswith(("xes", "ches", "shes", "sses")):
        last = last[:-2]
    elif last.endswith("s"):
        last = last[:-1]
    return words("_".join(w[:-1] + [last]))


def fk_of(t, name):
    for c in sc.T[t]["cols"]:
        if c["name"] == name and sc.fk_target(c):
            return sc.fk_target(c)[0]
    return None


def field(t, name, p, required, built, depth=0):
    p = dict(p)
    r = sc.ref(p)
    for part in p.get("allOf", []) or []:
        r = r or sc.ref(part)
    f = {"name": name, "label": label(name)}
    if required:
        f["required"] = True
    if p.get("readOnly"):
        f["readOnly"] = True
    ty = p.get("type")
    if isinstance(ty, list):
        ty = next((x for x in ty if x != "null"), "string") if len(ty) <= 2 else "json"
    fmt = p.get("format")
    tgt = fk_of(t, name) if t else None
    if r == "Money":
        f["kind"] = "money"
    elif p.get("enum"):
        f["kind"], f["options"] = "enum", p["enum"]
    elif tgt and tgt in built and tgt != "currencies" and "list" in sc.RES.get(tgt, {}) and "{" not in sc.RES[tgt]["list"]["path"]:
        f["kind"], f["ref"] = "ref", tgt.replace("_", "-")
    elif name.endswith("currency") or tgt == "currencies":
        f["kind"] = "currency"
    elif fmt == "date":
        f["kind"] = "date"
    elif fmt == "date-time":
        f["kind"] = "datetime"
    elif fmt == "decimal":
        f["kind"] = "decimal"
    elif fmt == "uuid":
        f["kind"] = "uuid"
    elif ty == "integer" or ty == "number":
        f["kind"] = "number"
    elif ty == "boolean":
        f["kind"] = "boolean"
    elif ty == "array":
        items = p.get("items", {})
        ir = sc.ref(items)
        if ir and ir in sc.SCH and depth == 0:
            sub, rq = sc.props(ir)
            f["kind"] = "rows"
            ct = sc.table_for(ir)
            f["fields"] = [field(ct, k, v, k in rq, built, 1) for k, v in sub.items()
                           if k not in sc.AUDIT_COLS and k != "links" and not v.get("readOnly")
                           and not (ct and sc.parent_fk(ct) and k == sc.parent_fk(ct)["name"])]
        elif items.get("format") == "uuid":
            f["kind"] = "list"
        else:
            f["kind"] = "list"
    elif ty == "object" or ty is None:
        f["kind"] = "json"
    elif name in ("description", "reason", "notes", "narrative", "comment", "details", "body", "summary", "findings"):
        f["kind"] = "textarea"
    else:
        f["kind"] = "text"
    if p.get("description"):
        f["hint"] = str(p["description"]).split("\n")[0][:110]
    if "pin" == name or name.endswith("_pin"):
        f["kind"] = "pin"
    if name == "signature_id":
        # The screen collects a PIN, and the server action turns it into a signature.
        f.update(kind="pin", label="Your 6-digit PIN")
        f.pop("hint", None)
    if p.get("format") == "password":
        f["kind"] = "password"
    return f


def fields_of(t, schema, built):
    if not schema or schema not in sc.SCH:
        return []
    p, req = sc.props(schema)
    has_ref = bool(t and sc.T[t]["prefix"])
    return [field(t, k, v, k in req, built) for k, v in p.items()
            if k not in ("id", "created_at", "updated_at", "created_by", "updated_by", "row_version", "links")
            and (k != "reference" or has_ref)]


def columns(fs):
    pick = []
    names = [f["name"] for f in fs]
    for n in ("reference", "code", "employee_number", "name", "title", "first_name", "last_name", "label", "key"):
        if n in names and len(pick) < 3:
            pick.append(n)
    rank = {"enum": 1, "ref": 2, "date": 3, "money": 3, "datetime": 5, "text": 6, "decimal": 6, "number": 6, "boolean": 7}
    rest = sorted([f for f in fs if f["name"] not in pick and f["kind"] in rank and f["name"] != "status" and not f.get("hint")],
                  key=lambda f: rank[f["kind"]])
    if "status" in names:
        pick.append("status")
    for f in rest:
        if len(pick) >= 5:
            break
        pick.append(f["name"])
    return pick


def action_key(path):
    return path.rsplit("/", 1)[1]


def emit(built, path):
    out = {}
    for t in built:
        r = sc.RES.get(t)
        if not r or "list" not in r:
            continue
        schema = r["schema"]
        lst = r["list"]
        fs = fields_of(t, schema, built)
        tag = (lst.get("tags") or ["?"])[0]
        item = {
            "slug": t.replace("_", "-"), "title": words(t), "singular": singular(t),
            "about": (sc.SCH[schema].get("description", "") or "").split("\n")[0],
            "group": tag, "tile": TILES.get(tag), "listPath": lst["path"], "readPermission": lst.get("x-permission", ""),
            "fields": fs, "columns": columns(fs), "restricted": bool(sc.T[t]["restricted"]),
        }
        if "get" in r:
            item["getPath"] = r["get"]["path"]
        if "create" in r:
            o = r["create"]
            item["create"] = {"path": o["path"], "permission": o.get("x-permission", ""),
                              "fields": [f for f in fields_of(t, o.get("req"), built) if not f.get("readOnly")]}
        if "update" in r:
            o = r["update"]
            item["update"] = {"path": o["path"], "permission": o.get("x-permission", ""),
                              "fields": [f for f in fields_of(t, o.get("req"), built) if not f.get("readOnly")]}
        status = next((f for f in fs if f["name"] == "status"), None)
        if status and status.get("options"):
            item["statuses"] = status["options"]
        actions = []
        own = r.get("get", {}).get("path")
        for o in sc.OPS:
            if o["method"] == "post" and own and o["path"].startswith(own + "/") and o["path"].count("/") == own.count("/") + 1 \
                    and not o["path"].endswith("}") and not o["operationId"].startswith("create"):
                source, _ = sc.workflows.infer(o["operationId"], action_key(o["path"]), sc.T[t]["status"], t)
                actions.append({"key": action_key(o["path"]), "label": o.get("summary", ""), "path": o["path"], "permission": o.get("x-permission", ""),
                                "operationId": o["operationId"], **({"from": source} if source else {}),
                                "fields": [f for f in fields_of(None, o.get("req"), built) if not f.get("readOnly")]})
        if actions:
            item["actions"] = actions
        if "{" in lst["path"]:
            parent = lst["path"].split("/{")[0]
            item["nestedUnder"] = next((x.replace("_", "-") for x in built if "list" in sc.RES.get(x, {}) and sc.RES[x]["list"]["path"] == parent), None)
        out[item["slug"]] = item
    src = ("// Generated by api/codegen/scaffold from openapi.yaml and data-model.md. Do not edit.\n\n"
           'import type { Resource } from "./resource-types";\n\n'
           "export const RESOURCES: Record<string, Resource> = "
           + json.dumps(out, indent=1, ensure_ascii=False) + ";\n")
    open(path, "w").write(src)
