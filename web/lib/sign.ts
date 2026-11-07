import "server-only";

import { serverApi } from "./api/server";
import { getSession } from "./session";

/**
 * Signs a record as the signed-in user and returns the signature id. A wrong
 * PIN withdraws the request, because an open one would block closing the record.
 */
export async function signRecord(type: string, id: string, pin: string, meaning: "approved" | "acknowledged" = "approved") {
  const { me } = await getSession();
  const request = (await serverApi.post("/signature-requests", {
    body: {
      subject_type: type, subject_id: id, meaning,
      signer_kind: "employee", signer_employee_id: me.employee.id,
    },
  })) as { id: string };
  const path = { signature_request_id: request.id };
  try {
    await serverApi.post("/signature-requests/{signature_request_id}/sign", { path, body: { pin } });
  } catch (error) {
    await serverApi.post("/signature-requests/{signature_request_id}/decline", {
      path,
      body: { reason: "PIN not confirmed" },
    });
    throw error;
  }
  return request.id;
}
