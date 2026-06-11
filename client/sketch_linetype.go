// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// SetCustomLineType loads a named line-type definition from an industry-standard
// .lin file onto the sketch and switches its lineType override to "custom".
//
//	lt, err := c.Sketch().SetCustomLineType(wire.SetSketchCustomLineTypeArgs{
//		SketchIndex: 0, FullFileName: "styles.lin", LineTypeName: "DASHDOT"})
func (s Sketch) SetCustomLineType(args wire.SetSketchCustomLineTypeArgs) (wire.SketchCustomLineTypeResult, error) {
	var r wire.SketchCustomLineTypeResult
	return r, s.c.call(wire.MethodSketchSetCustomLineType, args, &r)
}

// GetCustomLineType returns the sketch's loaded custom line-type definition, if any.
//
//	lt, err := c.Sketch().GetCustomLineType(0)
func (s Sketch) GetCustomLineType(index int) (wire.SketchCustomLineTypeResult, error) {
	var r wire.SketchCustomLineTypeResult
	return r, s.c.call(wire.MethodSketchGetCustomLineType, wire.SketchArgs{SketchIndex: index}, &r)
}
