# Oblikovati API for Rust

The Apache-2.0 automation contract for Oblikovati. Import the crate as
`oblikovati_api`. Modules: `types`, `contract`, `wire`, `client`, `ffi`, `json`.

```rust
use oblikovati_api::{types::Point, wire::CreateWorkPointArgs};
let point = Point { x: 1.0, y: 2.0, z: 3.0 };
let request = CreateWorkPointArgs::default();
```

Use `client::Client` with an injected `client::Caller`. Shared-library add-ins
can use `client::HostTransport` with the host's C callbacks. The transport
releases host-owned buffers on both successful and failed calls.

```sh
cargo test
cargo run --example version
cargo build --example c_addin
```

The checked-in declarations come from the Go contract in the same repository.
Building this crate does not require Go. The shared version and JSON fixtures
are checked by the repository gate. See the
[API wiki](https://github.com/Oblikovati/Oblikovati.API/wiki) for paired Go and
Rust examples and C ABI ownership rules.
