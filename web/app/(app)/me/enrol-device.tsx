import { cookies } from "next/headers";
import { revalidatePath } from "next/cache";

import { Submit } from "@/components/form";
import { serverApi } from "@/lib/api/server";
import { getSession } from "@/lib/session";

async function enrol() {
  "use server";
  const { me } = await getSession();
  const device = await serverApi.post("/devices", {
    body: { label: `${me.employee.site_id} device`, site_id: me.employee.site_id, platform: "web", shared: true },
  });
  (await cookies()).set("foodies_device", String(device?.id), {
    httpOnly: true, sameSite: "lax", secure: process.env.NODE_ENV === "production", path: "/", maxAge: 60 * 60 * 24 * 365,
  });
  revalidatePath("/me");
}

export async function EnrolDevice() {
  const enrolled = (await cookies()).has("foodies_device");
  return (
    <section className="panel stack">
      <h2>This device</h2>
      {enrolled ? (
        <p className="muted">Enrolled for offline checklists.</p>
      ) : (
        <>
          <p className="muted">Enrol it so branch staff can capture checklists without a connection.</p>
          <form action={enrol}><Submit variant="ghost">Enrol this device</Submit></form>
        </>
      )}
    </section>
  );
}
