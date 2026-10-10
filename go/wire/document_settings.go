// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Per-document settings (#147): the Document Settings dialog's persisted defaults, addressed by the
// document's session id (from documents.list). This starts with the Sketch tab — the constraint-
// inference preferences the sketch tools read — via document.getSketchSettings / setSketchSettings;
// 3D-sketch and modeling settings are separate future methods on the same document.

// GetSketchSettingsArgs is the request of [MethodDocumentGetSketchSettings]: the document whose
// sketch settings to read.
type GetSketchSettingsArgs struct {
	Document uint64 `json:"document"`
}

// SetSketchSettingsArgs is the request of [MethodDocumentSetSketchSettings]: replace the document's
// sketch settings with the given values.
type SetSketchSettingsArgs struct {
	Document uint64               `json:"document"`
	Settings types.SketchSettings `json:"settings"`
}

// SketchSettingsResult is the response of [MethodDocumentGetSketchSettings] and
// [MethodDocumentSetSketchSettings]: the document's current sketch settings.
type SketchSettingsResult struct {
	Settings types.SketchSettings `json:"settings"`
}
