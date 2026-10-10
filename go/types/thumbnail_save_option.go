// SPDX-License-Identifier: Apache-2.0

package types

// ThumbnailSaveOption controls whether and how a preview thumbnail is captured
// when a document is saved (M03-F09, Oblikovati/Oblikovati#610). Thumbnails are
// git-ignored sidecar images, never document content (ADR-0020).
//
// The values are a frozen block matching the reference API's thumbnail-save
// enum; never renumber them. A host may implement a subset and reject the
// rest when the save options are written.
type ThumbnailSaveOption int32

const (
	// ThumbnailNone captures no thumbnail.
	ThumbnailNone ThumbnailSaveOption = 79873
	// ThumbnailIsoViewOnSave captures the isometric view at save time.
	ThumbnailIsoViewOnSave ThumbnailSaveOption = 79874
	// ThumbnailActiveWindowOnSave captures the active viewport at save time.
	ThumbnailActiveWindowOnSave ThumbnailSaveOption = 79875
	// ThumbnailActiveWindow captures the active viewport immediately.
	ThumbnailActiveWindow ThumbnailSaveOption = 79876
	// ThumbnailImportFromFile uses a caller-supplied image file.
	ThumbnailImportFromFile ThumbnailSaveOption = 79877
)

// thumbnailSaveOptionNames are the frozen wire spellings.
var thumbnailSaveOptionNames = map[ThumbnailSaveOption]string{
	ThumbnailNone:               "none",
	ThumbnailIsoViewOnSave:      "isoViewOnSave",
	ThumbnailActiveWindowOnSave: "activeWindowOnSave",
	ThumbnailActiveWindow:       "activeWindow",
	ThumbnailImportFromFile:     "importFromFile",
}

// String returns the thumbnail option's wire spelling.
func (o ThumbnailSaveOption) String() string { return enumName(thumbnailSaveOptionNames, o, "enum(?)") }

// ParseThumbnailSaveOption resolves a wire spelling back to its option.
func ParseThumbnailSaveOption(s string) (ThumbnailSaveOption, bool) {
	return enumFromName(thumbnailSaveOptionNames, s)
}
