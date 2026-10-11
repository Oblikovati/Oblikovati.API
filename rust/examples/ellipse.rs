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
