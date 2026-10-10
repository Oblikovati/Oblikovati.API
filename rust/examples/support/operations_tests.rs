// SPDX-License-Identifier: Apache-2.0

use super::*;
use serde_json::{Value, json};

struct OperationHost {
    method: &'static str,
    request: Value,
    response: &'static str,
}

impl Caller for OperationHost {
    fn call(&self, method: &str, request: &[u8]) -> Result<Vec<u8>, Error> {
        assert_eq!(method, self.method);
        assert_eq!(
            serde_json::from_slice::<Value>(request).unwrap(),
            self.request
        );
        Ok(self.response.as_bytes().to_vec())
    }
}

#[test]
fn create_part_sends_a_typed_document_request() {
    let host = OperationHost {
        method: wire::METHOD_DOCUMENTS_CREATE,
        request: json!({"type":"part", "name":"Bracket"}),
        response: r#"{"id":7,"name":"Bracket","type":"part"}"#,
    };
    assert_eq!(create_part(host, "Bracket").unwrap().id, 7);
}

#[test]
fn extrude_profile_uses_typed_feature_arguments() {
    let host = OperationHost {
        method: wire::METHOD_FEATURES_ADD,
        request: json!({"kind":"extrude", "args":{"sketchIndex":2,"profileIndex":0,"distance":"50 mm","operation":"new"}}),
        response: "{}",
    };
    assert!(extrude_profile(host, 2, "50 mm").is_ok());
}
