"use client";

import { useFormStatus } from "react-dom";

import type { FormState } from "@/lib/form-state";

export function Submit({
  children,
  pendingText,
  variant = "",
  full = false,
  size = "",
}: {
  children: React.ReactNode;
  pendingText?: string;
  variant?: "" | "ghost" | "danger";
  full?: boolean;
  size?: "" | "small";
}) {
  const { pending } = useFormStatus();
  return (
    <button
      type="submit"
      className={`btn ${variant} ${size} ${full ? "full" : ""}`}
      disabled={pending}
    >
      {pending ? (pendingText ?? "Working…") : children}
    </button>
  );
}

export function FormNotice({ state }: { state: FormState }) {
  if (state.ok) {
    return (
      <p className="notice good" role="status">
        {state.ok}
      </p>
    );
  }
  if (!state.error) return null;
  return (
    <div className="notice bad" role="alert">
      {state.error}
      {state.blockers && state.blockers.length > 0 && (
        <ul>
          {state.blockers.map((b) => (
            <li key={b}>{b}</li>
          ))}
        </ul>
      )}
    </div>
  );
}

export function Field({
  label,
  name,
  error,
  hint,
  ...props
}: {
  label: string;
  name: string;
  error?: string;
  hint?: string;
} & React.InputHTMLAttributes<HTMLInputElement>) {
  const id = `f-${name}`;
  return (
    <div className="field">
      <label htmlFor={id}>{label}</label>
      <input
        id={id}
        name={name}
        aria-invalid={error ? true : undefined}
        aria-describedby={error ? `${id}-err` : undefined}
        {...props}
      />
      {hint && !error && <small>{hint}</small>}
      {error && (
        <span className="err" id={`${id}-err`}>
          {error}
        </span>
      )}
    </div>
  );
}
