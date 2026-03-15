# AGENTS.md

AI agent guidelines for the Station-Manager codebase.

## Architecture Overview

Station-Manager is a **Go workspace** (see `go.work`) with ~20 modules for ham radio QSO logging. Key layers:

- **Desktop apps** (`logbook-app`, `config-app`, `logging-app`): [Wails v2](https://wails.io/) + SvelteKit 5 frontends. Go backend exposes a `facade` service that Wails auto-generates TypeScript bindings for (output to `frontend/src/lib/wailsjs/`).
- **Core services**: `config`, `logging`, `database`, `cat` (rig control via serial), `email`, `forwarding` — each follows a `Service` struct pattern with `Initialize()`, dependency injection via `di.inject` tags.
- **Shared packages**: `types` (zero internal deps — shared data structures), `errors` (custom `DetailedError` with `Op` chain), `adapters` (struct-to-struct mapping), `iocdi` (DI container), `enums/*`.
- **Server** (`server/`): REST API with API-key auth for syncing QSOs to a central PostgreSQL instance.

## Build & Task Runner

All modules use [Task](https://taskfile.dev/) (not Make). Run from module directory:

```bash
task              # default dev build: vet → build → short tests
task prod         # bump version, full tests, tag, push to GitHub
task docker:start # (database/) start Postgres container
```

Wails apps: `wails dev` for hot-reload, `wails build -o <name>` for production.

## Code Patterns

### Service Initialization & DI

Services use `iocdi` for wiring. Tag fields with `di.inject:"<id>"`:
```go
type Service struct {
    ConfigService *config.Service `di.inject:"configservice"`
    Logger        *logging.Service `di.inject:"loggingservice"`
}
```
Services implement `Initialize()` called once after injection; use `sync.Once` and `atomic.Bool` guards.

### Error Handling (`errors` package)

Always wrap with operation context:
```go
const op errors.Op = "package.Function"
if err != nil {
    return errors.New(op).Err(err).Msg("human message")
}
```

### Struct Mapping (`adapters` package)

Map between `types.*` and DB models (`database/sqlite/models`, `database/postgres/models`):
```go
adapter := adapters.New()
adapter.Into(&dst, &src)  // destination-first
```
Use `json:"..."` tags for field matching. Mark overflow fields with `adapter:"additional"` for JSON spillover.

### Frontend (SvelteKit 5 + Tailwind 4)

- Use Svelte 5 runes: `$state`, `$derived`, `$props`.
- Shared utilities in `shared-utils/` (npm package `@station-manager/shared-utils`).
- Wails bindings auto-generated to `frontend/src/lib/wailsjs/`.

## Database

Dual-database support:
- **SQLite** (`database/sqlite/`): local desktop storage
- **PostgreSQL** (`database/postgres/`): server-side

Models generated via [SQLBoiler](https://github.com/volatiletech/sqlboiler). Regenerate:
```bash
# SQLite
task sqlite:db && cd database/sqlite && sqlboiler sqlite3

# Postgres
docker-compose up -d  # in database/postgres/
go run ./postgres/example/main.go
cd database/postgres && PSQL_PASS=... sqlboiler psql
```

## Testing

```bash
go test -race -run Test ./... -short   # unit tests
go test -race -run Test ./...          # full tests
```
Place tests in `*_test.go` in same package. Use `testify` assertions where available.

## Key Files

| Purpose | Location |
|---------|----------|
| Shared types (no deps) | `types/*.go` |
| DI container | `iocdi/` |
| Struct adapter | `adapters/adapter.go`, `adapters/DESIGN.md` |
| Error wrapper | `errors/errors.go` |
| API key auth design | `DESIGN_REVIEW.md`, `apikey/` |
| Rig CAT control | `cat/service.go` |
| Desktop app entry | `logbook-app/main.go` |

## Conventions

- **Module isolation**: `types` must have zero Station-Manager deps; other modules import only what's needed.
- **Config via JSON**: `build/config.json` at runtime; `config/` service loads and validates.
- **Context-aware CRUD**: database methods accept `context.Context` for timeouts.
- **Functional style preferred** over OOP where practical.

