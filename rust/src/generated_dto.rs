// SPDX-License-Identifier: Apache-2.0
//! Shared field rules for generated Go-compatible wire values.

// Each invocation expands to its own public struct. These rules centralize the
// null and omission semantics without sharing the state or identity of DTOs.
macro_rules! generated_dto {
    ($(#[$doc:meta])* pub struct $name:ident { $($fields:tt)* }) => {
        generated_dto!(@fields [$(#[$doc])*] [$name] [] $($fields)*);
    };
    (@fields [$($doc:tt)*] [$name:ident] [$($result:tt)*]) => {
        $($doc)*
        #[derive(Debug, Clone, PartialEq, Default, serde::Serialize, serde::Deserialize)]
        #[serde(default)]
        pub struct $name { $($result)* }
    };
    (@fields [$($doc:tt)*] [$name:ident] [$($result:tt)*]
     $(#[$field_doc:meta])* pub $field:ident : $ty:ty => required($rename:literal); $($rest:tt)*) => {
        generated_dto!(@fields [$($doc)*] [$name] [$($result)*
            $(#[$field_doc])*
            #[serde(rename = $rename, deserialize_with = "crate::json::null_default")]
            pub $field: $ty,
        ] $($rest)*);
    };
    (@fields [$($doc:tt)*] [$name:ident] [$($result:tt)*]
     $(#[$field_doc:meta])* pub $field:ident : $ty:ty => omitted($rename:literal, $skip:literal); $($rest:tt)*) => {
        generated_dto!(@fields [$($doc)*] [$name] [$($result)*
            $(#[$field_doc])*
            #[serde(rename = $rename, skip_serializing_if = $skip, deserialize_with = "crate::json::null_default")]
            pub $field: $ty,
        ] $($rest)*);
    };
    (@fields [$($doc:tt)*] [$name:ident] [$($result:tt)*]
     $(#[$field_doc:meta])* pub $field:ident : $ty:ty => finite($rename:literal); $($rest:tt)*) => {
        generated_dto!(@fields [$($doc)*] [$name] [$($result)*
            $(#[$field_doc])*
            #[serde(rename = $rename, serialize_with = "crate::json::finite::serialize", deserialize_with = "crate::json::null_default")]
            pub $field: $ty,
        ] $($rest)*);
    };
    (@fields [$($doc:tt)*] [$name:ident] [$($result:tt)*]
     $(#[$field_doc:meta])* pub $field:ident : $ty:ty => omitted_finite($rename:literal, $skip:literal); $($rest:tt)*) => {
        generated_dto!(@fields [$($doc)*] [$name] [$($result)*
            $(#[$field_doc])*
            #[serde(rename = $rename, skip_serializing_if = $skip, serialize_with = "crate::json::finite::serialize", deserialize_with = "crate::json::null_default")]
            pub $field: $ty,
        ] $($rest)*);
    };
    (@fields [$($doc:tt)*] [$name:ident] [$($result:tt)*]
     $(#[$field_doc:meta])* pub $field:ident : $ty:ty => raw($($attr:tt)*); $($rest:tt)*) => {
        generated_dto!(@fields [$($doc)*] [$name] [$($result)*
            $(#[$field_doc])*
            #[serde($($attr)*)]
            pub $field: $ty,
        ] $($rest)*);
    };
}

#[cfg(test)]
mod tests {
    generated_dto! {
        pub struct Example {
            pub id: i64 => required("id");
            pub label: String => omitted("label", "crate::json::is_zero");
            pub size: f64 => finite("size");
            pub gain: f64 => omitted_finite("gain", "crate::json::is_zero");
            pub custom: Option<String> => raw(rename = "custom");
        }
    }

    #[test]
    fn field_rules_keep_null_defaults_omission_and_finite_checks() {
        let value: Example = serde_json::from_str(
            r#"{"id":null,"label":null,"size":null,"gain":null,"custom":null}"#,
        )
        .unwrap();
        let json = serde_json::to_value(&value).unwrap();
        assert_eq!(json["id"], 0);
        assert_eq!(json["size"], 0.0);
        assert_eq!(json["custom"], serde_json::Value::Null);
        assert!(json.get("label").is_none());
        assert!(json.get("gain").is_none());
        assert!(
            serde_json::to_value(Example {
                size: f64::INFINITY,
                ..value
            })
            .is_err()
        );
    }
}
