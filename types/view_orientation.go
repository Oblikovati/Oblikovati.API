// SPDX-License-Identifier: Apache-2.0

package types

// ViewOrientationTypeEnum names a standard camera orientation — the front/top/iso… presets a
// view cube or the named-view picker restores. The numeric ids are stable, frozen values
// (10753–10773).
//
// This is the canonical Apache-2.0 definition; the GPL implementation aliases it.
type ViewOrientationTypeEnum int32

const (
	// DefaultViewOrientation is the document's default orientation (10753).
	DefaultViewOrientation ViewOrientationTypeEnum = 10753
	// TopViewOrientation looks straight down (10754).
	TopViewOrientation ViewOrientationTypeEnum = 10754
	// RightViewOrientation looks from the right (10755).
	RightViewOrientation ViewOrientationTypeEnum = 10755
	// BackViewOrientation looks from behind (10756).
	BackViewOrientation ViewOrientationTypeEnum = 10756
	// BottomViewOrientation looks straight up (10757).
	BottomViewOrientation ViewOrientationTypeEnum = 10757
	// LeftViewOrientation looks from the left (10758).
	LeftViewOrientation ViewOrientationTypeEnum = 10758
	// IsoTopRightViewOrientation is the top-right isometric (10759).
	IsoTopRightViewOrientation ViewOrientationTypeEnum = 10759
	// IsoTopLeftViewOrientation is the top-left isometric (10760).
	IsoTopLeftViewOrientation ViewOrientationTypeEnum = 10760
	// IsoBottomRightViewOrientation is the bottom-right isometric (10761).
	IsoBottomRightViewOrientation ViewOrientationTypeEnum = 10761
	// IsoBottomLeftViewOrientation is the bottom-left isometric (10762).
	IsoBottomLeftViewOrientation ViewOrientationTypeEnum = 10762
	// ArbitraryViewOrientation is a non-standard orientation (10763).
	ArbitraryViewOrientation ViewOrientationTypeEnum = 10763
	// FrontViewOrientation looks from the front (10764).
	FrontViewOrientation ViewOrientationTypeEnum = 10764
	// CurrentViewOrientation keeps the current orientation (10765).
	CurrentViewOrientation ViewOrientationTypeEnum = 10765
	// SavedCameraViewOrientation restores a saved camera (10766).
	SavedCameraViewOrientation ViewOrientationTypeEnum = 10766
	// FlatPivotRightViewOrientation pivots flat to the right (10767).
	FlatPivotRightViewOrientation ViewOrientationTypeEnum = 10767
	// FlatPivotLeftViewOrientation pivots flat to the left (10768).
	FlatPivotLeftViewOrientation ViewOrientationTypeEnum = 10768
	// FlatPivot180ViewOrientation pivots flat 180° (10769).
	FlatPivot180ViewOrientation ViewOrientationTypeEnum = 10769
	// FlatBacksideViewOrientation flips to the flat backside (10770).
	FlatBacksideViewOrientation ViewOrientationTypeEnum = 10770
	// FlatBacksidePivotRightViewOrientation is flat backside pivoted right (10771).
	FlatBacksidePivotRightViewOrientation ViewOrientationTypeEnum = 10771
	// FlatBacksidePivotLeftViewOrientation is flat backside pivoted left (10772).
	FlatBacksidePivotLeftViewOrientation ViewOrientationTypeEnum = 10772
	// FlatBacksidePivot180ViewOrientation is flat backside pivoted 180° (10773).
	FlatBacksidePivot180ViewOrientation ViewOrientationTypeEnum = 10773
)

var viewOrientationNames = map[ViewOrientationTypeEnum]string{
	DefaultViewOrientation:                "Default",
	TopViewOrientation:                    "Top",
	RightViewOrientation:                  "Right",
	BackViewOrientation:                   "Back",
	BottomViewOrientation:                 "Bottom",
	LeftViewOrientation:                   "Left",
	IsoTopRightViewOrientation:            "Iso Top Right",
	IsoTopLeftViewOrientation:             "Iso Top Left",
	IsoBottomRightViewOrientation:         "Iso Bottom Right",
	IsoBottomLeftViewOrientation:          "Iso Bottom Left",
	ArbitraryViewOrientation:              "Arbitrary",
	FrontViewOrientation:                  "Front",
	CurrentViewOrientation:                "Current",
	SavedCameraViewOrientation:            "Saved Camera",
	FlatPivotRightViewOrientation:         "Flat Pivot Right",
	FlatPivotLeftViewOrientation:          "Flat Pivot Left",
	FlatPivot180ViewOrientation:           "Flat Pivot 180",
	FlatBacksideViewOrientation:           "Flat Backside",
	FlatBacksidePivotRightViewOrientation: "Flat Backside Pivot Right",
	FlatBacksidePivotLeftViewOrientation:  "Flat Backside Pivot Left",
	FlatBacksidePivot180ViewOrientation:   "Flat Backside Pivot 180",
}

// String returns the view-orientation's user-facing name.
func (v ViewOrientationTypeEnum) String() string {
	return enumName(viewOrientationNames, v, "viewOrientation(?)")
}

// IsValid reports whether v is a defined view orientation.
func (v ViewOrientationTypeEnum) IsValid() bool {
	return enumValid(viewOrientationNames, v)
}

// AllViewOrientations returns every defined view orientation, in the standard picker order.
func AllViewOrientations() []ViewOrientationTypeEnum {
	return []ViewOrientationTypeEnum{
		FrontViewOrientation, BackViewOrientation, TopViewOrientation, BottomViewOrientation,
		LeftViewOrientation, RightViewOrientation, IsoTopRightViewOrientation, IsoTopLeftViewOrientation,
		IsoBottomRightViewOrientation, IsoBottomLeftViewOrientation, DefaultViewOrientation,
		CurrentViewOrientation, ArbitraryViewOrientation, SavedCameraViewOrientation,
		FlatPivotRightViewOrientation, FlatPivotLeftViewOrientation, FlatPivot180ViewOrientation,
		FlatBacksideViewOrientation, FlatBacksidePivotRightViewOrientation,
		FlatBacksidePivotLeftViewOrientation, FlatBacksidePivot180ViewOrientation,
	}
}
