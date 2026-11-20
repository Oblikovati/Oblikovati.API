// SPDX-License-Identifier: Apache-2.0

use oblikovati_api::{
    Error,
    client::{Caller, Client},
    wire,
};
use serde_json::json;

struct ExpressionHost(&'static str);

impl Caller for ExpressionHost {
    fn call(&self, method: &str, request: &[u8]) -> Result<Vec<u8>, Error> {
        assert_eq!(method, self.0);
        let value: serde_json::Value = serde_json::from_slice(request).unwrap();
        for (key, expected) in [
            ("radius", "bore_r + 2 mm"),
            ("width", "slot_width"),
            ("majorRadius", "major_r"),
            ("minorRadius", "minor_r"),
            ("distance2", "chamfer_b"),
        ] {
            assert_eq!(value[key], json!(expected));
        }
        Ok(br#"{"entityId":1,"kind":"circle","pointIds":[]}"#.to_vec())
    }
}

#[test]
fn sketch_entity_client_preserves_size_expressions() {
    let request = wire::AddSketchEntityArgs {
        kind: "circle".into(),
        radius: "bore_r + 2 mm".into(),
        width: "slot_width".into(),
        major_radius: "major_r".into(),
        minor_radius: "minor_r".into(),
        distance2: "chamfer_b".into(),
        ..Default::default()
    };
    Client::new(ExpressionHost(wire::METHOD_SKETCH_ADD_ENTITY))
        .sketch()
        .add_entity(&request)
        .unwrap();
    Client::new(ExpressionHost(wire::METHOD_SKETCH_ADD_ASSOCIATIVE_ENTITY))
        .sketch()
        .add_associative_entity(&request)
        .unwrap();
}

struct LegacyHost;

impl Caller for LegacyHost {
    fn call(&self, method: &str, _request: &[u8]) -> Result<Vec<u8>, Error> {
        assert_eq!(method, wire::METHOD_SKETCH_ADD_ASSOCIATIVE_ENTITY);
        Err(Error("unknown method".into()))
    }
}

#[test]
fn associative_entity_preserves_legacy_host_error() {
    let request = wire::AddSketchEntityArgs {
        kind: "circle".into(),
        radius: "bore_r".into(),
        ..Default::default()
    };
    let error = Client::new(LegacyHost)
        .sketch()
        .add_associative_entity(&request)
        .unwrap_err();
    assert!(error.to_string().contains("unknown method"));
}
