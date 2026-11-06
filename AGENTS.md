# Project overview
Go REST API (`/api`) + Next.js App Router frontend (`/web`).

## Source of truth
- `api/openapi.yaml` defines the API. Never edit generated files.
- Do NOT change openapi.yaml without asking me first. If a change is
  needed, propose the diff and wait for approval.
- After any spec or SQL change, run `make gen`.

## Backend conventions
- Handlers in internal/http are thin: parse, call service, map errors.
- Business logic lives in internal/service; DB access only via sqlc.
- Errors: wrap with %w, map to the error schema defined in the spec.
- Table-driven tests; integration tests use testcontainers.
- context.Context is the first parameter of anything doing I/O.

## Frontend conventions
- Use generated types from lib/api/schema.d.ts; no hand-written API types.
- Server components by default; "use client" only when needed.
- All API calls go through lib/api/client.ts.

## Workflow
- Work in vertical slices (one resource at a time).
- Write tests first, then implement, then run `make check`.
- Commit after each passing slice with a conventional commit message.
- Do not add dependencies without asking.