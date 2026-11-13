// SPDX-License-Identifier: Apache-2.0
//! The move-policy wire fields remain optional for existing clients.
use oblikovati_api::{
    types::{SKETCH_MOVE_DIMENSION_IF_NO_EQUATION, SKETCH_MOVE_GEOMETRIC_PROMPT},
    wire::{SketchMoveDecision, TransformSketchArgs, TransformSketchResult},
};

#[test]
fn move_policies_round_trip_and_empty_policies_are_omitted() {
    let mut request = TransformSketchArgs {
        op: "move".into(),
        entities: Some(vec![1]),
        vector: Some(vec![1., 2.]),
        ..Default::default()
    };
    let old = serde_json::to_value(&request).unwrap();
    assert!(old.get("dimensionPolicy").is_none());
    assert!(old.get("geometricPolicy").is_none());
    request.dimension_policy = SKETCH_MOVE_DIMENSION_IF_NO_EQUATION;
    request.geometric_policy = SKETCH_MOVE_GEOMETRIC_PROMPT;
    let json = serde_json::to_value(&request).unwrap();
    assert_eq!(json["dimensionPolicy"], "ifNoEquation");
    assert_eq!(json["geometricPolicy"], "prompt");
    assert_eq!(
        serde_json::from_value::<TransformSketchArgs>(json).unwrap(),
        request
    );
    let old_request: TransformSketchArgs = serde_json::from_value(old).unwrap();
    assert_eq!(old_request.dimension_policy, Default::default());
    assert_eq!(old_request.geometric_policy, Default::default());
}

#[test]
fn move_decision_is_typed_and_distinct_from_a_committed_copy() {
    let reply = TransformSketchResult {
        decision_required: Some(Box::new(SketchMoveDecision {
            dimension: true,
            geometric: false,
        })),
        ..Default::default()
    };
    let json = serde_json::to_value(&reply).unwrap();
    assert_eq!(json["decisionRequired"]["dimension"], true);
    assert!(json.get("created").is_none());
    assert_eq!(
        serde_json::from_value::<TransformSketchResult>(json).unwrap(),
        reply
    );
    assert!(
        serde_json::to_value(TransformSketchResult::default())
            .unwrap()
            .get("decisionRequired")
            .is_none()
    );
}
