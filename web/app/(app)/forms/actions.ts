"use server";

import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";

import { serverApi } from "@/lib/api/server";
import type { FormState } from "@/lib/form-state";
import { formError } from "@/lib/session";

type Input = {
  version: string;
  site: string;
  task?: string;
  id?: string;
  rowVersion?: number;
  answers: Record<string, unknown>;
};

function today() {
  return new Intl.DateTimeFormat("en-CA", { timeZone: "Africa/Harare" }).format(new Date());
}

/** Saves the answers, then submits them, so one tap does both. */
export async function submitForm(input: Input): Promise<FormState> {
  let id = input.id;
  let rowVersion = input.rowVersion;
  const answers = input.answers as Record<string, never>;
  try {
    if (id) {
      const saved = await serverApi.patch("/form-submissions/{form_submission_id}", {
        path: { form_submission_id: id },
        headers: { "If-Match": String(rowVersion) },
        body: { answers },
      });
      rowVersion = saved?.row_version;
    } else {
      const created = await serverApi.post("/form-submissions", {
        body: {
          template_version_id: input.version, site_id: input.site, business_date: today(), task_id: input.task,
          captured_at: new Date().toISOString(), answers,
        },
      });
      id = created?.id as string;
      rowVersion = created?.row_version;
    }
    await serverApi.post("/form-submissions/{form_submission_id}/submit", { path: { form_submission_id: id } });
  } catch (error) {
    // The draft now exists, so a retry has to update it, not add another.
    return { ...formError(error), values: id ? { id, rowVersion: String(rowVersion) } : undefined };
  }
  revalidatePath("/", "layout");
  redirect(`/forms/${id}`);
}
