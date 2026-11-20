// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// ValidateSketchTextFormatting checks AddTextArgs against its content.
// For example, validate a decoded add request before applying it to the sketch.
func (args AddTextArgs) ValidateSketchTextFormatting() error {
	return types.ValidateSketchTextRuns(&args.Text, args.FormattingRuns)
}

// ValidateSketchTextFormatting checks EditTextArgs before the entity is loaded.
// For example, validate first, then check the resolved content for a format-only edit.
func (args EditTextArgs) ValidateSketchTextFormatting() error {
	if args.FormattingRuns == nil {
		return nil
	}
	return types.ValidateSketchTextRuns(args.Text, *args.FormattingRuns)
}

// ValidateSketchTextFormatting checks a resolved style against its content.
// For example, validate a host result before passing it to a text editor.
func (result SketchTextResult) ValidateSketchTextFormatting() error {
	return types.ValidateSketchTextRuns(&result.Style.Content, result.Style.FormattingRuns)
}
