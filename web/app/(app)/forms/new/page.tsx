import type { Metadata } from "next";

import { FormRunner, type FormDefinition } from "@/components/form-runner";
import { serverApi } from "@/lib/api/server";

export const metadata: Metadata = { title: "Checklist" };

type Props = { searchParams: Promise<{ template?: string; site?: string; task?: string }> };

export default async function NewFormPage({ searchParams }: Props) {
  const { template, site, task } = await searchParams;
  if (!template || !site) return <section className="panel empty"><strong>Open a checklist from a task.</strong></section>;
  const t = await serverApi.get("/form-templates/{form_template_id}", { path: { form_template_id: template } });
  const versionId = t?.current_version_id;
  if (!versionId) return <section className="panel empty"><strong>This checklist has no published version yet.</strong></section>;
  const v = await serverApi.get("/form-templates/{form_template_id}/versions/{form_template_version_id}", {
    path: { form_template_id: template, form_template_version_id: versionId },
  });
  return (
    <>
      <section className="panel page-head">
        <h1>{t?.title}</h1>
        <p className="muted">Version {v?.version}. Items marked * are required.</p>
      </section>
      <FormRunner definition={v?.definition as unknown as FormDefinition} version={versionId} site={site} task={task} />
    </>
  );
}
