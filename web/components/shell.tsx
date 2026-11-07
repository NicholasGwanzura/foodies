import Link from "next/link";

import { Icon } from "@/components/icons";
import { NavLink } from "@/components/nav-link";
import { MODULES } from "@/lib/modules";
import { getSession } from "@/lib/session";

export async function Shell({ children }: { children: React.ReactNode }) {
  const session = await getSession();
  const { employee } = session.me;
  return (
    <>
      <header className="arch">
        <div className="arch-bar">
          <Link href="/" className="logo">
            foodies
          </Link>
          <div className="arch-actions">
            <Link
              href="/me"
              className="round-btn"
              aria-label={`Your account, ${employee.first_name}`}
            >
              <Icon name="user" size={20} />
            </Link>
          </div>
        </div>
      </header>
      <div className="frame">
        <nav className="side" aria-label="Main">
          <NavLink href="/" exact>
            <Icon name="home" size={20} /> My Day
          </NavLink>
          <hr />
          {MODULES.map((m) => (
            <NavLink key={m.slug} href={`/m/${m.slug}`}>
              <span className="num-tag">{m.number}</span> {m.title}
            </NavLink>
          ))}
        </nav>
        <main className="stack">{children}</main>
      </div>
      <nav className="bottom-nav" aria-label="Main">
        <NavLink href="/" exact>
          <Icon name="home" /> My Day
        </NavLink>
        <NavLink href="/modules">
          <Icon name="grid" /> Modules
        </NavLink>
        <NavLink href="/me">
          <Icon name="user" /> Account
        </NavLink>
      </nav>
    </>
  );
}
