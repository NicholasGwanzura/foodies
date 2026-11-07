import type { FieldMeta } from "./resource-types.ts";

/** Why a form value could not be sent, by field name. */
export type CoerceErrors = Record<string, string>;

function one(field: FieldMeta, raw: string, errors: CoerceErrors): unknown {
  switch (field.kind) {
    case "number": {
      const n = Number(raw);
      if (Number.isNaN(n)) errors[field.name] = "Enter a number.";
      return n;
    }
    case "list":
      return raw.split(",").map((s) => s.trim()).filter(Boolean);
    case "json":
      try {
        return JSON.parse(raw);
      } catch {
        errors[field.name] = "This must be valid JSON.";
        return undefined;
      }
    case "datetime":
      return new Date(raw).toISOString();
    case "boolean":
      return raw === "on" || raw === "true";
    default:
      return raw;
  }
}

/** Turns submitted form values into an API body, typed as the fields say. */
export function coerce(
  fields: FieldMeta[],
  get: (name: string) => string | undefined,
): { body: Record<string, unknown>; errors: CoerceErrors } {
  const body: Record<string, unknown> = {};
  const errors: CoerceErrors = {};
  for (const field of fields) {
    if (field.kind === "boolean") {
      body[field.name] = get(field.name) === "on";
      continue;
    }
    if (field.kind === "money") {
      const amount = get(`${field.name}.amount`)?.trim();
      if (amount) {
        body[field.name] = { amount, currency: get(`${field.name}.currency`) || "USD" };
      } else if (field.required) {
        errors[field.name] = "This field is required.";
      }
      continue;
    }
    const raw = get(field.name)?.trim();
    if (!raw) {
      if (field.required) errors[field.name] = "This field is required.";
      continue;
    }
    if (field.kind === "rows") {
      body[field.name] = JSON.parse(raw);
      continue;
    }
    const value = one(field, raw, errors);
    if (value !== undefined) body[field.name] = value;
  }
  return { body, errors };
}
