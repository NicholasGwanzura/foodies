export type FieldKind =
  | "text" | "textarea" | "number" | "decimal" | "money" | "currency" | "date"
  | "datetime" | "boolean" | "enum" | "ref" | "uuid" | "list" | "rows" | "json"
  | "pin" | "password";

export type FieldMeta = {
  name: string;
  label: string;
  kind: FieldKind;
  required?: boolean;
  readOnly?: boolean;
  hint?: string;
  options?: string[];
  /** Slug of the resource this field points to. */
  ref?: string;
  /** For `rows`: the columns of each row. */
  fields?: FieldMeta[];
};

export type Write = { path: string; permission: string; fields: FieldMeta[] };

export type ActionMeta = Write & {
  key: string;
  label: string;
  operationId: string;
  /** Statuses it can run from, when the workflow is declared. */
  from?: string[];
};

export type Resource = {
  slug: string;
  title: string;
  singular: string;
  about: string;
  group: string;
  tile?: string | null;
  listPath: string;
  getPath?: string;
  readPermission: string;
  fields: FieldMeta[];
  columns: string[];
  restricted: boolean;
  statuses?: string[];
  create?: Write;
  update?: Write;
  actions?: ActionMeta[];
  nestedUnder?: string | null;
};
