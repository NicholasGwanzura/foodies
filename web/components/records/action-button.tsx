"use client";

import { useActionState } from "react";

import { FormNotice, Submit } from "@/components/form";
import { FieldInput } from "@/components/records/field-input";
import type { FormState } from "@/lib/form-state";
import type { RefOptions } from "@/lib/resources";
import type { ActionMeta } from "@/lib/resource-types";

const DANGER = /reject|cancel|revoke|disable|void|suspend|lock$|decline|retire|delete/i;

export function ActionButton({
  meta,
  action,
  refs,
  currencies,
}: {
  meta: ActionMeta;
  action: (state: FormState, form: FormData) => Promise<FormState>;
  refs: RefOptions;
  currencies: string[];
}) {
  const [state, formAction] = useActionState(action, {});
  const variant = DANGER.test(meta.key) ? "danger" : "";
  if (meta.fields.length === 0) {
    return (
      <form action={formAction} className="stack-s">
        <Submit variant={variant || "ghost"} size="small">{meta.label}</Submit>
        <FormNotice state={state} />
      </form>
    );
  }
  return (
    <details className="action-details">
      <summary className={`btn small ${variant || "ghost"}`}>{meta.label}</summary>
      <form action={formAction} className="stack panel" style={{ boxShadow: "none", border: "1px solid var(--line)", marginTop: 8 }}>
        <FormNotice state={state} />
        <div className="form-grid">
          {meta.fields.map((f) => (
            <FieldInput key={f.name} field={f} error={state.fields?.[f.name]} refs={refs} currencies={currencies} />
          ))}
        </div>
        <div><Submit variant={variant}>{meta.label}</Submit></div>
      </form>
    </details>
  );
}
