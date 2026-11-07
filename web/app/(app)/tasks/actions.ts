"use server";

import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";

import { serverApi } from "@/lib/api/server";
import type { FormState } from "@/lib/form-state";
import { formError } from "@/lib/session";

export async function claimTask(id: string, _: FormState): Promise<FormState> {
  try {
    await serverApi.post("/tasks/{task_id}/claim", { path: { task_id: id } });
  } catch (error) {
    return formError(error);
  }
  revalidatePath(`/tasks/${id}`);
  return { ok: "This task is yours." };
}

export async function completeTask(id: string, _: FormState): Promise<FormState> {
  try {
    await serverApi.post("/tasks/{task_id}/complete", { path: { task_id: id } });
  } catch (error) {
    return formError(error);
  }
  revalidatePath("/", "layout");
  redirect("/tasks");
}

export async function skipTask(id: string, _: FormState, form: FormData): Promise<FormState> {
  try {
    await serverApi.post("/tasks/{task_id}/skip", {
      path: { task_id: id },
      body: {
        reason: String(form.get("reason") ?? ""),
        alternative_control: String(form.get("alternative_control") ?? "") || undefined,
      },
    });
  } catch (error) {
    return formError(error);
  }
  revalidatePath("/", "layout");
  redirect("/tasks");
}
