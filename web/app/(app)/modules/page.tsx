import type { Metadata } from "next";

import { ModuleTiles } from "@/components/module-tiles";

export const metadata: Metadata = { title: "Modules" };

export default function ModulesPage() {
  return (
    <section className="panel stack">
      <h1>Modules</h1>
      <ModuleTiles />
    </section>
  );
}
