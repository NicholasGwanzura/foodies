# `make help` lists the targets. Run from the repository root.
.DEFAULT_GOAL := help

API_TOOL := go tool -modfile=tools/go.mod
WEB_DEPS := web/node_modules/.modules.yaml

# Matches compose.yaml. Override for any other database.
DATABASE_URL ?= postgres://foodies:foodies@localhost:55440/foodies?sslmode=disable
GOOSE := GOOSE_DRIVER=postgres GOOSE_DBSTRING='$(DATABASE_URL)' GOOSE_MIGRATION_DIR=db/migrations $(API_TOOL) goose

# Everything `make gen` writes. CI fails if a commit leaves these stale.
GENERATED := api/internal/db api/internal/http/apigen web/lib/api/schema.d.ts

.PHONY: help
help: ## List targets
	@grep -hE '^[a-z-]+:.*## ' $(MAKEFILE_LIST) | awk -F ':.*## ' '{printf "  %-16s %s\n", $$1, $$2}'

$(WEB_DEPS): web/package.json web/pnpm-lock.yaml
	cd web && pnpm install --frozen-lockfile
	@touch $@

# Code generation -------------------------------------------------------------

.PHONY: gen gen-api gen-sql gen-web gen-check
gen: gen-api gen-sql gen-web ## Regenerate all code from the spec and SQL

gen-api:
	cd api && $(API_TOOL) oapi-codegen -config codegen/oapi-codegen.yaml openapi.yaml

gen-sql:
	cd api && $(API_TOOL) sqlc generate

gen-web: $(WEB_DEPS)
	cd web && pnpm gen

gen-check: gen ## Fail if generated code does not match the committed files
	@git add --intent-to-add $(GENERATED)
	@git diff --exit-code --stat -- $(GENERATED) || \
		{ echo 'Generated code is stale. Run make gen and commit the result.'; exit 1; }

# Checks ----------------------------------------------------------------------

.PHONY: check check-api check-web
check: check-api check-web ## Lint, typecheck, test and build everything

check-api:
	cd api && go build ./...
	cd api && go vet ./...
	cd api && $(API_TOOL) golangci-lint run ./...
	cd api && $(API_TOOL) sqlc compile
	cd api && go test ./...

check-web: $(WEB_DEPS)
	cd web && pnpm lint
	cd web && pnpm typecheck
	cd web && pnpm test
	cd web && pnpm build

# Database --------------------------------------------------------------------

.PHONY: db-up db-down migrate-up migrate-down migrate-status migrate-new migrate-check
db-up: ## Start the local Postgres
	docker compose up --detach --wait

db-down: ## Stop the local Postgres, keeping its data
	docker compose down

migrate-up: ## Apply all pending migrations
	cd api && $(GOOSE) up

migrate-down: ## Roll back the latest migration
	cd api && $(GOOSE) down

migrate-status: ## Show which migrations have run
	cd api && $(GOOSE) status

migrate-new: ## Create a migration: make migrate-new name=add_sites
	@test -n '$(name)' || { echo 'Usage: make migrate-new name=add_sites'; exit 1; }
	cd api && $(GOOSE) -s create $(name) sql

migrate-check: ## Prove every migration applies and rolls back
	cd api && $(GOOSE) up
	cd api && $(GOOSE) reset
	cd api && $(GOOSE) up

# Running ---------------------------------------------------------------------

.PHONY: run-api run-web
run-api: ## Run the API against the local Postgres
	cd api && DATABASE_URL='$(DATABASE_URL)' go run ./cmd/server

run-web: $(WEB_DEPS) ## Run the web app in dev mode
	cd web && pnpm dev
