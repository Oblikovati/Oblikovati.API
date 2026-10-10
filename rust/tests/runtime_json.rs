// SPDX-License-Identifier: Apache-2.0

use oblikovati_api::{
    json::{ByteSlice, Bytes, JsonValue, Timestamp, finite::Finite},
    types::{Point, UnitVector, UnitVector2d, Variant, VariantValue},
};
use serde_json::{from_str, json, to_string};

#[test]
fn bytes_preserve_null_empty_and_base64() {
    for (text, expected) in [
        ("null", None),
        ("\"\"", Some(vec![])),
        ("\"AAH/\"", Some(vec![0, 1, 255])),
    ] {
        let bytes: Bytes = from_str(text).unwrap();
        assert_eq!(bytes.0, expected);
        assert_eq!(to_string(&bytes).unwrap(), text);
    }
    assert_eq!(from_str::<Bytes>("[0,255]").unwrap().0, Some(vec![0, 255]));
    assert_eq!(
        from_str::<Bytes>("\"AA\\r\\nH/\"").unwrap().0,
        Some(vec![0, 1, 255])
    );
    assert!(from_str::<Bytes>("\"bad base64\"").is_err());
    assert!(from_str::<Bytes>("[256]").is_err());
    assert!(Bytes(None).is_empty() && Bytes(Some(vec![])).is_empty());
    assert!(!Bytes(Some(vec![1])).is_empty());
}

#[test]
fn byte_enums_use_the_byte_slice_encoding() {
    let encoded = ByteSlice(Some(vec![2u8, 255]));
    assert_eq!(to_string(&encoded).unwrap(), "\"Av8=\"");
    assert_eq!(from_str::<ByteSlice<u8>>("\"Av8=\"").unwrap(), encoded);
    assert!(ByteSlice::<u8>(None).is_empty());
    assert!(ByteSlice::<u8>(Some(vec![])).is_empty());
    assert!(!encoded.is_empty());
}

#[test]
fn raw_json_distinguishes_absent_null_and_large_identifiers() {
    assert!(JsonValue::default().is_empty());
    let null = JsonValue::parse("null").unwrap();
    assert!(!null.is_empty());
    assert_eq!(to_string(&null).unwrap(), "null");
    let text = "{\"id\":18446744073709551615}";
    assert_eq!(to_string(&JsonValue::parse(text).unwrap()).unwrap(), text);
    assert!(JsonValue::parse("broken").unwrap_err().0.contains("broken"));
}

#[test]
fn timestamps_keep_nanoseconds_and_validate_the_shape() {
    for text in [
        "2026-10-10T12:30:00.123456789Z",
        "0001-01-01T00:00:00Z",
        "2026-10-10T12:30:00+02:00",
    ] {
        assert_eq!(
            to_string(&Timestamp::parse(text).unwrap()).unwrap(),
            to_string(text).unwrap()
        );
    }
    assert_eq!(from_str::<Timestamp>("null").unwrap(), Timestamp::default());
    assert!(
        Timestamp::parse("yesterday")
            .unwrap_err()
            .0
            .contains("yesterday")
    );
    assert!(from_str::<Timestamp>("\"2026-99-10\"").is_err());
}

#[test]
fn geometry_uses_ordered_arrays_and_unit_directions() {
    let point = Point {
        x: 1.0,
        y: -2.0,
        z: 3.5,
    };
    assert_eq!(to_string(&point).unwrap(), "[1.0,-2.0,3.5]");
    assert_eq!(from_str::<Point>("[1,-2,3.5]").unwrap(), point);
    assert!(from_str::<Point>("[1,2]").is_err());
    assert_eq!(UnitVector::new(0.0, 3.0, 4.0).unwrap().z, 0.8);
    assert_eq!(UnitVector2d::new(3.0, 4.0).unwrap().x, 0.6);
    assert!(UnitVector::new(0.0, 0.0, 0.0).is_err());
    assert!(UnitVector2d::new(f64::INFINITY, 0.0).is_err());
    assert!(UnitVector2d::new(1e-13, 0.0).is_err());
    assert!(UnitVector2d::new(1e-12, 0.0).is_ok());
    assert!(
        to_string(&Point {
            x: f64::NAN,
            ..point
        })
        .is_err()
    );
}

#[test]
fn nonfinite_numbers_are_rejected_through_nested_containers() {
    assert!(
        Some(Box::new(vec![1.0_f32, f32::INFINITY]))
            .check_finite()
            .is_err()
    );
    assert!(Some(Box::new(vec![1.0_f64, 2.0])).check_finite().is_ok());
    assert!(None::<f64>.check_finite().is_ok());
    assert!([0.0, f64::NEG_INFINITY].check_finite().is_err());
}

#[test]
fn variants_preserve_every_tag_and_units() {
    for value in [
        VariantValue::Integer(i64::MAX),
        VariantValue::Double(2.5),
        VariantValue::String("日本語".into()),
        VariantValue::Boolean(true),
        VariantValue::Bytes(Bytes(Some(vec![1, 2]))),
    ] {
        let variant = Variant {
            value,
            unit: "mm".into(),
        };
        let text = to_string(&variant).unwrap();
        assert_eq!(from_str::<Variant>(&text).unwrap(), variant);
    }
    let null: Variant = from_str(r#"{"type":"integer","value":null}"#).unwrap();
    assert_eq!(null.value, VariantValue::Integer(0));
    assert!(from_str::<Variant>(r#"{"type":"integer","value":"wrong"}"#).is_err());
    assert!(from_str::<Variant>(r#"{"type":"unknown","value":1}"#).is_err());
    assert_eq!(
        serde_json::to_value(Variant::default()).unwrap(),
        json!({"type":"valueType(?)", "value":false})
    );
    assert!(
        to_string(&Variant {
            value: VariantValue::Double(f64::NAN),
            unit: String::new()
        })
        .is_err()
    );
}
