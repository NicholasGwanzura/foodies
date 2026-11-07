import type { Metadata } from "next";

import { StatusPill } from "@/components/records/status-pill";
import { serverApi } from "@/lib/api/server";
import { items, type Escalation } from "@/lib/api/types";
import { humanise, when } from "@/lib/format";
import { can, getSession } from "@/lib/session";

import { Ack } from "./ack";

export const metadata: Metadata = { title: "Escalations" };

export default async function EscalationsPage() {
  const session = await getSession();
  if (!can(session, "exceptions.escalation.read")) {
    return <section className="panel empty"><strong>Escalations are not part of your role</strong></section>;
  }
  const list = await serverApi.get("/escalations", { query: { status: ["open", "breached", "acknowledged"], limit: 50 } });
  const rows = items<Escalation>(list);
  return (
    <section className="panel stack">
      <div className="page-head">
        <h1>Escalations</h1>
        <p className="muted">Anything not acknowledged in time moves up a level.</p>
      </div>
      {rows.length === 0 ? (
        <div className="empty"><strong>Nothing is escalated right now</strong></div>
      ) : (
        <ul className="list">
          {rows.map((e) => (
            <li key={e.id}>
              <div className="item">
                <span>
                  <strong>Level {e.severity_level}: {humanise(e.domain ?? "general")}</strong>
                  <small className="muted">{humanise(e.subject_type ?? "")}, raised {when(e.triggered_at)}, step {e.current_step}</small>
                </span>
                <span className="row">
                  <StatusPill status={e.status} />
                  {e.status !== "acknowledged" && <Ack id={e.id as string} />}
                </span>
              </div>
            </li>
          ))}
        </ul>
      )}
    </section>
  );
}
