import type { Metadata } from "next";
import Link from "next/link";

import { Icon } from "@/components/icons";
import { RESOURCES } from "@/lib/resources.gen";
import { can, getSession } from "@/lib/session";

export const metadata: Metadata = { title: "Setup" };

export default async function SetupPage() {
  const session = await getSession();
  const items = Object.values(RESOURCES).filter((r) => !r.tile && !r.nestedUnder && can(session, r.readPermission));
  const groups = [...new Set(items.map((r) => r.group))];
  return (
    <section className="panel stack">
      <div className="page-head">
        <h1>Setup</h1>
        <p className="muted">Rules, approval routes, schedules, audit and connections.</p>
      </div>
      {can(session, "rules.get_readiness") && (
        <div className="tiles">
          <Link className="tile" href="/readiness">
            <span className="sq"><Icon name="clipboard" /></span>
            <span className="tx"><strong>Readiness</strong><small>What stops each workflow going live</small></span>
          </Link>
        </div>
      )}
      {groups.map((g) => (
        <div className="stack-s" key={g}>
          <h2>{g}</h2>
          <div className="tiles">
            {items.filter((r) => r.group === g).map((r) => (
              <Link key={r.slug} className="tile" href={`/r/${r.slug}` as never}>
                <span className="sq"><Icon name="gear" /></span>
                <span className="tx"><strong>{r.title}</strong><small>{r.about.split(".")[0]}</small></span>
              </Link>
            ))}
          </div>
        </div>
      ))}
    </section>
  );
}
