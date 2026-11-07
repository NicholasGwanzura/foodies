import { serverApi } from "@/lib/api/server";
import { type Attachment, type Comment, items } from "@/lib/api/types";
import { when } from "@/lib/format";
import { can, type Session } from "@/lib/session";

import { EvidenceForms } from "./evidence-forms";

/** EVD-01: files and comments on any record. */
export async function Evidence({ type, id, session }: { type: string; id: string; session: Session }) {
  const canRead = can(session, "evidence.attachment.read");
  if (!canRead) return null;
  const query = { subject_type: type, subject_id: id, limit: 50 };
  const [files, comments] = await Promise.all([
    serverApi.get("/attachments", { query }),
    serverApi.get("/comments", { query }),
  ]);
  const rowsF = items<Attachment>(files);
  const rowsC = items<Comment>(comments);
  return (
    <section className="panel stack">
      <h2>Evidence and comments</h2>
      {rowsF.length > 0 && (
        <ul className="list">
          {rowsF.map((f) => (
            <li key={f.id}>
              <div className="item">
                <span><strong>{f.filename}</strong><small className="muted">{when(f.created_at)}</small></span>
                <a className="btn ghost small" href={`/files/${f.id}`}>Open</a>
              </div>
            </li>
          ))}
        </ul>
      )}
      {rowsC.length > 0 && (
        <ul className="list">
          {rowsC.map((c) => (
            <li key={c.id}><div className="item"><span><strong style={{ fontWeight: 600 }}>{c.body}</strong><small className="muted">{when(c.created_at)}</small></span></div></li>
          ))}
        </ul>
      )}
      <EvidenceForms type={type} id={id} />
    </section>
  );
}
