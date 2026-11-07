"use client";

import { useActionState } from "react";

import { Field, FormNotice } from "@/components/form";
import { useFormStatus } from "react-dom";

import { decide } from "../actions";

function Choice({ value, children, variant = "" }: { value: string; children: string; variant?: string }) {
  const { pending } = useFormStatus();
  return <button type="submit" name="decision" value={value} className={`btn ${variant}`} disabled={pending}>{children}</button>;
}

export function Decision({ id, type, subject }: { id: string; type: string; subject: string }) {
  const [state, action] = useActionState(decide.bind(null, { id, type, subject }), {});
  return (
    <form action={action} className="stack">
      <h2>Your decision</h2>
      <FormNotice state={state} />
      <div className="field">
        <label htmlFor="comment">Comment</label>
        <textarea id="comment" name="comment" />
        <small>Needed if you reject or return this.</small>
      </div>
      <Field label="Your 6-digit PIN" name="pin" type="password" className="pin" inputMode="numeric" maxLength={6} autoComplete="off" hint="Needed to approve or reject." />
      <div className="row">
        <Choice value="approve">Approve</Choice>
        <Choice value="return" variant="ghost">Return for changes</Choice>
        <Choice value="reject" variant="danger">Reject</Choice>
      </div>
    </form>
  );
}
