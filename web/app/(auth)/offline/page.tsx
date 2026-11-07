import type { Metadata } from "next";

export const metadata: Metadata = { title: "Offline" };

export default function OfflinePage() {
  return (
    <div className="stack">
      <h1>You are offline</h1>
      <p className="muted">
        This page needs a connection. Checklists you have already opened still
        work, and anything you submit is sent when you are back online.
      </p>
    </div>
  );
}
