// SPDX-License-Identifier: Apache-2.0

package types

// AddInLoadBehavior is when (and whether) the host activates an installed add-in —
// the AddInLoadBehaviorEnum equivalent, narrowed to the behaviors the host honors
// today. Per-document-type loading (load-with-parts/assemblies/…) is intentionally
// not modeled until the host has per-document applets to defer to (M05-F01, #245).
//
// The zero value is LoadOnStartup so a freshly discovered add-in activates without
// any stored preference.
type AddInLoadBehavior uint8

const (
	// LoadOnStartup activates the add-in when the host starts (the default).
	LoadOnStartup AddInLoadBehavior = 0
	// LoadOnDemand registers the add-in but defers activation until something —
	// the user, another add-in, or a script — asks for it (addins.activate).
	LoadOnDemand AddInLoadBehavior = 1
	// LoadDisabled registers the add-in so it is listed, but the host never
	// activates it automatically and addins.activate refuses to.
	LoadDisabled AddInLoadBehavior = 2
)

var addInLoadBehaviorNames = map[AddInLoadBehavior]string{
	LoadOnStartup: "startup", LoadOnDemand: "demand", LoadDisabled: "disabled",
}

// String returns the behavior's stable name ("startup", "demand", "disabled").
func (b AddInLoadBehavior) String() string {
	if name, ok := addInLoadBehaviorNames[b]; ok {
		return name
	}
	return "addInLoadBehavior(?)"
}

// ParseAddInLoadBehavior maps a stable name back to its behavior; ok is false for an
// unknown name (callers keep their current value rather than guessing).
//
//	b, ok := types.ParseAddInLoadBehavior("demand") // LoadOnDemand, true
func ParseAddInLoadBehavior(name string) (AddInLoadBehavior, bool) {
	for b, n := range addInLoadBehaviorNames {
		if n == name {
			return b, true
		}
	}
	return LoadOnStartup, false
}
