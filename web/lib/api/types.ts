import type { components } from "./schema";

type Schemas = components["schemas"];

/** Names for generated schemas, used where a list's items need naming. */
export type Task = Schemas["Task"];
export type ApprovalRequest = Schemas["ApprovalRequest"];
export type Escalation = Schemas["Escalation"];
export type Notification = Schemas["Notification"];
export type Attachment = Schemas["Attachment"];
export type Comment = Schemas["Comment"];
export type Site = Schemas["Site"];

/**
 * The spec composes each list with a shared envelope whose `data` is untyped, so
 * the generated type loses the item type. This puts it back.
 */
export const items = <T>(list?: { data: unknown } | null): T[] =>
  (list?.data ?? []) as T[];
