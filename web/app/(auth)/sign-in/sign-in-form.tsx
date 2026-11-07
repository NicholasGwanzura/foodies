"use client";

import { useActionState } from "react";

import { Field, FormNotice, Submit } from "@/components/form";

import { signIn } from "./actions";

export function SignInForm() {
  const [state, action] = useActionState(signIn, {});
  return (
    <form action={action} className="stack">
      <h1>Sign in</h1>
      <FormNotice state={state} />
      <Field
        label="Employee number"
        name="employee_number"
        defaultValue={state.values?.employee_number}
        autoComplete="username"
        autoCapitalize="characters"
        autoFocus
        required
      />
      <Field
        label="Password"
        name="password"
        type="password"
        autoComplete="current-password"
        required
      />
      <Submit full pendingText="Signing in…">
        Sign in
      </Submit>
      <p className="small muted">
        Locked out or forgotten your password? Ask ICT to reset it.
      </p>
    </form>
  );
}
