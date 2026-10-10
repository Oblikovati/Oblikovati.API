// SPDX-License-Identifier: Apache-2.0

use serde::{Serialize, de::DeserializeOwned};

/// The only host dependency. Implement with C callbacks or a named test fake.
pub trait Caller {
    /// Send one request. Example: `transport.call("application.apiVersion", &[])`.
    fn call(&self, method: &str, request: &[u8]) -> Result<Vec<u8>, crate::Error>;
}

/// Typed access to the host through an injected transport.
pub struct Client<T: Caller> {
    transport: T,
}

impl<T: Caller> Client<T> {
    /// Bind a transport. Example: `let client = Client::new(transport)`.
    pub fn new(transport: T) -> Self {
        Self { transport }
    }

    pub(super) fn invoke<R: DeserializeOwned, Q: Serialize>(
        &self,
        method: &str,
        request: Option<&Q>,
    ) -> Result<R, crate::Error> {
        let response = self.send(method, request)?;
        serde_json::from_slice(&response).map_err(|error| {
            crate::Error(format!(
                "method {method:?} response {:?}: expected its documented JSON result: {error}",
                String::from_utf8_lossy(&response)
            ))
        })
    }

    pub(super) fn invoke_void<Q: Serialize>(
        &self,
        method: &str,
        request: Option<&Q>,
    ) -> Result<(), crate::Error> {
        self.send(method, request).map(|_| ())
    }

    fn send<Q: Serialize>(
        &self,
        method: &str,
        request: Option<&Q>,
    ) -> Result<Vec<u8>, crate::Error> {
        let bytes = request
            .map(serde_json::to_vec)
            .transpose()
            .map_err(|error| {
                crate::Error(format!(
                    "method {method:?}: expected a JSON-serializable request: {error}"
                ))
            })?
            .unwrap_or_default();
        self.transport.call(method, &bytes)
    }
}
