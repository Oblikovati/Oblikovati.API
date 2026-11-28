# WebAssembly Add-In ABI (`js-wasm`)

Browser builds of Oblikovati run on WebAssembly + WebGPU inside an Orchestrator Web Worker.
Because browsers cannot load native shared libraries (`.so`, `.dylib`, `.dll`), add-ins target
the `js-wasm` platform (`GOOS=js GOARCH=wasm`, `CGO_ENABLED=0`) and run inside isolated Web Workers.

## Bundle Format

A `js-wasm` bundle uses the same zip structure as native bundles:

- `manifest.json` — standard add-in manifest with `apiCompat` stamped to `MAJOR.MINOR`.
- `icon.svg` — add-in icon.
- `<addin-name>.wasm` — Go WebAssembly binary built with `GOOS=js GOARCH=wasm`.
- Optional images declared in `manifest.json`.

## Worker Host Contract (`globalThis.__obkHost`)

Before instantiating and running the `.wasm` module with Go's `wasm_exec.js`, the host worker
exposes `globalThis.__obkHost`:

- `register(exports)` — called by `oblikovati.org/api/wasmaddin.Serve` with:
  - `id() -> string`
  - `manifest() -> string`
  - `apiMajor() -> number`
  - `apiMinor() -> number`
  - `hasAutomation() -> boolean`
  - `activate() -> Promise<number>` (`0` = `OBK_OK`, `1` = `OBK_ERR`)
  - `deactivate() -> Promise<number>`
  - `notify(eventBytes: Uint8Array) -> Promise<number>`
  - `automation(method: string, reqBytes: Uint8Array) -> Promise<{ ok: boolean, data: Uint8Array }>`
- `call(method: string, reqBytes: Uint8Array) -> { ok: boolean, data: Uint8Array }` — synchronous
  RPC to the Oblikovati Orchestrator Worker backed by `SharedArrayBuffer` and `Atomics.wait`.

## Version Handshake

The host checks `apiMajor()` and `apiMinor()` before calling `activate()`, applying the same
rule as native add-ins: matching major version, and add-in minor version less than or equal to
the host minor version.
