import assert from "node:assert/strict";
import { describe, it } from "node:test";

import { coerce } from "./coerce.ts";
import type { FieldMeta } from "./resource-types.ts";

const from = (values: Record<string, string>) => (k: string) => values[k];

describe("coerce", () => {
  const fields: FieldMeta[] = [
    { name: "name", label: "Name", kind: "text", required: true },
    { name: "count", label: "Count", kind: "number" },
    { name: "active", label: "Active", kind: "boolean" },
    { name: "limit", label: "Limit", kind: "money" },
    { name: "tags", label: "Tags", kind: "list" },
  ];

  it("types values and folds money", () => {
    const { body, errors } = coerce(
      fields,
      from({ name: " A ", count: "3", active: "on", "limit.amount": "10.50", "limit.currency": "ZWG", tags: "x, y" }),
    );
    assert.deepEqual(errors, {});
    assert.deepEqual(body, {
      name: "A", count: 3, active: true,
      limit: { amount: "10.50", currency: "ZWG" }, tags: ["x", "y"],
    });
  });

  it("leaves out empty optional fields, and an unticked box is false", () => {
    const { body } = coerce(fields, from({ name: "A" }));
    assert.deepEqual(body, { name: "A", active: false });
  });

  it("reports required and malformed fields", () => {
    const { errors } = coerce(fields, from({ count: "many" }));
    assert.equal(errors.name, "This field is required.");
    assert.equal(errors.count, "Enter a number.");
  });
});
