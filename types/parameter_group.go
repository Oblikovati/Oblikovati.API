// SPDX-License-Identifier: Apache-2.0

package types

// ParameterGroupKey is the immutable, locale-stable internal name that
// identifies a custom parameter group for its whole life (M02-F05,
// Oblikovati/Oblikovati#604). The editable, localizable display name lives
// beside it; persistence, wire methods, and membership all address a group by
// this key, so renaming the display never breaks references.
//
//	key := types.ParameterGroupKey("com.example.gears:ratios")
type ParameterGroupKey string

// String returns the key's raw spelling.
func (k ParameterGroupKey) String() string { return string(k) }
