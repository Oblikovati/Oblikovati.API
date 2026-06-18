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
	var r wire.MassPropertiesResult
	return r, a.c.call(wire.MethodAnalysisMassProperties, args, &r)
}
