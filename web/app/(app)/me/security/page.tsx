import type { Metadata } from "next";
import Link from "next/link";

import { PasswordForm, PinForm } from "./forms";

export const metadata: Metadata = { title: "Password and PIN" };

export default function SecurityPage() {
  return (
    <>
      <Link className="crumbs" href="/me">
        My account
      </Link>
      <section className="panel stack">
        <h1>Change your password</h1>
        <PasswordForm />
      </section>
      <section className="panel stack">
        <div className="page-head">
          <h2>Set your signing PIN</h2>
          <p className="muted">
            Your 6-digit PIN confirms sign-offs, in the app and on WhatsApp.
            Five wrong tries lock it until ICT resets it.
          </p>
        </div>
        <PinForm />
      </section>
    </>
  );
}
