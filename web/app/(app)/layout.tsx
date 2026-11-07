import { Suspense } from "react";

import { Shell } from "@/components/shell";

// Every page here reads the signed-in session, so none can be instant.
export const instant = false;

function Loading() {
  return (
    <>
      <header className="arch">
        <div className="arch-bar">
          <span className="logo">foodies</span>
        </div>
      </header>
      <div className="frame" style={{ display: "block" }}>
        <div className="skeleton" style={{ height: 320, borderRadius: 22 }} />
      </div>
    </>
  );
}

export default function AppLayout({ children }: LayoutProps<"/">) {
  return (
    <Suspense fallback={<Loading />}>
      <Shell>{children}</Shell>
    </Suspense>
  );
}
