import { type NextRequest, NextResponse } from "next/server";

import type { components } from "./lib/api/schema";

const ACCESS_COOKIE = "foodies_at";
const REFRESH_COOKIE = "foodies_rt";
const PUBLIC_PATHS = ["/sign-in"];
/** Refresh this long before expiry, so a page never renders with a dead token. */
const EARLY_SECONDS = 30;

const baseUrl =
  process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080/api/v1";

type TokenPair = components["schemas"]["TokenPair"];

function expiresAt(token: string): number {
  try {
    const payload = token.split(".")[1];
    return JSON.parse(Buffer.from(payload, "base64url").toString()).exp ?? 0;
  } catch {
    return 0;
  }
}

/**
 * Refresh tokens are single use and a reused one ends the session, yet a page
 * load fires several requests at once. They share one refresh here.
 */
const inFlight = new Map<string, Promise<TokenPair | null>>();

function refresh(token: string): Promise<TokenPair | null> {
  const running = inFlight.get(token);
  if (running) return running;
  const call = fetch(`${baseUrl}/auth/refresh`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ refresh_token: token }),
  })
    .then(async (r) => (r.ok ? ((await r.json()) as TokenPair) : null))
    .catch(() => null);
  inFlight.set(token, call);
  setTimeout(() => inFlight.delete(token), 30_000);
  return call;
}

export async function proxy(request: NextRequest) {
  const { pathname } = request.nextUrl;
  const isPublic = PUBLIC_PATHS.some((p) => pathname.startsWith(p));
  const access = request.cookies.get(ACCESS_COOKIE)?.value;
  const refreshToken = request.cookies.get(REFRESH_COOKIE)?.value;
  const now = Date.now() / 1000;

  if (!access && !refreshToken) {
    if (isPublic) return NextResponse.next();
    return NextResponse.redirect(new URL("/sign-in", request.url));
  }

  if (refreshToken && (!access || expiresAt(access) - EARLY_SECONDS < now)) {
    const pair = await refresh(refreshToken);
    if (!pair?.access_token || !pair.refresh_token) {
      const response = isPublic
        ? NextResponse.next()
        : NextResponse.redirect(new URL("/sign-in", request.url));
      response.cookies.delete(ACCESS_COOKIE);
      response.cookies.delete(REFRESH_COOKIE);
      return response;
    }

    // Server components read cookies from the request, so update it as well.
    request.cookies.set(ACCESS_COOKIE, pair.access_token);
    request.cookies.set(REFRESH_COOKIE, pair.refresh_token);
    const headers = new Headers(request.headers);
    headers.set("cookie", request.cookies.toString());
    const response = NextResponse.next({ request: { headers } });
    const options = {
      httpOnly: true,
      sameSite: "lax",
      secure: process.env.NODE_ENV === "production",
      path: "/",
      maxAge: 60 * 60 * 24 * 30,
    } as const;
    response.cookies.set(ACCESS_COOKIE, pair.access_token, options);
    response.cookies.set(REFRESH_COOKIE, pair.refresh_token, options);
    return response;
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/((?!_next/|favicon.ico|icon|manifest).*)"],
};
