import assert from "node:assert/strict";
import { describe, it } from "node:test";

import { dropFromQueue, enqueue, readQueue, type QueuedForm } from "./offline-queue.ts";

const memory = () => {
  const data = new Map<string, string>();
  return { getItem: (k: string) => data.get(k) ?? null, setItem: (k: string, v: string) => void data.set(k, v) };
};

const form = (id: string): QueuedForm => ({
  clientRecordId: id, capturedAt: "2026-11-07T06:00:00Z", version: "v", site: "s", businessDate: "2026-11-07", answers: {},
});

describe("offline queue", () => {
  it("keeps items until they are dropped", () => {
    const store = memory();
    enqueue(store, form("a"));
    enqueue(store, form("b"));
    dropFromQueue(store, ["a"]);
    assert.deepEqual(readQueue(store).map((q) => q.clientRecordId), ["b"]);
  });

  it("survives a corrupt store", () => {
    const store = memory();
    store.setItem("foodies.queue", "not json");
    assert.deepEqual(readQueue(store), []);
  });
});
