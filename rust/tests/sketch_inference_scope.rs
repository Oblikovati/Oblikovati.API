// SPDX-License-Identifier: Apache-2.0
use oblikovati_api::{types, wire::SketchInferenceScopeView};

#[test]
fn selected_scope_round_trips_the_sketch_and_candidate_ids() {
    let value = serde_json::json!({
        "sketchIndex": 2,
        "scope": "selectedGeometry",
        "entities": [4, 9],
    });
    let scope: SketchInferenceScopeView = serde_json::from_value(value.clone()).unwrap();
    assert_eq!(scope.sketch_index, 2);
    assert_eq!(scope.scope, types::INFERENCE_SCOPE_SELECTED_GEOMETRY);
    assert_eq!(scope.entities.as_deref(), Some(&[4, 9][..]));
    assert_eq!(serde_json::to_value(scope).unwrap(), value);
}

#[test]
fn all_and_current_modes_have_distinct_wire_values() {
    assert_eq!(types::INFERENCE_SCOPE_CURRENT_COMMAND.0, "currentCommand");
    assert_eq!(types::INFERENCE_SCOPE_ALL_GEOMETRY.0, "allGeometry");
}
