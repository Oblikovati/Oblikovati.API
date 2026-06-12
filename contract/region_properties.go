// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// RegionProperties is the scalar view of a closed profile region's section
// properties (M06-F08, Oblikovati/Oblikovati#623). All values are in database
// units (cm-based) on the profile's plane; the inertia values are about the
// region's centroid. The host's model/sketch region calculator satisfies this
// via a compile-time assertion.
type RegionProperties interface {
	// Accuracy is the computational accuracy the values were computed at.
	Accuracy() types.Accuracy
	// Area is the enclosed area in cm² (holes subtracted).
	Area() float64
	// Perimeter is the total boundary length in cm (holes' rims included).
	Perimeter() float64
	// Centroid is the area centroid [x, y] in sketch-plane cm.
	Centroid() (x, y float64)
	// MomentsOfInertia are the centroidal second moments Ixx, Iyy and the
	// product Ixy, in cm⁴.
	MomentsOfInertia() (ixx, iyy, ixy float64)
	// PrincipalMoments are the principal second moments I1 ≥ I2 in cm⁴.
	PrincipalMoments() (i1, i2 float64)
	// RotationAngle is the CCW angle in radians from the sketch X axis to the
	// first principal axis.
	RotationAngle() float64
}
