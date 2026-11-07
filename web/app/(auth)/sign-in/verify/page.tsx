"use client";

import { useActionState } from "react";

import { Field, FormNotice, Submit } from "@/components/form";

import { verifyCode } from "../actions";

export default function VerifyPage() {
  const [state, action] = useActionState(verifyCode, {});
  return (
    <form action={action} className="stack">
      <h1>Enter your code</h1>
      <p className="muted">
        Open your authenticator app and enter the 6-digit code for Foodies.
      </p>
      <FormNotice state={state} />
      <Field
        label="6-digit code"
        name="code"
        className="pin"
        inputMode="numeric"
        autoComplete="one-time-code"
        pattern="\d{6}"
        maxLength={6}
        autoFocus
        required
      />
      <Submit full pendingText="Checking…">
        Continue
      </Submit>
    </form>
  );
}
