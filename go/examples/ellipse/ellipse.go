// SPDX-License-Identifier: Apache-2.0

// Package ellipse demonstrates independent ellipse radius dimensions.
package ellipse

import (
	"oblikovati.org/api/client"
	"oblikovati.org/api/wire"
)

// AddMinorRadius drives an existing owned ellipse's minor radius.
// Example: AddMinorRadius(host, 0, ellipseID, "5 mm").
func AddMinorRadius(host client.Caller, sketchIndex int, ellipseID uint64, expression string) (wire.AddDimensionResult, error) {
	return client.New(host).Sketch().Dimension(sketchIndex).EllipseMinorRadius(ellipseID, expression)
}
