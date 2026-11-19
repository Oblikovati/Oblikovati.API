// SPDX-License-Identifier: Apache-2.0
//! Verify the typed client preserves instance identity and resulting references.
use oblikovati_api::{
    Error,
    client::{Caller, Client},
    wire,
};

struct ExplodeHost;
impl Caller for ExplodeHost {
    fn call(&self, method: &str, request: &[u8]) -> Result<Vec<u8>, Error> {
        assert_eq!(method, wire::METHOD_SKETCH_EXPLODE_BLOCK_INSTANCE);
        let value: serde_json::Value = serde_json::from_slice(request).unwrap();
        assert_eq!(
            value,
            serde_json::json!({"sketchIndex": 2, "instanceId": 41})
        );
        Ok(br#"{"entityRefs":[17,18,19]}"#.to_vec())
    }
}

#[test]
fn typed_explode_preserves_the_instance_and_new_entity_identities() {
    let result = Client::new(ExplodeHost)
        .sketch()
        .explode_block_instance(&wire::ExplodeSketchBlockArgs {
            sketch_index: 2,
            instance_id: 41,
        })
        .unwrap();
    assert_eq!(result.entity_refs, Some(vec![17, 18, 19]));
}
