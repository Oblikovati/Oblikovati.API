// SPDX-License-Identifier: Apache-2.0

package operations

import (
	"oblikovati.org/api/client"
	"oblikovati.org/api/wire"
)

// ListAllAxes includes construction axes in the active local model.
// Example: ListAllAxes(host). Use Occurrence only when the host supports that scope.
func ListAllAxes(host client.Caller) (wire.ListWorkAxesResult, error) {
	return client.New(host).WorkAxes().ListWithOptions(wire.ListWorkAxesArgs{IncludeConstruction: true})
}

// ListAllPoints includes construction points without changing their visibility.
// Example: ListAllPoints(host).
func ListAllPoints(host client.Caller) (wire.ListWorkPointsResult, error) {
	return client.New(host).WorkPoints().ListWithOptions(wire.ListWorkPointsArgs{IncludeConstruction: true})
}
