import type { Metadata } from "next";
import Link from "next/link";

import { ModuleTiles } from "@/components/module-tiles";
import { StatusPill } from "@/components/records/status-pill";
import { serverApi } from "@/lib/api/server";
import { items, type Task } from "@/lib/api/types";
import { dueText } from "@/lib/format";
import { getSession } from "@/lib/session";

export const metadata: Metadata = { title: "My Day" };

function greeting() {
  const hour = Number(
    new Intl.DateTimeFormat("en-ZW", { hour: "numeric", hour12: false, timeZone: "Africa/Harare" }).format(new Date()),
  );
  if (hour < 12) return "Good morning";
  return hour < 17 ? "Good afternoon" : "Good evening";
}

export default async function MyDayPage() {
  const session = await getSession();
  const { me } = session;
  const [board, tasks] = await Promise.all([
    serverApi.get("/dashboards/{dashboard_key}", { path: { dashboard_key: "my-day" } }),
    serverApi.get("/tasks", { query: { assignee: "me", status: ["open", "in_progress"], limit: 5 } }),
  ]);
  const summary = (board?.tiles ?? []).filter((t) => ["tasks_due", "approvals_waiting", "open_escalations"].includes(t.key));
  const open = items<Task>(tasks);
  return (
    <>
      <section className="panel stack">
        <div className="page-head">
          <h1>{greeting()}, {me.employee.first_name}</h1>
          <p className="muted">{me.employee.job_title}</p>
        </div>
        <div className="summary">
          {summary.map((t) => (
            <Link key={t.key} href={(t.drilldown ?? "/") as never}>
              <b>{String(t.value)}</b>
              <small>{t.label}</small>
            </Link>
          ))}
        </div>
      </section>
      <section className="panel stack" aria-labelledby="next-h">
        <div className="spread">
          <h2 id="next-h">Next up</h2>
          <Link className="crumbs" href="/tasks">All tasks</Link>
        </div>
        {open.length === 0 ? (
          <p className="muted">Nothing is waiting on you.</p>
        ) : (
          <ul className="list">
            {open.map((t) => (
              <li key={t.id}>
                <Link href={`/tasks/${t.id}` as never}>
                  <span>
                    <strong>{t.title}</strong>
                    <small className={t.overdue ? "late" : "muted"}>{t.overdue ? "Overdue" : dueText(t.due_at)}</small>
                  </span>
                  <StatusPill status={t.status} />
                </Link>
              </li>
            ))}
          </ul>
        )}
      </section>
      <section className="panel stack" aria-labelledby="modules-h">
        <h2 id="modules-h">Modules</h2>
        <ModuleTiles session={session} />
      </section>
    </>
  );
}
