// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// ValidateSketchTextFormatting checks AddTextArgs against its content.
// For example, validate a decoded add request before applying it to the sketch.
func (args AddTextArgs) ValidateSketchTextFormatting() error {
	return sketchTextFormatChecks{
		content: &args.Text, spacing: args.LineSpacing, runs: args.FormattingRuns,
		fractions: args.StackedFractions, fields: args.LinkedFields,
	}.validate()
}

// ValidateSketchTextFormatting checks EditTextArgs before the entity is loaded.
// For example, validate first, then check the resolved content for a format-only edit.
func (args EditTextArgs) ValidateSketchTextFormatting() error {
	return sketchTextFormatChecks{
		content: args.Text, spacing: args.LineSpacing, runs: sketchTextOptionalSlice(args.FormattingRuns),
		fractions: sketchTextOptionalSlice(args.StackedFractions), fields: sketchTextOptionalSlice(args.LinkedFields),
	}.validate()
}

// ValidateSketchTextFormatting checks a resolved style against its content.
// For example, validate a host result before passing it to a text editor.
func (result SketchTextResult) ValidateSketchTextFormatting() error {
	style := result.Style
	return sketchTextFormatChecks{
		content: &style.Content, spacing: style.LineSpacing, runs: style.FormattingRuns,
		fractions: style.StackedFractions, fields: style.LinkedFields,
	}.validate()
}

type sketchTextFormatChecks struct {
	content   *string
	spacing   *types.SketchTextLineSpacing
	runs      []types.SketchTextRun
	fractions []types.SketchTextFraction
	fields    []types.SketchTextField
}

func (checks sketchTextFormatChecks) validate() error {
	if err := checks.spacing.Validate(); err != nil {
		return err
	}
	if err := types.ValidateSketchTextFractions(checks.content, checks.fractions); err != nil {
		return err
	}
	if err := types.ValidateSketchTextFields(checks.content, checks.fields, checks.fractions); err != nil {
		return err
	}
	return types.ValidateSketchTextRuns(checks.content, checks.runs)
}

func sketchTextOptionalSlice[T any](value *[]T) []T {
	if value == nil {
		return nil
	}
	return *value
}
