// SPDX-License-Identifier: Apache-2.0

package wire

// Highlight sets (Oblikovati/Oblikovati#157). A highlight set is a named group of model
// references (the face/vertex key strings model.referenceKeys / model.selection report) that
// the viewport outlines in a colour WITHOUT selecting them — an add-in uses it to guide the
// user (Inventor HighlightSet.AddItem / Color). Colours are "#rrggbb" hex.

// CreateHighlightSetArgs is the request of [MethodModelHighlightSetCreate]: a unique Name and a
// "#rrggbb" Color. The set starts empty; add references with [MethodModelHighlightSetAddItems].
type CreateHighlightSetArgs struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// HighlightSetRefArgs addresses a highlight set by Name ([MethodModelHighlightSetDelete]).
type HighlightSetRefArgs struct {
	Name string `json:"name"`
}

// HighlightSetItemsArgs is the request of [MethodModelHighlightSetAddItems]: add the given
// model references (from model.referenceKeys) to the named set.
type HighlightSetItemsArgs struct {
	Name string   `json:"name"`
	Refs []string `json:"refs"`
}

// SetHighlightSetColorArgs is the request of [MethodModelHighlightSetSetColor]: re-colour the
// named set to a "#rrggbb" Color.
type SetHighlightSetColorArgs struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// HighlightSetInfo is one highlight set's summary: its name, "#rrggbb" colour, and how many
// references it holds. The response of create/addItems/setColor and a row of the list.
type HighlightSetInfo struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Count int    `json:"count"`
}

// ListHighlightSetsResult is the response of [MethodModelHighlightSetList].
type ListHighlightSetsResult struct {
	Sets []HighlightSetInfo `json:"sets"`
}
