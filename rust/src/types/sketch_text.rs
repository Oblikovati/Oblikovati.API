// SPDX-License-Identifier: Apache-2.0

use super::{OVERRIDE_COLOR_SOURCE, SketchTextFraction, SketchTextLineSpacing, SketchTextRun};
use crate::Error;

#[path = "sketch_text_fields.rs"]
mod sketch_text_fields;

impl SketchTextLineSpacing {
    /// Validate optional spacing, for example `SketchTextLineSpacing::validate(Some(&spacing))?`.
    pub fn validate(spacing: Option<&Self>) -> Result<(), Error> {
        let Some(spacing) = spacing else {
            return Ok(());
        };
        let expected = match spacing.mode.0.as_ref() {
            "single" | "oneAndHalf" | "double" if spacing.value == 0. => return Ok(()),
            "single" | "oneAndHalf" | "double" => "zero for a fixed mode",
            "multiple" | "exactly" if spacing.value > 0. && spacing.value.is_finite() => {
                return Ok(());
            }
            "multiple" | "exactly" => "a positive finite value",
            mode => {
                return Err(Error(format!(
                    "line spacing mode {mode:?}: expected single, oneAndHalf, double, multiple, or exactly"
                )));
            }
        };
        Err(Error(format!(
            "line spacing {:?} value {:?}: expected {expected}",
            spacing.mode.0, spacing.value
        )))
    }
}

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

impl SketchTextFraction {
    /// Check ordered, non-overlapping Unicode ranges and their selected content.
    /// For example, validate a selection in `Some("⌀1/2")` before saving.
    /// `None` defers the content check to the host for a format-only edit.
    pub fn validate(content: Option<&str>, fractions: &[Self]) -> Result<(), Error> {
        if fractions.is_empty() {
            return Ok(());
        }
        let scalars = content.map(|text| text.chars().collect::<Vec<_>>());
        let mut previous_end = 0;
        for (index, fraction) in fractions.iter().enumerate() {
            fraction.validate_range(index, previous_end)?;
            if let Some(text) = scalars.as_deref() {
                fraction.validate_content(index, text)?;
            }
            previous_end = fraction.end;
        }
        Ok(())
    }

    fn validate_range(&self, index: usize, previous_end: u32) -> Result<(), Error> {
        if self.start < previous_end || self.end.saturating_sub(self.start) < 3 {
            return Err(Error(format!(
                "stacked fraction {index} [{},{}): expected start >= {previous_end} and at least three Unicode scalars",
                self.start, self.end
            )));
        }
        Ok(())
    }

    fn validate_content(&self, index: usize, text: &[char]) -> Result<(), Error> {
        let selected = text
            .get(self.start as usize..self.end as usize)
            .ok_or_else(|| {
                Error(format!(
                    "stacked fraction {index} end {}: expected at most {} Unicode scalars",
                    self.end,
                    text.len()
                ))
            })?;
        let value: String = selected.iter().collect();
        if is_single_line_fraction(&value) {
            return Ok(());
        }
        Err(Error(format!(
            "stacked fraction {index} content {value:?}: expected one slash between two non-empty values on one line"
        )))
    }
}

fn is_single_line_fraction(value: &str) -> bool {
    value.split_once('/').is_some_and(|(top, bottom)| {
        !top.trim().is_empty()
            && !bottom.trim().is_empty()
            && !bottom.contains('/')
            && !value.contains(['\r', '\n'])
    })
}
