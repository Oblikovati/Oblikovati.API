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
