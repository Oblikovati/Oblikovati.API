// SPDX-License-Identifier: Apache-2.0

package types

// Environment is a ribbon context within a document. The base environment is always shown; a
// contextual environment (e.g. sketch editing) contributes its tabs only while it is active —
// the mechanism behind a contextual tab like Inventor's Sketch tab (RibbonUI_Overview). An
// add-in scopes a control to an environment so it appears only in that context.
//
// This is the canonical, Apache-2.0 definition; the GPL implementation aliases it
// (app.Environment) so existing call sites are unaffected.
type Environment uint8

const (
	// BaseEnvironment is the document's normal environment; its controls always show.
	BaseEnvironment Environment = 0
	// SketchEnvironment is active while a sketch is open for editing; its controls form the
	// contextual Sketch tab and show only then.
	SketchEnvironment Environment = 1
)

var environmentNames = map[Environment]string{
	BaseEnvironment: "base", SketchEnvironment: "sketch",
}

// String returns the environment's stable name.
func (e Environment) String() string {
	if name, ok := environmentNames[e]; ok {
		return name
	}
	return "environment(?)"
}
