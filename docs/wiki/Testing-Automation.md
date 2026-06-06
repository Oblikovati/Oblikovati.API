# Testing Automation

How to test an add-in at two levels — fast unit tests with no host, and integration tests
against the live application — while keeping the **Apache-2.0 / GPL-2.0 license boundary**
intact.

Tests should be **F.I.R.S.T**: Fast, Independent, Repeatable, Self-validating, Timely. Mock
external I/O (here, the host) with **named fake types**, not inline stubs.

## The license boundary, in one rule

> The **shipped** add-in library links **only** `oblikovati/api` (Apache-2.0). It must never
> link the GPL application.

Add-in logic talks to the host through the `client.Transport` interface, so it has **no compile
dependency on the host at all** — which is exactly what makes it unit-testable with a fake. The
host only appears when you want a real end-to-end test, and then strictly as a **test-scope**
dependency that never reaches the shipped artifact.

## Level 1 — Unit tests (no host, link only the contract)

Because your logic depends on `client.Transport`, you can drive it with a fake that records calls
and returns canned replies. No process, no model, milliseconds to run.

```go
package main

import (
	"encoding/json"
	"testing"

	"oblikovati/api/wire"
)

// fakeTransport is a named test double for client.Transport: it records every method
// call and returns a scripted reply per method. (Mock I/O with a named fake, not an
// inline closure.)
type fakeTransport struct {
	calls   []string                  // method names, in order
	replies map[string][]byte         // method -> canned JSON reply
}

func (f *fakeTransport) Call(method string, _ []byte) ([]byte, error) {
	f.calls = append(f.calls, method)
	if r, ok := f.replies[method]; ok {
		return r, nil
	}
	return []byte("{}"), nil
}

func TestMakeCubeDrivesTheModel(t *testing.T) {
	sketchReply, _ := json.Marshal(wire.CreateSketchResult{SketchIndex: 0, Plane: "XY"})
	ft := &fakeTransport{replies: map[string][]byte{
		wire.MethodSketchCreate: sketchReply,
	}}

	a, err := NewAddIn(ft) // registers the button via commands.create
	if err != nil {
		t.Fatalf("NewAddIn: %v", err)
	}
	if err := a.makeCube(); err != nil {
		t.Fatalf("makeCube: %v", err)
	}

	// Assert the add-in issued the operations in the right order.
	want := []string{
		wire.MethodCommandsCreate, // from NewAddIn
		wire.MethodDocumentsCreate,
		wire.MethodSketchCreate,
		wire.MethodSketchRectangle,
		wire.MethodFeaturesAdd,
	}
	if got := ft.calls; !equal(got, want) {
		t.Errorf("call sequence =\n  %v\nwant\n  %v", got, want)
	}
}
```

This pins your add-in's **intent** — which methods it calls, with what arguments, in what order —
without any geometry. Every new behavior gets a test like this; every bug fix gets a regression
test that reproduces the bad call sequence first.

What to assert at this level:

- the **method sequence** and the **arguments** marshalled onto each call;
- error handling when the transport returns an error (return early, don't panic);
- that you read schemas via `client.Features().List()` rather than hard-coding shapes that could
  drift.

## Level 2 — Integration tests (against the live host)

Unit tests prove you *call* the host correctly; integration tests prove the host *does the right
thing* — the sketch solves, the feature builds, the geometry is valid. These need a real host, so
they pull the GPL application in as a **test-only** dependency.

Keep them in a separate, clearly-scoped place (a test build tag or a dedicated test module) so the
GPL requirement is visibly test-scope and never part of the shipped graph. The reference
add-in does exactly this: its integration tests drive the live host router, while the shipped
library's dependency graph stays contract-only.

A typical integration test:

1. start (or connect to) a host session and obtain a `Transport`;
2. run your add-in action (`makeCube`);
3. read back the result and assert **topology + geometry**: the body is manifold and closed, and
   its volume/area match the expected values within a tolerance (tessellation makes geometry
   approximate, so compare with relative error, not equality);
4. optionally change a parameter and re-read to prove the **parametric** path recomputes.

```go
// pseudocode — wiring depends on how you host the session in tests
func TestMakeCubeProducesA50mmCube(t *testing.T) {
	tr := startHostSessionForTest(t)        // GPL host, TEST-SCOPE ONLY
	a, _ := NewAddIn(tr)
	if err := a.makeCube(); err != nil {
		t.Fatal(err)
	}
	props := getPhysicalProperties(t, tr)   // via client.Model().PhysicalProperties()
	const want = 5.0 * 5.0 * 5.0            // cm³ (50 mm cube)
	if relErr(props.Volume, want) > 1e-3 {
		t.Errorf("volume = %.4f cm³, want ~%.1f", props.Volume, want)
	}
}
```

## Verify the boundary in CI

Make the license boundary a **test that fails the build**, not a convention. The shipped
package's dependency graph must contain no application packages:

```sh
# From the add-in's module: expect NO output.
go list -deps . | grep -E '^oblikovati(/|$)'
```

If your integration tests require the GPL host, guard them so this check runs on the shipped
build only (e.g. the integration tests live behind a build tag, or in a separate module whose
`require` on the application is plainly test-scope). The contract module itself enforces the
mirror-image rule: `oblikovati/api` is checked to never import the application.

## Testing the contract itself

If you contribute to `oblikovati/api`, note how it is tested:

- `types` and `wire` are **pure data** — tested by pinning frozen ids and JSON round-trips.
- `client` is tested with a **fake transport** (the same technique as above), so it needs no host.

```sh
cd Oblikovati.API
go build ./... && go vet ./... && go test ./...
```

## Checklist

- [ ] Add-in logic depends on `client.Transport`, never on the host directly.
- [ ] Every behavior has a fake-transport unit test; every bug fix has a regression test.
- [ ] Integration tests assert valid topology **and** geometry within tolerance.
- [ ] `go list -deps` on the shipped build shows **no** application packages.
- [ ] Sketches are asserted **fully constrained** (0 DOF) before a feature consumes them.
