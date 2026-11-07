import type { Metadata } from "next";
import Link from "next/link";
import { notFound } from "next/navigation";

import { RecordForm } from "@/components/records/record-form";
import { loadCurrencies, loadRefs, resourceOr404 } from "@/lib/resources";

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
  const [refs, currencies] = await Promise.all([loadRefs(fields), loadCurrencies()]);
  return (
    <section className="panel stack">
      <div className="page-head">
        <Link className="crumbs" href={`/r/${resource.slug}` as never}>{resource.title}</Link>
        <h1>New {resource.singular.toLowerCase()}</h1>
      </div>
      <RecordForm
        fields={fields}
        action={createRecord.bind(null, resource.slug, parent)}
        refs={refs}
        currencies={currencies}
        submitLabel={`Add ${resource.singular.toLowerCase()}`}
      />
    </section>
  );
}
