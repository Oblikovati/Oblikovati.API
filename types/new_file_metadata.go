// SPDX-License-Identifier: Apache-2.0

package types

// NewFileMetadata is the typed "new file" descriptor consumed by operations
// that mint a file the user never explicitly created — save-copy-as targets
// today; model-state and member exports later (M03-F09,
// Oblikovati/Oblikovati#610). Zero-value fields inherit from the source
// document of the operation.
type NewFileMetadata struct {
	// FileName is the target full file name, when the operation does not
	// already carry one.
	FileName string `json:"fileName,omitempty"`
	// DisplayName overrides the new file's display name.
	DisplayName string `json:"displayName,omitempty"`
	// TemplateFile seeds the new file from a template package.
	TemplateFile string `json:"templateFile,omitempty"`
	// SubType stamps the new file with a flavored document subtype id.
	SubType string `json:"subType,omitempty"`
}
