// SPDX-License-Identifier: Apache-2.0
//! Shared definitions for generated transparent scalar types.

// Go named scalar types must stay distinct Rust tuple structs. Centralize the
// shared representation so every generated type has identical wire semantics.
macro_rules! generated_scalar {
    (integer, $(#[$meta:meta])* pub struct $name:ident(pub $inner:ty);) => {
        $(#[$meta])*
        #[derive(Debug, Clone, Default, PartialEq, PartialOrd, serde::Serialize, serde::Deserialize, Eq, Ord, Hash, Copy)]
        #[repr(transparent)]
        #[serde(transparent)]
        pub struct $name(pub $inner);
    };
    (float, $(#[$meta:meta])* pub struct $name:ident(pub $inner:ty);) => {
        $(#[$meta])*
        #[derive(Debug, Clone, Default, PartialEq, PartialOrd, serde::Serialize, serde::Deserialize, Copy)]
        #[repr(transparent)]
        #[serde(transparent)]
        pub struct $name(pub $inner);
    };
    (string, $(#[$meta:meta])* pub struct $name:ident(pub $inner:ty);) => {
        $(#[$meta])*
        #[derive(Debug, Clone, Default, PartialEq, PartialOrd, serde::Serialize, serde::Deserialize, Eq, Ord, Hash)]
        #[repr(transparent)]
        #[serde(transparent)]
        pub struct $name(pub $inner);
    };
}

#[cfg(test)]
mod tests {
    generated_scalar! { float, pub struct TestFloat(pub f64); }

    #[test]
    fn floating_scalar_retains_copy_layout_and_json() {
        fn require_copy<T: Copy + PartialOrd + Default>() {}
        require_copy::<TestFloat>();
        assert_eq!(std::mem::size_of::<TestFloat>(), std::mem::size_of::<f64>());
        assert_eq!(serde_json::to_string(&TestFloat(1.25)).unwrap(), "1.25");
        assert_eq!(serde_json::from_str::<TestFloat>("1.25").unwrap().0, 1.25);
    }
}
