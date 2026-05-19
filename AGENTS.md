# AGENTS.md

## Project Shape
- This is a Go 1.21 library module: `github.com/studiolambda/golidate`.
- Keep the repo Go-only; do not reintroduce Node/npm/VitePress/docs-site tooling.
- Public packages are the root `golidate` package plus `rule`, `format`, `translate`, and `translate/language`.
- Root package owns orchestration and result types; `rule/` owns rule constructors; `format/` owns final message string formatters; `translate/` owns translation entries and dictionaries.

## Verification Commands
- Full local verification should match CI: `go test ./...`, then `go vet ./...`, then `go test -race ./...`.
- CI runs the same checks on Go 1.21 in `.github/workflows/go.yml`.
- Focus one package with `go test ./rule`, `go test ./format`, or `go test ./translate/...`.
- Focus one test with `go test ./rule -run TestName` or `go test . -run TestName`.

## Code And Behavior Gotchas
- `Value` recursively validates values implementing `Validator`; `Self` skips validator recursion and applies direct rules to the value.
- Pointer-receiver validators must remain discoverable: do not blindly dereference values that implement `Validator`.
- Map-derived validation output is intentionally deterministic; preserve sorting in `pending.go`, `rule/map_keys.go`, and `rule/map_values.go`.
- `Results.Passed`, `Results.Failed`, and `Results.PassesAny`/`PassesAll` are child-aware via nested result state.
- Membership is intentionally strict: `rule.In("1")(1)` and cross-width numeric matches such as `int64(1)` vs `int(1)` should fail.
- `rule.Equal` uses deep equality so slices and maps are safe inputs.
- Reflection-heavy rules should fail safely on nil or unsupported values, not panic.
- `rule.Min` and `rule.Max` support integer and floating numeric inputs without truncating decimals.
- Length rules support only arrays, channels, maps, slices, and strings.
- Translation dictionaries are merged once in `Results.Translate`; keep override order stable and avoid per-result remapping.
- Placeholder replacement in `translate.Simple` is deterministic for overlapping metadata keys; preserve longest-key-first behavior.

## Documentation Expectations
- The docs site was removed; keep documentation in Go comments, examples, and `README.md`.
- README examples should use `go get` only and must not mention Node/npm tooling.
- This repo intentionally documents exported and unexported identifiers; keep comments accurate when changing behavior.

## Dependencies
- Runtime dependency: `github.com/spf13/cast` for translation placeholder coercion.
- Test dependency: `github.com/stretchr/testify`; keep test additions within the existing Go test setup.
