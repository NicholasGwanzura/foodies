import Link from "next/link";

import { Icon } from "@/components/icons";
import { MODULES } from "@/lib/modules";
import { RESOURCES } from "@/lib/resources.gen";
import { can, type Session } from "@/lib/session";

/** The tiles a role can use. A tile with nothing readable behind it is hidden. */
export function visibleModules(session: Session) {
  const resources = Object.values(RESOURCES);
  return MODULES.filter((m) =>
    resources.some((r) => r.tile === m.slug && !r.nestedUnder && can(session, r.readPermission)),
  );
}

export function ModuleTiles({ session }: { session: Session }) {
  return (
    <div className="tiles">
      {visibleModules(session).map((m) => (
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
