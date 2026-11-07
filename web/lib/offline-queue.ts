/** A checklist waiting to send. Stored on this device until it syncs. */
export type QueuedForm = {
  clientRecordId: string;
  capturedAt: string;
  version: string;
  site: string;
  task?: string;
  businessDate: string;
  answers: Record<string, unknown>;
};

const KEY = "foodies.queue";

type Store = Pick<Storage, "getItem" | "setItem">;

export function readQueue(store: Store): QueuedForm[] {
  try {
    return JSON.parse(store.getItem(KEY) ?? "[]") as QueuedForm[];
  } catch {
    return [];
  }
}

export function enqueue(store: Store, item: QueuedForm) {
  store.setItem(KEY, JSON.stringify([...readQueue(store), item]));
}

export function dropFromQueue(store: Store, ids: string[]) {
  store.setItem(KEY, JSON.stringify(readQueue(store).filter((q) => !ids.includes(q.clientRecordId))));
}
