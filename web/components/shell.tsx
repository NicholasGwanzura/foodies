import Link from "next/link";

import { Icon } from "@/components/icons";
import { visibleModules } from "@/components/module-tiles";
import { NavLink } from "@/components/nav-link";
import { serverApi } from "@/lib/api/server";
import { RESOURCES } from "@/lib/resources.gen";
import { can, getSession } from "@/lib/session";

async function counts() {
  try {
    const board = await serverApi.get("/dashboards/{dashboard_key}", { path: { dashboard_key: "my-day" } });
    const by = Object.fromEntries((board?.tiles ?? []).map((t) => [t.key, Number(t.value)]));
    return { tasks: by.tasks_due ?? 0, approvals: by.approvals_waiting ?? 0, unread: by.unread ?? 0 };
  } catch {
    return { tasks: 0, approvals: 0, unread: 0 };
  }
}

export async function Shell({ children }: { children: React.ReactNode }) {
  const session = await getSession();
  const { employee } = session.me;
  const { tasks, approvals, unread } = await counts();
  const setup = Object.values(RESOURCES).some((r) => !r.tile && !r.nestedUnder && can(session, r.readPermission));
  return (
    <>
      <header className="arch">
        <div className="arch-bar">
          <Link href="/" className="logo">
            foodies
          </Link>
          <div className="arch-actions">
            <Link href="/search" className="round-btn" aria-label="Find a record">
              <Icon name="search" size={20} />
            </Link>
            <Link href="/notifications" className="round-btn" aria-label={`Notifications, ${unread} unread`}>
              <Icon name="bell" size={20} />
              {unread > 0 && <span className="badge">{unread > 99 ? "99+" : unread}</span>}
            </Link>
            <Link href="/me" className="round-btn" aria-label={`Your account, ${employee.first_name}`}>
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
          <NavLink href="/tasks">
            <Icon name="tasks" size={20} /> Tasks
            {tasks > 0 && <span className="count">{tasks}</span>}
          </NavLink>
          <NavLink href="/approvals">
            <Icon name="approve" size={20} /> Approvals
            {approvals > 0 && <span className="count">{approvals}</span>}
          </NavLink>
          <NavLink href="/escalations">
            <Icon name="alert" size={20} /> Escalations
          </NavLink>
          <hr />
          {visibleModules(session).map((m) => (
            <NavLink key={m.slug} href={`/m/${m.slug}`}>
              <span className="num-tag">{m.number}</span> {m.title}
            </NavLink>
          ))}
          {setup && (
            <>
              <hr />
              <NavLink href="/setup">
                <Icon name="gear" size={20} /> Setup
              </NavLink>
            </>
          )}
        </nav>
        <main className="stack">{children}</main>
      </div>
      <nav className="bottom-nav" aria-label="Main">
        <NavLink href="/" exact>
          <Icon name="home" /> My Day
        </NavLink>
        <NavLink href="/tasks">
          <Icon name="tasks" /> Tasks
          {tasks > 0 && <span className="badge">{tasks}</span>}
        </NavLink>
        <NavLink href="/approvals">
          <Icon name="approve" /> Approvals
          {approvals > 0 && <span className="badge">{approvals}</span>}
        </NavLink>
        <NavLink href="/modules">
          <Icon name="grid" /> Modules
        </NavLink>
      </nav>
    </>
  );
}
