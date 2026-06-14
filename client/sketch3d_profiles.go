// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Profiles enumerates the closed, planar loops of the 3D sketch (area + plane normal +
// vertex count); each Index can feed a planar-section feature.
//
// mcp:tool list_sketch3d_profiles
// mcp:summary Enumerate a 3D sketch's closed profiles.
func (s Sketch3D) Profiles(index int) (wire.ListProfiles3DResult, error) {
	var r wire.ListProfiles3DResult
	return r, s.c.call(wire.MethodSketch3DProfiles, wire.Sketch3DArgs{SketchIndex: index}, &r)
}

// Paths enumerates the connected line/arc chains of the 3D sketch (the sweep/loft rails),
// each with its closed flag and vertex count.
//
// mcp:tool list_sketch3d_paths
// mcp:summary Enumerate a 3D sketch's open paths — the sweep/loft path candidates.
func (s Sketch3D) Paths(index int) (wire.ListPaths3DResult, error) {
	var r wire.ListPaths3DResult
	return r, s.c.call(wire.MethodSketch3DPaths, wire.Sketch3DArgs{SketchIndex: index}, &r)
}
