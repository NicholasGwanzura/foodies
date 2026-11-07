"use server";

import { revalidatePath } from "next/cache";

import { serverApi } from "@/lib/api/server";

export async function readAll() {
  await serverApi.post("/notifications/read-all");
  revalidatePath("/", "layout");
}
