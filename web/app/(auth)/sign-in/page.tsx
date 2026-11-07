import type { Metadata } from "next";
import { Suspense } from "react";

import { SignInForm } from "./sign-in-form";

export const metadata: Metadata = { title: "Sign in" };

async function Notice({ searchParams }: { searchParams: Promise<{ done?: string }> }) {
  const { done } = await searchParams;
  if (!done) return null;
  return <p className="notice good" role="status">All set. Sign in with your new details.</p>;
}

export default function SignInPage({ searchParams }: { searchParams: Promise<{ done?: string }> }) {
  return (
    <>
      <Suspense fallback={null}>
        <Notice searchParams={searchParams} />
      </Suspense>
      <SignInForm />
    </>
  );
}
