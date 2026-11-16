// SPDX-License-Identifier: Apache-2.0
//! Validation of the optional startup sketch plane wire value.

impl crate::types::StartupSketchPlane {
    /// Reports whether the host can accept this startup choice.
    ///
    /// ```
    /// use oblikovati_api::types::STARTUP_SKETCH_XY;
    /// assert!(STARTUP_SKETCH_XY.is_valid());
    /// ```
    pub fn is_valid(&self) -> bool {
        [
            &crate::types::STARTUP_SKETCH_NONE,
            &crate::types::STARTUP_SKETCH_XY,
            &crate::types::STARTUP_SKETCH_YZ,
            &crate::types::STARTUP_SKETCH_XZ,
        ]
        .contains(&self)
    }
}
