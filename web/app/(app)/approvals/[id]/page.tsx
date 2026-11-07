import type { Metadata } from "next";
import Link from "next/link";

import { StatusPill } from "@/components/records/status-pill";
import { serverApi } from "@/lib/api/server";
import { humanise, money, when } from "@/lib/format";

import { Decision } from "./decision";

export const metadata: Metadata = { title: "Approval" };

type Props = { params: Promise<{ id: string }> };

export default async function ApprovalPage({ params }: Props) {
  const { id } = await params;
  const r = await serverApi.get("/approval-requests/{approval_request_id}", { path: { approval_request_id: id } });
  if (!r) return null;
  const type = r.subject_type ?? "";
  return (
    <>
      <Link className="crumbs" href="/approvals">Approvals</Link>
      <section className="panel stack">
        <div className="page-head">
          <h1>{r.summary}</h1>
          <div className="row">
            <StatusPill status={r.status} />
            {r.amount && <span className="ref">{money(r.amount)}</span>}
          </div>
        </div>
        <dl className="dl">
          <div><dt>Workflow</dt><dd>{humanise(r.workflow_key ?? "")}</dd></div>
          <div><dt>Requested</dt><dd>{when(r.created_at)}</dd></div>
          <div><dt>Step</dt><dd>{r.current_step} of {r.steps?.length ?? 1}</dd></div>
        </dl>
        {r.subject_id && <p><Link className="crumbs" href={`/r/${type.replaceAll("_", "-")}/${r.subject_id}` as never}>Open the record</Link></p>}
        {r.status === "pending" && r.subject_id && <Decision id={id} type={type} subject={r.subject_id} />}
      </section>
    </>
  );
}
