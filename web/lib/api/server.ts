import "server-only";
import { cookies } from "next/headers";

import { createApi } from "./client";

export const ACCESS_COOKIE = "foodies_at";
export const REFRESH_COOKIE = "foodies_rt";

/** The API as the signed-in user. Tokens stay in httpOnly cookies. */
export const serverApi = createApi(async () => {
  const token = (await cookies()).get(ACCESS_COOKIE)?.value;
  const headers: Record<string, string> = {};
  if (token) headers.Authorization = `Bearer ${token}`;
  return headers;
});
