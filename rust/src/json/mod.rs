// SPDX-License-Identifier: Apache-2.0

//! JSON rules shared by all DTOs. These preserve Go's null and empty semantics.

mod bytes;
pub mod finite;
mod value;
pub use bytes::{ByteSlice, Bytes};
pub use value::{JsonValue, Timestamp};

/// The JSON object used by methods that explicitly send an empty request.
#[derive(Debug, Clone, Default, PartialEq, serde::Serialize, serde::Deserialize)]
pub struct EmptyObject {}

pub(crate) fn is_zero<T: Default + PartialEq>(value: &T) -> bool {
    value == &T::default()
}

pub(crate) fn slice_is_empty<T>(value: &Option<Vec<T>>) -> bool {
    value.as_ref().is_none_or(Vec::is_empty)
}

pub(crate) fn map_is_empty<K, V>(value: &Option<std::collections::BTreeMap<K, V>>) -> bool {
    value
        .as_ref()
        .is_none_or(std::collections::BTreeMap::is_empty)
}

pub(crate) fn null_default<
    'de,
    T: serde::Deserialize<'de> + Default,
    D: serde::Deserializer<'de>,
>(
    deserializer: D,
) -> Result<T, D::Error> {
    Ok(Option::<T>::deserialize(deserializer)?.unwrap_or_default())
}

use serde::Deserialize;
