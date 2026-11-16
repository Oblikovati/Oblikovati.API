// SPDX-License-Identifier: Apache-2.0
use oblikovati_api::{types, wire};

#[test]
fn plan_uses_explicit_policies_and_distinct_wire_methods() {
    let args = wire::PlanSketchMoveArgs {
        sketch_index: 1,
        entities: Some(vec![3]),
        vector: Some(vec![1.0, 2.0]),
        dimension_policy: types::SKETCH_MOVE_DIMENSION_PROMPT,
        geometric_policy: types::SKETCH_MOVE_GEOMETRIC_NEVER,
    };
    let json = serde_json::to_value(args).unwrap();
    assert_eq!(json["dimensionPolicy"], "prompt");
    assert_eq!(json["geometricPolicy"], "never");
    assert_ne!(wire::METHOD_SKETCH_PLAN_MOVE, wire::METHOD_SKETCH_TRANSFORM);
    assert_ne!(
        wire::METHOD_SKETCH_COMMIT_MOVE,
        wire::METHOD_SKETCH_TRANSFORM
    );
}

#[test]
fn impact_and_approval_round_trip() {
    let plan: wire::PlanSketchMoveResult = serde_json::from_str(
        r#"{"token":"opaque","dimensions":[{"index":2,"beforeExpression":"10 mm","afterExpression":"12 mm","approvalRequired":true}],"geometric":[]}"#,
    )
    .unwrap();
    assert_eq!(plan.token, "opaque");
    assert!(plan.dimensions.unwrap()[0].approval_required);
    let commit = wire::CommitSketchMoveArgs {
        token: "opaque".into(),
        approve_dimensions: true,
        approve_geometric: false,
    };
    let decoded: wire::CommitSketchMoveArgs =
        serde_json::from_value(serde_json::to_value(&commit).unwrap()).unwrap();
    assert_eq!(decoded, commit);
    assert!(
        !serde_json::from_str::<wire::CommitSketchMoveResult>("{}")
            .unwrap()
            .committed
    );
}
