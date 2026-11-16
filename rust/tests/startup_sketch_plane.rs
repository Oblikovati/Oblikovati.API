// SPDX-License-Identifier: Apache-2.0

use oblikovati_api::{types, wire::PartOptionsView};

#[test]
fn startup_sketch_choices_round_trip_and_old_payload_is_unchanged() {
    for plane in [
        types::STARTUP_SKETCH_NONE,
        types::STARTUP_SKETCH_XY,
        types::STARTUP_SKETCH_YZ,
        types::STARTUP_SKETCH_XZ,
    ] {
        assert!(plane.is_valid());
        let input = PartOptionsView {
            startup_sketch_plane: Some(Box::new(plane.clone())),
            ..Default::default()
        };
        let encoded = serde_json::to_value(&input).unwrap();
        let decoded: PartOptionsView = serde_json::from_value(encoded).unwrap();
        assert_eq!(decoded.startup_sketch_plane.as_deref(), Some(&plane));
    }
    let old: PartOptionsView = serde_json::from_str(r#"{"chamferFlatCorners":true}"#).unwrap();
    assert!(old.startup_sketch_plane.is_none());
    assert!(
        serde_json::to_value(old)
            .unwrap()
            .get("startupSketchPlane")
            .is_none()
    );
}

#[test]
fn invalid_plane_values_are_not_valid_startup_choices() {
    for value in ["", "XY", "bogus"] {
        let plane = types::StartupSketchPlane(value.into());
        assert!(!plane.is_valid(), "unexpected valid plane: {value:?}");
    }
}

#[test]
fn null_and_unknown_planes_remain_distinct() {
    let null: PartOptionsView = serde_json::from_str(r#"{"startupSketchPlane":null}"#).unwrap();
    assert!(null.startup_sketch_plane.is_none());
    let unknown: PartOptionsView =
        serde_json::from_str(r#"{"startupSketchPlane":"bogus"}"#).unwrap();
    assert!(!unknown.startup_sketch_plane.unwrap().is_valid());
}

#[test]
fn part_sketch_preferences_remain_independent() {
    for enabled in [false, true] {
        let input = PartOptionsView {
            chamfer_flat_corners: true,
            startup_sketch_plane: Some(Box::new(types::STARTUP_SKETCH_YZ.clone())),
            look_at_sketch_plane: Some(Box::new(enabled)),
        };
        let encoded = serde_json::to_vec(&input).unwrap();
        let output: PartOptionsView = serde_json::from_slice(&encoded).unwrap();
        assert_eq!(
            output.startup_sketch_plane.as_deref(),
            Some(&types::STARTUP_SKETCH_YZ)
        );
        assert_eq!(output.look_at_sketch_plane.as_deref(), Some(&enabled));
    }
}
