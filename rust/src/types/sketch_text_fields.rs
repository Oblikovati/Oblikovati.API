// SPDX-License-Identifier: Apache-2.0

use crate::{
    Error,
    types::{SKETCH_TEXT_FIELD_PLACEHOLDER, SketchTextField, SketchTextFraction},
};

impl SketchTextField {
    /// Validate fields against separately validated fractions and optional source content.
    /// For example, validate fields in Some("⌀\u{fffc}") before a linked text add.
    pub fn validate(
        content: Option<&str>,
        fields: &[Self],
        fractions: &[SketchTextFraction],
    ) -> Result<(), Error> {
        if fields.is_empty() {
            return Ok(());
        }
        let scalars = content.map(|text| text.chars().collect::<Vec<_>>());
        for field in fields.windows(2) {
            if field[1].at <= field[0].at {
                return Err(Error(format!(
                    "sketch text field at {}: expected a scalar offset greater than {}",
                    field[1].at, field[0].at
                )));
            }
        }
        for field in fields {
            field.validate_source()?;
            field.validate_location(scalars.as_deref(), fractions)?;
        }
        Ok(())
    }

    fn validate_source(&self) -> Result<(), Error> {
        let expected = match self.kind.0.as_ref() {
            "modelParameter"
                if self.parameter.is_empty() || self.parameter.trim() != self.parameter =>
            {
                "a nonempty parameter name without surrounding spaces"
            }
            "modelParameter" if self.precision.as_deref().is_none_or(|&digits| digits <= 8) => {
                return Ok(());
            }
            "modelParameter" => "zero through eight decimal places",
            "fileName" if self.parameter.is_empty() && self.precision.is_none() => return Ok(()),
            "fileName" => "no parameter name or precision",
            _ => "kind modelParameter or fileName",
        };
        Err(Error(format!(
            "sketch text field {self:?}: expected {expected}"
        )))
    }

    fn validate_location(
        &self,
        scalars: Option<&[char]>,
        fractions: &[SketchTextFraction],
    ) -> Result<(), Error> {
        let placeholder = SKETCH_TEXT_FIELD_PLACEHOLDER.chars().next();
        if scalars.is_some_and(|text| text.get(self.at as usize).copied() != placeholder) {
            return Err(Error(format!(
                "sketch text field at {}: expected an object-replacement scalar U+FFFC",
                self.at
            )));
        }
        let index = fractions.partition_point(|range| range.end <= self.at);
        if let Some(range) = fractions.get(index).filter(|range| range.start <= self.at) {
            return Err(Error(format!(
                "sketch text field at {} in stacked range {range:?}: expected a location outside all stacked fractions",
                self.at
            )));
        }
        Ok(())
    }
}
