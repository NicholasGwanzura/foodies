"use client";

import { useEffect, useState } from "react";

import { dropFromQueue, readQueue } from "@/lib/offline-queue";

const EVERY = 30_000;

/** Registers the service worker, and sends queued work when it can. */
export function OfflineSync() {
  const [waiting, setWaiting] = useState(0);
  const [blocked, setBlocked] = useState(false);

  useEffect(() => {
    // Only in production: dev chunks change on every compile, so a cached copy
    // would break hydration.
    if (process.env.NODE_ENV === "production" && "serviceWorker" in navigator) {
      navigator.serviceWorker.register("/sw.js").catch(() => undefined);
    }

    async function flush() {
      const queue = readQueue(localStorage);
      setWaiting(queue.length);
      if (queue.length === 0 || !navigator.onLine) return;
      const res = await fetch("/sync", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(queue),
      }).catch(() => undefined);
      if (res?.status === 409) setBlocked(true);
      if (res?.ok) {
        const { delivered } = (await res.json()) as { delivered: string[] };
        dropFromQueue(localStorage, delivered);
        setWaiting(readQueue(localStorage).length);
      }
    }

    void flush();
    const timer = setInterval(() => void flush(), EVERY);
    window.addEventListener("online", flush);
    return () => {
      clearInterval(timer);
      window.removeEventListener("online", flush);
    };
  }, []);

  if (waiting === 0) return null;
  return (
    <p className="queue-note" role="status">
      {waiting} checklist{waiting > 1 ? "s" : ""} waiting to send
      {blocked ? ". This device is not enrolled, so ask ICT." : "."}
    </p>
  );
}
