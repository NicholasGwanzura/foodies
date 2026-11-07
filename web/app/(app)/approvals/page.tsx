import type { Metadata } from "next";
import Link from "next/link";

import { serverApi } from "@/lib/api/server";
import { items, type ApprovalRequest } from "@/lib/api/types";
import { money, when } from "@/lib/format";

export const metadata: Metadata = { title: "Approvals" };

export default async function ApprovalsPage() {
  const inbox = await serverApi.get("/approval-requests", { query: { assignee: "me", limit: 50 } });
  const rows = items<ApprovalRequest>(inbox);
  return (
    <section className="panel stack">
      <h1>Approvals</h1>
      {rows.length === 0 ? (
        <div className="empty"><strong>Nothing is waiting for your decision</strong></div>
      ) : (
        <ul className="list">
          {rows.map((r) => (
            <li key={r.id}>
              <Link href={`/approvals/${r.id}` as never}>
                <span>
                  <strong>{r.summary}</strong>
                  <small className="muted">Requested {when(r.created_at)}</small>
                </span>
                {r.amount && <span className="ref">{money(r.amount)}</span>}
              </Link>
            </li>
          ))}
        </ul>
      )}
    </section>
  );
}
