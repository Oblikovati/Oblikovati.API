// SPDX-License-Identifier: Apache-2.0

package types

// PromptRestriction is whether the user may suppress a prompt by remembering their
// answer — the PromptMessageRestrictionsEnum equivalent, narrowed to the two
// behaviors the host implements (M05-F09, #616).
type PromptRestriction uint8

const (
	// PromptAlwaysAsk shows the prompt every time (the zero value).
	PromptAlwaysAsk PromptRestriction = 0
	// PromptAllowRemember offers a "remember my answer" choice; a remembered
	// prompt resolves instantly from the stored answer until it is reset.
	PromptAllowRemember PromptRestriction = 1
)

var promptRestrictionNames = map[PromptRestriction]string{
	PromptAlwaysAsk: "always-ask", PromptAllowRemember: "allow-remember",
}

// String returns the restriction's stable name.
func (p PromptRestriction) String() string {
	if name, ok := promptRestrictionNames[p]; ok {
		return name
	}
	return "promptRestriction(?)"
}
