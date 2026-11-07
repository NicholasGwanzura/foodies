import Link from "next/link";

import { Icon } from "@/components/icons";
import { MODULES } from "@/lib/modules";

export function ModuleTiles() {
  return (
    <div className="tiles">
      {MODULES.map((m) => (
        <Link key={m.slug} className="tile" href={`/m/${m.slug}` as never}>
          <span className="sq">
            <Icon name={m.icon} />
          </span>
          <span className="tx">
            <i>{m.number}</i>
            <strong>{m.title}</strong>
          </span>
        </Link>
      ))}
    </div>
  );
}
