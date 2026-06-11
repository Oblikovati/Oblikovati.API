// SPDX-License-Identifier: Apache-2.0

package types

// StartupActionType is what the application opens with — the StartupActionTypeEnum
// equivalent, narrowed to the actions the host implements (M05-F11, #618). The zero
// value preserves the historical behavior: a fresh part document ready to model in.
type StartupActionType uint8

const (
	// StartupNewPart opens a new part document at launch (the default).
	StartupNewPart StartupActionType = 0
	// StartupEmptyWorkspace opens with no document — the Get Started (ZeroDoc)
	// ribbon, for users who always begin from File ▸ Open.
	StartupEmptyWorkspace StartupActionType = 1
)

var startupActionNames = map[StartupActionType]string{
	StartupNewPart: "new-part", StartupEmptyWorkspace: "empty",
}

// String returns the action's stable name ("new-part", "empty").
func (a StartupActionType) String() string {
	if name, ok := startupActionNames[a]; ok {
		return name
	}
	return "startupAction(?)"
}
