import type { Metadata } from "next";
import Link from "next/link";

import { StatusPill } from "@/components/records/status-pill";
import { serverApi } from "@/lib/api/server";
import { humanise } from "@/lib/format";
import { can, getSession } from "@/lib/session";

export const metadata: Metadata = { title: "Find a record" };

type Props = { searchParams: Promise<{ q?: string }> };

export default async function SearchPage({ searchParams }: Props) {
  const { q } = await searchParams;
  const session = await getSession();
  const allowed = can(session, "reporting.search");
  const hits = q && allowed ? await serverApi.get("/search", { query: { q } }) : undefined;
  return (
    <section className="panel stack">
      <h1>Find a record</h1>
      <form className="row" role="search">
        <input type="search" name="q" defaultValue={q} autoFocus placeholder="Reference, name or title" aria-label="Search" style={{ flex: 1 }} />
        <button className="btn" type="submit">Find</button>
      </form>
      {!allowed && <p className="muted">Search is not part of your role.</p>}
      {hits && hits.data.length === 0 && <div className="empty"><strong>Nothing matches “{q}”</strong></div>}
      {hits && hits.data.length > 0 && (
        <ul className="list">
          {hits.data.map((h) => (
            <li key={h.id}>
              <Link href={`/r/${h.type.replaceAll("_", "-")}/${h.id}` as never}>
                <span><strong>{h.title}</strong><small className="muted">{humanise(h.type)}{h.reference ? `, ${h.reference}` : ""}</small></span>
                {h.status && <StatusPill status={h.status} />}
              </Link>
            </li>
          ))}
        </ul>
      )}
    </section>
  );
}
