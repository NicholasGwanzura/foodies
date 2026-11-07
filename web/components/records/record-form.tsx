"use client";

import { useActionState } from "react";

import { FormNotice, Submit } from "@/components/form";
import { FieldInput } from "@/components/records/field-input";
import type { FormState } from "@/lib/form-state";
import type { RefOptions, Row } from "@/lib/resources";
import type { FieldMeta } from "@/lib/resource-types";

export function RecordForm({
  fields,
  action,
  row,
  refs,
  currencies,
  submitLabel,
}: {
  fields: FieldMeta[];
  action: (state: FormState, form: FormData) => Promise<FormState>;
  row?: Row;
  refs: RefOptions;
  currencies: string[];
  submitLabel: string;
}) {
  const [state, formAction] = useActionState(action, {});
  return (
    <form action={formAction} className="stack">
      <FormNotice state={state} />
      <div className="form-grid">
        {fields.map((f) => (
          <FieldInput key={f.name} field={f} value={row?.[f.name]} error={state.fields?.[f.name]} refs={refs} currencies={currencies} />
        ))}
      </div>
      <div className="row">
        <Submit pendingText="Saving…">{submitLabel}</Submit>
      </div>
    </form>
  );
}
