// SPDX-License-Identifier: Apache-2.0

use serde::{Serialize, Serializer};

/// Go rejects non-finite numbers instead of silently encoding them as null.
pub trait Finite {
    fn check_finite(&self) -> Result<(), crate::Error>;
}

impl Finite for f64 {
    fn check_finite(&self) -> Result<(), crate::Error> {
        if self.is_finite() {
            return Ok(());
        }
        Err(crate::Error(format!(
            "JSON number {self}: expected a finite number"
        )))
    }
}

impl Finite for f32 {
    fn check_finite(&self) -> Result<(), crate::Error> {
        f64::from(*self).check_finite()
    }
}

impl<T: Finite, const N: usize> Finite for [T; N] {
    fn check_finite(&self) -> Result<(), crate::Error> {
        self.iter().try_for_each(Finite::check_finite)
    }
}

impl<T: Finite> Finite for Vec<T> {
    fn check_finite(&self) -> Result<(), crate::Error> {
        self.iter().try_for_each(Finite::check_finite)
    }
}

impl<T: Finite> Finite for Option<T> {
    fn check_finite(&self) -> Result<(), crate::Error> {
        self.as_ref().map_or(Ok(()), Finite::check_finite)
    }
}

impl<T: Finite> Finite for Box<T> {
    fn check_finite(&self) -> Result<(), crate::Error> {
        (**self).check_finite()
    }
}

pub(crate) fn serialize<T: Finite + Serialize, S: Serializer>(
    value: &T,
    serializer: S,
) -> Result<S::Ok, S::Error> {
    value.check_finite().map_err(serde::ser::Error::custom)?;
    value.serialize(serializer)
}
