// SPDX-License-Identifier: Apache-2.0

package types

// ParameterKind is a parameter's category (Inventor's ParameterTypeEnum), with
// stable explicit ids. Model/User/Table parameters are user-editable; Reference
// (driven by geometry) and Derived (linked from another document) are read-only.
//
// This is the canonical, Apache-2.0 definition; the GPL implementation aliases it
// (model/param.ParameterKind) so existing call sites are unaffected.
type ParameterKind uint8

const (
	ModelParam     ParameterKind = 1
	UserParam      ParameterKind = 2
	ReferenceParam ParameterKind = 3
	DerivedParam   ParameterKind = 4
	TableParam     ParameterKind = 5
)

var parameterKindNames = map[ParameterKind]string{
	ModelParam: "model", UserParam: "user", ReferenceParam: "reference",
	DerivedParam: "derived", TableParam: "table",
}

// String returns the kind's name.
func (k ParameterKind) String() string {
	if name, ok := parameterKindNames[k]; ok {
		return name
	}
	return "kind(?)"
}

// Editable reports whether a user may set this kind's expression/value (Model,
// User and Table parameters are editable; Reference and Derived are read-only).
func (k ParameterKind) Editable() bool {
	return k == ModelParam || k == UserParam || k == TableParam
}
