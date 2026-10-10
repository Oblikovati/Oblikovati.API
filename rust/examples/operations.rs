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
