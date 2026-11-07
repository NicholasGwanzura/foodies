"use server";

import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";

import { serverApi } from "@/lib/api/server";
import type { FormState } from "@/lib/form-state";
import { formError } from "@/lib/session";
import { signRecord } from "@/lib/sign";

type Subject = { id: string; type: string; subject: string };

export async function decide(subject: Subject, _: FormState, form: FormData): Promise<FormState> {
  const decision = String(form.get("decision"));
  const comment = String(form.get("comment") ?? "").trim();
  const path = { approval_request_id: subject.id };
  try {
    if (decision === "return") {
      await serverApi.post("/approval-requests/{approval_request_id}/return", { path, body: { comment } });
    } else {
      const signature_id = await signRecord(subject.type, subject.subject, String(form.get("pin") ?? ""));
      if (decision === "reject") {
        await serverApi.post("/approval-requests/{approval_request_id}/reject", { path, body: { comment, signature_id } });
      } else {
        await serverApi.post("/approval-requests/{approval_request_id}/approve", {
          path, body: { comment: comment || undefined, signature_id },
        });
      }
    }
  } catch (error) {
    return formError(error);
  }
  revalidatePath("/", "layout");
  redirect("/approvals");
}
