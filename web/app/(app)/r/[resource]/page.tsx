import type { Metadata } from "next";
import Link from "next/link";

import { Icon } from "@/components/icons";
import { RecordTable } from "@/components/records/table";
import { humanise } from "@/lib/format";
import { listRows, loadRefs, resourceOr404 } from "@/lib/resources";
import { can, getSession } from "@/lib/session";

type Props = {
  params: Promise<{ resource: string }>;
  searchParams: Promise<Record<string, string | undefined>>;
};

export async function generateMetadata({ params }: Props): Promise<Metadata> {
  return { title: resourceOr404((await params).resource).title };
}

export default async function ListPage({ params, searchParams }: Props) {
  const resource = resourceOr404((await params).resource);
  const { q, status, cursor, parent } = await searchParams;
  const session = await getSession();
  const { data, page } = await listRows(resource, { q, status, cursor }, parent);
  const fields = resource.columns.flatMap((n) => resource.fields.filter((f) => f.name === n));
  const refs = await loadRefs(fields);
  const canAdd = resource.create && can(session, resource.create.permission);
  const url = (extra: Record<string, string | undefined>) => {
    const search = new URLSearchParams();
    for (const [k, v] of Object.entries({ q, status, parent, ...extra })) if (v) search.set(k, v);
    const text = search.toString();
    return `/r/${resource.slug}${text ? `?${text}` : ""}` as never;
  };

  return (
    <>
      <section className="panel stack">
        <div className="spread">
          <div className="page-head">
            {resource.tile && <Link className="crumbs" href={`/m/${resource.tile}` as never}>{resource.group}</Link>}
            <h1>{resource.title}</h1>
            <p className="muted">{resource.about}</p>
          </div>
          {canAdd && (
            <Link className="btn" href={`/r/${resource.slug}/new${parent ? `?parent=${parent}` : ""}` as never}>
              <Icon name="plus" size={18} /> New {resource.singular.toLowerCase()}
            </Link>
          )}
        </div>
        <form className="row" role="search">
          <input name="q" type="search" defaultValue={q} placeholder={`Search ${resource.title.toLowerCase()}`} aria-label={`Search ${resource.title}`} style={{ flex: 1, minWidth: 180 }} />
          {status && <input type="hidden" name="status" value={status} />}
          {parent && <input type="hidden" name="parent" value={parent} />}
          <button className="btn ghost" type="submit">Search</button>
        </form>
        {resource.statuses && (
          <nav className="tabs" aria-label="Filter by status">
            <Link href={url({ status: undefined, cursor: undefined })} aria-current={!status ? "page" : undefined}>All</Link>
            {resource.statuses.map((s) => (
              <Link key={s} href={url({ status: s, cursor: undefined })} aria-current={status === s ? "page" : undefined}>{humanise(s)}</Link>
            ))}
          </nav>
        )}
        {data.length === 0 ? (
          <div className="empty">
            <strong>{q || status ? "Nothing matches" : `No ${resource.title.toLowerCase()} yet`}</strong>
            {canAdd && !q && !status && <Link className="btn ghost" href={`/r/${resource.slug}/new` as never}>Add the first one</Link>}
          </div>
        ) : (
          <RecordTable resource={resource} rows={data} refs={refs} parent={parent} />
        )}
        {(cursor || page.has_more) && (
          <div className="spread">
            {cursor ? <Link className="btn ghost small" href={url({ cursor: undefined })}>Back to start</Link> : <span />}
            {page.has_more && page.next_cursor && <Link className="btn ghost small" href={url({ cursor: page.next_cursor })}>Next page</Link>}
          </div>
        )}
      </section>
    </>
  );
}
