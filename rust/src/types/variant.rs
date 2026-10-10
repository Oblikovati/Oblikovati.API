// SPDX-License-Identifier: Apache-2.0

use crate::json::{Bytes, finite::Finite};
use serde::{Deserialize, Deserializer, Serialize, Serializer};

/// Typed payload for a property bag value.
#[derive(Debug, Clone, PartialEq, Default)]
pub enum VariantValue {
    Integer(i64),
    Double(f64),
    String(String),
    Bytes(Bytes),
    Boolean(bool),
    /// Go's uninitialized Variant has no valid type tag.
    #[default]
    Uninitialized,
}

/// A tagged value and optional unit expression.
/// Example: `Variant { value: VariantValue::Double(2.5), unit: "mm".into() }`.
#[derive(Debug, Clone, PartialEq, Default)]
pub struct Variant {
    pub value: VariantValue,
    pub unit: String,
}

#[derive(Serialize, Deserialize)]
struct VariantWire {
    #[serde(rename = "type")]
    tag: String,
    value: serde_json::Value,
    #[serde(default, skip_serializing_if = "String::is_empty")]
    unit: String,
}

impl Serialize for Variant {
    fn serialize<S: Serializer>(&self, serializer: S) -> Result<S::Ok, S::Error> {
        let (tag, value) = variant_payload(&self.value).map_err(serde::ser::Error::custom)?;
        VariantWire {
            tag: tag.into(),
            value,
            unit: self.unit.clone(),
        }
        .serialize(serializer)
    }
}

impl<'de> Deserialize<'de> for Variant {
    fn deserialize<D: Deserializer<'de>>(deserializer: D) -> Result<Self, D::Error> {
        let encoded = VariantWire::deserialize(deserializer)?;
        let value =
            decode_payload(&encoded.tag, encoded.value).map_err(serde::de::Error::custom)?;
        Ok(Self {
            value,
            unit: encoded.unit,
        })
    }
}

fn variant_payload(
    value: &VariantValue,
) -> Result<(&'static str, serde_json::Value), crate::Error> {
    match value {
        VariantValue::Integer(value) => Ok(("integer", (*value).into())),
        VariantValue::Double(value) => {
            value.check_finite()?;
            Ok(("double", (*value).into()))
        }
        VariantValue::String(value) => Ok(("string", value.clone().into())),
        VariantValue::Boolean(value) => Ok(("boolean", (*value).into())),
        VariantValue::Bytes(value) => serde_json::to_value(value)
            .map(|value| ("bytes", value))
            .map_err(|error| crate::Error(error.to_string())),
        VariantValue::Uninitialized => Ok(("valueType(?)", false.into())),
    }
}

fn decode_payload(tag: &str, value: serde_json::Value) -> Result<VariantValue, crate::Error> {
    match tag {
        "integer" => decode_value(value).map(VariantValue::Integer),
        "double" => decode_value(value).map(VariantValue::Double),
        "string" => decode_value(value).map(VariantValue::String),
        "boolean" => decode_value(value).map(VariantValue::Boolean),
        "bytes" => decode_value(value).map(VariantValue::Bytes),
        _ => Err(crate::Error(format!(
            "Variant type {tag:?}: expected integer|double|string|bytes|boolean"
        ))),
    }
}

fn decode_value<T: serde::de::DeserializeOwned + Default>(
    value: serde_json::Value,
) -> Result<T, crate::Error> {
    if value.is_null() {
        return Ok(T::default());
    }
    serde_json::from_value(value.clone()).map_err(|error| {
        crate::Error(format!(
            "Variant value {value}: expected the tagged payload shape: {error}"
        ))
    })
}
