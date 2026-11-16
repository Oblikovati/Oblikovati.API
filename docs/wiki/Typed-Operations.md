# Typed Operations

Use the typed client for host calls. Inject a `Caller` in Go or Rust. In an add-in,
the caller uses the host callbacks supplied at activation. Tests can supply a
named fake caller.

Both examples create a part and provide a function to extrude a profile. The
extrude function requires an active part with a closed, fully constrained sketch.
Create and constrain the sketch before you call it. Dimensions use expressions,
such as `50 mm` or the name of a document parameter.

## Express assemblies and save warnings

Set `Express` in `OpenDocumentArgs` to request an assembly's validated saved
presentation. The response's `Express` field reports whether component geometry
is deferred. A missing or stale presentation is an error; the host also rejects
`Express` together with `DeferContent`. Call the typed `LoadFull` operation with
the returned document ID to load the components. If loading fails, the assembly
remains in Express mode. These fields are omitted when false so older requests
and responses retain their meaning.

Save, batch-save, and save-on-close results can include `SaveWarning` values.
For `assemblyCacheUnavailable`, the document was saved but its optional Express
presentation could not be qualified. Inspect the warning detail before relying
on a later Express open. A warning does not turn a successful save into an error.

## Part sketch preferences

Part options may carry `startupSketchPlane` (`none`, `xy`, `yz`, or `xz`) and
`lookAtSketchPlane` (a Boolean). An omitted or null field leaves the existing
preference unchanged on write and means unavailable on read. An explicit
`false` remains distinct from omission. The typed write clients reject unknown
startup-plane values before transport; hosts must also validate requests from
all callers. JSON decoding retains unknown values for forward-compatible reads.

## Go

<!-- example: go/examples/operations/operations.go -->
```go
// SPDX-License-Identifier: Apache-2.0

// Package operations demonstrates typed requests through an injected host.
package operations

import (
	"encoding/json"
	"oblikovati.org/api/client"
	"oblikovati.org/api/wire"
	"oblikovati.org/api/wire/featureargs"
)

// CreatePart opens a part document through the host. Example: CreatePart(host, "Bracket").
func CreatePart(host client.Caller, name string) (wire.DocumentInfo, error) {
	return client.New(host).Documents().Create(wire.CreateDocumentArgs{Type: "part", Name: name})
}

// ExtrudeProfile extrudes a closed, fully constrained sketch profile in the active part.
// Example: ExtrudeProfile(host, 0, "50 mm"). Create the sketch before this call.
func ExtrudeProfile(host client.Caller, sketchIndex int, distance string) (json.RawMessage, error) {
	args := featureargs.Extrude{SketchIndex: sketchIndex, Distance: distance, Operation: "new"}
	return client.AddFeature(client.New(host).Features(), args)
}
```
<!-- /example -->

## Rust

<!-- example: rust/examples/operations.rs -->
```rust
// SPDX-License-Identifier: Apache-2.0

use oblikovati_api::{
    Error,
    client::{Caller, Client},
    json::JsonValue,
    wire::{self, featureargs},
};

/// Open a part through the host. Example: `create_part(host, "Bracket")?`.
pub fn create_part<T: Caller>(host: T, name: &str) -> Result<wire::DocumentInfo, Error> {
    Client::new(host)
        .documents()
        .create(&wire::CreateDocumentArgs {
            type_: "part".into(),
            name: name.into(),
            ..Default::default()
        })
}

/// Extrude an existing closed, fully constrained profile in the active part.
/// Example: `extrude_profile(host, 0, "50 mm")?`. Create the sketch first.
pub fn extrude_profile<T: Caller>(
    host: T,
    sketch_index: i64,
    distance: &str,
) -> Result<JsonValue, Error> {
    let args = featureargs::Extrude {
        sketch_index,
        distance: distance.into(),
        operation: "new".into(),
        ..Default::default()
    };
    Client::new(host).features().add_feature(&args)
}

fn main() {
    println!("Call create_part or extrude_profile with your add-in's HostTransport.");
}

#[cfg(test)]
#[path = "support/operations_tests.rs"]
mod tests;
```
<!-- /example -->

`AddFeature` and `add_feature` take a typed `featureargs` value. Its `Kind` or
`kind` method supplies the operation name. The client builds the wire envelope.
The feature result is an opaque JSON value because each operation has a different
result shape.

Rust client methods take references to the wire request structs. Go provides some
scalar convenience methods as well. Use the language reference to find the exact
signature. The method name and JSON field names on the wire are the same.

## Datum list options

`WorkPlanes.ListWithOptions`, `WorkAxes.ListWithOptions`, and
`WorkPoints.ListWithOptions` in Go correspond to the Rust `list_with_options`
methods on `work_planes()`, `work_axes()`, and `work_points()`.
Use the existing typed list request to include construction datums or select an
occurrence. The host must support the requested occurrence scope. The Rust
application currently supports the active local part only.

The no-argument list methods keep their empty request and default filtering.
Including construction rows does not change visibility. Keep each returned
index and reference; filtering can leave gaps in collection indices.

### Go datum lists

<!-- example: go/examples/operations/datums.go -->
```go
// SPDX-License-Identifier: Apache-2.0

package operations

import (
	"oblikovati.org/api/client"
	"oblikovati.org/api/wire"
)

// ListAllAxes includes construction axes in the active local model.
// Example: ListAllAxes(host). Use Occurrence only when the host supports that scope.
func ListAllAxes(host client.Caller) (wire.ListWorkAxesResult, error) {
	return client.New(host).WorkAxes().ListWithOptions(wire.ListWorkAxesArgs{IncludeConstruction: true})
}

// ListAllPoints includes construction points without changing their visibility.
// Example: ListAllPoints(host).
func ListAllPoints(host client.Caller) (wire.ListWorkPointsResult, error) {
	return client.New(host).WorkPoints().ListWithOptions(wire.ListWorkPointsArgs{IncludeConstruction: true})
}

// ListAllPlanes includes construction planes in the active local model.
// Example: ListAllPlanes(host).
func ListAllPlanes(host client.Caller) (wire.ListWorkPlanesResult, error) {
	return client.New(host).WorkPlanes().ListWithOptions(wire.ListWorkPlanesArgs{IncludeConstruction: true})
}
```
<!-- /example -->

### Rust datum lists

<!-- example: rust/examples/datums.rs -->
```rust
// SPDX-License-Identifier: Apache-2.0
use oblikovati_api::{
    Error,
    client::{Caller, Client},
    wire,
};

/// Include construction axes in the active local model. Example: `list_all_axes(host)?`.
pub fn list_all_axes<T: Caller>(host: T) -> Result<wire::ListWorkAxesResult, Error> {
    Client::new(host)
        .work_axes()
        .list_with_options(&wire::ListWorkAxesArgs {
            include_construction: true,
            ..Default::default()
        })
}

/// Include construction points without changing visibility. Example: `list_all_points(host)?`.
pub fn list_all_points<T: Caller>(host: T) -> Result<wire::ListWorkPointsResult, Error> {
    Client::new(host)
        .work_points()
        .list_with_options(&wire::ListWorkPointsArgs {
            include_construction: true,
            ..Default::default()
        })
}

/// Include construction planes in the active local model. Example: `list_all_planes(host)?`.
pub fn list_all_planes<T: Caller>(host: T) -> Result<wire::ListWorkPlanesResult, Error> {
    Client::new(host)
        .work_planes()
        .list_with_options(&wire::ListWorkPlanesArgs {
            include_construction: true,
            ..Default::default()
        })
}

fn main() {
    println!("Call the datum list examples with your add-in's HostTransport.");
}

#[cfg(test)]
#[path = "support/datums_tests.rs"]
mod tests;
```
<!-- /example -->

## Planning a constraint-aware sketch move

`Sketch.PlanMove` and `Sketch.CommitMove` are distinct from the legacy
`sketch.transform` operation. An older host must return an unknown-method error;
the client must not fall back to the legacy operation with policy fields that
the host could ignore. These API methods describe a contract; a host that does
not implement a policy must return an error.

A plan request supplies a nonempty selection, a finite sketch-plane vector in
centimetres, and both explicit policies. It does not alter the document. The
reply identifies all proposed dimension-expression changes and geometric
relation removals, including indirect effects, and marks the changes that need
approval. The opaque token is bound to the exact plan and document state.

The caller may inspect the effects and either discard the token or commit it.
Commit requires approval for every marked effect, and must refuse a missing,
expired, used or stale token, a changed document, unsupported policy, or any
invalid result without changing the document or history. A successful commit
is one undoable edit: Undo and Redo restore geometry, constraints and
dimensions together. A noninteractive caller chooses explicit non-prompt
policies or makes an explicit decision after reading the plan; the host does
not infer permission to remove a constraint.

`Never` preserves its constraint category. `If No Equation` preserves
dimensions that depend on other dimensions; unknown dependencies are refused.
Dimensional `Always` permits recomputing dimensions that cross the selection
boundary, while retaining dimensions within the selection. Geometric `Always`
permits removal of fixed constraints on selected geometry and constraints that
cross the selection boundary, except parallel, perpendicular and equal.
Relations inside the selection remain. The plan must enumerate those effects.
`Prompt` requires
approval for every proposed change in that category. The source scenarios do
not establish every policy combination or GUI prompt time; this contract does
not prescribe either. The [Inventor Move reference](https://help.autodesk.com/cloudhelp/2024/ENU/Inventor-Help/files/GUID-F8114270-18F7-4C7A-947E-AC489EDACFA9.htm)
describes its selection-specific policy behavior and provides design context.

## Events and identifiers

The C host sends events through `ObkAddInNotify`. Decode the event with the
corresponding `wire` event type in Go or Rust. Use the event constants from that
package to select its type. Do not retain a pointer to the notification buffer.

Document and entity identifiers use integers. Keep the full integer value; do
not convert an identifier to a floating-point number. Unknown numeric enum values
remain representable in Rust so a newer host can return an added value.

See [[C ABI Add-ins]] for buffer ownership and [[Testing Automation]] for the
tests that check both language examples.
