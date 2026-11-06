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
    if (value !== undefined && value !== null) {
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

async function send(
  method: Method,
  path: string,
  options: RawOptions = {},
): Promise<unknown> {
  const { path: pathParams, query, body, headers, signal } = options;

  const requestHeaders = new Headers(headers);
  if (body !== undefined) {
    requestHeaders.set("Content-Type", "application/json");
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

const forMethod = <M extends Method>(method: M) =>
  ((path: string, options?: RawOptions) =>
    send(method, path, options)) as unknown as MethodClient<M>;

/** Every API call goes through here. Failures throw `ApiError`. */
export const api = {
  get: forMethod("get"),
  post: forMethod("post"),
  put: forMethod("put"),
  patch: forMethod("patch"),
  delete: forMethod("delete"),
};
