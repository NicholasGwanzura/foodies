"use server";

import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";

import { serverApi } from "@/lib/api/server";
import type { FormState } from "@/lib/form-state";
import { formError, getSession } from "@/lib/session";

type Subject = { id: string; type: string; subject: string };

/** Signing is what makes a decision: the PIN confirms this version of the record. */
async function signFor(subject: Subject, pin: string) {
  const { me } = await getSession();
  const request = await serverApi.post("/signature-requests", {
    body: {
      subject_type: subject.type, subject_id: subject.subject, meaning: "approved",
      signer_kind: "employee", signer_employee_id: me.employee.id,
    },
  });
  const id = request?.id as string;
  await serverApi.post("/signature-requests/{signature_request_id}/sign", {
    path: { signature_request_id: id },
    body: { pin },
  });
  return id;
}

export async function decide(subject: Subject, _: FormState, form: FormData): Promise<FormState> {
  const decision = String(form.get("decision"));
  const comment = String(form.get("comment") ?? "").trim();
  const path = { approval_request_id: subject.id };
  try {
    if (decision === "return") {
      await serverApi.post("/approval-requests/{approval_request_id}/return", { path, body: { comment } });
    } else {
      const signature_id = await signFor(subject, String(form.get("pin") ?? ""));
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
