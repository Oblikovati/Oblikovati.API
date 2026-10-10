// SPDX-License-Identifier: Apache-2.0

use base64::{
    Engine, alphabet,
    engine::{GeneralPurpose, GeneralPurposeConfig},
};
use serde::{Deserialize, Deserializer, Serialize, Serializer};

// Go's StdEncoding accepts nonzero unused padding bits; match that decoder contract.
const GO_BASE64: GeneralPurpose = GeneralPurpose::new(
    &alphabet::STANDARD,
    GeneralPurposeConfig::new().with_decode_allow_trailing_bits(true),
);

/// A Go byte slice: null when absent, otherwise a base64 JSON string.
#[derive(Debug, Clone, Default, PartialEq, Eq)]
pub struct Bytes(pub Option<Vec<u8>>);

impl Bytes {
    /// Test the Go `omitempty` rule. Example: `assert!(Bytes::default().is_empty())`.
    pub fn is_empty(&self) -> bool {
        self.0.as_ref().is_none_or(Vec::is_empty)
    }
}

impl Serialize for Bytes {
    fn serialize<S: Serializer>(&self, serializer: S) -> Result<S::Ok, S::Error> {
        match &self.0 {
            None => serializer.serialize_none(),
            Some(bytes) => serializer.serialize_str(&GO_BASE64.encode(bytes)),
        }
    }
}

impl<'de> Deserialize<'de> for Bytes {
    fn deserialize<D: Deserializer<'de>>(deserializer: D) -> Result<Self, D::Error> {
        #[derive(Deserialize)]
        #[serde(untagged)]
        enum Encoded {
            Text(String),
            Array(Vec<u8>),
        }
        match Option::<Encoded>::deserialize(deserializer)? {
            None => Ok(Self(None)),
            Some(Encoded::Array(bytes)) => Ok(Self(Some(bytes))),
            Some(Encoded::Text(text)) => GO_BASE64
                .decode(text.replace(['\r', '\n'], ""))
                .map(|bytes| Self(Some(bytes)))
                .map_err(serde::de::Error::custom),
        }
    }
}

/// A slice of byte-valued enum members, also base64 in Go's JSON encoding.
#[derive(Debug, Clone, Default, PartialEq, Eq)]
pub struct ByteSlice<T>(pub Option<Vec<T>>);

impl<T> ByteSlice<T> {
    /// Test the Go `omitempty` rule. Example: `ByteSlice::<u8>(None).is_empty()`.
    pub fn is_empty(&self) -> bool {
        self.0.as_ref().is_none_or(Vec::is_empty)
    }
}

impl<T: Copy + Into<u8>> Serialize for ByteSlice<T> {
    fn serialize<S: Serializer>(&self, serializer: S) -> Result<S::Ok, S::Error> {
        let bytes = self
            .0
            .as_ref()
            .map(|values| values.iter().copied().map(Into::into).collect());
        Bytes(bytes).serialize(serializer)
    }
}

impl<'de, T: From<u8>> Deserialize<'de> for ByteSlice<T> {
    fn deserialize<D: Deserializer<'de>>(deserializer: D) -> Result<Self, D::Error> {
        let bytes = Bytes::deserialize(deserializer)?;
        Ok(Self(
            bytes
                .0
                .map(|values| values.into_iter().map(T::from).collect()),
        ))
    }
}
