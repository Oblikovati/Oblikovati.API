// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// RegionProperties computes the full section property set (area, perimeter,
// centroid, moments of inertia, principal axes) of a closed profile at the
// given accuracy (M06-F08, Oblikovati/Oblikovati#623).
//
// mcp:tool sketch_region_properties
// mcp:summary Computes the full section property set (area, perimeter, centroid, moments of inertia, principal axes) of a closed profile at the given accuracy (M06-F08, Oblikovati/Oblikovati#623).
func (s Sketch) RegionProperties(index, profileIndex int, accuracy types.Accuracy) (wire.RegionPropertiesResult, error) {
	return call[wire.RegionPropertiesResult](s.c, wire.MethodSketchRegionProperties, wire.RegionPropertiesArgs{
		SketchIndex: index, ProfileIndex: profileIndex, Accuracy: accuracy.String(),
	})
}

// RegionProperties computes the section property set of a planar closed 3D
// profile, reported in the profile plane's coordinates (M06-F08).
//
// mcp:tool sketch3d_region_properties
// mcp:summary Computes the section property set of a planar closed 3D profile, reported in the profile plane's coordinates (M06-F08).
func (s Sketch3D) RegionProperties(index, profileIndex int, accuracy types.Accuracy) (wire.RegionPropertiesResult, error) {
	return call[wire.RegionPropertiesResult](s.c, wire.MethodSketch3DRegionProperties, wire.RegionPropertiesArgs{
		SketchIndex: index, ProfileIndex: profileIndex, Accuracy: accuracy.String(),
	})
}
