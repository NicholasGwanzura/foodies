import type { components, paths } from "./schema";

export type Problem = components["schemas"]["Problem"];

const baseUrl =
  process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080/api/v1";

type Method = "get" | "post" | "put" | "patch" | "delete";
type ApiPath = keyof paths;

/** Paths that define `M`. The schema marks the others `?: never`. */
type PathWith<M extends Method> = {
  [P in ApiPath]: paths[P][M] extends { responses: unknown } ? P : never;
}[ApiPath];

type Operation<M extends Method, P> = P extends ApiPath
  ? Extract<paths[P][M], { responses: unknown }>
  : never;

type PathPart<O> = O extends { parameters: { path: infer T } }
  ? { path: T }
  : { path?: never };

type QueryPart<O> = O extends { parameters: { query: infer T } }
  ? { query: T }
  : O extends { parameters: { query?: infer T } }
    ? [NonNullable<T>] extends [never]
      ? { query?: never }
      : { query?: NonNullable<T> }
    : { query?: never };

type BodyPart<O> = O extends {
  requestBody: { content: { "application/json": infer T } };
}
  ? { body: T }
  : { body?: never };

/** Only accepts what the operation does, and requires what the spec requires. */
export type RequestOptions<O> = PathPart<O> &
  QueryPart<O> &
  BodyPart<O> & {
    headers?: HeadersInit;
    signal?: AbortSignal;
  };

/** Union of the JSON bodies on the operation's 2xx responses. */
type SuccessBody<O> = O extends { responses: infer R }
  ? {
      [S in keyof R & (200 | 201 | 202 | 204)]: R[S] extends {
        content: { "application/json": infer T };
      }
        ? T
        : undefined;
    }[keyof R & (200 | 201 | 202 | 204)]
  : never;

/** RequestOptions once the generics are spent, for building the request. */
type RawOptions = {
  path?: Record<string, unknown>;
  query?: Record<string, unknown>;
  body?: unknown;
  headers?: HeadersInit;
  signal?: AbortSignal;
};

export class ApiError extends Error {
  readonly problem: Problem;

  constructor(problem: Problem) {
    super(problem.detail ?? problem.title);
    this.name = "ApiError";
    this.problem = problem;
  }
}

function buildUrl(
  path: string,
  pathParams?: Record<string, unknown>,
  query?: Record<string, unknown>,
): string {
  const filled = path.replace(/\{(\w+)\}/g, (_, name: string) =>
    encodeURIComponent(String(pathParams?.[name])),
  );

  const search = new URLSearchParams();
  for (const [key, value] of Object.entries(query ?? {})) {
    if (Array.isArray(value)) {
      // The spec's array parameters are `explode: false`.
      if (value.length > 0) search.set(key, value.join(","));
    } else if (value !== undefined && value !== null) {
      search.set(key, String(value));
    }
  }

  const suffix = search.size > 0 ? `?${search}` : "";
  return `${baseUrl}${filled}${suffix}`;
}

async function toApiError(response: Response): Promise<ApiError> {
  const contentType = response.headers.get("Content-Type") ?? "";
  if (contentType.startsWith("application/problem+json")) {
    return new ApiError((await response.json()) as Problem);
  }

  // Something in front of the API answered, so there is no Problem to read.
  return new ApiError({
    type: "about:blank",
    title: response.statusText || "Request failed",
    status: response.status,
    code: "internal_error",
    trace_id: response.headers.get("X-Request-Id") ?? "",
  });
}

/** Writes the API exempts from Idempotency-Key (docs/spec.md, 12.6). */
const noIdempotencyKey = /^\/(auth|webhooks)\/|^\/notifications\/read-all$/;

async function send(
  method: Method,
  path: string,
  options: RawOptions = {},
  defaults?: HeadersProvider,
): Promise<unknown> {
  const { path: pathParams, query, body, headers, signal } = options;

  const requestHeaders = new Headers(await defaults?.());
  new Headers(headers).forEach((value, key) => requestHeaders.set(key, value));
  if (body !== undefined) {
    requestHeaders.set("Content-Type", "application/json");
  }
  if (
    method === "post" &&
    !noIdempotencyKey.test(path) &&
    !requestHeaders.has("Idempotency-Key")
  ) {
    requestHeaders.set("Idempotency-Key", crypto.randomUUID());
  }

  const response = await fetch(buildUrl(path, pathParams, query), {
    method: method.toUpperCase(),
    headers: requestHeaders,
    body: body === undefined ? undefined : JSON.stringify(body),
    signal,
  });

  if (!response.ok) {
    throw await toApiError(response);
  }

  const text = await response.text();
  return text === "" ? undefined : JSON.parse(text);
}

type Args<O> = object extends RequestOptions<O>
  ? [RequestOptions<O>?]
  : [RequestOptions<O>];

/** Fixing the method here is what keeps the full API inside compiler limits. */
type MethodClient<M extends Method> = <P extends PathWith<M>>(
  path: P,
  ...options: Args<Operation<M, P>>
) => Promise<SuccessBody<Operation<M, P>>>;

type HeadersProvider = () => HeadersInit | Promise<HeadersInit>;

const forMethod = <M extends Method>(method: M, defaults?: HeadersProvider) =>
  ((path: string, options?: RawOptions) =>
    send(method, path, options, defaults)) as unknown as MethodClient<M>;

/** A client that adds `defaults` to each request, such as the bearer token. */
export function createApi(defaults?: HeadersProvider) {
  return {
    get: forMethod("get", defaults),
    post: forMethod("post", defaults),
    put: forMethod("put", defaults),
    patch: forMethod("patch", defaults),
    delete: forMethod("delete", defaults),
    /**
     * For screens driven by a path known only at run time. Prefer the typed
     * methods anywhere the path is written in the code.
     */
    request: (method: Method, path: string, options?: RawOptions) =>
      send(method, path, options, defaults),
  };
}

/** Every API call goes through here. Failures throw `ApiError`. */
export const api = createApi();
