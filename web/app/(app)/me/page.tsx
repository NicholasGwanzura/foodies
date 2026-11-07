import type { Metadata } from "next";
import Link from "next/link";

import { signOut } from "@/app/(auth)/sign-in/actions";
import { Submit } from "@/components/form";
import { serverApi } from "@/lib/api/server";
import { can, getSession } from "@/lib/session";

import { EnrolDevice } from "./enrol-device";

export const metadata: Metadata = { title: "My account" };

const when = new Intl.DateTimeFormat("en-ZW", {
  dateStyle: "medium",
  timeStyle: "short",
  timeZone: "Africa/Harare",
});

export default async function AccountPage() {
  const session = await getSession();
  const { me } = session;
  const sessions = await serverApi.get("/me/sessions");
  return (
    <>
      <section className="panel stack">
        <div className="page-head">
          <h1>
            {me.employee.first_name} {me.employee.last_name}
          </h1>
          <p className="muted">
            {me.employee.job_title}, employee {me.employee.employee_number}
          </p>
        </div>
        <dl className="dl">
          <div>
            <dt>Roles</dt>
            <dd>{me.roles.map((r) => r.name).join(", ") || "None yet"}</dd>
          </div>
          <div>
            <dt>Two-factor sign-in</dt>
            <dd>{me.user.mfa_required ? "Required" : "Not required for your role"}</dd>
          </div>
        </dl>
        <div className="row">
          <Link className="btn ghost" href="/me/security">
            Password and PIN
          </Link>
          <form action={signOut}>
            <Submit variant="ghost" pendingText="Signing out…">
              Sign out
            </Submit>
          </form>
        </div>
      </section>
      {can(session, "iam.device.write") && <EnrolDevice />}
      <section className="panel stack">
        <h2>Signed in on</h2>
        <ul className="stack-s" style={{ listStyle: "none", padding: 0 }}>
          {sessions?.data.map((s) => (
            <li key={s.id} className="spread">
              <span>{s.channel === "whatsapp" ? "WhatsApp" : "Web app"}</span>
              <span className="muted small">
                Last active {when.format(new Date(s.last_seen_at))}
              </span>
            </li>
          ))}
        </ul>
      </section>
    </>
  );
}
