// SPDX-License-Identifier: Apache-2.0

use crate::{
    Error,
    types::{SketchTextFraction, SketchTextLineSpacing, SketchTextRun},
    wire,
};
use serde::de::DeserializeOwned;

pub(super) fn validate_request(method: &str, bytes: &[u8]) -> Result<(), Error> {
    match method {
        wire::METHOD_SKETCH_ADD_TEXT => {
            let args: wire::AddTextArgs = decode(method, "request", bytes)?;
            SketchTextRun::validate(
                Some(&args.text),
                args.formatting_runs.as_deref().unwrap_or(&[]),
            )
            .and_then(|()| SketchTextLineSpacing::validate(args.line_spacing.as_deref()))
            .and_then(|()| {
                SketchTextFraction::validate(
                    Some(&args.text),
                    args.stacked_fractions.as_deref().unwrap_or(&[]),
                )
            })
        }
        wire::METHOD_SKETCH_EDIT_TEXT => {
            let args: wire::EditTextArgs = decode(method, "request", bytes)?;
            let runs = args.formatting_runs.as_deref().and_then(Option::as_deref);
            SketchTextRun::validate(
                args.text.as_deref().map(String::as_str),
                runs.unwrap_or(&[]),
            )
            .and_then(|()| SketchTextLineSpacing::validate(args.line_spacing.as_deref()))
            .and_then(|()| {
                let fractions = args.stacked_fractions.as_deref().and_then(Option::as_deref);
                SketchTextFraction::validate(
                    args.text.as_deref().map(String::as_str),
                    fractions.unwrap_or(&[]),
                )
            })
        }
        _ => return Ok(()),
    }
    .map_err(|error| Error(format!("method {method:?} request: {error}")))
}

pub(super) fn validate_response(method: &str, bytes: &[u8]) -> Result<(), Error> {
    if !matches!(
        method,
        wire::METHOD_SKETCH_EDIT_TEXT | wire::METHOD_SKETCH_GET_TEXT
    ) {
        return Ok(());
    }
    let result: wire::SketchTextResult = decode(method, "response", bytes)?;
    let runs = result.style.formatting_runs.as_deref().unwrap_or(&[]);
    SketchTextRun::validate(Some(&result.style.content), runs)
        .and_then(|()| SketchTextLineSpacing::validate(result.style.line_spacing.as_deref()))
        .and_then(|()| {
            SketchTextFraction::validate(
                Some(&result.style.content),
                result.style.stacked_fractions.as_deref().unwrap_or(&[]),
            )
        })
        .map_err(|error| Error(format!("method {method:?} response: {error}")))
}

fn decode<T: DeserializeOwned>(method: &str, role: &str, bytes: &[u8]) -> Result<T, Error> {
    serde_json::from_slice(bytes).map_err(|error| {
        Error(format!(
            "method {method:?} {role}: expected text JSON: {error}"
        ))
    })
}
