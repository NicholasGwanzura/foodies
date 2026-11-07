import "server-only";
import { notFound } from "next/navigation";

import { ApiError } from "./api/client";
import { serverApi } from "./api/server";
import { recordLabel } from "./format";
import type { FieldMeta, Resource } from "./resource-types";
import { RESOURCES } from "./resources.gen";

export type Row = Record<string, unknown> & { id: string };
export type ListResult = {
  data: Row[];
  page: { limit: number; has_more: boolean; next_cursor?: string | null };
};
export type RefOptions = Record<string, { value: string; label: string }[]>;

export function resourceOr404(slug: string): Resource {
  const resource = RESOURCES[slug];
  if (!resource) notFound();
  return resource;
}

/** Fills `{id}` style path parameters: the last is the record, earlier ones its parents. */
export function fillPath(path: string, id?: string, parent?: string): string {
  const names = [...path.matchAll(/\{(\w+)\}/g)].map((m) => m[0]);
  let out = path;
  if (names.length === 1) out = out.replace(names[0], id ?? parent ?? "");
  if (names.length === 2) {
    out = out.replace(names[0], parent ?? "").replace(names[1], id ?? "");
  }
  return out;
}

export async function listRows(
  resource: Resource,
  query: Record<string, unknown> = {},
  parent?: string,
): Promise<ListResult> {
  return (await serverApi.request("get", fillPath(resource.listPath, undefined, parent), {
    query,
  })) as ListResult;
}

/** Id to label for every ref field, so tables and forms show names, not uuids. */
export async function loadRefs(fields: FieldMeta[]): Promise<RefOptions> {
  const slugs = [...new Set(allFields(fields).flatMap((f) => (f.ref ? [f.ref] : [])))];
  const entries = await Promise.all(
    slugs.map(async (slug) => {
      const target = RESOURCES[slug];
      if (!target) return [slug, []] as const;
      try {
        const { data } = await listRows(target, { limit: 100 });
        return [slug, data.map((r) => ({ value: r.id, label: recordLabel(r) }))] as const;
      } catch (error) {
        // A role may fill a form without being able to browse what it links to.
        if (error instanceof ApiError && error.problem.status < 500) return [slug, []] as const;
        throw error;
      }
    }),
  );
  return Object.fromEntries(entries);
}

function allFields(fields: FieldMeta[]): FieldMeta[] {
  return fields.flatMap((f) => [f, ...(f.fields ?? [])]);
}

export async function loadCurrencies(): Promise<string[]> {
  try {
    const res = await serverApi.get("/currencies", { query: { limit: 20 } });
    const rows = (res?.data ?? []) as { code: string }[];
    return rows.length > 0 ? rows.map((c) => c.code) : ["USD", "ZWG"];
  } catch {
    return ["USD", "ZWG"];
  }
}
