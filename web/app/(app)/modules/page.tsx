import type { Metadata } from "next";

import { ModuleTiles } from "@/components/module-tiles";
import { getSession } from "@/lib/session";

export const metadata: Metadata = { title: "Modules" };

export default async function ModulesPage() {
  return (
    <section className="panel stack">
      <h1>Modules</h1>
      <ModuleTiles session={await getSession()} />
    </section>
  );
}
