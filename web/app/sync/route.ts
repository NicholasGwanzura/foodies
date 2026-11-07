import { cookies } from "next/headers";

import { serverApi } from "@/lib/api/server";
import type { QueuedForm } from "@/lib/offline-queue";

const DEVICE_COOKIE = "foodies_device";

type Item = { client_record_id: string; status: string; record_id?: string };

/**
 * Sends queued checklists through /sync/batches: first the record, then its
 * submit, which needs the id the server gave it. A row the server refused goes
 * to a supervisor as a conflict, so it still counts as delivered.
 */
export async function POST(request: Request) {
  const device = (await cookies()).get(DEVICE_COOKIE)?.value;
  if (!device) return Response.json({ error: "not_enrolled" }, { status: 409 });
  const forms = (await request.json()) as QueuedForm[];
  const batch = async (items: unknown[]) => {
    const result = await serverApi.request("post", "/sync/batches", { body: { device_id: device, items } });
    return (result as { items: Item[] }).items;
  };
  const created = await batch(
    forms.map((f) => ({
      client_record_id: f.clientRecordId, operation_id: "createFormSubmission", captured_at: f.capturedAt,
      payload: {
        template_version_id: f.version, site_id: f.site, business_date: f.businessDate, task_id: f.task,
        captured_at: f.capturedAt, client_record_id: f.clientRecordId, answers: f.answers,
      },
    })),
  );
  const toSubmit = created.filter((c) => c.record_id);
  if (toSubmit.length > 0) {
    await batch(
      toSubmit.map((c) => ({
        client_record_id: crypto.randomUUID(), operation_id: "submitFormSubmission", captured_at: new Date().toISOString(),
        path_params: { form_submission_id: c.record_id }, payload: {},
      })),
    );
  }
  return Response.json({ delivered: created.map((c) => c.client_record_id) });
}
