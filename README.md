# Learning Go

A structured repository for learning Go, from fundamentals through distributed systems.

## Structure

```
exercises/      # Small, focused exercises tied to specific concepts
experiments/    # Quick throwaway code to test ideas
notes/          # Concept notes and things that tripped me up
projects/       # Larger builds (CLI Calculator, Web Scraper, REST API, etc.)
PLAN.md         # Full learning plan with phases and milestones
```

## Phases

| Phase | Focus                                                | Key Projects                                 |
| ----- | ---------------------------------------------------- | -------------------------------------------- |
| 0     | Environment & Toolchain                              | —                                            |
| 1     | Fundamentals (types, errors, interfaces, testing)    | CLI Calculator, File Parser                  |
| 2     | Concurrency & HTTP                                   | Web Scraper, Rate Limiter, Middleware Chain  |
| 3     | Backend Engineering (architecture, DB, auth, Docker) | REST API, Auth Service, URL Shortener        |
| 4     | Advanced (perf, gRPC, distributed systems)           | Task Queue, Chat Server, Distributed Workers |

## Running

```sh
go test ./...          # run all tests
go test -race ./...    # run with race detector
go vet ./...           # static analysis
```

## Rules

- Attempt every problem manually first.
- Write tests before implementation.
- Use AI for review, not generation.
- Rewrite any AI-provided code in my own style.
- Log difficult concepts in PLAN.md.
