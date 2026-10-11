# Ellipse Dimensions

Use `ellipseRadius` for the major radius and `ellipseMinorRadius` for the minor
radius. Both kinds take one owned ellipse entity ID and a length expression.
The major radius must remain greater than or equal to the positive minor radius.
A radius crossing is an error; the host must preserve the axis identities.

Both kinds use `sketch.addDimension`. Set `driven` to create a reference dimension
that measures geometry without constraining it. Annotation placement, deletion,
and dimension drives use the existing dimension methods. A dimension drive
resolves an expression once; use `parameters.set` on its named Model parameter
to retain a live formula.

The examples require an existing ellipse in the selected sketch. They create its
minor-radius dimension through the typed client. This contract addition does not
prove that a particular host implements it. An older host can refuse the new kind;
it must not silently create a major-radius dimension.

## Go

<!-- example: go/examples/ellipse/ellipse.go -->
```go
// SPDX-License-Identifier: Apache-2.0

// Package ellipse demonstrates independent ellipse radius dimensions.
package ellipse

import (
	"oblikovati.org/api/client"
	"oblikovati.org/api/wire"
)

// AddMinorRadius drives an existing owned ellipse's minor radius.
// Example: AddMinorRadius(host, 0, ellipseID, "5 mm").
func AddMinorRadius(host client.Caller, sketchIndex int, ellipseID uint64, expression string) (wire.AddDimensionResult, error) {
	return client.New(host).Sketch().Dimension(sketchIndex).EllipseMinorRadius(ellipseID, expression)
}
```
<!-- /example -->

## Rust

<!-- example: rust/examples/ellipse.rs -->
```rust
// SPDX-License-Identifier: Apache-2.0
use oblikovati_api::{
    Error,
    client::{Caller, Client},
    types, wire,
};

/// Drive the minor radius of an existing owned ellipse. Example: `add_minor_radius(host, 0, 42, "5 mm")?`.
pub fn add_minor_radius<T: Caller>(
    host: T,
    sketch_index: i64,
    ellipse_id: u64,
    expression: &str,
) -> Result<wire::AddDimensionResult, Error> {
    Client::new(host).dimension().add(&wire::AddDimensionArgs {
        sketch_index,
        kind: types::DIM_CONSTRAINT_ELLIPSE_MINOR_RADIUS.0.to_string(),
        entities: Some(vec![ellipse_id]),
        expression: expression.into(),
        ..Default::default()
    })
}
fn main() {
    println!("Call add_minor_radius with your add-in's HostTransport and an owned ellipse.");
}

#[cfg(test)]
#[path = "support/ellipse_tests.rs"]
mod tests;
```
<!-- /example -->
