# Project instructions

These instructions apply to every task and all changes in this repository.

## Required guidelines

- Before making any change, read and apply `.agents/skills/karpathy-guidelines/SKILL.md`. Follow it for planning, implementation, refactoring, reviews, tests, and configuration changes.
- Apply `.agents/skills/ponytail/SKILL.md` to coding tasks, using its default `full` level unless the user chooses another level or turns it off.
- Make the smallest complete change, state material assumptions, preserve existing conventions, and verify the behavior affected by the change. Keep unrelated work intact.
- Treat specialized skills as guidance within the requested scope; their examples do not require adding dependencies, abstractions, or features.

## Go changes

Read the applicable skills before editing or reviewing Go code:

| Work | Skill under `.agents/skills/` |
| --- | --- |
| Idiomatic code, package layout, consumer-owned interfaces, useful zero values | `golang-patterns/SKILL.md` |
| Service architecture and broader Go design decisions | `golang-pro/SKILL.md` |
| Goroutines, channels, synchronization, cancellation, lifecycle ownership | `golang-concurrency/SKILL.md` |
| Error wrapping, sentinel and typed errors, structured failure handling | `golang-error-handling/SKILL.md` |
| Table-driven tests, subtests, HTTP tests, golden files, fuzzing | `golang-testing/SKILL.md` |
| Benchmarks, allocations, pprof, execution traces, benchstat | `golang-benchmark/SKILL.md` |
| Review of a Go diff, including the final review of Go changes | `go-code-review/SKILL.md` |

- Use the Go version declared in `backend/go.mod`; check version requirements before copying skill examples.
- Run Go commands from `backend/`. For Go changes, check formatting with `gofmt -l .`, run `go vet ./...` and `go test ./...`, and run `golangci-lint run ./...` when available. Correct formatting in changed files only.
- Run `go test -race ./...` for concurrency changes. For performance changes, measure allocations and compare repeated baseline/candidate benchmark runs with `benchstat`; run measurements serially.
- Review context propagation, goroutine exit paths, channel ownership, error handling, and API/interface design in the changed code. Use the review skill's bundled script when Bash is available, or equivalent native commands on Windows.
- Report checks that failed or could not run; do not claim an unavailable tool passed.

## Angular changes

- Use the `angular-cli` MCP server configured in `.codex/config.toml` for Angular workspace information, best practices, and official documentation when working on the frontend.
- The server uses the Angular CLI already declared in `frontend/package.json`. Install frontend dependencies with `npm ci` from `frontend/` if they are missing.
- Run the frontend's relevant tests and build from `frontend/` after Angular behavior changes.
