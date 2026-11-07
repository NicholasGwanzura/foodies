import { redirect } from "next/navigation";

import { serverApi } from "@/lib/api/server";

/** Trades a record id for a short-lived signed link, as the signed-in user. */
export async function GET(_: Request, { params }: { params: Promise<{ id: string }> }) {
  const { id } = await params;
  const file = await serverApi.get("/attachments/{attachment_id}", { path: { attachment_id: id } });
  if (!file?.download_url) return new Response("Not available", { status: 404 });
  redirect(file.download_url);
}
