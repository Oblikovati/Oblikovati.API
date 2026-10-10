# Oblikovati API Architecture

The Go and Rust bindings describe one automation contract. Go definitions
are the declaration source. `make generate` uses the Go type checker to
produce the Rust definitions, typed client methods and parity registrations.
The generated Rust source is checked in, so Rust consumers do not need Go.

## Repository layout

| Responsibility | Go | Rust |
| --- | --- | --- |
| Enums, IDs and values | `go/types` | `types` |
| In-process interfaces | `go/contract` | `contract` traits |
| Method names and JSON DTOs | `go/wire` | `wire` |
| Typed feature arguments | `go/wire/featureargs` | `wire::featureargs` |
| Transport-backed clients | `go/client` | `client` |
| C boundary | `include/oblikovati_addin.h` | `ffi` and `client::HostTransport` |

In-process interfaces remain within their language runtime. A Go interface
or Rust trait object is not a C ABI type. Add-ins use the typed clients and
JSON DTOs across the shared C callback interface.

## One injected transport

In Go, `client.Caller` supplies:

```go
Call(method string, request []byte) ([]byte, error)
```

In Rust, `client::Caller` supplies:

```rust
fn call(&self, method: &str, request: &[u8]) -> Result<Vec<u8>, Error>;
```

Production add-ins use host callbacks; unit tests use named fake hosts.
See [[First Steps]] for complete compiled examples in both languages.

Go convenience methods can take scalar arguments. Rust methods take the
canonical request DTO, for example `client.documents().close(&args)`.
Both select the same wire method and encode the same request shape.

## JSON rules

The wire names remain unchanged when Rust uses snake_case field names.
Enums are transparent numeric or string types, with their original values.
Numeric enum types retain unknown values so a newer peer can extend an enum.

| Go representation | Rust representation | JSON |
| --- | --- | --- |
| A nil slice | `None` | `null`, or omitted with `omitempty` |
| An empty allocated slice | `Some(vec![])` | `[]`, or omitted with `omitempty` |
| A pointer to a value | `Option<Box<T>>` | Value or null |
| `[]byte` | `json::Bytes` | Base64 string or null |
| Slice of byte-valued enums | `json::ByteSlice<T>` | Base64 string or null |
| Point or vector | Named coordinate struct | Fixed numeric array |
| `Variant` | `types::Variant` | Type tag, value and optional unit |
| `json.RawMessage` | `json::JsonValue` | JSON at a documented extension point |

The fixtures are produced by Go's encoder, decoded and encoded in Rust,
and compared with Go's result. This includes zero, populated and empty
collection cases for every value type. Additional tests cover errors and
buffer ownership. Non-finite numeric inputs are rejected.

## Add or change a definition

1. Edit the Go type, contract interface, wire DTO and typed client method.
2. Run `make generate` to update Rust declarations, shared version values,
   JSON fixtures and wiki examples.
3. Add or update the behavior tests and the paired language examples.
4. Run `make gate`. Commit both language definitions together.

Generated files identify their source tool. Edit the generator for a type
mapping change. Do not hand-edit a generated binding. Manual Rust modules
implement JSON rules and the C callback adapter; the parity suite tests them.

The crate must not depend on the GPL application. `make gate` checks both
dependency graphs. Shared C header changes also require both C ABI examples
to pass the same compiled C host test.

## Releases and synchronization

`go/version.go` is the version source. Generation updates `rust/Cargo.toml`
and `rust/src/version.rs`. CI rejects stale output, missing client methods,
version drift and failed JSON parity tests.

The release workflow derives the next version from commit messages, updates
both bindings, and runs the full gate again. It pushes the version commit,
`vX.Y.Z` tag and `go/vX.Y.Z` tag atomically. Rust consumers can select the
shared `vX.Y.Z` tag; Go resolves the `go/` module through `go/vX.Y.Z`.
See `RELEASING.md` for the import metadata needed by Go.

The C load handshake uses the same major and minor in both languages:
`api.Major()` / `api.Minor()` in Go and `API_MAJOR` / `API_MINOR` in Rust.
The host requires equal major versions and an add-in minor version no newer
than its own. The contract remains in initial development while major is 0.
