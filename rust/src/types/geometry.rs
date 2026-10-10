// SPDX-License-Identifier: Apache-2.0

//! Array encodings are part of the protocol; named coordinate fields remain useful to callers.

use crate::json::finite;
use serde::{Deserialize, Deserializer, Serialize, Serializer};

macro_rules! coordinate_record {
    ($name:ident, $($field:ident),+) => {
        #[derive(Debug, Clone, Copy, Default, PartialEq)]
        pub struct $name { $(pub $field: f64),+ }

        impl Serialize for $name {
            fn serialize<S: Serializer>(&self, serializer: S) -> Result<S::Ok, S::Error> {
                finite::serialize(&[$(self.$field),+], serializer)
            }
        }
    };
}

macro_rules! coordinate_value {
    ($name:ident, $n:literal, $($field:ident),+) => {
        coordinate_record!($name, $($field),+);

        impl<'de> Deserialize<'de> for $name {
            fn deserialize<D: Deserializer<'de>>(deserializer: D) -> Result<Self, D::Error> {
                let [$($field),+] = deserialize_cells::<D, $n>(deserializer)?;
                Ok(Self { $($field),+ })
            }
        }
    };
}

coordinate_value!(Point, 3, x, y, z);
coordinate_value!(Point2d, 2, x, y);
coordinate_value!(Vector, 3, x, y, z);
coordinate_value!(Vector2d, 2, x, y);

macro_rules! unit_value {
    ($name:ident, $n:literal, $($field:ident),+) => {
        coordinate_record!($name, $($field),+);

        impl $name {
            /// Normalize a direction. Example: `UnitVector::new(1.0, 0.0, 0.0)?`.
            pub fn new($($field: f64),+) -> Result<Self, crate::Error> {
                let length = (0.0_f64 $(+ $field * $field)+).sqrt();
                // Match Go's public near-zero normalization boundary exactly.
                if length < 1e-12 || !length.is_finite() {
                    return Err(crate::Error(format!("direction {:?}: expected finite length >= 1e-12", [$($field),+])));
                }
                Ok(Self { $($field: $field / length),+ })
            }
        }

        impl<'de> Deserialize<'de> for $name {
            fn deserialize<D: Deserializer<'de>>(deserializer: D) -> Result<Self, D::Error> {
                let [$($field),+] = deserialize_cells::<D, $n>(deserializer)?;
                Self::new($($field),+).map_err(serde::de::Error::custom)
            }
        }
    };
}

unit_value!(UnitVector, 3, x, y, z);
unit_value!(UnitVector2d, 2, x, y);

fn deserialize_cells<'de, D, const N: usize>(deserializer: D) -> Result<[f64; N], D::Error>
where
    D: Deserializer<'de>,
    [Option<f64>; N]: Deserialize<'de>,
{
    // Go's []float64 decoder treats a null cell as zero but still checks array length.
    <[Option<f64>; N]>::deserialize(deserializer).map(|cells| cells.map(Option::unwrap_or_default))
}

/// A row-major affine 3D transform; JSON is the flat 16-cell array.
#[derive(Debug, Clone, Copy, Default, PartialEq, Serialize, Deserialize)]
#[serde(transparent)]
pub struct Matrix {
    #[serde(
        serialize_with = "finite::serialize",
        deserialize_with = "deserialize_cells"
    )]
    pub cells: [f64; 16],
}

/// A row-major affine 2D transform; JSON is the flat nine-cell array.
#[derive(Debug, Clone, Copy, Default, PartialEq, Serialize, Deserialize)]
#[serde(transparent)]
pub struct Matrix2d {
    #[serde(
        serialize_with = "finite::serialize",
        deserialize_with = "deserialize_cells"
    )]
    pub cells: [f64; 9],
}
