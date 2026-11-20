// SPDX-License-Identifier: Apache-2.0
use oblikovati_api::types::{
    SKETCH_INFER_HORIZONTAL, SKETCH_INFER_TANGENT, SKETCH_RELAX_FIX, SKETCH_RELAX_PARALLEL,
    SketchInferenceTypes, SketchRelaxTypes, SketchSettings,
};

#[test]
fn explicit_choices_round_trip_without_changing_legacy_defaults() {
    let legacy: SketchSettings = serde_json::from_str("{}").unwrap();
    assert!(!legacy.custom_inference_types && !legacy.custom_relax_types);
    let selected = SketchSettings {
        custom_inference_types: true,
        inference_types: SketchInferenceTypes(SKETCH_INFER_HORIZONTAL.0 | SKETCH_INFER_TANGENT.0),
        custom_relax_types: true,
        relax_types: SketchRelaxTypes(SKETCH_RELAX_PARALLEL.0 | SKETCH_RELAX_FIX.0),
        ..legacy
    };
    let json = serde_json::to_string(&selected).unwrap();
    assert_eq!(
        serde_json::from_str::<SketchSettings>(&json).unwrap(),
        selected
    );
}

#[test]
fn explicit_empty_selection_is_distinct_from_legacy_policy() {
    let selected: SketchSettings =
        serde_json::from_str(r#"{"customInferenceTypes":true,"customRelaxTypes":true}"#).unwrap();
    assert!(selected.custom_inference_types && selected.custom_relax_types);
    assert_eq!(selected.inference_types.0, 0);
    assert_eq!(selected.relax_types.0, 0);
}
