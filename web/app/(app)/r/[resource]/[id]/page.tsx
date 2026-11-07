import type { Metadata } from "next";
import Link from "next/link";
import { notFound } from "next/navigation";

import { Icon } from "@/components/icons";
import { ActionButton } from "@/components/records/action-button";
import { StatusPill } from "@/components/records/status-pill";
import { Evidence } from "@/components/records/evidence";
import { RecordTable } from "@/components/records/table";
import { serverApi } from "@/lib/api/server";
import { formatValue, recordLabel } from "@/lib/format";
import { fillPath, listRows, loadCurrencies, loadRefs, resourceOr404, type Row } from "@/lib/resources";
import { RESOURCES } from "@/lib/resources.gen";
import type { FieldMeta } from "@/lib/resource-types";
import { can, getSession } from "@/lib/session";

import { runAction } from "../../actions";

type Props = {
  params: Promise<{ resource: string; id: string }>;
  searchParams: Promise<{ parent?: string }>;
};

export async function generateMetadata({ params }: Props): Promise<Metadata> {
  return { title: resourceOr404((await params).resource).singular };
}

function Value({ field, row, refs }: { field: FieldMeta; row: Row; refs: Awaited<ReturnType<typeof loadRefs>> }) {
  const value = row[field.name];
  if (value === null || value === undefined || value === "") return <span className="muted">Not set</span>;
  if (field.ref) {
    const label = refs[field.ref]?.find((o) => o.value === value)?.label;
    return <Link href={`/r/${field.ref}/${String(value)}` as never}>{label ?? "View"}</Link>;
  }
  if (field.kind === "rows" && Array.isArray(value)) {
    const columns = field.fields ?? [];
    return (
      <div className="table-wrap">
        <table className="data">
          <thead><tr>{columns.map((c) => <th key={c.name}>{c.label}</th>)}</tr></thead>
          <tbody>
            {(value as Record<string, unknown>[]).map((r, i) => (
              <tr key={i}>{columns.map((c) => <td key={c.name}>{formatValue(c, r[c.name])}</td>)}</tr>
            ))}
          </tbody>
        </table>
      </div>
    );
  }
  return formatValue(field, value);
}

export default async function DetailPage({ params, searchParams }: Props) {
  const { resource: slug, id } = await params;
  const resource = resourceOr404(slug);
  const { parent } = await searchParams;
  if (!resource.getPath) notFound();
  const session = await getSession();
  const row = (await serverApi.request("get", fillPath(resource.getPath, id, parent))) as Row;
  const actions = (resource.actions ?? []).filter(
    (a) => can(session, a.permission) && (!a.from || a.from.includes(String(row.status))),
  );
  const nested = Object.values(RESOURCES).filter((r) => r.nestedUnder === slug && can(session, r.readPermission));
  const shown = resource.fields.filter((f) => !["status", "reference"].includes(f.name));
  const simple = shown.filter((f) => f.kind !== "rows");
  const tables = shown.filter((f) => f.kind === "rows" && Array.isArray(row[f.name]));
  const actionFields = actions.flatMap((a) => a.fields);
  const [refs, currencies, children] = await Promise.all([
    loadRefs([...shown, ...actionFields]),
    loadCurrencies(),
    Promise.all(nested.map(async (n) => {
      const list = await listRows(n, { limit: 25 }, id);
      const columns = n.columns.flatMap((c) => n.fields.filter((f) => f.name === c));
      return { resource: n, rows: list.data, refs: await loadRefs(columns) };
    })),
  ]);

  return (
    <>
      <section className="panel stack">
        <div className="spread">
          <div className="page-head">
            <Link className="crumbs" href={`/r/${slug}${parent ? `?parent=${parent}` : ""}` as never}>{resource.title}</Link>
            <h1>{recordLabel(row)}</h1>
            <div className="row">
              {row.reference && row.reference !== recordLabel(row) ? <span className="ref">{String(row.reference)}</span> : null}
              {typeof row.status === "string" && <StatusPill status={row.status} />}
            </div>
          </div>
          {resource.update && can(session, resource.update.permission) && (
            <Link className="btn ghost" href={`/r/${slug}/${id}/edit${parent ? `?parent=${parent}` : ""}` as never}>Edit</Link>
          )}
        </div>
        {actions.length > 0 && (
          <div className="row" style={{ alignItems: "flex-start" }}>
            {actions.map((a) => (
              <ActionButton key={a.key} meta={a} action={runAction.bind(null, slug, id, parent, a.key)} refs={refs} currencies={currencies} />
            ))}
          </div>
        )}
        <dl className="dl">
          {simple.map((f) => (
            <div key={f.name}>
              <dt>{f.label}</dt>
              <dd><Value field={f} row={row} refs={refs} /></dd>
            </div>
          ))}
        </dl>
        {tables.map((f) => (
          <div className="stack-s" key={f.name}>
            <h2>{f.label}</h2>
            <Value field={f} row={row} refs={refs} />
          </div>
        ))}
      </section>
      {children.map(({ resource: n, rows, refs: childRefs }) => (
        <section key={n.slug} className="panel stack">
          <div className="spread">
            <h2>{n.title}</h2>
            {n.create && can(session, n.create.permission) && (
              <Link className="btn ghost small" href={`/r/${n.slug}/new?parent=${id}` as never}>
                <Icon name="plus" size={16} /> Add
              </Link>
            )}
          </div>
          {rows.length === 0 ? <p className="muted">None yet.</p> : <RecordTable resource={n} rows={rows} refs={childRefs} parent={id} />}
        </section>
      ))}
      <Evidence type={slug.replaceAll("-", "_")} id={id} session={session} />
    </>
  );
}
