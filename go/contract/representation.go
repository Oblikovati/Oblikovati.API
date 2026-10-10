// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The scalar read surface of the assembly representation system (M12-F04, Oblikovati/
// Oblikovati#361/#367): the three representation families and the model states that select one
// of each. An in-proc consumer reads a representation's identity and active state directly;
// every capture/activate/override travels over api/wire (designReps.*/positionalReps.*/
// lodReps.*/modelStates.*). The host implementations live in /source (model/assembly).

// Representation is the identity common to every representation family: a session id, a
// display name, its kind, and whether it is the active representation of its family.
type Representation interface {
	// ID is the representation's session id, unique within its family's collection.
	ID() uint64
	// Name is the representation's display name (e.g. "Exploded").
	Name() string
	// Kind is the representation family.
	Kind() types.RepresentationKind
	// Active reports whether this is the active representation of its family.
	Active() bool
}

// DesignViewRepresentation overrides visibility, appearance, section planes, and the camera.
type DesignViewRepresentation interface {
	Representation
	// SectionPlanes returns the representation's section/clipping planes.
	SectionPlanes() []types.SectionPlane
}

// PositionalRepresentation overrides constraint/joint values (and per-occurrence flexibility);
// activating it re-solves the assembly to the alternate position.
type PositionalRepresentation interface {
	Representation
	// OverrideCount returns the number of constraint/joint value overrides it carries.
	OverrideCount() int
}

// LevelOfDetailRepresentation suppresses occurrences for performance on large assemblies.
type LevelOfDetailRepresentation interface {
	Representation
	// SuppressedCount returns the number of occurrences the representation suppresses.
	SuppressedCount() int
}

// DesignViewRepresentations / PositionalRepresentations / LevelOfDetailRepresentations are the
// per-family collections in creation order (host: assembly.Representations).
type (
	DesignViewRepresentations    = Enumerable[DesignViewRepresentation]
	PositionalRepresentations    = Enumerable[PositionalRepresentation]
	LevelOfDetailRepresentations = Enumerable[LevelOfDetailRepresentation]
)

// ModelState is a named tuple selecting one representation of each family — the single switch
// users flip. An empty family name means "leave that family's active representation unchanged".
type ModelState interface {
	// ID is the model state's session id.
	ID() uint64
	// Name is the model state's display name.
	Name() string
	// DesignViewName / PositionalName / LevelOfDetailName are the selected representations'
	// names (empty when the family is left unchanged).
	DesignViewName() string
	PositionalName() string
	LevelOfDetailName() string
	// Active reports whether this model state is the active one.
	Active() bool
}

// ModelStates is the assembly's model-state collection (host: assembly.ModelStateSet).
type ModelStates = Enumerable[ModelState]

// RepresentationsManager is the assembly's representation hub: the three family collections
// plus the model states (the reference API's RepresentationsManager).
type RepresentationsManager interface {
	// DesignViews / Positionals / LevelsOfDetail return the three family collections.
	DesignViews() DesignViewRepresentations
	Positionals() PositionalRepresentations
	LevelsOfDetail() LevelOfDetailRepresentations
	// ModelStates returns the model-state collection.
	ModelStates() ModelStates
}
