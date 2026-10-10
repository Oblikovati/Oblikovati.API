// SPDX-License-Identifier: Apache-2.0

mod client_generated;

struct RecordingCaller {
    method: &'static str,
    request: Vec<u8>,
}

impl RecordingCaller {
    fn new(method: &'static str, request: Vec<u8>) -> Self {
        Self { method, request }
    }
}

impl oblikovati_api::client::Caller for RecordingCaller {
    fn call(&self, method: &str, request: &[u8]) -> Result<Vec<u8>, oblikovati_api::Error> {
        assert_eq!(method, self.method);
        assert_eq!(request, self.request);
        Err(oblikovati_api::Error("recorded".into()))
    }
}
