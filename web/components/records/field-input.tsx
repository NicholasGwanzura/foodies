"use client";

import { useState } from "react";

import { humanise } from "@/lib/format";
import type { RefOptions } from "@/lib/resources";
import type { FieldMeta } from "@/lib/resource-types";

type Common = {
  field: FieldMeta;
  value?: unknown;
  error?: string;
  refs: RefOptions;
  currencies: string[];
};

function toDateTimeLocal(value: unknown) {
  if (!value) return undefined;
  const d = new Date(String(value));
  return new Date(d.getTime() + 2 * 3600_000).toISOString().slice(0, 16);
}

function Control({ field, value, refs, currencies, invalid, describedBy, id, name }: Omit<Common, "error"> & {
  invalid: boolean;
  describedBy?: string;
  id: string;
  name: string;
}) {
  const shared = {
    id,
    name,
    required: field.required,
    "aria-invalid": invalid || undefined,
    "aria-describedby": describedBy,
  };
  const text = value === null || value === undefined ? undefined : String(value);
  switch (field.kind) {
    case "enum":
      return (
        <select {...shared} defaultValue={text ?? ""}>
          <option value="">Choose…</option>
          {field.options?.map((o) => <option key={o} value={o}>{humanise(o)}</option>)}
        </select>
      );
    case "ref": {
      const options = refs[field.ref ?? ""] ?? [];
      if (options.length === 0) {
        return (
          <select {...shared} defaultValue="" disabled={!field.required}>
            <option value="">{field.required ? "Nothing to choose yet" : "Not available"}</option>
          </select>
        );
      }
      return (
        <select {...shared} defaultValue={text ?? ""}>
          <option value="">Choose…</option>
          {options.map((o) => <option key={o.value} value={o.value}>{o.label}</option>)}
        </select>
      );
    }
    case "currency":
      return (
        <select {...shared} defaultValue={text ?? currencies[0]}>
          {currencies.map((c) => <option key={c}>{c}</option>)}
        </select>
      );
    case "textarea":
      return <textarea {...shared} defaultValue={text} />;
    case "json":
      return <textarea {...shared} defaultValue={value === undefined || value === null ? undefined : JSON.stringify(value, null, 2)} placeholder='{ "key": "value" }' />;
    case "date":
      return <input {...shared} type="date" defaultValue={text} />;
    case "datetime":
      return <input {...shared} type="datetime-local" defaultValue={toDateTimeLocal(value)} />;
    case "number":
      return <input {...shared} type="number" inputMode="numeric" defaultValue={text} />;
    case "decimal":
      return <input {...shared} inputMode="decimal" pattern="-?\d+(\.\d+)?" defaultValue={text} />;
    case "pin":
      return <input {...shared} type="password" className="pin" inputMode="numeric" pattern="\d{6}" maxLength={6} autoComplete="off" />;
    case "password":
      return <input {...shared} type="password" autoComplete="off" />;
    case "list":
      return <input {...shared} defaultValue={Array.isArray(value) ? value.join(", ") : undefined} placeholder="Separate with commas" />;
    default:
      return <input {...shared} defaultValue={text} />;
  }
}

type RowValue = Record<string, string>;

function RowsField({ field, value, error, refs, currencies }: Common) {
  const columns = (field.fields ?? []).filter((c) => !c.readOnly);
  const start = Array.isArray(value) ? (value as RowValue[]) : [];
  const [rows, setRows] = useState<RowValue[]>(start.length > 0 ? start : [{}]);
  const set = (i: number, key: string, v: string) =>
    setRows(rows.map((r, n) => (n === i ? { ...r, [key]: v } : r)));
  const payload = rows
    .filter((r) => Object.values(r).some(Boolean))
    .map((r) => {
      const out: Record<string, unknown> = {};
      for (const c of columns) {
        const v = r[c.name];
        if (v === undefined || v === "") continue;
        out[c.name] = c.kind === "number" ? Number(v) : c.kind === "boolean" ? v === "on" : v;
      }
      return out;
    });
  return (
    <fieldset className="stack-s" style={{ gridColumn: "1 / -1", border: 0, padding: 0 }}>
      <legend className="field-label">{field.label}{field.required ? " *" : ""}</legend>
      <input type="hidden" name={field.name} value={JSON.stringify(payload)} />
      {rows.map((row, i) => (
        <div key={i} className="form-grid panel" style={{ boxShadow: "none", border: "1px solid var(--line)", padding: 10 }}>
          {columns.map((c) => (
            <div className="field" key={c.name}>
              <label htmlFor={`${field.name}-${i}-${c.name}`}>{c.label}</label>
              {c.kind === "enum" || c.kind === "ref" || c.kind === "currency" ? (
                <select
                  id={`${field.name}-${i}-${c.name}`}
                  value={row[c.name] ?? ""}
                  onChange={(e) => set(i, c.name, e.target.value)}
                >
                  <option value="">Choose…</option>
                  {(c.kind === "currency"
                    ? currencies.map((x) => ({ value: x, label: x }))
                    : c.kind === "ref"
                      ? (refs[c.ref ?? ""] ?? [])
                      : (c.options ?? []).map((x) => ({ value: x, label: humanise(x) }))
                  ).map((o) => <option key={o.value} value={o.value}>{o.label}</option>)}
                </select>
              ) : (
                <input
                  id={`${field.name}-${i}-${c.name}`}
                  type={c.kind === "date" ? "date" : c.kind === "number" ? "number" : "text"}
                  inputMode={c.kind === "decimal" ? "decimal" : undefined}
                  value={row[c.name] ?? ""}
                  onChange={(e) => set(i, c.name, e.target.value)}
                />
              )}
            </div>
          ))}
          {rows.length > 1 && (
            <button type="button" className="btn ghost small" onClick={() => setRows(rows.filter((_, n) => n !== i))}>
              Remove row
            </button>
          )}
        </div>
      ))}
      <div>
        <button type="button" className="btn ghost small" onClick={() => setRows([...rows, {}])}>
          Add row
        </button>
      </div>
      {error && <span className="err">{error}</span>}
    </fieldset>
  );
}

export function FieldInput(props: Common) {
  const { field, value, error, currencies } = props;
  const id = `f-${field.name}`;
  if (field.kind === "rows") return <RowsField {...props} />;
  if (field.kind === "boolean") {
    return (
      <label className="check">
        <input type="checkbox" name={field.name} defaultChecked={value === true || (value === undefined && field.name === "active")} />
        {field.label}
      </label>
    );
  }
  if (field.kind === "money") {
    const m = value as { amount: string; currency: string } | undefined;
    return (
      <div className="field">
        <label htmlFor={id}>{field.label}{field.required ? " *" : ""}</label>
        <div className="row" style={{ flexWrap: "nowrap" }}>
          <select name={`${field.name}.currency`} defaultValue={m?.currency ?? currencies[0]} aria-label={`${field.label} currency`} style={{ width: 96, flex: "none" }}>
            {currencies.map((c) => <option key={c}>{c}</option>)}
          </select>
          <input id={id} name={`${field.name}.amount`} inputMode="decimal" pattern="-?\d+(\.\d{1,4})?" defaultValue={m?.amount} aria-invalid={error ? true : undefined} />
        </div>
        {error && <span className="err">{error}</span>}
      </div>
    );
  }
  return (
    <div className="field" style={field.kind === "textarea" || field.kind === "json" ? { gridColumn: "1 / -1" } : undefined}>
      <label htmlFor={id}>{field.label}{field.required ? " *" : ""}</label>
      <Control {...props} id={id} name={field.name} invalid={Boolean(error)} describedBy={error ? `${id}-err` : undefined} />
      {field.hint && !error && <small>{field.hint}</small>}
      {error && <span className="err" id={`${id}-err`}>{error}</span>}
    </div>
  );
}
