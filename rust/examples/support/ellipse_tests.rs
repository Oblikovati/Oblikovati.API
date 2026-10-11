// SPDX-License-Identifier: Apache-2.0
use super::*;
struct MinorRadiusHost;
impl Caller for MinorRadiusHost {
    fn call(&self, method: &str, payload: &[u8]) -> Result<Vec<u8>, Error> {
        assert_eq!(method, wire::METHOD_SKETCH_ADD_DIMENSION);
        let actual: serde_json::Value = serde_json::from_slice(payload).unwrap();
        assert_eq!(
            actual,
            serde_json::json!({"sketchIndex":2,"kind":"ellipseMinorRadius","entities":[42],"expression":"5 mm"})
        );
        Ok(br#"{"kind":"ellipseMinorRadius","value":0.5,"dof":0}"#.to_vec())
    }
}
#[test]
fn minor_radius_example_preserves_owner_kind_and_length_expression() {
    let reply = add_minor_radius(MinorRadiusHost, 2, 42, "5 mm").unwrap();
    assert_eq!(reply.kind, "ellipseMinorRadius");
    assert_eq!(reply.value, 0.5);
    assert_eq!(reply.dof, 0);
    assert_eq!(types::DIM_CONSTRAINT_ELLIPSE_RADIUS.0, "ellipseRadius");
}
