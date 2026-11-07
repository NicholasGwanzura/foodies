"use client";

import { useActionState } from "react";

import { Field, FormNotice, Submit } from "@/components/form";

import { changePassword } from "../actions";

export default function PasswordPage() {
  const [state, action] = useActionState(changePassword, {});
  return (
    <form action={action} className="stack">
      <h1>Choose a new password</h1>
      <FormNotice state={state} />
      <Field
        label="Current password"
        name="current_password"
        type="password"
        autoComplete="current-password"
        error={state.fields?.current_password}
        required
      />
      <Field
        label="New password"
        name="new_password"
        type="password"
        autoComplete="new-password"
        hint="At least 8 characters."
        error={state.fields?.new_password}
        required
      />
      <Field
        label="Repeat new password"
        name="confirm"
        type="password"
        autoComplete="new-password"
        error={state.fields?.confirm}
        required
      />
      <Submit full pendingText="Saving…">
        Save password
      </Submit>
    </form>
  );
}
