// SPDX-License-Identifier: Apache-2.0

package types

// The geometry-query result vocabulary (M01-F06, #603). Numeric values are
// FROZEN to the reference SolutionNatureEnum / ContainmentEnum.

// SolutionNature classifies the solution set of a geometric query such as
// closest-point: how many equally good answers exist.
type SolutionNature int32

const (
	// UnknownSolutionNature means no solution strategy could be determined.
	UnknownSolutionNature SolutionNature = 0
	// UniqueSolution means there is exactly one solution.
	UniqueSolution SolutionNature = 1
	// DistinctlyManySolutions means finitely many, equally good solutions exist.
	DistinctlyManySolutions SolutionNature = 2
	// InfinitelyManySolutions means a continuum of equally good solutions exists
	// (e.g. the closest point to a circle queried from its center).
	InfinitelyManySolutions SolutionNature = 3
	// NoSolution means no solution exists.
	NoSolution SolutionNature = 4
)

var solutionNatureNames = map[SolutionNature]string{
	UnknownSolutionNature: "unknown", UniqueSolution: "unique",
	DistinctlyManySolutions: "distinctlyMany", InfinitelyManySolutions: "infinitelyMany",
	NoSolution: "none",
}

// String returns the classification's stable name.
func (n SolutionNature) String() string {
	if name, ok := solutionNatureNames[n]; ok {
		return name
	}
	return "solutionNature(?)"
}

// Containment is the result vocabulary for box/region containment queries.
type Containment int32

const (
	UnknownContainment Containment = 30977
	InsideContainment  Containment = 30978
	OnContainment      Containment = 30979
	OutsideContainment Containment = 30980
)

var containmentNames = map[Containment]string{
	UnknownContainment: "unknown", InsideContainment: "inside",
	OnContainment: "on", OutsideContainment: "outside",
}

// String returns the containment's stable name.
func (c Containment) String() string {
	if name, ok := containmentNames[c]; ok {
		return name
	}
	return "containment(?)"
}

// ParamAnomaly reports the irregularities of one parameter direction: whether
// the parameterization wraps (and with what period), hits singular points
// (a sphere pole, a cone apex, a polyline corner), or runs unbounded.
type ParamAnomaly struct {
	Periodic  bool    `json:"periodic,omitempty"`
	Period    float64 `json:"period,omitempty"`
	Singular  bool    `json:"singular,omitempty"`
	Unbounded bool    `json:"unbounded,omitempty"`
}

// ContinuityInfinite is the Continuity() order reported by analytic geometry,
// which is smooth to every order (C∞).
const ContinuityInfinite int = 1 << 30
