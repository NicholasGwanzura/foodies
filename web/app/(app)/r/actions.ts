"use server";

import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";

import { serverApi } from "@/lib/api/server";
import { coerce } from "@/lib/coerce";
import type { FormState } from "@/lib/form-state";
import { fillPath, resourceOr404, type Row } from "@/lib/resources";
import { formError } from "@/lib/session";
import { signRecord } from "@/lib/sign";

const reader = (form: FormData) => (name: string) => {
  const value = form.get(name);
  return typeof value === "string" ? value : undefined;
};

export async function createRecord(
  slug: string,
  parent: string | undefined,
  _: FormState,
  form: FormData,
): Promise<FormState> {
  const resource = resourceOr404(slug);
  if (!resource.create) return { error: "This record cannot be added here." };
  const { body, errors } = coerce(resource.create.fields, reader(form));
  if (Object.keys(errors).length > 0) {
    return { error: "Some fields need attention.", fields: errors };
  }
  let row: Row;
  try {
    row = (await serverApi.request("post", fillPath(resource.create.path, undefined, parent), { body })) as Row;
  } catch (error) {
    return formError(error);
  }
  const query = parent ? `?parent=${parent}` : "";
  redirect(resource.getPath ? `/r/${slug}/${row.id}${query}` : `/r/${slug}${query}`);
}

export async function updateRecord(
  slug: string,
  id: string,
  parent: string | undefined,
  version: number,
  _: FormState,
  form: FormData,
): Promise<FormState> {
  const resource = resourceOr404(slug);
  if (!resource.update) return { error: "This record cannot be changed." };
  const { body, errors } = coerce(resource.update.fields, reader(form));
  if (Object.keys(errors).length > 0) {
    return { error: "Some fields need attention.", fields: errors };
  }
  try {
    await serverApi.request("patch", fillPath(resource.update.path, id, parent), {
      body,
      headers: { "If-Match": String(version) },
    });
  } catch (error) {
    return formError(error);
  }
  redirect(`/r/${slug}/${id}${parent ? `?parent=${parent}` : ""}`);
}

export async function runAction(
  slug: string,
  id: string,
  parent: string | undefined,
  key: string,
  _: FormState,
  form: FormData,
): Promise<FormState> {
  const resource = resourceOr404(slug);
  const action = resource.actions?.find((a) => a.key === key);
  if (!action) return { error: "This action is not available." };
  const { body, errors } = coerce(action.fields, reader(form));
  if (Object.keys(errors).length > 0) {
    return { error: "Some fields need attention.", fields: errors };
  }
  try {
    // A signed decision: the PIN confirms this exact version of the record.
    if (typeof body.signature_id === "string") {
      body.signature_id = await signRecord(slug.replaceAll("-", "_"), id, body.signature_id);
    }
    await serverApi.request("post", fillPath(action.path, id, parent), {
      body: action.fields.length > 0 ? body : undefined,
    });
  } catch (error) {
    return formError(error);
  }
  revalidatePath("/", "layout");
  return { ok: `${action.label}: done.` };
}
