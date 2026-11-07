"use client";

export default function ErrorPage({ error, reset }: { error: Error; reset: () => void }) {
  return (
    <section className="panel empty">
      <strong>That did not load</strong>
      <p>{error.message.includes("role") ? error.message : "Something went wrong on our side. Try again."}</p>
      <button type="button" className="btn ghost" onClick={reset}>Try again</button>
    </section>
  );
}
