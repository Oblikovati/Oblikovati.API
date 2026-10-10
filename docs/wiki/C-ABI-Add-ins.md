# C ABI Add-ins

Activation runs on the host session thread. Start host calls on a worker and
return from activation before waiting for their results. The host frame loop
must run to process those calls. Deactivation joins the worker before the host
can release the callback library. The examples report a failed version query
through the deactivation result.

These two examples start a worker that queries the host API version. Both are
compiled and loaded by the same C host test. They return no add-in-owned
buffers. They release all callbacks when deactivation returns.

The host resolves the eight required exports declared by
`include/oblikovati_addin.h`. The ID and manifest strings remain valid for
the lifetime of the add-in. The version exports use the API package version.

## Go shared library

The entry points:

<!-- example: go/examples/c_addin/main.go -->
```go
// SPDX-License-Identifier: Apache-2.0

// Package main builds the small Go add-in used by the shared C ABI test.
package main

/*
#cgo CFLAGS: -I${SRCDIR}/../../include -DOBK_BUILDING_ADDIN
#include <stdlib.h>
#include "oblikovati_addin.h"
static char *example_id(void) { return "com.oblikovati.example.go"; }
static char *example_manifest(void) { return "{\"id\":\"com.oblikovati.example.go\",\"displayName\":\"Go API example\",\"version\":\"1.0.0\",\"capabilities\":[]}"; }
*/
import "C"

import (
	"log/slog"
	"oblikovati.org/api"
	"oblikovati.org/api/client"
	"os"
	"sync"
)

// C exports have no instance parameter. Keep only the worker's lifecycle at this boundary.
var activationMu sync.Mutex
var versionDone chan C.int

//export ObkAddInId
func ObkAddInId() *C.char { return C.example_id() }

//export ObkAddInManifest
func ObkAddInManifest() *C.char { return C.example_manifest() }

//export ObkAddInApiMajor
func ObkAddInApiMajor() C.int { return C.int(api.Major()) }

//export ObkAddInApiMinor
func ObkAddInApiMinor() C.int { return C.int(api.Minor()) }

//export ObkAddInActivate
func ObkAddInActivate(call C.ObkHostCall, free C.ObkHostFree) C.int {
	if call == nil || free == nil {
		return C.OBK_ERR
	}
	activationMu.Lock()
	defer activationMu.Unlock()
	if versionDone != nil {
		return C.OBK_ERR
	}
	versionDone = make(chan C.int, 1)
	// Activation runs on the host session thread. Return so it can drain queued calls.
	go probeVersion(cgoCaller{call, free}, versionDone)
	return C.OBK_OK
}

func probeVersion(host cgoCaller, done chan<- C.int) {
	version, err := client.New(host).Application().ApiVersion()
	if err != nil || version.Major != api.Major() || version.Minor < api.Minor() {
		done <- C.OBK_ERR
		return
	}
	slog.New(slog.NewJSONHandler(os.Stderr, nil)).Info("api-version", "language", "go", "version", version.Version)
	done <- C.OBK_OK
}

//export ObkAddInDeactivate
func ObkAddInDeactivate() C.int {
	activationMu.Lock()
	done := versionDone
	versionDone = nil
	activationMu.Unlock()
	if done == nil {
		return C.OBK_OK
	}
	return <-done
}

//export ObkAddInNotify
func ObkAddInNotify(event *C.uint8_t, length C.int) C.int { return C.OBK_OK }

// No add-in-owned buffers are returned by this example.
//
//export ObkFree
func ObkFree(_ *C.uint8_t) {}

func main() {}
```
<!-- /example -->

The transport adapter:

<!-- example: go/examples/c_addin/caller.go -->
```go
// SPDX-License-Identifier: Apache-2.0

package main

/*
#cgo CFLAGS: -I${SRCDIR}/../../include -DOBK_BUILDING_ADDIN
#include <stdlib.h>
#include <limits.h>
#include "oblikovati_addin.h"
static int invoke_host(ObkHostCall call, const char *method, const uint8_t *request, int length, uint8_t **response, int *size) { return call(method, request, length, response, size); }
static void release_host(ObkHostFree free, uint8_t *response) { free(response); }
*/
import "C"

import (
	"fmt"
	"strings"
	"unsafe"
)

type cgoCaller struct {
	call C.ObkHostCall
	free C.ObkHostFree
}

func (host cgoCaller) Call(method string, request []byte) ([]byte, error) {
	if strings.ContainsRune(method, 0) || len(request) > int(C.INT_MAX) {
		return nil, fmt.Errorf("request %q length %d: expected a NUL-free method and C int length", method, len(request))
	}
	name := C.CString(method)
	defer C.free(unsafe.Pointer(name))
	var response *C.uint8_t
	var length C.int
	status := C.invoke_host(host.call, name, (*C.uint8_t)(unsafe.Pointer(unsafe.SliceData(request))), C.int(len(request)), &response, &length)
	if response != nil {
		defer C.release_host(host.free, response)
	}
	if length < 0 || (length > 0 && response == nil) {
		return nil, fmt.Errorf("response %p length %d: expected allocated bytes and a nonnegative length", response, length)
	}
	bytes := C.GoBytes(unsafe.Pointer(response), length)
	if status != C.OBK_OK {
		return nil, fmt.Errorf("host method %q status %d: %s", method, status, bytes)
	}
	return bytes, nil
}
```
<!-- /example -->

Build from the API repository:

```sh
cd go
GOWORK=off CGO_ENABLED=1 go build -buildmode=c-shared -o ../build/go-example.so ./examples/c_addin
```

Use `.dylib` on macOS and `.dll` on Windows. Keep the shared header on your
C include path when adapting the example to your own repository.

## Rust shared library

The example target in `rust/Cargo.toml` has `crate-type = ["cdylib"]`:

<!-- example: rust/examples/c_addin.rs -->
```rust
// SPDX-License-Identifier: Apache-2.0

//! Minimal shared-library add-in: query the host from a worker after activation returns.

use oblikovati_api::{
    API_MAJOR, API_MINOR,
    client::{Client, HostTransport},
    ffi::*,
};
use std::ffi::{c_char, c_int};
use std::sync::Mutex;
use std::thread::JoinHandle;

// C exports have no instance parameter. The transport itself stays inside the worker.
static VERSION_WORKER: Mutex<Option<JoinHandle<c_int>>> = Mutex::new(None);

#[unsafe(no_mangle)]
pub extern "C" fn ObkAddInId() -> *const c_char {
    c"com.oblikovati.example.rust".as_ptr()
}

#[unsafe(no_mangle)]
pub extern "C" fn ObkAddInManifest() -> *const c_char {
    c"{\"id\":\"com.oblikovati.example.rust\",\"displayName\":\"Rust API example\",\"version\":\"1.0.0\",\"capabilities\":[]}".as_ptr()
}

#[unsafe(no_mangle)]
pub extern "C" fn ObkAddInApiMajor() -> c_int {
    API_MAJOR
}

#[unsafe(no_mangle)]
pub extern "C" fn ObkAddInApiMinor() -> c_int {
    API_MINOR
}

/// Start a version query without blocking the host session. Example: load this library in the host.
///
/// # Safety
/// Both callback pointers must remain valid until `ObkAddInDeactivate` returns.
#[unsafe(no_mangle)]
pub unsafe extern "C" fn ObkAddInActivate(
    call: Option<ObkHostCall>,
    free: Option<ObkHostFree>,
) -> c_int {
    if call.is_none() || free.is_none() {
        return OBK_ERR;
    }
    let Ok(mut worker) = VERSION_WORKER.lock() else {
        return OBK_ERR;
    };
    if worker.is_some() {
        return OBK_ERR;
    }
    // Activation runs on the session thread; it must return before the host drains calls.
    match std::thread::Builder::new()
        .name("api-version".into())
        .spawn(move || probe_version(call, free))
    {
        Ok(started) => {
            *worker = Some(started);
            OBK_OK
        }
        Err(_) => OBK_ERR,
    }
}

fn probe_version(call: Option<ObkHostCall>, free: Option<ObkHostFree>) -> c_int {
    // SAFETY: activation's contract keeps callbacks alive until deactivation joins this worker.
    let Ok(transport) = (unsafe { HostTransport::new(call, free) }) else {
        return OBK_ERR;
    };
    let client = Client::new(transport);
    match client.application().api_version() {
        Ok(version)
            if version.major == i64::from(API_MAJOR) && version.minor >= i64::from(API_MINOR) =>
        {
            eprintln!(
                "{}",
                serde_json::json!({"msg":"api-version", "language":"rust", "version":version.version})
            );
            OBK_OK
        }
        _ => OBK_ERR,
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn ObkAddInDeactivate() -> c_int {
    let Ok(mut worker) = VERSION_WORKER.lock() else {
        return OBK_ERR;
    };
    worker
        .take()
        .map_or(OBK_OK, |started| started.join().unwrap_or(OBK_ERR))
}

#[unsafe(no_mangle)]
pub extern "C" fn ObkAddInNotify(_event: *const u8, _length: c_int) -> c_int {
    OBK_OK
}

// This example never returns an add-in-owned buffer, so there is nothing to release.
#[unsafe(no_mangle)]
pub extern "C" fn ObkFree(_pointer: *mut u8) {}
```
<!-- /example -->

```sh
cargo build --manifest-path rust/Cargo.toml --example c_addin
```

The result is under `rust/target/debug/examples/`. Cargo selects the platform
library extension. `HostTransport` is tied to the callback lifetime and cannot
move between threads. Callbacks must stay valid for every call made through it.
Do not unwind a Rust panic across a C entry point in an application add-in.

## Memory and event rules

- Host replies belong to the host. Free them only with `ObkHostFree`.
- Add-in buffers belong to the add-in. The host releases them with `ObkFree`.
- ID and manifest strings remain owned by the add-in.
- Decode pushed events using the `wire` event structs and event constants.
- An automation callback must not make a synchronous host call while the
  host dispatcher is waiting for that callback.

Run the shared C test with `make abi-test` on Linux or macOS. It checks export
names, version handshake, a typed call, host failures, malformed replies and
exactly one buffer release on each allocated-response path. See
[[Testing Automation]] for the other checks.
