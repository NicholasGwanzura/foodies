import Link from "next/link";

export default function NotFound() {
  return (
    <section className="panel empty">
      <strong>That page does not exist</strong>
      <Link className="btn ghost" href="/">Back to My Day</Link>
    </section>
  );
}
