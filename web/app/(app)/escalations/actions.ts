"use server";

import { revalidatePath } from "next/cache";

import { serverApi } from "@/lib/api/server";
import type { FormState } from "@/lib/form-state";
import { formError } from "@/lib/session";

export async function acknowledge(id: string, _: FormState): Promise<FormState> {
  try {
    await serverApi.post("/escalations/{escalation_id}/acknowledge", { path: { escalation_id: id }, body: {} });
  } catch (error) {
    return formError(error);
  }
  revalidatePath("/", "layout");
  return { ok: "Acknowledged." };
}
