// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"

	"oblikovati.org/api/wire"
)

// PatternCircular replicates one or more existing features in a circular array about an axis
// (#189) — the typed form of an add_feature call with kind patternCircular. The count may be a
// parameter expression (args.CountExpr, e.g. "slots"), enabling the canonical parametric
// workflow: model one slot, cut it, then circular-pattern it N = slots.
//
// mcp:tool pattern_circular
// mcp:summary Circular-pattern existing features about an axis. {sourceFeatures:[names], count or countExpr, angle:"360 deg", axisPoint?, axisDir?}.
func (f Features) PatternCircular(args wire.CircularPatternFeatureArgs) (json.RawMessage, error) {
	return addPattern(f, wire.FeatureKindPatternCircular, args)
}

// PatternRectangular replicates features on a rectangular grid (#189); CountX/CountY may be given
// as parameter expressions (CountXExpr/CountYExpr).
//
// mcp:tool pattern_rectangular
// mcp:summary Rectangular-pattern existing features on a grid. {sourceFeatures:[names], countX/countY or *Expr, stepX:[x,y,z] cm, stepY?}.
func (f Features) PatternRectangular(args wire.RectangularPatternFeatureArgs) (json.RawMessage, error) {
	return addPattern(f, wire.FeatureKindPatternRectangular, args)
}

// MirrorFeatures mirrors features across a plane (#189).
//
// mcp:tool mirror_features
// mcp:summary Mirror existing features across a plane. {sourceFeatures:[names], normal:[x,y,z], origin?}.
func (f Features) MirrorFeatures(args wire.MirrorFeatureArgs) (json.RawMessage, error) {
	return addPattern(f, wire.FeatureKindMirror, args)
}

// addPattern marshals a typed pattern-args struct and dispatches it through features.add, so the
// three pattern helpers share one encode-and-call path (the house no-duplication rule). It is a
// free function because Go methods cannot carry type parameters.
func addPattern[A wire.CircularPatternFeatureArgs | wire.RectangularPatternFeatureArgs |
	wire.MirrorFeatureArgs](f Features, kind string, args A) (json.RawMessage, error) {
	raw, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	return f.Add(wire.AddFeatureArgs{Kind: kind, Args: raw})
}
