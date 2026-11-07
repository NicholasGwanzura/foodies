"use server";

import { createHash } from "node:crypto";

import { revalidatePath } from "next/cache";

import { serverApi } from "@/lib/api/server";
import type { FormState } from "@/lib/form-state";
import { formError } from "@/lib/session";

export async function addComment(type: string, id: string, _: FormState, form: FormData): Promise<FormState> {
  try {
    await serverApi.post("/comments", {
      body: { subject_type: type, subject_id: id, body: String(form.get("body") ?? "") },
    });
  } catch (error) {
    return formError(error);
  }
  revalidatePath(`/r/${type.replaceAll("_", "-")}/${id}`);
  return {};
}

/** The three steps in docs/spec.md 12.7. */
export async function uploadFile(type: string, id: string, _: FormState, form: FormData): Promise<FormState> {
  const file = form.get("file");
  if (!(file instanceof File) || file.size === 0) return { error: "Choose a file first." };
  const bytes = Buffer.from(await file.arrayBuffer());
  try {
    const record = await serverApi.post("/attachments", {
      body: {
        filename: file.name, content_type: file.type || "application/octet-stream", size_bytes: bytes.length,
        sha256: createHash("sha256").update(bytes).digest("hex"), subject_type: type, subject_id: id,
      },
    });
    if (!record?.upload_url) return { error: "The upload could not start." };
    const put = await fetch(record.upload_url, { method: "PUT", body: bytes });
    if (!put.ok) return { error: "The upload failed. Try again." };
    await serverApi.post("/attachments/{attachment_id}/complete", { path: { attachment_id: record.id as string } });
  } catch (error) {
    return formError(error);
  }
  revalidatePath(`/r/${type.replaceAll("_", "-")}/${id}`);
  return { ok: "Uploaded." };
}
