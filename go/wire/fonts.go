// SPDX-License-Identifier: Apache-2.0

package wire

// FontFace is one selectable font in the picker: a family + style, its Source ("embedded" for
// an application-bundled face, "system" for a host-installed one), and — for a system face —
// the file Path whose bytes are embedded into the document when it is chosen (ADR-0031).
type FontFace struct {
	Family string `json:"family"`
	Style  string `json:"style,omitempty"`
	Source string `json:"source"`         // "embedded" | "system"
	Path   string `json:"path,omitempty"` // system fonts only: the file to embed on select
}

// ListFontsResult is the response of [MethodFontsList]: every face the picker can offer — the
// application's bundled faces plus the fonts installed on the host.
type ListFontsResult struct {
	Faces []FontFace `json:"faces"`
}

// SetTextFontArgs is the request of [MethodSketchSetTextFont]: choose the font of the sketch
// TEXT entity EntityID in sketch SketchIndex. Provide Path for a host font (its bytes are
// embedded into the document) OR Family for a bundled face (recorded without bytes); if both
// are set, Path wins. The font becomes a document resource the text/emboss resolves by, so the
// document is self-contained (ADR-0031).
type SetTextFontArgs struct {
	SketchIndex int    `json:"sketchIndex"`
	EntityID    uint64 `json:"entityId"`
	Family      string `json:"family,omitempty"`
	Path        string `json:"path,omitempty"`
}

// SetTextFontResult reports the document resource UUID the text now cites and the resolved
// family name (for the picker's label).
type SetTextFontResult struct {
	Resource string `json:"resource"`
	Family   string `json:"family"`
}
