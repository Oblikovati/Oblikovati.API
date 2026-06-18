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
// mcp:summary Compute the active part's mass properties — volume (mm³), surface area (mm²), centre of mass (mm) and mass (g) — over all its solid bodies. densityGCm3 is the material density in g/cm³ (0 ⇒ 1.0).
func (a Analysis) MassProperties(args wire.MassPropertiesArgs) (wire.MassPropertiesResult, error) {
	var r wire.MassPropertiesResult
	return r, a.c.call(wire.MethodAnalysisMassProperties, args, &r)
}
