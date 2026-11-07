import Link from "next/link";

import { StatusPill } from "@/components/records/status-pill";
import { formatValue } from "@/lib/format";
import type { RefOptions, Row } from "@/lib/resources";
import type { FieldMeta, Resource } from "@/lib/resource-types";

function Cell({
  field,
  row,
  refs,
}: {
  field: FieldMeta;
  row: Row;
  refs: RefOptions;
}) {
  const value = row[field.name];
  if (field.name === "status" && typeof value === "string") {
    return <StatusPill status={value} />;
  }
  if (field.ref && typeof value === "string") {
    return refs[field.ref]?.find((o) => o.value === value)?.label ?? "";
  }
  return formatValue(field, value);
}

export function RecordTable({
  resource,
  rows,
  refs,
  parent,
}: {
  resource: Resource;
  rows: Row[];
  refs: RefOptions;
  parent?: string;
}) {
  const fields = resource.columns.flatMap((name) => {
    const field = resource.fields.find((f) => f.name === name);
    return field ? [field] : [];
  });
  const href = (row: Row) =>
    `/r/${resource.slug}/${row.id}${parent ? `?parent=${parent}` : ""}`;
  return (
    <div className="table-wrap">
      <table className="data cards">
        <thead>
          <tr>
            {fields.map((f) => (
              <th key={f.name} className={f.kind === "money" ? "money" : undefined}>
                {f.label}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {rows.map((row) => (
            <tr key={row.id}>
              {fields.map((f, i) => (
                <td key={f.name} data-label={f.label} className={f.kind === "money" ? "money" : undefined}>
                  {i === 0 && resource.getPath ? (
                    <Link href={href(row) as never}>
                      <Cell field={f} row={row} refs={refs} />
                    </Link>
                  ) : (
                    <Cell field={f} row={row} refs={refs} />
                  )}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
