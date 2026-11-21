// SPDX-License-Identifier: Apache-2.0

use crate::{
    Error,
    types::{SketchTextField, SketchTextFraction, SketchTextLineSpacing, SketchTextRun},
    wire,
};
use serde::de::DeserializeOwned;

pub(super) fn validate_request(method: &str, bytes: &[u8]) -> Result<(), Error> {
    match method {
        wire::METHOD_SKETCH_ADD_TEXT => validate_add(decode(method, "request", bytes)?),
        wire::METHOD_SKETCH_EDIT_TEXT => validate_edit(decode(method, "request", bytes)?),
        _ => return Ok(()),
    }
    .map_err(|error| Error(format!("method {method:?} request: {error}")))
}

fn validate_add(args: wire::AddTextArgs) -> Result<(), Error> {
    Formatting {
        content: Some(&args.text),
        runs: args.formatting_runs.as_deref().unwrap_or_default(),
        spacing: args.line_spacing.as_deref(),
        fractions: args.stacked_fractions.as_deref().unwrap_or_default(),
        fields: args.linked_fields.as_deref().unwrap_or_default(),
    }
    .validate()
}

fn validate_edit(args: wire::EditTextArgs) -> Result<(), Error> {
    Formatting {
        content: args.text.as_deref().map(String::as_str),
        runs: args
            .formatting_runs
            .as_deref()
            .and_then(Option::as_deref)
            .unwrap_or_default(),
        spacing: args.line_spacing.as_deref(),
        fractions: args
            .stacked_fractions
            .as_deref()
            .and_then(Option::as_deref)
            .unwrap_or_default(),
        fields: args
            .linked_fields
            .as_deref()
            .and_then(Option::as_deref)
            .unwrap_or_default(),
    }
    .validate()
}

pub(super) fn validate_response(method: &str, bytes: &[u8]) -> Result<(), Error> {
    if !matches!(
        method,
        wire::METHOD_SKETCH_EDIT_TEXT | wire::METHOD_SKETCH_GET_TEXT
    ) {
        return Ok(());
    }
    let result: wire::SketchTextResult = decode(method, "response", bytes)?;
    let style = result.style;
    Formatting {
        content: Some(&style.content),
        runs: style.formatting_runs.as_deref().unwrap_or_default(),
        spacing: style.line_spacing.as_deref(),
        fractions: style.stacked_fractions.as_deref().unwrap_or_default(),
        fields: style.linked_fields.as_deref().unwrap_or_default(),
    }
    .validate()
    .map_err(|error| Error(format!("method {method:?} response: {error}")))
}

struct Formatting<'a> {
    content: Option<&'a str>,
    runs: &'a [SketchTextRun],
    spacing: Option<&'a SketchTextLineSpacing>,
    fractions: &'a [SketchTextFraction],
    fields: &'a [SketchTextField],
}

impl Formatting<'_> {
    fn validate(self) -> Result<(), Error> {
        SketchTextLineSpacing::validate(self.spacing)?;
        SketchTextFraction::validate(self.content, self.fractions)?;
        SketchTextField::validate(self.content, self.fields, self.fractions)?;
        SketchTextRun::validate(self.content, self.runs)
    }
}

fn decode<T: DeserializeOwned>(method: &str, role: &str, bytes: &[u8]) -> Result<T, Error> {
    serde_json::from_slice(bytes).map_err(|error| {
        Error(format!(
            "method {method:?} {role}: expected text JSON: {error}"
        ))
    })
}
