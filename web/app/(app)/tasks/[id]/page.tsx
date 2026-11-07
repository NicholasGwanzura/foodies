import type { Metadata } from "next";
import Link from "next/link";

import { StatusPill } from "@/components/records/status-pill";
import { serverApi } from "@/lib/api/server";
import { dueText, humanise, when } from "@/lib/format";

import { TaskActions } from "./task-actions";

export const metadata: Metadata = { title: "Task" };

type Props = { params: Promise<{ id: string }> };

export default async function TaskPage({ params }: Props) {
  const { id } = await params;
  const task = await serverApi.get("/tasks/{task_id}", { path: { task_id: id } });
  if (!task) return null;
  const schedule = task.schedule_id
    ? await serverApi.get("/schedules/{schedule_id}", { path: { schedule_id: task.schedule_id } }).catch(() => undefined)
    : undefined;
  const template = schedule?.form_template_id;
  const open = task.status === "open" || task.status === "in_progress";
  return (
    <>
      <Link className="crumbs" href="/tasks">Tasks</Link>
      <section className="panel stack">
        <div className="page-head">
          <h1>{task.title}</h1>
          <div className="row">
            <StatusPill status={task.status} />
            <span className={task.overdue ? "late" : "muted"}>{task.overdue ? "Overdue" : dueText(task.due_at)}</span>
          </div>
        </div>
        {task.skip_reason && <p className="notice warn">Skipped: {task.skip_reason}</p>}
        {task.completed_at && <p className="notice good">Finished {when(task.completed_at)}</p>}
        {template && open && (
          <div>
            <Link className="btn" href={`/forms/new?template=${template}&site=${task.site_id}&task=${task.id}` as never}>
              Fill in the checklist
            </Link>
          </div>
        )}
        {task.subject_type && task.subject_id && (
          <p><Link className="crumbs" href={`/r/${task.subject_type.replaceAll("_", "-")}/${task.subject_id}` as never}>Open the {humanise(task.subject_type).toLowerCase()}</Link></p>
        )}
        <TaskActions id={task.id as string} status={task.status} hasForm={Boolean(template)} />
      </section>
    </>
  );
}
