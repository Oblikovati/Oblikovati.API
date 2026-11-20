// SPDX-License-Identifier: Apache-2.0
use oblikovati_api::types::{ASideFaceStatus, Accuracy, AddInKind, AssetSource};
use std::any::TypeId;
use std::borrow::Cow;
use std::hash::Hash;
use std::mem::{align_of, size_of};

fn require_integer_traits<T: Copy + Eq + Ord + Hash + Default>() {}

#[test]
fn generated_integers_keep_distinct_types_and_transparent_wire_values() {
    require_integer_traits::<ASideFaceStatus>();
    require_integer_traits::<Accuracy>();
    require_integer_traits::<AddInKind>();
    assert_ne!(TypeId::of::<ASideFaceStatus>(), TypeId::of::<Accuracy>());
    assert_eq!(size_of::<ASideFaceStatus>(), size_of::<i32>());
    assert_eq!(align_of::<ASideFaceStatus>(), align_of::<i32>());
    assert_eq!(serde_json::to_string(&ASideFaceStatus(17)).unwrap(), "17");
    assert_eq!(serde_json::from_str::<ASideFaceStatus>("17").unwrap().0, 17);
    assert_eq!(u8::from(AddInKind::from(7)), 7);
}

#[test]
fn generated_string_keeps_transparent_cow_representation() {
    let value = AssetSource(Cow::Borrowed("library"));
    assert_eq!(size_of::<AssetSource>(), size_of::<Cow<'static, str>>());
    assert_eq!(serde_json::to_string(&value).unwrap(), "\"library\"");
    assert_eq!(
        serde_json::from_str::<AssetSource>("\"library\"").unwrap(),
        value
    );
}
