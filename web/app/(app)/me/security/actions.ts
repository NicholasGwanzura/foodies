"use server";

import { serverApi } from "@/lib/api/server";
import type { FormState } from "@/lib/form-state";
import { formError } from "@/lib/session";

export async function savePassword(
  _: FormState,
  form: FormData,
): Promise<FormState> {
  try {
    await serverApi.post("/auth/password/change", {
      body: {
        current_password: String(form.get("current_password") ?? ""),
        new_password: String(form.get("new_password") ?? ""),
      },
    });
    return { ok: "Password saved." };
  } catch (error) {
    return formError(error);
  }
}

export async function savePin(
  _: FormState,
  form: FormData,
): Promise<FormState> {
  try {
    await serverApi.post("/auth/pin/set", {
      body: {
        pin: String(form.get("pin") ?? ""),
        current_secret: String(form.get("current_secret") ?? ""),
      },
    });
    return { ok: "PIN saved." };
  } catch (error) {
    return formError(error);
  }
}
