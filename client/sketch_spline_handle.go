// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// SetSplineHandle activates, edits, or deactivates the tangency handle on one
// fit point of a 2D interpolation spline (M06-F11, Oblikovati/Oblikovati#626).
func (s Sketch) SetSplineHandle(args wire.SetSplineHandleArgs) (wire.SplineHandleInfo, error) {
	var r wire.SplineHandleInfo
	return r, s.c.call(wire.MethodSketchSetSplineHandle, args, &r)
}

// SetSplineHandle activates, edits, or deactivates the tangency handle on one
// fit point of a 3D interpolation spline (M06-F11).
func (s Sketch3D) SetSplineHandle(args wire.SetSplineHandleArgs) (wire.SplineHandleInfo, error) {
	var r wire.SplineHandleInfo
	return r, s.c.call(wire.MethodSketch3DSetSplineHandle, args, &r)
}
