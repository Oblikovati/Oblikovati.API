// SPDX-License-Identifier: Apache-2.0

package contract

import (
	"reflect"
	"testing"
)

// The fat contract interfaces are kept as unions of embedded capability families (audit I9):
// segregating the width lets a consumer depend only on the family it uses, while the union
// keeps every existing implementer and caller source-compatible. These tests are the guard
// that the union never drifts away from the sum of its families — in particular that a new
// method cannot be added to a fat name directly, bypassing a family (the drift the split
// closes). A union method that belongs to no family fails here unless it is a declared
// one-off exception.

// familyUnion pairs a fat union interface with the families that must cover it. exceptions are
// union method names deliberately declared on the union directly (a one-off not worth its own
// family); the split here has none, but the mechanism is kept so an intentional one-off is a
// recorded decision rather than a silent gap.
type familyUnion struct {
	name       string
	union      reflect.Type
	families   []reflect.Type
	exceptions map[string]bool
}

func iface[T any]() reflect.Type { return reflect.TypeOf((*T)(nil)).Elem() }

var contractFamilyUnions = []familyUnion{
	{name: "TransientGeometry", union: iface[TransientGeometry](), families: []reflect.Type{
		iface[TransientCurves3d](), iface[TransientCurves2d](), iface[TransientSurfaces](),
	}},
	{name: "Document", union: iface[Document](), families: []reflect.Type{
		iface[DocumentIdentity](), iface[DocumentDirtyState](), iface[DocumentLifecycle](),
	}},
	{name: "DisplaySettings", union: iface[DisplaySettings](), families: []reflect.Type{
		iface[BackgroundDisplaySettings](), iface[EdgeDisplaySettings](),
		iface[WindowDisplaySettings](), iface[GroundShadowSettings](),
	}},
	{name: "SurfaceEvaluator", union: iface[SurfaceEvaluator](), families: []reflect.Type{
		iface[SurfaceExtents](), iface[SurfaceDifferential](),
		iface[SurfaceProjection](), iface[SurfaceIso](),
	}},
	{name: "FileDescriptor", union: iface[FileDescriptor](), families: []reflect.Type{
		iface[FileReferenceIdentity](), iface[FileReferenceStatus](), iface[FileReferenceRepair](),
	}},
}

// methodNames returns the set of an interface type's method names.
func methodNames(t reflect.Type) map[string]bool {
	names := make(map[string]bool, t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		names[t.Method(i).Name] = true
	}
	return names
}

// TestUnionEqualsFamilies asserts each fat interface's method set is exactly the union of its
// families' method sets (plus any declared exception): no union method is uncovered by a family
// (the key guard), and no family declares a method the union lacks.
func TestUnionEqualsFamilies(t *testing.T) {
	for _, u := range contractFamilyUnions {
		union := methodNames(u.union)
		fams := map[string]bool{}
		for _, f := range u.families {
			for name := range methodNames(f) {
				fams[name] = true
			}
		}
		for name := range union {
			if !fams[name] && !u.exceptions[name] {
				t.Errorf("%s: union method %q is covered by no family — add it to a family or declare it an exception", u.name, name)
			}
		}
		for name := range fams {
			if !union[name] {
				t.Errorf("%s: family declares method %q absent from the union", u.name, name)
			}
		}
	}
}

// TestFamiliesAreDisjoint asserts the families of each union do not overlap (a method in two
// embedded families would merge only while their signatures stay identical — a latent hard
// compile error if one ever drifts, the shared-method-name embedding trap). A shared method
// belongs in one tiny shared interface embedded once, not duplicated across families.
func TestFamiliesAreDisjoint(t *testing.T) {
	for _, u := range contractFamilyUnions {
		seen := map[string]string{}
		for _, f := range u.families {
			for name := range methodNames(f) {
				if prev, dup := seen[name]; dup {
					t.Errorf("%s: method %q declared in both %s and %s — factor it into one shared interface",
						u.name, name, prev, f.Name())
				}
				seen[name] = f.Name()
			}
		}
	}
}
