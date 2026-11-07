import type { Metadata } from "next";
import Link from "next/link";

import { StatusPill } from "@/components/records/status-pill";
import { serverApi } from "@/lib/api/server";
import { items, type Task } from "@/lib/api/types";
import { dueText } from "@/lib/format";

export const metadata: Metadata = { title: "Tasks" };

type Props = { searchParams: Promise<{ view?: string; overdue?: string }> };

export default async function TasksPage({ searchParams }: Props) {
  const { view, overdue } = await searchParams;
  const done = view === "done";
  const tasks = await serverApi.get("/tasks", {
    query: {
      assignee: "me", limit: 50,
      status: done ? ["done", "skipped"] : ["open", "in_progress"],
      overdue: overdue === "true" ? true : undefined,
    },
  });
  const rows = items<Task>(tasks);
  return (
    <section className="panel stack">
      <h1>Tasks</h1>
      <nav className="tabs" aria-label="Filter tasks">
        <Link href="/tasks" aria-current={!done && !overdue ? "page" : undefined}>To do</Link>
        <Link href="/tasks?overdue=true" aria-current={overdue ? "page" : undefined}>Overdue</Link>
        <Link href="/tasks?view=done" aria-current={done ? "page" : undefined}>Finished</Link>
      </nav>
      {rows.length === 0 ? (
        <div className="empty">
          <strong>{done ? "Nothing finished yet" : "You are all caught up"}</strong>
          <p>New tasks appear here when a schedule or a manager assigns them.</p>
        </div>
      ) : (
        <ul className="list">
          {rows.map((t) => (
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
  );
}
