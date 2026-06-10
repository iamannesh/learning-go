# Go Learning Plan

## Goal

Become comfortable building production-quality backend systems in Go with strong fundamentals in:

- Concurrency
- Networking
- APIs
- Testing
- Performance
- Clean Architecture

---

## AI Usage Rules

- First attempt every problem manually.
- Use AI for hints before asking for full solutions.
- Write the tests first, then ask AI to review your implementation.
- Ask AI to review code instead of generating entire implementations.
- Always rewrite AI-generated code in your own style.
- Keep notes on concepts that were difficult.

---

## Resources

| Type   | Resource                                                  | When to Use              |
| ------ | --------------------------------------------------------- | ------------------------ |
| Book   | _The Go Programming Language_ — Donovan & Kernighan       | Phase 1–2 companion      |
| Book   | _100 Go Mistakes and How to Avoid Them_ — Teiva Harsanyi  | Phase 2–3 companion      |
| Book   | _Cloud Native Go_ — Matthew Titmus                        | Phase 3–4 companion      |
| Site   | [go.dev/tour](https://go.dev/tour)                        | Start here               |
| Site   | [gobyexample.com](https://gobyexample.com)                | Quick reference          |
| Site   | [pkg.go.dev](https://pkg.go.dev)                          | Standard library docs    |
| Site   | [effectivego](https://go.dev/doc/effective_go)            | Best practices reference |
| Videos | [GopherCon talks](https://www.youtube.com/@GopherAcademy) | Phase 3–4                |

---

## Phase 0 — Environment & Toolchain

> Get comfortable with Go's toolchain before writing a single line of real code. It is a feature of the language.

### Tooling

- [ ] Install Go and understand `$GOPATH` vs module mode
- [ ] `go mod init`, `go mod tidy`, `go get`
- [ ] `go build`, `go run`, `go test`, `go install`
- [ ] `gofmt` and `goimports` — auto-formatting
- [ ] `go vet` — static analysis
- [ ] `golangci-lint` — linter suite
- [ ] `gopls` — language server (editor integration)
- [ ] `delve` — debugger (`dlv debug`, `dlv test`)

### Setup

- [ ] Configure editor (VS Code or GoLand) with gopls
- [ ] Create your learning repo with a clean module structure
- [ ] Write a `Makefile` with targets: `build`, `test`, `lint`, `fmt`

---

## Phase 1 — Go Fundamentals

### Concepts

- [ ] Variables, constants, and zero values
- [ ] Basic types: `int`, `string`, `bool`, `float64`, `byte`, `rune`
- [ ] Slices — creation, append, copy, 3-index slicing, gotchas
- [ ] Maps — creation, iteration, zero values, concurrency caveats
- [ ] Structs and struct embedding
- [ ] Methods and pointer receivers vs value receivers
- [ ] Interfaces and implicit implementation
- [ ] Functions — multiple return values, named returns
- [ ] Variadic functions
- [ ] Closures
- [ ] Pointers — when and why
- [ ] `defer` — ordering and use cases
- [ ] `panic` and `recover`
- [ ] Error handling — the `error` interface
- [ ] Error wrapping — `fmt.Errorf %w`, `errors.Is`, `errors.As`
- [ ] Type assertions and type switches
- [ ] Packages and visibility (exported vs unexported)
- [ ] Modules and dependency management

### Testing Basics

- [ ] `testing` package — `TestXxx(t *testing.T)`
- [ ] Table-driven tests
- [ ] `t.Run` for subtests
- [ ] `go test ./...`, `-v`, `-run` flags

### Projects

- [ ] **CLI Calculator** — arithmetic ops, error handling, input parsing
- [ ] **File Parser** — read CSV or JSON, transform data, write output

---

## Phase 2 — Intermediate Go

### Concurrency

- [ ] Goroutines — creation, lifecycle
- [ ] Channels — unbuffered vs buffered
- [ ] `select` statement
- [ ] `sync.WaitGroup`
- [ ] `sync.Mutex` and `sync.RWMutex`
- [ ] `sync.Once`
- [ ] `sync/atomic`
- [ ] Context — cancellation, timeouts, propagation
- [ ] Worker pool pattern
- [ ] Common concurrency mistakes (race conditions, goroutine leaks)
- [ ] `-race` flag

### I/O and HTTP

- [ ] `io.Reader` and `io.Writer` — the backbone of Go I/O
- [ ] `bufio`, `bytes`, `strings` packages
- [ ] JSON encoding and decoding (`encoding/json`)
- [ ] `net/http` — building a basic HTTP server
- [ ] HTTP handlers, `ServeMux`, middleware pattern
- [ ] Making HTTP requests with `http.Client`

### Testing Intermediate

- [ ] Mocking with interfaces (no magic — just pass fakes)
- [ ] `httptest.NewRecorder` and `httptest.NewServer`
- [ ] `testify` — `assert` and `require`
- [ ] Test helpers and `t.Helper()`
- [ ] Test coverage — `go test -cover`

### Projects

- [ ] **Concurrent Web Scraper** — goroutines, channels, context timeouts
- [ ] **Rate Limiter** — token bucket or leaky bucket, concurrency-safe
- [ ] **HTTP Middleware Chain** — logging, auth, recovery middlewares

---

## Phase 3 — Backend Engineering

### Architecture

- [ ] Clean Architecture / Hexagonal Architecture
- [ ] Dependency injection (manual, no framework)
- [ ] Repository pattern
- [ ] Service layer pattern
- [ ] `internal/` package convention

### Data & Storage

- [ ] `database/sql` — the standard interface
- [ ] `sqlx` — extensions for struct scanning
- [ ] `pgx` — PostgreSQL driver (direct)
- [ ] GORM — ORM tradeoffs vs raw SQL
- [ ] Database migrations — `golang-migrate` or `goose`
- [ ] Redis — `go-redis` client, caching patterns, TTL
- [ ] Connection pooling and timeouts

### Production Concerns

- [ ] Authentication — JWT, session tokens, bcrypt for passwords
- [ ] Middleware — auth, logging, rate limiting, CORS, recovery
- [ ] Structured logging — `log/slog` (stdlib, Go 1.21+) or `zap` / `zerolog`
- [ ] Configuration management — environment variables, `viper`, 12-factor
- [ ] Graceful shutdown — `os.Signal`, `context.WithCancel`, server drain
- [ ] Docker — write a `Dockerfile`, multi-stage builds

### Testing Advanced

- [ ] Integration tests with a real database (use `testcontainers-go`)
- [ ] Fuzz testing with `go test -fuzz`
- [ ] Contract testing for HTTP APIs

### Projects

- [ ] **REST API** — full CRUD, auth, DB, middleware, structured logs, Docker
- [ ] **Auth Service** — registration, login, JWT issue/refresh/revoke
- [ ] **URL Shortener** — redirect service, Redis for caching, Postgres for persistence

---

## Phase 4 — Advanced Topics

### Performance

- [ ] `pprof` — `net/http/pprof`, `go tool pprof`, flame graphs
- [ ] Benchmarks — `BenchmarkXxx`, `-benchmem`, `benchstat`
- [ ] Memory optimization — escape analysis, `go build -gcflags="-m"`
- [ ] Reducing allocations — `sync.Pool`, pre-allocation patterns
- [ ] Tracing — `runtime/trace`, `go tool trace`

### Go Language Depth

- [ ] Generics (Go 1.18+) — type parameters, constraints, when to use vs avoid
- [ ] Reflection — `reflect` package, when it is and isn't appropriate
- [ ] `unsafe` — awareness and why to avoid it
- [ ] Build tags and conditional compilation
- [ ] `cgo` — awareness and tradeoffs

### Distributed Systems

- [ ] gRPC — Protobuf, services, streaming, interceptors
- [ ] Message queues — basic producer/consumer (NATS or RabbitMQ)
- [ ] Distributed tracing — OpenTelemetry basics
- [ ] Kubernetes — pods, deployments, services, configmaps, health checks
- [ ] 12-factor application principles

### Projects

- [ ] **Task Queue** — job scheduling, workers, retries, dead-letter queue
- [ ] **Chat Server** — WebSockets, rooms, presence, broadcast
- [ ] **Distributed Worker System** — gRPC coordination, multiple workers, task dispatch

---

## Milestone Roadmap

> This is a rough guide, not a strict deadline. Adjust based on your pace.

| Week    | Focus                                               | Deliverable                         |
| ------- | --------------------------------------------------- | ----------------------------------- |
| Week 1  | Phase 0 + Phase 1 concepts (types through errors)   | Toolchain configured, notes written |
| Week 2  | Phase 1 remaining concepts + testing basics         | CLI Calculator                      |
| Week 3  | Phase 1 wrap-up                                     | File Parser                         |
| Week 4  | Phase 2 — Goroutines, channels, select, WaitGroup   | Notes + small experiments           |
| Week 5  | Phase 2 — Context, Mutex, worker pools, race flag   | Concurrent Web Scraper              |
| Week 6  | Phase 2 — I/O, JSON, `net/http`                     | Rate Limiter                        |
| Week 7  | Phase 2 — Middleware + testing intermediate         | HTTP Middleware Chain               |
| Week 8  | Phase 3 — Clean architecture, DI, repo pattern      | Architecture skeleton               |
| Week 9  | Phase 3 — `database/sql`, migrations, Redis         | DB-backed service                   |
| Week 10 | Phase 3 — Auth, middleware, structured logging      | Auth Service                        |
| Week 11 | Phase 3 — Config, graceful shutdown, Docker         | URL Shortener                       |
| Week 12 | Phase 3 — Full REST API, integration tests          | REST API (production-ready)         |
| Week 13 | Phase 4 — pprof, benchmarks, memory analysis        | Performance audit of prior project  |
| Week 14 | Phase 4 — gRPC, generics, reflection basics         | gRPC service                        |
| Week 15 | Phase 4 — Message queues, tracing, Kubernetes intro | Task Queue                          |
| Week 16 | Phase 4 — Distributed systems                       | Distributed Worker System           |

---

## Notes Log

> Keep a running log of hard concepts as you hit them.

| Date | Concept | What Confused Me | How I Resolved It |
| ---- | ------- | ---------------- | ----------------- |
|      |         |                  |                   |
