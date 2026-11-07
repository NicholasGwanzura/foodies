"use client";

import { useActionState } from "react";

import { FormNotice, Submit } from "@/components/form";

import { acknowledge } from "./actions";

export function Ack({ id }: { id: string }) {
  const [state, action] = useActionState(acknowledge.bind(null, id), {});
  return <form action={action} className="stack-s"><Submit size="small">Acknowledge</Submit><FormNotice state={state} /></form>;
}
