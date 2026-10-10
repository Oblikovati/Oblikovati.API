# First Steps

The public API has Go and Rust bindings in the same repository. Both use the
same C ABI and JSON protocol. An add-in depends only on the Apache-2.0 API.

## Choose a language

| Language | Source | Package |
| --- | --- | --- |
| Go | `go/` | `oblikovati.org/api` |
| Rust | `rust/` | `oblikovati-api`, imported as `oblikovati_api` |
| C interface | `include/oblikovati_addin.h` | Shared by both languages |

Clone the API repository beside your add-in:

```sh
git clone https://github.com/Oblikovati/Oblikovati.API.git
```

For Go, use the Go version specified by `go/go.mod` and point the local
workspace to the module's new location:

```sh
cd my-addin
go work init .
go work edit -replace oblikovati.org/api=../Oblikovati.API/go
```

For Rust, add a local dependency:

```toml
[dependencies]
oblikovati-api = { path = "../Oblikovati.API/rust" }
```

A Go release is selected with `require oblikovati.org/api vX.Y.Z`; the matching
repository tag is `go/vX.Y.Z`. A Rust Git dependency can use tag `vX.Y.Z` and
package `oblikovati-api`. Both tags point to the same release commit.

## Run a typed request without the application

These complete examples use named fake hosts. They verify the version request
and return the API version from the package. Replace the fake with a host
transport when you build an add-in.

### Go

<!-- example: go/examples/version/main.go -->
```go
// SPDX-License-Identifier: Apache-2.0

// Command version demonstrates a typed client with a named fake host.
package main

import (
	"encoding/json"
	"fmt"
	"oblikovati.org/api"
	"oblikovati.org/api/client"
	"oblikovati.org/api/wire"
)

type exampleHost struct{}

func (exampleHost) Call(method string, request []byte) ([]byte, error) {
	if method != wire.MethodApplicationApiVersion || len(request) != 0 {
		return nil, fmt.Errorf("method %q payload %q: expected application.apiVersion without a body", method, request)
	}
	return json.Marshal(wire.ApplicationApiVersionResult{Version: api.Version, Major: api.Major(), Minor: api.Minor()})
}

func main() {
	version, err := client.New(exampleHost{}).Application().ApiVersion()
	if err != nil {
		panic(err)
	}
	fmt.Println(version.Version)
}
```
<!-- /example -->

```sh
cd Oblikovati.API/go
GOWORK=off go run ./examples/version
```

### Rust

<!-- example: rust/examples/version.rs -->
```rust
// SPDX-License-Identifier: Apache-2.0

use oblikovati_api::{
    API_MAJOR, API_MINOR, Error, VERSION,
    client::{Caller, Client},
    wire,
};

struct ExampleHost;

impl Caller for ExampleHost {
    fn call(&self, method: &str, request: &[u8]) -> Result<Vec<u8>, Error> {
        if method != wire::METHOD_APPLICATION_API_VERSION || !request.is_empty() {
            return Err(Error(format!(
                "method {method:?} payload {request:?}: expected application.apiVersion without a body"
            )));
        }
        let reply = wire::ApplicationApiVersionResult {
            version: VERSION.into(),
            major: API_MAJOR.into(),
            minor: API_MINOR.into(),
        };
        serde_json::to_vec(&reply).map_err(|error| Error(error.to_string()))
    }
}

fn main() -> Result<(), Error> {
    let version = Client::new(ExampleHost).application().api_version()?;
    println!("{}", version.version);
    Ok(())
}
```
<!-- /example -->

```sh
cd Oblikovati.API
cargo run --manifest-path rust/Cargo.toml --example version
```

## Connect to the application

Go implements `client.Caller` with a small cgo adapter. Rust uses
`client::HostTransport` with the callbacks supplied at activation. See
[[C ABI Add-ins]] for the complete, compiled shared-library examples.

A host callback returns a host-owned buffer. Release that buffer with the
host's `ObkHostFree` callback, including when the call returns an error.
The Rust transport performs this release automatically.

Use [[Typed Operations]] for document, parameter, sketch and feature requests.
Use [[Testing Automation]] for fake-host tests and the shared C ABI test.
