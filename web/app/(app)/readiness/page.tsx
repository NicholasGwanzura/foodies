import type { Metadata } from "next";
import Link from "next/link";

import { StatusPill } from "@/components/records/status-pill";
import { serverApi } from "@/lib/api/server";
import { items, type Site } from "@/lib/api/types";
import { humanise } from "@/lib/format";

export const metadata: Metadata = { title: "Readiness" };

export default async function ReadinessPage() {
  const [summary, sites] = await Promise.all([
    serverApi.get("/readiness"),
    serverApi.get("/sites", { query: { limit: 100 } }),
  ]);
  const rows = summary?.items ?? [];
  const names = new Map(items<Site>(sites).map((s) => [s.id, s.code]));
  const siteIds = [...new Set(rows.map((i) => i.site_id))];
  const workflows = [...new Set(rows.map((i) => i.workflow_key))];
  const missing = [...new Set(rows.flatMap((i) => i.missing_rule_keys ?? []))].sort();
  return (
    <>
      <section className="panel stack">
        <div className="page-head">
          <h1>Readiness</h1>
          <p className="muted">A workflow can only go live at a site once every rule it depends on has an approved value.</p>
        </div>
        <div className="table-wrap">
          <table className="data matrix">
            <thead>
              <tr><th>Workflow</th>{siteIds.map((s) => <th key={s}>{String(names.get(s) ?? "Site")}</th>)}</tr>
            </thead>
            <tbody>
              {workflows.map((w) => (
                <tr key={w}>
                  <td><strong>{humanise(w.replace(".", ": "))}</strong></td>
                  {siteIds.map((s) => {
                    const item = rows.find((i) => i.workflow_key === w && i.site_id === s);
                    return <td key={s}>{item && <StatusPill status={item.status} />}</td>;
                  })}
                </tr>
              ))}
            </tbody>
          </table>
        </div>
        <p><Link className="crumbs" href="/r/workflow-activations">Activate or suspend a workflow</Link></p>
      </section>
      <section className="panel stack">
        <h2>Rules still needing an approved value</h2>
        {missing.length === 0 ? <p className="muted">None.</p> : (
          <ul className="list">
            {missing.map((k) => (
              <li key={k}><Link href={`/r/rule-definitions?q=${k}` as never}><span><strong>{k}</strong></span></Link></li>
            ))}
          </ul>
        )}
      </section>
    </>
  );
}
