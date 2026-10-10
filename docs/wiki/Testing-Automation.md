# Testing Automation

Run the API checks from the repository root:

```sh
make gate
```

The gate checks generation drift, builds and tests Go and Rust, runs Go vet
and Rust Clippy, checks formatting and SPDX headers, checks the dependency
boundary, checks wiki examples, and runs the compiled C ABI host test. It also
copies the Go module to an isolated directory and runs all Go tests there, so a
missing file outside the published module cannot pass unnoticed.

## Named fake hosts in both languages

The [[First Steps]] examples define `exampleHost` in Go and `ExampleHost` in
Rust. Each validates the method and request before returning a typed version
reply. Both are compiled by the gate. A fake must fail on an unexpected
method or request, so a passing test verifies the operation it claims to test.

Go tests can call `client.New(fake).Application().ApiVersion()`. Rust tests
can call `Client::new(fake).application().api_version()`. Use the method
constants from `wire`; do not duplicate protocol strings in application code.

Run each language independently:

```sh
(cd go && GOWORK=off go test ./...)
cargo test --manifest-path rust/Cargo.toml
```

## JSON parity

The Go fixture producer uses the real API types and Go JSON encoder. It
covers zero values, populated values, and empty allocated collections.
Additional cases check byte padding, large identifiers, null coordinates,
normalization boundaries, and invalid timestamps and variant tags.
The Rust tests deserialize these payloads, serialize them again, and compare
with Go's decode-and-encode result. Go rejections must also fail in Rust.

The generator checks every wire method against the Go client's typed calls.
Every generated Rust method has a test that checks its method name and
serialized request. Separate transport tests check successful replies,
malformed replies, host errors and non-finite request numbers.

```sh
make generate
make check-generated
```

`check-generated` fails if the generated declarations, version, fixtures or
wiki examples differ. Commit the refreshed output with its source change.
Rust package consumers run these tests with Cargo alone; the fixture files
are checked in.

## Shared C ABI test

```sh
make abi-test
```

The test compiles `tests/c_abi_host.c` against the public C header, builds
both example add-ins and loads them. It checks the same exports and callback
behavior for each language. Missing callbacks, host errors and malformed
JSON are tested as well as the successful call. Allocated replies must be
freed once on every path.

This test requires a C compiler and Linux or macOS. It does not require the
CAD application or a display. Rust unit tests also exercise invalid pointer
and length combinations without reading invalid memory.

## Tests against the application

Build the windowed host and MCP bridge against `Oblikovati.API/go`. Build the
two example libraries with `make abi-test`, then put them beside the bridge in
an isolated add-ins directory. Launch the host in the foreground with
`OBK_ADDINS_DIR` set to that directory and `OBK_MCP_ADDR=127.0.0.1:7815`.
Capture its standard output and error in `/tmp/oblikovati-api-live.log`.

```sh
python3 tests/live_api.py \
  --endpoint http://127.0.0.1:7815/mcp \
  --log /tmp/oblikovati-api-live.log \
  --version "$(cd go && GOWORK=off go run ./examples/version)" \
  --screenshot build/live/api-viewport.png
```

This test uses the MCP bridge to activate and stop each example 20 times.
It waits for each worker's structured success record, checks the host version,
and captures the viewport with the MCP screenshot API. Inspect the saved image.
The add-ins are inactive at the end of the test.

An add-in feature still needs a live application test. Drive the feature
through its typed API calls, then check the exact B-rep's validity, per-face
properties, stable references and parameter recompute. For sketch-based
features, check that the sketch has zero degrees of freedom before building
the feature. Use the viewport to check the rendered result as an additional
visual check. A mesh or image does not prove exact geometry.

Keep the application out of the shipped API and add-in dependency graph.
The contract gate verifies this boundary for both language bindings.
