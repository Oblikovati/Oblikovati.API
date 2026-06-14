// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Offset offsets a single line/circle/arc (by entity id) by a signed unit-bearing
// distance, returning the new entity's id and kind.
//
// mcp:tool offset_sketch
// mcp:summary Offset a sketch curve (entity id), a line chain (entities), or a whole closed region (profileIndex) by a distance expression — region offset is OpenSCAD offset(r): +grows/−shrinks with rounded convex corners.
func (s Sketch) Offset(index int, entity uint64, distance string) (wire.OffsetSketchResult, error) {
	var r wire.OffsetSketchResult
	args := wire.OffsetSketchArgs{SketchIndex: index, Entity: entity, Distance: distance}
	return r, s.c.call(wire.MethodSketchOffset, args, &r)
}

// OffsetChain offsets a connected chain of lines (ids in order) by a signed unit-bearing
// distance, mitring the joins; returns the created line ids.
//
// mcp:tool offset_sketch
// mcp:summary Offset a sketch curve (entity id), a line chain (entities), or a whole closed region (profileIndex) by a distance expression — region offset is OpenSCAD offset(r): +grows/−shrinks with rounded convex corners.
func (s Sketch) OffsetChain(index int, lines []uint64, distance string) (wire.OffsetSketchResult, error) {
	var r wire.OffsetSketchResult
	args := wire.OffsetSketchArgs{SketchIndex: index, Entities: lines, Distance: distance}
	return r, s.c.call(wire.MethodSketchOffset, args, &r)
}

// AutoDimension fully constrains the sketch (grounds free geometry to 0 DOF), returning
// the number of constraints added and the resulting DOF.
//
// mcp:tool auto_dimension_sketch
// mcp:summary Fully constrain a sketch automatically with dimensions and constraints; reports any remaining DOF.
func (s Sketch) AutoDimension(index int) (wire.AutoDimensionResult, error) {
	var r wire.AutoDimensionResult
	return r, s.c.call(wire.MethodSketchAutoDimension, wire.SketchArgs{SketchIndex: index}, &r)
}

// Project projects part edges/vertices (by reference-key string) onto the sketch plane as
// associative reference geometry; mode "include" projects them as ordinary geometry.
//
// mcp:tool project_geometry
// mcp:summary Project part edges/vertices/faces (by reference key) onto a sketch as reference geometry: {sketchIndex, refs:[…], mode}.
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
//
// mcp:tool add_sketch_image
// mcp:summary Place a raster image into a sketch (reference, anchor, size, rotation, opacity).
func (s Sketch) AddImage(index int, ref string, anchor []float64, width, height, rotation string, opacity float64) (wire.AddSketchImageResult, error) {
	var r wire.AddSketchImageResult
	args := wire.AddSketchImageArgs{
		SketchIndex: index, Ref: ref, Anchor: anchor,
		Width: width, Height: height, Rotation: rotation, Opacity: opacity,
	}
	return r, s.c.call(wire.MethodSketchAddImage, args, &r)
}

// AddFillRegion fills the closed region containing seed ([x,y] cm) with the named style.
//
// mcp:tool add_fill_region
// mcp:summary Add a fill/hatch region to a sketch, seeded at a point [x,y] inside a closed loop.
func (s Sketch) AddFillRegion(index int, seed []float64, style string) (wire.AddEntityIDResult, error) {
	var r wire.AddEntityIDResult
	args := wire.AddFillRegionArgs{SketchIndex: index, Seed: seed, Style: style}
	return r, s.c.call(wire.MethodSketchAddFillRegion, args, &r)
}

// AddText places sketch text at anchor ([x,y] cm) with a unit-bearing height; rotation and
// justify ("left"|"center"|"right") are optional.
//
// mcp:tool add_sketch_text
// mcp:summary Add a text box to a sketch at an anchor [x,y] with a string, height and optional rotation/justify.
func (s Sketch) AddText(index int, anchor []float64, text, height, rotation, justify string) (wire.AddEntityIDResult, error) {
	var r wire.AddEntityIDResult
	args := wire.AddTextArgs{SketchIndex: index, Anchor: anchor, Text: text, Height: height, Rotation: rotation, Justify: justify}
	return r, s.c.call(wire.MethodSketchAddText, args, &r)
}

// AddTextWith places sketch text with the full field set (font family/size + vertical
// alignment), so the text can be embossed/extruded by reference.
//
// mcp:tool add_sketch_text
// mcp:summary Add a text box to a sketch at an anchor [x,y] with a string, height and optional rotation/justify.
func (s Sketch) AddTextWith(args wire.AddTextArgs) (wire.AddEntityIDResult, error) {
	var r wire.AddEntityIDResult
	return r, s.c.call(wire.MethodSketchAddText, args, &r)
}

// SetTextFont sets the font of the sketch text entity entityID: pass a system font file path
// (its bytes are embedded into the document) or a bundled face family. The font becomes a
// document resource the text/emboss resolves by, so the document stays self-contained (ADR-0031).
//
// mcp:tool sketch_set_text_font
// mcp:summary Sets the font of the sketch text entity entityID: pass a system font file path (its bytes are embedded into the document) or a bundled face family.
func (s Sketch) SetTextFont(args wire.SetTextFontArgs) (wire.SetTextFontResult, error) {
	var r wire.SetTextFontResult
	return r, s.c.call(wire.MethodSketchSetTextFont, args, &r)
}

// EditText applies a partial edit to an existing sketch text entity (only the set fields),
// returning the entity's resolved style. Editing re-derives the text's geometry, so any
// emboss referencing it recomputes.
//
// mcp:tool sketch_edit_text
// mcp:summary Applies a partial edit to an existing sketch text entity (only the set fields), returning the entity's resolved style.
func (s Sketch) EditText(args wire.EditTextArgs) (wire.SketchTextResult, error) {
	var r wire.SketchTextResult
	return r, s.c.call(wire.MethodSketchEditText, args, &r)
}

// GetText reads back a sketch text entity's style.
//
// mcp:tool sketch_get_text
// mcp:summary Reads back a sketch text entity's style.
func (s Sketch) GetText(index int, entity uint64) (wire.SketchTextResult, error) {
	var r wire.SketchTextResult
	args := wire.GetTextArgs{SketchIndex: index, EntityID: entity}
	return r, s.c.call(wire.MethodSketchGetText, args, &r)
}
