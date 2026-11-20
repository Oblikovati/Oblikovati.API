// SPDX-License-Identifier: Apache-2.0

use super::{OVERRIDE_COLOR_SOURCE, SketchTextRun};
use crate::Error;

impl SketchTextRun {
    /// Validate style runs against Unicode scalar positions. `None` defers the length check.
    /// For example, call `SketchTextRun::validate(Some("a界🙂"), &runs)` before saving.
    pub fn validate(content: Option<&str>, runs: &[Self]) -> Result<(), Error> {
        if runs.is_empty() {
            return Ok(());
        }
        let limit = content.map(|text| text.chars().count() as u64);
        let mut previous_end = 0;
        for (index, run) in runs.iter().enumerate() {
            run.validate_range(index, previous_end, limit)?;
            run.validate_style(index)?;
            previous_end = run.end;
        }
        Ok(())
    }

    fn validate_range(
        &self,
        index: usize,
        previous_end: u32,
        limit: Option<u64>,
    ) -> Result<(), Error> {
        if self.start < previous_end || self.start >= self.end {
            return Err(Error(format!(
                "sketch text run {index} [{},{}): expected {previous_end} <= start < end",
                self.start, self.end
            )));
        }
        if let Some(limit) = limit.filter(|&limit| u64::from(self.end) > limit) {
            return Err(Error(format!(
                "sketch text run {index} end {}: expected at most {limit} Unicode scalars",
                self.end
            )));
        }
        Ok(())
    }

    fn validate_style(&self, index: usize) -> Result<(), Error> {
        if let Some(width) = self
            .width_percent
            .as_deref()
            .copied()
            .filter(|&width| width <= 0. || !width.is_finite())
        {
            return Err(Error(format!(
                "sketch text run {index} width {width}: expected a positive finite percentage"
            )));
        }
        if let Some(color) = self.color.as_ref().filter(|color| {
            color.source != OVERRIDE_COLOR_SOURCE || !(0. ..=1.).contains(&color.opacity)
        }) {
            return Err(Error(format!(
                "sketch text run {index} color {color:?}: expected override color with opacity in [0,1]"
            )));
        }
        Ok(())
    }
}
