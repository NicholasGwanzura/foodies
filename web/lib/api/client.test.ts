import assert from "node:assert/strict";
import { afterEach, describe, it, mock } from "node:test";

import { ApiError, api } from "./client.ts";

const json = (status: number, body: unknown, contentType = "application/json") =>
  new Response(JSON.stringify(body), {
    status,
    headers: { "Content-Type": contentType },
  });

const stubFetch = (response: Response) =>
  mock.method(globalThis, "fetch", async () => response);

describe("api", () => {
  afterEach(() => mock.restoreAll());

  it("fills path parameters and skips empty query values", async () => {
    const fetchMock = stubFetch(new Response(null, { status: 204 }));

    await api.delete("/me/sessions/{session_id}", {
      path: { session_id: "a/b" },
    });
    assert.equal(
      String(fetchMock.mock.calls[0].arguments[0]),
      "http://localhost:8080/api/v1/me/sessions/a%2Fb",
    );

    await api.get("/tasks", {
      query: { limit: 10, cursor: undefined },
    });
    assert.equal(
      String(fetchMock.mock.calls[1].arguments[0]),
      "http://localhost:8080/api/v1/tasks?limit=10",
    );
  });

  it("sends JSON bodies and headers", async () => {
    const fetchMock = stubFetch(json(200, {}));

    await api.post("/auth/login", {
      body: { employee_number: "E1", password: "test-only" },
      headers: { "X-Test": "1" },
    });

    const init = fetchMock.mock.calls[0].arguments[1] as RequestInit;
    const headers = new Headers(init.headers);
    assert.equal(init.method, "POST");
    assert.equal(headers.get("Content-Type"), "application/json");
    assert.equal(headers.get("X-Test"), "1");
    assert.deepEqual(JSON.parse(init.body as string), {
      employee_number: "E1",
      password: "test-only",
    });
  });

  it("returns undefined for an empty response", async () => {
    stubFetch(new Response(null, { status: 204 }));

    const result = await api.delete("/me/sessions/{session_id}", {
      path: { session_id: "1" },
    });
    assert.equal(result, undefined);
  });

  it("throws ApiError carrying the problem", async () => {
    const problem = {
      type: "/problems/not-found",
      title: "Not found",
      status: 404,
      code: "not_found",
      trace_id: "T1",
    };
    stubFetch(json(404, problem, "application/problem+json"));

    await assert.rejects(
      api.get("/me"),
      (error: unknown) =>
        error instanceof ApiError &&
        error.problem.code === "not_found" &&
        error.problem.trace_id === "T1",
    );
  });

  it("throws ApiError with internal_error for a non-problem failure", async () => {
    stubFetch(new Response("bad gateway", { status: 502 }));

    await assert.rejects(
      api.get("/me"),
      (error: unknown) =>
        error instanceof ApiError &&
        error.problem.status === 502 &&
        error.problem.code === "internal_error",
    );
  });
});

// Compile-time checks, enforced by `pnpm typecheck`.
export const typeChecks = async () => {
  // @ts-expect-error unknown path
  await api.get("/does-not-exist");
  // @ts-expect-error /auth/login has no GET
  await api.get("/auth/login");
  // @ts-expect-error path parameters are required
  await api.delete("/me/sessions/{session_id}");
  // @ts-expect-error body is required
  await api.post("/auth/login");
  // @ts-expect-error password is required
  await api.post("/auth/login", { body: { employee_number: "E1" } });

  const me = await api.get("/me");
  // @ts-expect-error not a field on the response
  void me.nope;
};

describe("api writes", () => {
  afterEach(() => mock.restoreAll());

  it("sends array query values comma-separated", async () => {
    const fetchMock = stubFetch(json(200, {}));
    await api.request("get", "/tasks", { query: { status: ["open", "done"] } });
    assert.equal(
      String(fetchMock.mock.calls[0].arguments[0]),
      "http://localhost:8080/api/v1/tasks?status=open%2Cdone",
    );
  });

  it("adds an Idempotency-Key to record writes, but not to sign-in", async () => {
    const fetchMock = mock.method(globalThis, "fetch", async () => json(200, {}));
    await api.request("post", "/complexes", { body: {} });
    await api.request("post", "/auth/login", { body: {} });
    const header = (i: number) =>
      new Headers(
        (fetchMock.mock.calls[i].arguments[1] as RequestInit).headers,
      ).get("Idempotency-Key");
    assert.ok(header(0));
    assert.equal(header(1), null);
  });

  it("keeps a caller's Idempotency-Key", async () => {
    const fetchMock = stubFetch(json(200, {}));
    await api.request("post", "/complexes", {
      body: {},
      headers: { "Idempotency-Key": "form-1" },
    });
    assert.equal(
      new Headers(
        (fetchMock.mock.calls[0].arguments[1] as RequestInit).headers,
      ).get("Idempotency-Key"),
      "form-1",
    );
  });
});
