"use client";

import Link from "next/link";
import { useState, useTransition } from "react";

import { submitForm } from "@/app/(app)/forms/actions";
import { FormNotice } from "@/components/form";
import type { FormState } from "@/lib/form-state";
import { enqueue } from "@/lib/offline-queue";

export type FormFieldDef = {
  key: string;
  label: string;
  type: string;
  required?: boolean;
  options?: string[];
  max?: number;
  min?: number;
};
export type FormDefinition = { sections: { title: string; fields: FormFieldDef[] }[] };
type Answers = Record<string, string>;

const CHOICES: Record<string, [string, string][]> = {
  pass_fail: [["pass", "Pass"], ["fail", "Fail"], ["na", "N/A"]],
  ok_nok: [["ok", "OK"], ["nok", "Not OK"]],
};

function Answer({ field, value, onChange, error }: { field: FormFieldDef; value?: string; onChange: (v: string) => void; error?: string }) {
  const options = CHOICES[field.type] ?? (field.type === "choice" ? (field.options ?? []).map((o): [string, string] => [o, o]) : null);
  const limit = field.max !== undefined ? `Limit ${field.max}` : undefined;
  const out = field.max !== undefined && value !== undefined && value !== "" && Number(value) > field.max;
  return (
    <div className="question">
      <span className="q" id={`q-${field.key}`}>{field.label}{field.required ? " *" : ""}</span>
      {options ? (
        <div className="choice" role="radiogroup" aria-labelledby={`q-${field.key}`}>
          {options.map(([v, label]) => (
            <label key={v}>
              <input type="radio" name={field.key} value={v} checked={value === v} onChange={() => onChange(v)} />
              <span>{label}</span>
            </label>
          ))}
        </div>
      ) : (
        <input
          aria-labelledby={`q-${field.key}`}
          type={field.type === "date" ? "date" : "text"}
          inputMode={field.type === "number" || field.type === "reading" ? "decimal" : undefined}
          value={value ?? ""}
          onChange={(e) => onChange(e.target.value)}
        />
      )}
      {limit && <small className={out ? "late" : "muted"}>{out ? `Over the limit of ${field.max}. This will be flagged.` : limit}</small>}
      {error && <span className="err" style={{ color: "var(--bad)", fontWeight: 700, fontSize: 13 }}>{error}</span>}
    </div>
  );
}

export function FormRunner({ definition, version, site, task, id, rowVersion, initial }: {
  definition: FormDefinition;
  version: string;
  site: string;
  task?: string;
  id?: string;
  rowVersion?: number;
  initial?: Answers;
}) {
  const [answers, setAnswers] = useState<Answers>(initial ?? {});
  const [state, setState] = useState<FormState>({});
  const [pending, start] = useTransition();
  const [queued, setQueued] = useState(false);
  const errors = Object.fromEntries(Object.entries(state.fields ?? {}).map(([k, v]) => [k.replace("answers/", ""), v]));
  const numeric = new Set(definition.sections.flatMap((s) => s.fields).filter((f) => f.type === "number" || f.type === "reading").map((f) => f.key));

  function send() {
    const body = Object.fromEntries(Object.entries(answers).map(([k, v]) => [k, numeric.has(k) ? Number(v) : v]));
    start(async () => {
      const known = state.values?.rowVersion ? Number(state.values.rowVersion) : rowVersion;
      const draft = id ?? state.values?.id;
      try {
        setState(await submitForm({ version, site, task, id: draft, rowVersion: known, answers: body }));
      } catch (error) {
        // submitForm redirects by throwing, which is success.
        if (String((error as { digest?: string }).digest ?? "").startsWith("NEXT_REDIRECT")) throw error;
        if (draft) {
          setState({ error: "That did not send. Check your connection and try again." });
          return;
        }
        enqueue(localStorage, {
          clientRecordId: crypto.randomUUID(), capturedAt: new Date().toISOString(), version, site, task,
          businessDate: new Intl.DateTimeFormat("en-CA", { timeZone: "Africa/Harare" }).format(new Date()),
          answers: body,
        });
        setQueued(true);
      }
    });
  }

  if (queued) {
    return (
      <section className="panel stack">
        <h2>Saved on this device</h2>
        <p className="muted">It sends by itself when the connection is back. You can carry on.</p>
        <div><Link className="btn" href="/">Back to My Day</Link></div>
      </section>
    );
  }

  return (
    <div className="stack">
      <FormNotice state={state} />
      {definition.sections.map((section) => (
        <section className="panel" key={section.title}>
          <h2>{section.title}</h2>
          {section.fields.map((f) => (
            <Answer key={f.key} field={f} value={answers[f.key]} error={errors[f.key]} onChange={(v) => setAnswers((a) => ({ ...a, [f.key]: v }))} />
          ))}
        </section>
      ))}
      <div className="sticky-foot">
        <button type="button" className="btn" disabled={pending} onClick={send}>{pending ? "Submitting…" : "Submit checklist"}</button>
      </div>
    </div>
  );
}
