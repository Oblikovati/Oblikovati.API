// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The analysis operation group (M18-F01 #423): engineering analysis on the model — mass properties
// of the active part.

// Analysis is the analysis operation group.
type Analysis struct{ c *Client }

// Analysis returns the analysis operation group.
func (c *Client) Analysis() Analysis { return Analysis{c} }

// MassProperties computes the active part's mass properties (volume, surface area, centre of mass,
// mass) for the given material density.
//
// mcp:tool analysis_mass_properties
// mcp:summary Compute the active part's mass properties over all its solid bodies — volume (mm³), surface area (mm²), centre of mass (mm), mass (g), and mass moment of inertia about the centroid (g·mm²) with principal moments/axes. densityGCm3 overrides the material density (0 ⇒ the assigned material's, else 1.0); accuracy is low|medium|high.
func (a Analysis) MassProperties(args wire.MassPropertiesArgs) (wire.MassPropertiesResult, error) {
	return call[wire.MassPropertiesResult](a.c, wire.MethodAnalysisMassProperties, args)
}

// ModelHealth aggregates the active part's feature health — the overall status, the sick count, and
// every feature that is not OK.
//
// mcp:tool analysis_model_health
// mcp:summary Aggregate the active part's model health: the overall (worst) status across its features, the count of sick features, and every feature that is not "ok" (with its status and reason) so they can be listed for repair.
func (a Analysis) ModelHealth(args wire.ModelHealthArgs) (wire.ModelHealthResult, error) {
	return call[wire.ModelHealthResult](a.c, wire.MethodAnalysisModelHealth, args)
}

// Measure reports a geometric quantity of one or two of the active part's entities.
//
// mcp:tool analysis_measure
// mcp:summary Measure an entity of the active part's body (bodyIndex) by reference key: type "length" (edge keyA), "area" (face keyA), "distance" (between vertices keyA and keyB), "minDistance" (closest approach between two entities keyA and keyB, each a vertex/edge/face), "angle" (between two entities keyA and keyB, an edge direction or planar-face normal; or with keyC, the angle at apex vertex keyB between vertices keyA and keyC), or "loopLength" (the perimeter of face keyA). Returns the value with its unit (mm, mm² or deg).
func (a Analysis) Measure(args wire.MeasureArgs) (wire.MeasureResult, error) {
	return call[wire.MeasureResult](a.c, wire.MethodAnalysisMeasure, args)
}
