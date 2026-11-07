import type { Metadata } from "next";
import Link from "next/link";

import { Submit } from "@/components/form";
import { serverApi } from "@/lib/api/server";
import { items, type Notification } from "@/lib/api/types";
import { when } from "@/lib/format";

import { readAll } from "./actions";

export const metadata: Metadata = { title: "Notifications" };

const hrefFor = (type?: string | null, id?: string | null) => {
  if (type === "approval_requests") return `/approvals/${id}`;
  if (type === "tasks") return `/tasks/${id}`;
  if (type === "escalations") return "/escalations";
  return type && id ? `/r/${type.replaceAll("_", "-")}/${id}` : "/";
};

export default async function NotificationsPage() {
  const list = await serverApi.get("/notifications", { query: { limit: 50 } });
  const rows = items<Notification>(list);
  const unread = rows.some((n) => !n.read_at);
  return (
    <section className="panel stack">
      <div className="spread">
        <h1>Notifications</h1>
        {unread && <form action={readAll}><Submit variant="ghost" size="small">Mark all read</Submit></form>}
      </div>
      {rows.length === 0 ? (
        <div className="empty"><strong>No notifications yet</strong></div>
      ) : (
        <ul className="list">
          {rows.map((n) => (
            <li key={n.id} className={n.read_at ? undefined : "unread"}>
              <Link href={hrefFor(n.subject_type, n.subject_id) as never}>
                <span>
                  <strong>{n.title}</strong>
                  <small className="muted">{n.body}</small>
                </span>
                <small className="muted">{when(n.created_at)}</small>
              </Link>
            </li>
          ))}
        </ul>
      )}
    </section>
  );
}
