// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// SetCustomLineType loads a named line-type definition from an industry-standard
// .lin file onto the sketch and switches its lineType override to "custom".
//
//	lt, err := c.Sketch().SetCustomLineType(wire.SetSketchCustomLineTypeArgs{
//		SketchIndex: 0, FullFileName: "styles.lin", LineTypeName: "DASHDOT"})
//
// mcp:tool sketch_set_custom_line_type
// mcp:summary Loads a named line-type definition from an industry-standard .lin file onto the sketch and switches its lineType override to "custom".
func (s Sketch) SetCustomLineType(args wire.SetSketchCustomLineTypeArgs) (wire.SketchCustomLineTypeResult, error) {
	return call[wire.SketchCustomLineTypeResult](s.c, wire.MethodSketchSetCustomLineType, args)
}

// GetCustomLineType returns the sketch's loaded custom line-type definition, if any.
//
//	lt, err := c.Sketch().GetCustomLineType(0)
//
// mcp:tool sketch_get_custom_line_type
// mcp:summary Returns the sketch's loaded custom line-type definition, if any.
func (s Sketch) GetCustomLineType(index int) (wire.SketchCustomLineTypeResult, error) {
	return call[wire.SketchCustomLineTypeResult](s.c, wire.MethodSketchGetCustomLineType, wire.SketchArgs{SketchIndex: index})
}
