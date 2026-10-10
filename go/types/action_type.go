// SPDX-License-Identifier: Apache-2.0

package types

// ActionType is a bitmask of the edit actions that can be restricted on a
// parameter (parity: ActionTypeEnum, used by Parameter.DisabledActionTypes). A
// set bit means the corresponding action is DISABLED for that parameter.
//
// The reference API's ActionTypeEnum is a large flags enum whose numeric ids are
// not available in-repo; only the parameter-relevant subset is modelled here, so
// these values are iota-owned (single bits) rather than frozen at reference ids
// (contrast [ToleranceType]). The wire form is the list of set-bit spellings
// ([ActionType.Names]), never the raw mask, so the numeric layout stays internal.
type ActionType int32

const (
	// ActionNone is the empty mask: no action is disabled.
	ActionNone ActionType = 0
	// ActionEdit disables changing the parameter's expression or value.
	ActionEdit ActionType = 1 << 0 // 1
	// ActionRename disables renaming the parameter.
	ActionRename ActionType = 1 << 1 // 2
	// ActionDelete disables deleting the parameter.
	ActionDelete ActionType = 1 << 2 // 4
)

// actionTypeNames are the wire spellings of each single-bit action.
var actionTypeNames = map[ActionType]string{
	ActionEdit:   "edit",
	ActionRename: "rename",
	ActionDelete: "delete",
}

// actionTypeOrder fixes the emission order of [ActionType.Names] so the wire list
// is deterministic (map iteration is not).
var actionTypeOrder = []ActionType{ActionEdit, ActionRename, ActionDelete}

// Has reports whether action's bit is set in the mask.
func (a ActionType) Has(action ActionType) bool { return a&action != 0 }

// Names returns the wire spellings of the set bits, in a stable order; the empty
// mask returns nil (the wire omits it).
func (a ActionType) Names() []string {
	var names []string
	for _, bit := range actionTypeOrder {
		if a.Has(bit) {
			names = append(names, actionTypeNames[bit])
		}
	}
	return names
}

// ParseActionType resolves one action spelling to its single-bit value.
func ParseActionType(s string) (ActionType, bool) {
	return enumFromName(actionTypeNames, s)
}

// ActionTypeMask combines a list of action spellings into a mask, reporting the
// first unknown spelling (and the mask built so far) as false. Duplicate and
// out-of-order spellings are accepted; the empty list is [ActionNone].
func ActionTypeMask(names []string) (ActionType, bool) {
	mask := ActionNone
	for _, name := range names {
		bit, ok := ParseActionType(name)
		if !ok {
			return mask, false
		}
		mask |= bit
	}
	return mask, true
}
