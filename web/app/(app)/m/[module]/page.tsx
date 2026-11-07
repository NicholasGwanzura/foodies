import type { Metadata } from "next";
import Link from "next/link";
import { notFound } from "next/navigation";

import { Icon } from "@/components/icons";
import { MODULES } from "@/lib/modules";
import { RESOURCES } from "@/lib/resources.gen";
import { can, getSession } from "@/lib/session";

type Props = { params: Promise<{ module: string }> };

export async function generateMetadata({ params }: Props): Promise<Metadata> {
  const slug = (await params).module;
  return { title: MODULES.find((m) => m.slug === slug)?.title ?? "Module" };
}

export default async function ModulePage({ params }: Props) {
  const slug = (await params).module;
  const mod = MODULES.find((m) => m.slug === slug);
  if (!mod) notFound();
  const session = await getSession();
  const items = Object.values(RESOURCES).filter(
    (r) => r.tile === slug && !r.nestedUnder && can(session, r.readPermission),
  );
  const groups = [...new Set(items.map((r) => r.group))];
  return (
    <section className="panel stack">
      <div className="page-head">
        <Link className="crumbs" href="/modules">Modules</Link>
        <h1>{mod.title}</h1>
      </div>
      {items.length === 0 && (
        <div className="empty">
          <strong>Nothing here for your role yet</strong>
          <p>Ask your manager if you think you need access.</p>
        </div>
      )}
      {groups.map((g) => (
        <div className="stack-s" key={g}>
          {groups.length > 1 && <h2>{g}</h2>}
          <div className="tiles">
            {items.filter((r) => r.group === g).map((r) => (
              <Link key={r.slug} className="tile" href={`/r/${r.slug}` as never}>
                <span className="sq"><Icon name={mod.icon} /></span>
                <span className="tx">
                  <strong>{r.title}</strong>
                  <small>{r.about.split(".")[0]}</small>
                </span>
              </Link>
            ))}
          </div>
        </div>
      ))}
    </section>
  );
}
