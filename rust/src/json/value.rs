// SPDX-License-Identifier: Apache-2.0

use serde::{Deserialize, Deserializer, Serialize, Serializer};

/// Opaque JSON at API extension points, retaining absent versus explicit null.
#[derive(Debug, Clone, Default, PartialEq)]
pub struct JsonValue(pub Option<serde_json::Value>);

impl JsonValue {
    /// Parse a JSON extension value. Example: `JsonValue::parse("{\"size\":2}")?`.
    pub fn parse(text: &str) -> Result<Self, crate::Error> {
        serde_json::from_str(text).map_err(|error| {
            crate::Error(format!(
                "JSON {text:?}: expected a valid JSON value: {error}"
            ))
        })
    }

    /// Test whether no JSON was supplied. Example: `assert!(JsonValue::default().is_empty())`.
    pub fn is_empty(&self) -> bool {
        self.0.is_none()
    }
}

impl Serialize for JsonValue {
    fn serialize<S: Serializer>(&self, serializer: S) -> Result<S::Ok, S::Error> {
        self.0.serialize(serializer)
    }
}

impl<'de> Deserialize<'de> for JsonValue {
    fn deserialize<D: Deserializer<'de>>(deserializer: D) -> Result<Self, D::Error> {
        serde_json::Value::deserialize(deserializer).map(|value| Self(Some(value)))
    }
}

/// An RFC3339 timestamp, as emitted by Go's `time.Time` JSON representation.
#[derive(Debug, Clone, PartialEq, Eq, Serialize)]
#[serde(transparent)]
pub struct Timestamp(String);

impl Timestamp {
    /// Validate and normalize a timestamp. Example: `Timestamp::parse("2026-10-10T12:30:00Z")?`.
    pub fn parse(text: &str) -> Result<Self, crate::Error> {
        use time::{OffsetDateTime, format_description::well_known::Rfc3339};
        // time accepts leap seconds by rounding them; Go's time.Time rejects them.
        if text.get(17..19) == Some("60") {
            return Err(crate::Error(format!(
                "timestamp {text:?}: expected RFC3339 seconds 00..59"
            )));
        }
        let value = OffsetDateTime::parse(text, &Rfc3339).map_err(|error| {
            crate::Error(format!("timestamp {text:?}: expected RFC3339: {error}"))
        })?;
        value
            .format(&Rfc3339)
            .map(Self)
            .map_err(|error| crate::Error(format!("timestamp {text:?}: expected RFC3339: {error}")))
    }
}

impl<'de> Deserialize<'de> for Timestamp {
    fn deserialize<D: Deserializer<'de>>(deserializer: D) -> Result<Self, D::Error> {
        let value = Option::<String>::deserialize(deserializer)?;
        value.map_or_else(
            || Ok(Self::default()),
            |text| Self::parse(&text).map_err(serde::de::Error::custom),
        )
    }
}

impl Default for Timestamp {
    fn default() -> Self {
        Self("0001-01-01T00:00:00Z".into())
    }
}
