// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// Shell and wire topology contracts (M07-F06, Oblikovati/Oblikovati#629), and
// the body-level query surface (M07-F07, #630). In-proc scalar interfaces;
// the wire DTO forms live in api/wire.

// FaceShell is one connected face group of a body: the outer skin or an
// inner cavity (void) skin.
type FaceShell interface {
	// IsClosed reports whether the shell bounds a region; IsVoid whether it
	// is an inner cavity skin (its material-outward normals face inward).
	IsClosed() bool
	IsVoid() bool
	// Volume is the magnitude of the bounded region's volume (cm³).
	Volume() float64
	FaceCount() int
	EdgeCount() int
	// IsPointInside classifies a point against the shell's bounded region.
	IsPointInside(x, y, z float64) types.Containment
	// ReferenceKey returns the persistent reference key (M03 scheme);
	// TransientKey the session-stable id.
	ReferenceKey() []byte
	TransientKey() uint64
}

// FaceShells enumerates a body's shells.
type FaceShells interface {
	Count() int
	Item(index int) FaceShell
}

// Wire is an ordered, face-less edge chain on a body — the section/profile
// currency of ruled surfaces and silhouettes.
type Wire interface {
	IsClosed() bool
	IsPlanar() bool
	EdgeCount() int
	ReferenceKey() []byte
	TransientKey() uint64
}

// Wires enumerates a body's wires.
type Wires interface {
	Count() int
	Item(index int) Wire
}

// BodyQueries is the reference SurfaceBody query surface: point/ray location,
// containment, convexity, validity and transient-key binding (M07-F07).
type BodyQueries interface {
	// IsPointInside classifies a point against the body's material.
	IsPointInside(x, y, z float64) types.Containment
	// ConvexEdgeCount / ConcaveEdgeCount report the dihedral classification
	// totals (the EdgeCollection cardinalities).
	ConvexEdgeCount() int
	ConcaveEdgeCount() int
	// IsEntityValid checks the whole body at the given level (1 = topology,
	// 2 = + self-intersection).
	IsEntityValid(checkLevel int) bool
	// BindTransientKey reports the kind ("vertex"/"edge"/"face"/"shell"/
	// "wire") and persistent reference key behind a session transient key.
	BindTransientKey(key uint64) (kind string, referenceKey []byte, ok bool)
}
