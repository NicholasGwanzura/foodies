import type { Metadata } from "next";
import Link from "next/link";
import { notFound } from "next/navigation";

import { RecordForm } from "@/components/records/record-form";
import { loadCurrencies, loadRefs, resourceOr404, type Row } from "@/lib/resources";
import { getSession } from "@/lib/session";

import { createRecord } from "../../actions";

type Props = {
  params: Promise<{ resource: string }>;
  searchParams: Promise<{ parent?: string }>;
};

export async function generateMetadata({ params }: Props): Promise<Metadata> {
  return { title: `New ${resourceOr404((await params).resource).singular.toLowerCase()}` };
}

export default async function NewPage({ params, searchParams }: Props) {
  const resource = resourceOr404((await params).resource);
  const { parent } = await searchParams;
  if (!resource.create) notFound();
  const fields = resource.create.fields;
  const [refs, currencies, { me }] = await Promise.all([loadRefs(fields), loadCurrencies(), getSession()]);
  const defaults: Row = { id: "" };
  const today = new Intl.DateTimeFormat("en-CA", { timeZone: "Africa/Harare" }).format(new Date());
  for (const f of fields) {
    if (f.name === "site_id") defaults[f.name] = me.employee.site_id;
    if (f.kind === "date" && /business_date|_date$/.test(f.name) && !/expiry|due|end|required|delivery|until/.test(f.name)) defaults[f.name] = today;
    if (/^(requester|reporter|originator)_employee_id$/.test(f.name)) defaults[f.name] = me.employee.id;
  }
  return (
    <section className="panel stack">
      <div className="page-head">
        <Link className="crumbs" href={`/r/${resource.slug}` as never}>{resource.title}</Link>
        <h1>New {resource.singular.toLowerCase()}</h1>
      </div>
      <RecordForm
        fields={fields}
        row={defaults}
        action={createRecord.bind(null, resource.slug, parent)}
        refs={refs}
        currencies={currencies}
        submitLabel={`Add ${resource.singular.toLowerCase()}`}
      />
    </section>
  );
}
