import type { Metadata } from "next";

import { serverApi } from "@/lib/api/server";

import { ConfirmForm } from "./confirm-form";

export const metadata: Metadata = { title: "Set up two-factor" };

export default async function TwoFactorPage() {
  const enrolment = await serverApi.post("/auth/mfa/enrol");
  return (
    <div className="stack">
      <h1>Set up two-factor sign-in</h1>
      <p className="muted">
        Your role handles money or admin settings, so it needs a code from an
        authenticator app each time you sign in.
      </p>
      <ol className="stack-s" style={{ paddingLeft: 20 }}>
        <li>Open your authenticator app and add an account.</li>
        <li>
          Choose to enter a setup key, and type this key:
          <p
            className="ref"
            data-testid="totp-secret"
            style={{ overflowWrap: "anywhere", whiteSpace: "normal", marginTop: 4 }}
          >
            {enrolment?.secret}
          </p>
        </li>
        <li>Enter the 6-digit code the app shows.</li>
      </ol>
      <ConfirmForm />
    </div>
  );
}
