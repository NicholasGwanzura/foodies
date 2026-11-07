import type { Metadata } from "next";

import { ModuleTiles } from "@/components/module-tiles";
import { getSession } from "@/lib/session";

export const metadata: Metadata = { title: "My Day" };

function greeting() {
  const hour = Number(
    new Intl.DateTimeFormat("en-ZW", {
      hour: "numeric",
      hour12: false,
      timeZone: "Africa/Harare",
    }).format(new Date()),
  );
  if (hour < 12) return "Good morning";
  return hour < 17 ? "Good afternoon" : "Good evening";
}

export default async function MyDayPage() {
  const { me } = await getSession();
  return (
    <>
      <section className="panel stack">
        <div className="page-head">
          <h1>
            {greeting()}, {me.employee.first_name}
          </h1>
          <p className="muted">{me.employee.job_title}</p>
        </div>
      </section>
      <section className="panel stack" aria-labelledby="modules-h">
        <h2 id="modules-h">Modules</h2>
        <ModuleTiles />
      </section>
    </>
  );
}
