// SPDX-License-Identifier: Apache-2.0

package client

import "github.com/Oblikovati/api/wire"

// Offset offsets a single line/circle/arc (by entity id) by a signed unit-bearing
// distance, returning the new entity's id and kind.
func (s Sketch) Offset(index int, entity uint64, distance string) (wire.OffsetSketchResult, error) {
	var r wire.OffsetSketchResult
	args := wire.OffsetSketchArgs{SketchIndex: index, Entity: entity, Distance: distance}
	return r, s.c.call(wire.MethodSketchOffset, args, &r)
}

// OffsetChain offsets a connected chain of lines (ids in order) by a signed unit-bearing
// distance, mitring the joins; returns the created line ids.
func (s Sketch) OffsetChain(index int, lines []uint64, distance string) (wire.OffsetSketchResult, error) {
	var r wire.OffsetSketchResult
	args := wire.OffsetSketchArgs{SketchIndex: index, Entities: lines, Distance: distance}
	return r, s.c.call(wire.MethodSketchOffset, args, &r)
}

// AutoDimension fully constrains the sketch (grounds free geometry to 0 DOF), returning
// the number of constraints added and the resulting DOF.
func (s Sketch) AutoDimension(index int) (wire.AutoDimensionResult, error) {
	var r wire.AutoDimensionResult
	return r, s.c.call(wire.MethodSketchAutoDimension, wire.SketchArgs{SketchIndex: index}, &r)
}

// Project projects part edges/vertices (by reference-key string) onto the sketch plane as
// associative reference geometry; mode "include" projects them as ordinary geometry.
func (s Sketch) Project(index int, refs []string, mode string) (wire.ProjectGeometryResult, error) {
	var r wire.ProjectGeometryResult
	args := wire.ProjectGeometryArgs{SketchIndex: index, Refs: refs, Mode: mode}
	return r, s.c.call(wire.MethodSketchProject, args, &r)
}

// Include projects part topology as ordinary sketch geometry (a convenience for Project
// with mode "include").
func (s Sketch) Include(index int, refs []string) (wire.ProjectGeometryResult, error) {
	return s.Project(index, refs, "include")
}

// AddImage places a raster image (ref is a package-store reference) anchored at [x,y] cm
// with unit-bearing width/height; rotation and opacity are optional ("" / 0).
func (s Sketch) AddImage(index int, ref string, anchor []float64, width, height, rotation string, opacity float64) (wire.AddSketchImageResult, error) {
	var r wire.AddSketchImageResult
	args := wire.AddSketchImageArgs{
		SketchIndex: index, Ref: ref, Anchor: anchor,
		Width: width, Height: height, Rotation: rotation, Opacity: opacity,
	}
	return r, s.c.call(wire.MethodSketchAddImage, args, &r)
}

// AddFillRegion fills the closed region containing seed ([x,y] cm) with the named style.
func (s Sketch) AddFillRegion(index int, seed []float64, style string) (wire.AddEntityIDResult, error) {
	var r wire.AddEntityIDResult
	args := wire.AddFillRegionArgs{SketchIndex: index, Seed: seed, Style: style}
	return r, s.c.call(wire.MethodSketchAddFillRegion, args, &r)
}

// AddText places sketch text at anchor ([x,y] cm) with a unit-bearing height; rotation and
// justify ("left"|"center"|"right") are optional.
func (s Sketch) AddText(index int, anchor []float64, text, height, rotation, justify string) (wire.AddEntityIDResult, error) {
	var r wire.AddEntityIDResult
	args := wire.AddTextArgs{SketchIndex: index, Anchor: anchor, Text: text, Height: height, Rotation: rotation, Justify: justify}
	return r, s.c.call(wire.MethodSketchAddText, args, &r)
}
