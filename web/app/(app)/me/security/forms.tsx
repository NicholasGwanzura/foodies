"use client";

import { useActionState } from "react";

import { Field, FormNotice, Submit } from "@/components/form";

import { savePassword, savePin } from "./actions";

export function PasswordForm() {
  const [state, action] = useActionState(savePassword, {});
  return (
    <form action={action} className="stack" style={{ maxWidth: 420 }}>
      <FormNotice state={state} />
      <Field label="Current password" name="current_password" type="password" autoComplete="current-password" error={state.fields?.current_password} required />
      <Field label="New password" name="new_password" type="password" autoComplete="new-password" hint="At least 8 characters." error={state.fields?.new_password} required />
      <div>
        <Submit pendingText="Saving…">Save password</Submit>
      </div>
    </form>
  );
}

export function PinForm() {
  const [state, action] = useActionState(savePin, {});
  return (
    <form action={action} className="stack" style={{ maxWidth: 420 }}>
      <FormNotice state={state} />
      <Field label="New 6-digit PIN" name="pin" type="password" className="pin" inputMode="numeric" pattern="\d{6}" maxLength={6} autoComplete="off" error={state.fields?.pin} required />
      <Field label="Your password or current PIN" name="current_secret" type="password" autoComplete="current-password" error={state.fields?.current_secret} required />
      <div>
        <Submit pendingText="Saving…">Save PIN</Submit>
      </div>
    </form>
  );
}
