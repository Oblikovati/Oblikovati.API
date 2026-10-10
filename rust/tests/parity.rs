// SPDX-License-Identifier: Apache-2.0

use serde::{Deserialize, Serialize, de::DeserializeOwned};
use serde_json::Value;
use std::{collections::BTreeMap, sync::OnceLock};

mod fixtures;
mod parity_generated;

#[test]
fn go_byte_codec_edges() {
    check::<oblikovati_api::json::Bytes>("std.Bytes");
}

#[test]
fn go_timestamp_codec_edges() {
    check::<oblikovati_api::json::Timestamp>("std.Time");
}

#[derive(Deserialize)]
#[serde(rename_all = "PascalCase")]
struct GoFixture {
    name: String,
    case: String,
    payload: Value,
    canonical: Value,
    error: String,
}

fn fixtures() -> &'static BTreeMap<String, Vec<GoFixture>> {
    static FIXTURES: OnceLock<BTreeMap<String, Vec<GoFixture>>> = OnceLock::new();
    FIXTURES.get_or_init(|| {
        let mut grouped = BTreeMap::<String, Vec<GoFixture>>::new();
        for line in fixtures::CONTENTS.lines() {
            let fixture: GoFixture = serde_json::from_str(line).expect("valid Go fixture");
            grouped
                .entry(fixture.name.clone())
                .or_default()
                .push(fixture);
        }
        grouped
    })
}

fn check<T: DeserializeOwned + Serialize>(name: &str) {
    for fixture in &fixtures()[name] {
        let decoded = serde_json::from_value::<T>(fixture.payload.clone());
        if !fixture.error.is_empty() {
            assert!(
                decoded.is_err(),
                "{name}/{} must reject: {}",
                fixture.case,
                fixture.error
            );
            continue;
        }
        let decoded = decoded.unwrap_or_else(|error| panic!("{name}/{}: {error}", fixture.case));
        let encoded = serde_json::to_value(decoded).unwrap();
        assert!(
            json_equal(&encoded, &fixture.canonical),
            "{name}/{}\nRust: {encoded}\nGo: {}",
            fixture.case,
            fixture.canonical
        );
    }
}

fn json_equal(left: &Value, right: &Value) -> bool {
    match (left, right) {
        (Value::Number(left), Value::Number(right)) => number_equal(left, right),
        (Value::Array(left), Value::Array(right)) => {
            left.len() == right.len() && left.iter().zip(right).all(|(l, r)| json_equal(l, r))
        }
        (Value::Object(left), Value::Object(right)) => {
            left.len() == right.len()
                && left.iter().all(|(key, value)| {
                    right.get(key).is_some_and(|other| json_equal(value, other))
                })
        }
        _ => left == right,
    }
}

fn number_equal(left: &serde_json::Number, right: &serde_json::Number) -> bool {
    if left.is_f64() && right.is_f64() {
        return left.as_f64() == right.as_f64();
    }
    if left.is_f64() {
        return whole_float_equal(left, right);
    }
    if right.is_f64() {
        return whole_float_equal(right, left);
    }
    left == right
}

fn whole_float_equal(floating: &serde_json::Number, integer: &serde_json::Number) -> bool {
    let value = floating.as_f64().unwrap();
    value.fract() == 0.0 && format!("{value:.0}") == integer.to_string()
}

#[test]
fn large_identifiers_are_compared_without_rounding() {
    let left = serde_json::json!(u64::MAX);
    let right = serde_json::json!(u64::MAX - 1);
    assert!(!json_equal(&left, &right));
    assert!(!json_equal(&left, &serde_json::json!(u64::MAX as f64)));
    assert!(json_equal(&serde_json::json!(1), &serde_json::json!(1.0)));
}
