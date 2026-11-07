import { Suspense } from "react";

export default function SetupLayout({ children }: LayoutProps<"/">) {
  return (
    <div className="auth-page">
      <header className="arch">
        <div className="logo">foodies</div>
        <p className="tagline">One more step before you start.</p>
      </header>
      <main className="panel auth-card">
        <Suspense fallback={<div className="skeleton" style={{ height: 240 }} />}>
          {children}
        </Suspense>
      </main>
    </div>
  );
}
