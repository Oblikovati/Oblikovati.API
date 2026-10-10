// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// TestASideFaceStatusFrozenBlock pins the reference ids and wire spellings.
func TestASideFaceStatusFrozenBlock(t *testing.T) {
	want := map[ASideFaceStatus]string{
		106241: "default", 106242: "sick", 106243: "upToDate",
	}
	assertFrozenBlock(t, "ASideFaceStatus", want, aSideFaceStatusNames, ParseASideFaceStatus)
}
