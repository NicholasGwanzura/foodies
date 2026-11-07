"use client";

import { useActionState } from "react";

import { Field, FormNotice, Submit } from "@/components/form";

import { confirmTwoFactor } from "../actions";

export function ConfirmForm() {
  const [state, action] = useActionState(confirmTwoFactor, {});
  return (
    <form action={action} className="stack">
      <FormNotice state={state} />
      <Field
        label="6-digit code"
        name="code"
        className="pin"
        inputMode="numeric"
        autoComplete="one-time-code"
        pattern="\d{6}"
        maxLength={6}
        error={state.fields?.code}
        required
      />
      <Submit full pendingText="Checking…">
        Turn on two-factor
      </Submit>
    </form>
  );
}
