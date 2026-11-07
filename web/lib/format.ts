import type { FieldMeta } from "./resource-types.ts";

const date = new Intl.DateTimeFormat("en-ZW", { dateStyle: "medium", timeZone: "Africa/Harare" });
const dateTime = new Intl.DateTimeFormat("en-ZW", { dateStyle: "medium", timeStyle: "short", timeZone: "Africa/Harare" });

/** `pending_approval` becomes `Pending approval`. */
export function humanise(value: string): string {
  const text = value.replaceAll("_", " ");
  return text.charAt(0).toUpperCase() + text.slice(1);
}

export function money(value: { amount: string; currency: string }): string {
  const n = Number(value.amount);
  const amount = Number.isNaN(n)
    ? value.amount
    : n.toLocaleString("en-ZW", { minimumFractionDigits: 2, maximumFractionDigits: 2 });
  const symbol = value.currency === "USD" ? "US$" : value.currency === "ZWG" ? "ZiG " : `${value.currency} `;
  return `${symbol}${amount}`;
}

export function formatValue(field: Pick<FieldMeta, "kind">, value: unknown): string {
  if (value === null || value === undefined || value === "") return "";
  switch (field.kind) {
    case "money":
      return money(value as { amount: string; currency: string });
    case "date":
      return date.format(new Date(`${String(value)}T12:00:00Z`));
    case "datetime":
      return dateTime.format(new Date(String(value)));
    case "boolean":
      return value ? "Yes" : "No";
    case "enum":
      return humanise(String(value));
    case "list":
      return (value as unknown[]).join(", ");
    case "json":
    case "rows":
      return JSON.stringify(value);
    default:
      return String(value);
  }
}

const TONES: [RegExp, string][] = [
  [/^(approved|active|closed|completed|complete|posted|verified|reconciled|live|published|available|paid|passed|pass|released|effective|acknowledged|signed|received|authenticated|ok|ready|resolved|done)$/, "good"],
  [/^(rejected|failed|fail|exception|overdue|cancelled|suspended|breached|expired|revoked|disabled|locked|voided|declined|missed|not_ready|critical|defective|escalated|blocked)$/, "bad"],
  [/^(pending|draft|invited|returned|open|in_progress|on_hold|held|submitted|requested|claimed|paused|unset|manual_required|proposed)/, "warn"],
];

export function tone(status: string): string {
  return TONES.find(([re]) => re.test(status))?.[1] ?? "info";
}

/** The best short name for a record. */
export function recordLabel(row: Record<string, unknown>): string {
  if (row.first_name) return `${row.first_name} ${row.last_name ?? ""}`.trim();
  for (const key of ["reference", "name", "title", "label", "code", "key"]) {
    if (row[key]) return String(row[key]);
  }
  return String(row.id ?? "").slice(0, 8);
}

export function dueText(when?: string | null): string {
  if (!when) return "No due time";
  return `Due ${dateTime.format(new Date(when))}`;
}

export function when(value?: string | null): string {
  return value ? dateTime.format(new Date(value)) : "";
}
