# AGENTS.md

## Project Overview
`apigo` (`github.com/wiztools/apigo`) is a small Go utility library providing common helpers for building JSON HTTP APIs. It has no external dependencies beyond the Go standard library.

## Language & Build
- **Language:** Go (1.25+)
- **Module:** `github.com/wiztools/apigo`
- **Package:** `apigo`
- **Build:** `go build ./...`
- **Test:** `go test ./...`

## Project Structure
This is a single-package library with no subdirectories:
- `err.go` — `Err` type implementing the `error` interface with HTTP status codes, plus constructor and checker functions for common HTTP errors (400, 401, 403, 404, 409, 422, 429).
- `response.go` — `Bind` (JSON request body decoding) and `WriteResp` (JSON response writing).
- `response-gen.go` — Generic response model structs (`ModelID`, `ModelName`, etc.) and helper functions (`RespCause`, `RespID`, `RespRowsAffected`, etc.) for building JSON response maps.
- `pagination.go` — `Page` struct and `PaginationQuerier` for extracting `page`/`size` query parameters.
- `err_test.go` — Tests for `Err`.

## Coding Conventions
- No external dependencies — only use the Go standard library.
- Exported types use the `Model` prefix for response structs (e.g., `ModelID`, `ModelCause`).
- Exported helper functions use the `Resp` prefix for response map builders and `Err` prefix for error constructors/checkers.
- Error checker functions follow the pattern `ErrIs<Status>(err error) bool`.
- Error constructor functions follow the pattern `Err<Status>(cause string) *Err`.
- JSON struct tags use `snake_case`.
- Keep the library minimal — each file is focused on a single concern.

## Testing
- Tests use the standard `testing` package — no third-party test frameworks.
- Run tests: `go test ./...`
