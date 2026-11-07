"use client";

import { useActionState } from "react";

import { addComment, uploadFile } from "@/app/(app)/r/evidence-actions";
import { FormNotice, Submit } from "@/components/form";

export function EvidenceForms({ type, id }: { type: string; id: string }) {
  const [comment, commentAction] = useActionState(addComment.bind(null, type, id), {});
  const [upload, uploadAction] = useActionState(uploadFile.bind(null, type, id), {});
  return (
    <div className="form-grid">
      <form action={uploadAction} className="stack-s">
        <label className="field-label" htmlFor="ev-file">Add a photo or document</label>
        <input id="ev-file" type="file" name="file" accept="image/*,application/pdf" required />
        <FormNotice state={upload} />
        <div><Submit variant="ghost" size="small" pendingText="Uploading…">Upload</Submit></div>
      </form>
      <form action={commentAction} className="stack-s">
        <label className="field-label" htmlFor="ev-comment">Add a comment</label>
        <input id="ev-comment" name="body" required />
        <FormNotice state={comment} />
        <div><Submit variant="ghost" size="small">Post comment</Submit></div>
      </form>
    </div>
  );
}
