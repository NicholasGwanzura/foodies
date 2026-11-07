"use client";

import { useActionState } from "react";

import { Field, FormNotice, Submit } from "@/components/form";

import { claimTask, completeTask, skipTask } from "../actions";

export function TaskActions({ id, status, hasForm }: { id: string; status: string; hasForm: boolean }) {
  const [claimed, claim] = useActionState(claimTask.bind(null, id), {});
  const [finished, complete] = useActionState(completeTask.bind(null, id), {});
  const [skipped, skip] = useActionState(skipTask.bind(null, id), {});
  const open = status === "open" || status === "in_progress";
  if (!open) return null;
  return (
    <div className="stack-s">
      <FormNotice state={claimed} />
      <FormNotice state={finished} />
      <div className="row">
        {status === "open" && <form action={claim}><Submit variant={hasForm ? "ghost" : ""}>Claim task</Submit></form>}
        {!hasForm && <form action={complete}><Submit>Mark as done</Submit></form>}
      </div>
      <details>
        <summary className="crumbs" style={{ cursor: "pointer" }}>Cannot do this today?</summary>
        <form action={skip} className="stack" style={{ marginTop: 10, maxWidth: 460 }}>
          <FormNotice state={skipped} />
          <Field label="Why is it being skipped?" name="reason" required />
          <Field label="What will you do instead?" name="alternative_control" hint="A skipped control needs an alternative." />
          <div><Submit variant="ghost">Skip this task</Submit></div>
        </form>
      </details>
    </div>
  );
}
