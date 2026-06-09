# First Steps

In this walkthrough you build a complete add-in that:

1. registers a **ribbon button** in the add-ins area of the Part ribbon, and
2. when that button is clicked, **creates a sketch, draws a rectangle, and extrudes it into a
   cube**.

It is a single Go package compiled as a C-shared library. Everything it links is the
Apache-2.0 contract `oblikovati.org/api`.

> Everything you need ships in this repository: the Go contract and the C ABI header
> (`include/oblikovati_addin.h`). The snippets below are the essential pieces, so you
> can write your own add-in from scratch with nothing else checked out.

## 1. Project layout

```
my-addin/
├── go.mod                 # module my-addin
├── export.go             # the C ABI entry points (cgo //export)
├── hostcaller.go         # the Transport over the host callback
└── addin.go              # your logic: register the button, run the action
```

Your `go.mod` requires only the contract:

```go
module my-addin

go 1.22

require oblikovati.org/api v0.0.0
```

During local development, resolve the contract from a sibling checkout with a workspace
(`go.work` is for local builds only and is git-ignored):

```sh
git clone https://github.com/Oblikovati/Oblikovati.API.git   # the Apache-2.0 contract
cd my-addin
go work init .
go work edit -replace oblikovati.org/api=../Oblikovati.API
```

You also need the C ABI header `oblikovati_addin.h` on your cgo include path — it
defines the host↔add-in boundary and ships in this repo's
[`include/`](https://github.com/Oblikovati/Oblikovati.API/blob/develop/include/oblikovati_addin.h).
Copy it next to your sources (the snippets below include from `${SRCDIR}/include`):

```sh
mkdir -p include
cp ../Oblikovati.API/include/oblikovati_addin.h include/
```

## 2. The C ABI entry points (`export.go`)

The host loads your library and resolves six exported functions. Most are boilerplate; the two
that matter are **Activate** (wire yourself up) and **Notify** (react to host events).

```go
package main

/*
#cgo CFLAGS: -I${SRCDIR}/include -DOBK_BUILDING_ADDIN
#include <stdlib.h>
#include <stdint.h>
#include "oblikovati_addin.h"
*/
import "C"

import (
	"sync"
	"unsafe"
)

const addInID = "com.example.cube-button"

var (
	idC  = C.CString(addInID)
	manC = C.CString(manifestJSON) // {"id":...,"displayName":...,"version":...,"capabilities":[]}

	mu       sync.Mutex
	hostCall C.ObkHostCall // the host's RPC entry, captured on Activate
	hostFree C.ObkHostFree // frees host-owned reply buffers
	app      *AddIn        // your add-in state
)

//export ObkAddInId
func ObkAddInId() *C.char { return idC }

//export ObkAddInManifest
func ObkAddInManifest() *C.char { return manC }

//export ObkAddInActivate
func ObkAddInActivate(call C.ObkHostCall, freeFn C.ObkHostFree) C.int {
	mu.Lock()
	defer mu.Unlock()
	hostCall, hostFree = call, freeFn
	a, err := NewAddIn(cgoHostCaller{}) // builds the client and registers the button
	if err != nil {
		return C.OBK_ERR
	}
	app = a
	return C.OBK_OK
}

//export ObkAddInDeactivate
func ObkAddInDeactivate() C.int {
	mu.Lock()
	defer mu.Unlock()
	app, hostCall, hostFree = nil, nil, nil
	return C.OBK_OK
}

//export ObkAddInNotify
func ObkAddInNotify(ev *C.uint8_t, n C.int) C.int {
	mu.Lock()
	a := app
	mu.Unlock()
	if a != nil {
		a.OnEvent(C.GoBytes(unsafe.Pointer(ev), n))
	}
	return C.OBK_OK
}

//export ObkFree
func ObkFree(p *C.uint8_t) { C.free(unsafe.Pointer(p)) }

func main() {}
```

## 3. The transport (`hostcaller.go`)

The contract's `client.Client` needs a `Transport` — one method that sends a JSON request for a
named method and returns the JSON reply. You back it with the host callback captured on Activate.
A reply buffer is **owned by the host**; always release it with `ObkHostFree`.

```go
package main

/*
#cgo CFLAGS: -I${SRCDIR}/include -DOBK_BUILDING_ADDIN
#include <stdint.h>
#include "oblikovati_addin.h"

// cgo cannot call a Go-held C function pointer directly, so trampoline through C.
static int  host_call(ObkHostCall c, const char* m, const uint8_t* req, int n, uint8_t** resp, int* rl) { return c(m, req, n, resp, rl); }
static void host_free(ObkHostFree f, uint8_t* p) { f(p); }
*/
import "C"

import (
	"errors"
	"unsafe"
)

// cgoHostCaller implements client.Transport over the C-ABI host callbacks.
type cgoHostCaller struct{}

func (cgoHostCaller) Call(method string, req []byte) ([]byte, error) {
	mu.Lock()
	call, freeFn := hostCall, hostFree
	mu.Unlock()
	if call == nil {
		return nil, errors.New("host not connected")
	}
	cMethod := C.CString(method)
	defer C.free(unsafe.Pointer(cMethod))

	var reqPtr *C.uint8_t
	if len(req) > 0 {
		reqPtr = (*C.uint8_t)(unsafe.Pointer(&req[0]))
	}
	var resp *C.uint8_t
	var respLen C.int
	rc := C.host_call(call, cMethod, reqPtr, C.int(len(req)), &resp, &respLen)
	if resp == nil {
		if rc == C.OBK_OK {
			return []byte{}, nil
		}
		return nil, errors.New("host call failed")
	}
	out := C.GoBytes(unsafe.Pointer(resp), respLen)
	C.host_free(freeFn, resp) // host owns the buffer — always free it
	if rc != C.OBK_OK {
		return nil, errors.New(string(out)) // host-formatted error message
	}
	return out, nil
}
```

## 4. Register the button and run the action (`addin.go`)

This is the only file with domain logic. On construction it creates a typed `client.Client` and
registers a ribbon button. On the button's command-ended event, it builds the cube.

```go
package main

import (
	"encoding/json"

	"oblikovati.org/api/client"
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

const cubeCommandID = "com.example.cube-button.makeCube"

// AddIn holds the typed client over the host transport.
type AddIn struct{ c *client.Client }

// NewAddIn wires the client and places the ribbon button. Together
// Ribbon/Tab/Category decide where the button lands; Environment scopes it to a
// context (base = always shown).
func NewAddIn(t client.Transport) (*AddIn, error) {
	c := client.New(t)
	_, err := c.Commands().Create(wire.CreateCommandArgs{
		ID:          cubeCommandID,
		DisplayName: "Make Cube",
		Ribbon:      types.PartRibbon,   // the Part document ribbon
		Tab:         "Add-Ins",          // tab to place it on
		Category:    "Examples",         // panel within the tab
		Environment: types.BaseEnvironment,
		ButtonStyle: types.LargeIconButton,
		Tooltip:     "Sketch a rectangle and extrude a 50 mm cube",
	})
	if err != nil {
		return nil, err
	}
	return &AddIn{c: c}, nil
}

// OnEvent is called from ObkAddInNotify for every host event. When our command
// finishes (the user clicked the button), build the cube. The event is a JSON
// object whose command id identifies the finished command.
func (a *AddIn) OnEvent(ev []byte) {
	var e struct {
		Event   string `json:"event"`
		Command string `json:"command"`
	}
	if err := json.Unmarshal(ev, &e); err != nil {
		return
	}
	if e.Command == cubeCommandID {
		_ = a.makeCube()
	}
}

// makeCube creates a part, sketches a 50 mm square, and extrudes it 50 mm.
func (a *AddIn) makeCube() error {
	if _, err := a.c.Documents().Create(wire.CreateDocumentArgs{Type: "part", Name: "Cube"}); err != nil {
		return err
	}
	// 1) A sketch on the XY origin plane.
	sk, err := a.c.Sketch().Create(wire.CreateSketchArgs{Plane: "XY"})
	if err != nil {
		return err
	}
	// 2) A closed 50 mm × 50 mm rectangle (one profile).
	if _, err := a.c.Sketch().Rectangle(wire.SketchRectangleArgs{
		SketchIndex: sk.SketchIndex,
		Width:       "50 mm",
		Height:      "50 mm",
	}); err != nil {
		return err
	}
	// 3) Extrude that profile 50 mm into a new solid → a cube.
	extrude, _ := json.Marshal(map[string]any{
		"sketchIndex":  sk.SketchIndex,
		"profileIndex": 0,
		"distance":     "50 mm",
		"operation":    "new",
	})
	_, err = a.c.Features().Add(wire.AddFeatureArgs{Kind: "extrude", Args: extrude})
	return err
}
```

### Why the action lives in `OnEvent`

Registering a button does **not** attach host logic to it. When the user clicks, the host runs
the command (which does nothing of its own) and then fires a *command-ended* event. Your add-in
receives that event through `ObkAddInNotify`, matches it by command id, and runs the action —
typically more client calls, exactly as above.

> **Discovering the feature schema.** `extrude`'s arguments above
> (`sketchIndex`, `profileIndex`, `distance`, `operation`, …) are validated by the host. Call
> `client.Features().List()` at runtime to get every feature kind and its JSON Schema, so you
> never hard-code a shape you can't verify.

## 5. Build

```sh
go build -buildmode=c-shared -o cube-button.so .
```

Drop the resulting library where the host scans for add-ins (see your host's add-in directory),
restart, open a part, and you'll find **Make Cube** on the **Add-Ins** tab. Click it to get a
50 mm cube.

## Where to go next

- [[Oblikovati API Architecture]] — the full method surface and the typed client groups
  (`Documents`, `Parameters`, `Sketch`, `Features`, `Model`, `View`, `Materials`, …).
- [[Testing Automation]] — test `makeCube` with a fake transport (no host needed) and run a
  real integration test against the host.
- [[API Docs]] — the generated reference for every type and method used above.
