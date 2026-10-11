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

fn main() {
    println!("Call list_all_axes or list_all_points with your add-in's HostTransport.");
}

#[cfg(test)]
#[path = "support/datums_tests.rs"]
mod tests;
