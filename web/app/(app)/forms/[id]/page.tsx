import type { Metadata } from "next";
import Link from "next/link";

import type { FormDefinition } from "@/components/form-runner";
import { StatusPill } from "@/components/records/status-pill";
import { serverApi } from "@/lib/api/server";
import { items } from "@/lib/api/types";

export const metadata: Metadata = { title: "Submitted checklist" };

type Props = { params: Promise<{ id: string }> };

const ANSWER_LABELS: Record<string, string> = { pass: "Pass", fail: "Fail", na: "N/A", ok: "OK", nok: "Not OK" };
const answerLabel = (v: unknown) => ANSWER_LABELS[String(v)] ?? String(v);

async function templateDefinition(versionId: string): Promise<FormDefinition> {
  const list = await serverApi.get("/form-templates", { query: { limit: 100 } });
  for (const t of items<{ id: string }>(list)) {
    const v = await serverApi
      .get("/form-templates/{form_template_id}/versions/{form_template_version_id}", {
        path: { form_template_id: t.id, form_template_version_id: versionId },
      })
      .catch(() => undefined);
    if (v) return v.definition as unknown as FormDefinition;
  }
  return { sections: [] };
}

export default async function FormResultPage({ params }: Props) {
  const { id } = await params;
  const sub = await serverApi.get("/form-submissions/{form_submission_id}", { path: { form_submission_id: id } });
  if (!sub) return null;
  const failed = sub.result === "fail";
  const answers = (sub.answers ?? {}) as Record<string, unknown>;
  const definition = await templateDefinition(sub.template_version_id);
  const shown = definition.sections.flatMap((s) => s.fields).filter((f) => answers[f.key] !== undefined && answers[f.key] !== "");
  return (
    <section className="panel stack">
      <div className="page-head">
        <span className="ref">{sub.reference}</span>
        <h1>{failed ? "Submitted, with items to fix" : "Checklist submitted"}</h1>
        <div className="row"><StatusPill status={sub.status} />{sub.result && <StatusPill status={sub.result} />}</div>
      </div>
      {failed && <p className="notice warn">Failed answers have raised the follow-up work set on this checklist. The people responsible have been told.</p>}
      <dl className="dl">
        {shown.map((f) => (
          <div key={f.key}><dt>{f.label}</dt><dd>{answerLabel(answers[f.key])}</dd></div>
        ))}
      </dl>
      <div className="row">
        <Link className="btn" href="/">Back to My Day</Link>
        <Link className="btn ghost" href={`/r/form-submissions/${id}` as never}>Open record</Link>
      </div>
    </section>
  );
}
