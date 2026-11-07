import type { Metadata } from "next";
import Link from "next/link";
import { notFound } from "next/navigation";

import { RecordForm } from "@/components/records/record-form";
import { serverApi } from "@/lib/api/server";
import { fillPath, loadCurrencies, loadRefs, resourceOr404, type Row } from "@/lib/resources";

import { updateRecord } from "../../../actions";

type Props = {
  params: Promise<{ resource: string; id: string }>;
  searchParams: Promise<{ parent?: string }>;
};

export async function generateMetadata({ params }: Props): Promise<Metadata> {
  return { title: `Edit ${resourceOr404((await params).resource).singular.toLowerCase()}` };
}

export default async function EditPage({ params, searchParams }: Props) {
  const { resource: slug, id } = await params;
  const resource = resourceOr404(slug);
  const { parent } = await searchParams;
  if (!resource.update || !resource.getPath) notFound();
  const fields = resource.update.fields;
  const [row, refs, currencies] = await Promise.all([
    serverApi.request("get", fillPath(resource.getPath, id, parent)) as Promise<Row>,
    loadRefs(fields),
    loadCurrencies(),
  ]);
  return (
    <section className="panel stack">
      <div className="page-head">
        <Link className="crumbs" href={`/r/${slug}/${id}${parent ? `?parent=${parent}` : ""}` as never}>Back to record</Link>
        <h1>Edit {resource.singular.toLowerCase()}</h1>
      </div>
      <RecordForm
        fields={fields}
        action={updateRecord.bind(null, slug, id, parent, Number(row.row_version))}
        row={row}
        refs={refs}
        currencies={currencies}
        submitLabel="Save changes"
      />
    </section>
  );
}
