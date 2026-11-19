// SPDX-License-Identifier: Apache-2.0

package types

// PartAnnotationScaleMode identifies an automatic, fixed-ratio, or custom ratio scale.
// Example: `mode := types.PartAnnotationScaleAuto`.
type PartAnnotationScaleMode string

const (
	// PartAnnotationScaleAuto selects the host's automatic annotation scale.
	PartAnnotationScaleAuto PartAnnotationScaleMode = "auto"
	// PartAnnotationScaleFixed selects a preset ratio.
	PartAnnotationScaleFixed PartAnnotationScaleMode = "fixed"
	// PartAnnotationScaleCustom selects an authored ratio.
	PartAnnotationScaleCustom PartAnnotationScaleMode = "custom"
)

// PartAnnotationScale is Auto or a positive finite ratio of drawing to model units.
// Auto has numerator and denominator zero; Fixed and Custom require both positive.
// Equivalent ratios retain the mode and values chosen by the caller.
// Example: `scale := types.PartAnnotationScale{Mode: types.PartAnnotationScaleCustom, Numerator: 3, Denominator: 2}`.
type PartAnnotationScale struct {
	Mode        PartAnnotationScaleMode `json:"mode"`
	Numerator   uint32                  `json:"numerator"`
	Denominator uint32                  `json:"denominator"`
}

// PartDesignViewCamera is the complete viewport camera stored by a part view.
// Scale is the view height at Target. Near and Far are positive clipping distances.
// Eye, Target and Up form a usable finite camera frame. Projection selects the
// projection mode; a host declines projection modes it cannot represent.
// Example: `camera := types.PartDesignViewCamera{Scale: 20, Near: 1, Far: 100}`.
type PartDesignViewCamera struct {
	Eye        Point              `json:"eye"`
	Target     Point              `json:"target"`
	Up         Vector             `json:"up"`
	Scale      float64            `json:"scale"`
	Near       float64            `json:"near"`
	Far        float64            `json:"far"`
	Projection ProjectionTypeEnum `json:"projection"`
}
