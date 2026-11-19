// SPDX-License-Identifier: Apache-2.0
//! The typed client must carry an explicit default separately from an override.
use oblikovati_api::{
    Error,
    client::{Caller, Client},
    types, wire,
};

struct BlockColorHost;
impl Caller for BlockColorHost {
    fn call(&self, method: &str, request: &[u8]) -> Result<Vec<u8>, Error> {
        let args: serde_json::Value = serde_json::from_slice(request).unwrap();
        if method == wire::METHOD_SKETCH_SET_BLOCK_INSTANCE_COLOR {
            assert_eq!(
                args,
                serde_json::json!({"sketchIndex": 2, "instanceId": 42, "color": null})
            );
            return Ok(br#"{"id":42,"definition":"Rib"}"#.to_vec());
        }
        assert_eq!(method, wire::METHOD_SKETCH_SET_BLOCK_DEFINITION_COLOR);
        assert_eq!(args["color"]["g"], 255);
        Ok(
            br#"{"name":"Rib","color":{"r":0,"g":255,"b":255,"opacity":1,"source":79105}}"#
                .to_vec(),
        )
    }
}
#[test]
fn typed_color_commands_carry_cyan_and_null_default() {
    let client = Client::new(BlockColorHost);
    let cyan = types::Color {
        r: 0,
        g: 255,
        b: 255,
        opacity: 1.0,
        source: types::OVERRIDE_COLOR_SOURCE,
    };
    let definition = client
        .sketch()
        .set_block_definition_color(&wire::SetBlockDefinitionColorArgs {
            name: "Rib".into(),
            color: Some(Box::new(cyan.clone())),
        })
        .unwrap();
    assert_eq!(definition.color.as_deref(), Some(&cyan));
    let instance = client
        .sketch()
        .set_block_instance_color(&wire::SetBlockInstanceColorArgs {
            sketch_index: 2,
            instance_id: 42,
            color: None,
        })
        .unwrap();
    assert!(instance.color.is_none());
}
