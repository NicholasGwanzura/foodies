"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";

import { api } from "@/lib/api/client";
import { serverApi } from "@/lib/api/server";
import type { FormState } from "@/lib/form-state";
import { clearTokens, formError, saveTokens } from "@/lib/session";

const CHALLENGE_COOKIE = "foodies_mfa";

export async function signIn(
  _: FormState,
  form: FormData,
): Promise<FormState> {
  const employee_number = String(form.get("employee_number") ?? "").trim();
  const password = String(form.get("password") ?? "");
  if (!employee_number || !password) {
    return { error: "Enter your employee number and password." };
  }

  let destination = "/";
  try {
    const pair = await api.post("/auth/login", {
      body: { employee_number, password },
    });
    if (pair?.status === "mfa_required" && pair.mfa_challenge_id) {
      (await cookies()).set(CHALLENGE_COOKIE, pair.mfa_challenge_id, {
        httpOnly: true,
        sameSite: "lax",
        secure: process.env.NODE_ENV === "production",
        maxAge: 300,
        path: "/",
      });
      destination = "/sign-in/verify";
    } else if (pair) {
      await saveTokens(pair);
      if (pair.status === "password_change_required") {
        destination = "/account/password";
      }
    }
  } catch (error) {
    return { ...formError(error), values: { employee_number } };
  }
  redirect(destination);
}

export async function verifyCode(
  _: FormState,
  form: FormData,
): Promise<FormState> {
  const jar = await cookies();
  const challenge = jar.get(CHALLENGE_COOKIE)?.value;
  if (!challenge) redirect("/sign-in");
  try {
    const pair = await api.post("/auth/mfa/verify", {
      body: {
        mfa_challenge_id: challenge,
        code: String(form.get("code") ?? "").trim(),
      },
    });
    if (pair) await saveTokens(pair);
  } catch (error) {
    return formError(error);
  }
  jar.delete(CHALLENGE_COOKIE);
  redirect("/");
}

export async function signOut() {
  try {
    await serverApi.post("/auth/logout");
  } catch {
    // The cookies go either way, and the session expires on its own.
  }
  await clearTokens();
  redirect("/sign-in");
}
