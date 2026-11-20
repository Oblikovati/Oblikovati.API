// SPDX-License-Identifier: Apache-2.0
use oblikovati_api::wire::PartOptionsView;

#[test]
fn part_sketch_view_keeps_omitted_false_and_true_distinct() {
    for (json, expected) in [
        (r#"{"chamferFlatCorners":false}"#, None),
        (
            r#"{"chamferFlatCorners":false,"lookAtSketchPlane":false}"#,
            Some(false),
        ),
        (
            r#"{"chamferFlatCorners":false,"lookAtSketchPlane":true}"#,
            Some(true),
        ),
        (
            r#"{"chamferFlatCorners":false,"lookAtSketchPlane":null}"#,
            None,
        ),
    ] {
        let view: PartOptionsView = serde_json::from_str(json).unwrap();
        assert_eq!(view.look_at_sketch_plane.as_deref().copied(), expected);
        let saved = serde_json::to_value(view).unwrap();
        assert_eq!(
            saved
                .get("lookAtSketchPlane")
                .and_then(|value| value.as_bool()),
            expected
        );
    }
}
