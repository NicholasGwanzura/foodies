import { humanise, tone } from "@/lib/format";

export function StatusPill({ status }: { status: string }) {
  return <span className={`pill ${tone(status)}`}>{humanise(status)}</span>;
}
