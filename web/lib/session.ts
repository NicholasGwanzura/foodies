import "server-only";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import { cache } from "react";

import { ApiError, type Problem } from "./api/client";
import type { components } from "./api/schema";
import { ACCESS_COOKIE, REFRESH_COOKIE, serverApi } from "./api/server";
import type { FormState } from "./form-state";

export type Me = components["schemas"]["Me"];

const cookieOptions = {
  httpOnly: true,
  sameSite: "lax" as const,
  secure: process.env.NODE_ENV === "production",
  path: "/",
};

const THIRTY_DAYS = 60 * 60 * 24 * 30;

export async function saveTokens(pair: components["schemas"]["TokenPair"]) {
  const jar = await cookies();
  if (pair.access_token) {
    jar.set(ACCESS_COOKIE, pair.access_token, {
      ...cookieOptions,
      maxAge: THIRTY_DAYS,
    });
  }
  if (pair.refresh_token) {
    jar.set(REFRESH_COOKIE, pair.refresh_token, {
      ...cookieOptions,
      maxAge: THIRTY_DAYS,
    });
  }
}

export async function clearTokens() {
  const jar = await cookies();
  jar.delete(ACCESS_COOKIE);
  jar.delete(REFRESH_COOKIE);
}

function goToSignIn(): never {
  redirect("/sign-in");
}

/** Sends the user to whichever step of signing in they have not finished. */
function finishSignIn(me: Me): never {
  if (me.user.credential_status === "must_change") {
    redirect("/account/password");
  }
  redirect("/account/two-factor");
}

export type Session = { me: Me; permissions: string[] };

export const getSession = cache(async (): Promise<Session> => {
  let me: Me;
  try {
    me = (await serverApi.get("/me")) as Me;
  } catch (error) {
    if (error instanceof ApiError && error.problem.status === 401) goToSignIn();
    throw error;
  }
  try {
    const held = await serverApi.get("/me/permissions");
    return {
      me,
      permissions: (held?.permissions ?? []).map((p) => p.key),
    };
  } catch (error) {
    if (error instanceof ApiError && error.problem.code === "mfa_required") {
      finishSignIn(me);
    }
    if (error instanceof ApiError && error.problem.status === 401) goToSignIn();
    throw error;
  }
});

/** Mirrors the server's matching, to decide what to show. The API still checks. */
export function matches(pattern: string, key: string): boolean {
  if (pattern === "*" || pattern === key) return true;
  if (pattern.startsWith("*.")) return key.endsWith(pattern.slice(1));
  const want = pattern.split(".");
  const have = key.split(".");
  for (let i = 0; i < want.length; i++) {
    if (i === want.length - 1 && want[i] === "*") return have.length > i;
    if (i >= have.length || (want[i] !== "*" && want[i] !== have[i])) {
      return false;
    }
  }
  return want.length === have.length;
}

export const can = (session: Session, permission: string) =>
  session.permissions.some((p) => matches(p, permission));

/** Turns an API failure into something a form can show. */
export function formError(error: unknown): FormState {
  if (!(error instanceof ApiError)) throw error;
  const problem: Problem = error.problem;
  const fields: Record<string, string> = {};
  for (const issue of problem.errors ?? []) {
    fields[issue.pointer.replace(/^\//, "")] = issue.message;
  }
  return {
    error: problem.detail ?? problem.title,
    fields,
    blockers: problem.blockers?.map((b) => b.message),
  };
}
