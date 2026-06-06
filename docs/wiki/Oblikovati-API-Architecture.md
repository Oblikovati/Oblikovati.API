# Oblikovati API Architecture

The public automation contract is the standalone Go module **`oblikovati/api`**, licensed
**Apache-2.0**. It is the single source of truth for the API surface: the host implements it,
and add-ins build against it. This page explains its shape and how to use it well.

## Invariant: the contract never imports the application

`oblikovati/api` must **never** import the GPL application module. The dependency flows one way
only — the app depends on the contract, not the reverse — and CI fails the build if it is ever
violated. This is what lets a closed-source add-in link the contract without touching GPL code.

## The four packages

| Package | Holds | Use it for |
|---|---|---|
| `types` | enums, stable ids, value/option structs — **pure data** | the shared vocabulary (document kinds, parameter kinds, 2D points, ribbon keys, …) |
| `contract` | in-process Go interfaces (`Document`, `Parameter`, `Sketch`, …) | **first-party, in-process** code only (one Go runtime) |
| `wire` | method-name constants + JSON request/response DTOs | the host↔add-in JSON contract |
| `client` | a `Transport` interface + a typed client over it | **out-of-runtime add-ins** driving the host |

### `types` — the vocabulary

Pure data with stable JSON tags and frozen string/numeric values. Define a concept once here and
both sides speak it. Examples: `DocumentType`, `ParameterKind`, `RibbonKey`, `ButtonStyle`,
`Environment`, `WorkPlaneKind`, the graphics and lighting enums. The application aliases each
(`type X = types.X`) so it shares the exact definition.

### `contract` — in-process interfaces

Scalar Go interfaces (`Application`, `Document`, `PartDocument`, `Parameters`, `Parameter`,
`Sketch`, …) that the host satisfies with compile-time assertions. These are for code running in
**the same Go runtime as the host** (first-party). A C-ABI add-in runs its *own* runtime, so a
live Go interface value cannot cross to it — add-ins use `wire` + `client` instead.

### `wire` — the JSON contract

The bytes that actually cross the boundary:

- **Method-name constants** (`wire.MethodSketchCreate = "sketch.create"`, …). The host router
  keys its dispatch on these; treat the string values as **frozen**.
- **Request/response DTOs** (`CreateSketchArgs`, `AddFeatureArgs`, `DocumentInfo`, …) with stable
  JSON tags. A field rename is a breaking change.

Operations are grouped by prefix: `documents.*`, `parameters.*`, `model.*`, `sketch.*`,
`sketch3d.*`, `features.*`, `workPlanes.*`, `commands.*`, `ribbon.*`, `view.*`, `lighting.*`,
`environment.*`, `appearances.*`, `materials.*`, `clientGraphics.*`, `interactionGraphics.*`,
`transaction.*`, `logs.*`.

### `client` — the typed façade

`client.Client` wraps a `Transport` and gives you typed method groups instead of hand-rolled
JSON. Each group mirrors a wire prefix:

```go
c := client.New(transport)

c.Documents()   // create / list / activate documents
c.Parameters()  // add / get / set / list named parameters
c.Sketch()      // create sketches, add geometry, constrain, solve, list profiles
c.Sketch3D()    // the 3D-sketch equivalent
c.Features()    // list feature kinds (+schema) and add features
c.Model()       // model tree, selection, reference keys, physical properties
c.WorkPlanes()  // datum-plane construction
c.Commands()    // list / execute / create ribbon commands
c.Ribbon()      // discover the active ribbon's tabs/panels/controls
c.View()        // get/set display mode and shadows
c.Lighting()    // lighting style + discrete lights
c.Materials()   // material library + assignment
c.Theme(), c.Appearances(), c.Graphics(), c.Transactions() // …and more
```

## The one dependency: `Transport`

The client's only requirement is a `Transport`:

```go
type Transport interface {
	Call(method string, req []byte) ([]byte, error)
}
```

A C-ABI add-in backs it with the host's `ObkHostCall` callback (see [[First Steps]]); a test
backs it with a fake. Everything else in `client` is built on this single method.

## The two consumption paths

A C-shared add-in runs its **own Go runtime** — two runtimes in one process, so a live Go
pointer/interface must not cross the boundary. The contract therefore serves two audiences:

- **In-process / first-party** code uses `contract` interfaces + `types` directly.
- **Out-of-runtime add-ins** use `wire` + `client` over a `Transport`, where the only thing that
  crosses the boundary is JSON. A future gRPC/socket transport can sit behind the same `wire`
  surface unchanged.

```
First-party (same runtime)            Add-in (own runtime)
   contract + types                      wire + client
        │                                     │
        ▼                                     ▼
   ┌──────────────────────────────────────────────┐
   │              Oblikovati host                   │
   │   implements contract · serves wire methods    │
   └──────────────────────────────────────────────┘
```

## Adding to the API surface

Every change is **two parts, in this order**:

1. **Contract first, in `oblikovati/api`:**
   - enum / value type → `types` (define it **once** here),
   - in-proc Go interface → `contract`,
   - method-name constant + request/response DTOs → `wire`,
   - a typed method group → `client` for any new wire method.
2. **Implementation in the application:** build the behavior, satisfy the `contract` interface
   with a compile-time assertion (`var _ contract.X = (*impl.X)(nil)`), and wire the handler into
   the host router keyed on the `wire` method constant.

Rules of thumb:

- Never re-declare a DTO or method string outside `wire` — import it.
- Never call the host from an add-in with raw JSON — use `client`.
- Every exported `.go` file carries an `SPDX-License-Identifier: Apache-2.0` header.

## Conventions you'll rely on

- **Units are explicit.** Lengths are unit-bearing expression strings (`"40 mm"`, `"5 cm"`).
- **Database units are centimetres.** Physical properties come back in cm³ / cm² (and grams with
  a density).
- **References are strings.** Faces/edges/planes are addressed by reference keys the host tracks
  across recompute.
- **Stable identities.** Method strings and enum values are frozen — saved automations depend on
  them.
