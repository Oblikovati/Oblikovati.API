// SPDX-License-Identifier: Apache-2.0

use crate::{
    Error,
    json::JsonValue,
    wire::{AddFeatureArgs, featureargs::Arg},
};
use serde::Serialize;

impl<T: super::Caller> super::Features<'_, T> {
    /// Create a feature with its typed arguments and their declared operation kind.
    /// Example: `client.features().add_feature(&featureargs::Extrude::default())?`.
    pub fn add_feature<A: Arg + Serialize>(&self, args: &A) -> Result<JsonValue, Error> {
        let kind = args.kind();
        let payload = serde_json::to_value(args).map_err(|error| {
            Error(format!(
                "feature {kind:?}: expected serializable typed arguments: {error}"
            ))
        })?;
        self.add(&AddFeatureArgs {
            kind,
            args: JsonValue(Some(payload)),
        })
    }
}
