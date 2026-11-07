"use server";

import { redirect } from "next/navigation";

import { serverApi } from "@/lib/api/server";
import type { FormState } from "@/lib/form-state";
import { formError } from "@/lib/session";

export async function changePassword(
  _: FormState,
  form: FormData,
): Promise<FormState> {
  const next = String(form.get("new_password") ?? "");
  if (next !== String(form.get("confirm") ?? "")) {
    return { fields: { confirm: "The two passwords do not match." } };
  }
  try {
    await serverApi.post("/auth/password/change", {
      body: {
        current_password: String(form.get("current_password") ?? ""),
        new_password: next,
      },
    });
  } catch (error) {
    return formError(error);
  }
  redirect("/");
}

export async function confirmTwoFactor(
  _: FormState,
  form: FormData,
): Promise<FormState> {
  try {
    await serverApi.post("/auth/mfa/confirm", {
      body: { code: String(form.get("code") ?? "").trim() },
    });
  } catch (error) {
    return formError(error);
  }
  redirect("/");
}
