// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// The assembly representation surface (M12-F04, Oblikovati/Oblikovati#361/#367): capture,
// activate, and edit the three representation families — design-view (visibility/appearance/
// section/camera), positional (constraint/joint value overrides) and level-of-detail
// (occurrence suppression) — plus model states selecting one of each. Representations are
// named override layers over an immutable base; capture snapshots the current scene into a
// new representation, activate applies a representation's overrides, and the set* operations
// edit individual overrides. Overrides are addressed by occurrence/relationship id; the host
// stores them by stable occurrence path so a representation survives recompute and reload.

// RepresentationInfo is the identity row common to every representation family.
type RepresentationInfo struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Kind   string `json:"kind"`
	Active bool   `json:"active,omitempty"`
}

// DesignViewInfo is one design-view representation: its identity, the count of hidden
// occurrences and appearance overrides, its section planes, and the captured camera (if any).
type DesignViewInfo struct {
	RepresentationInfo
	HiddenCount     int                  `json:"hiddenCount,omitempty"`
	AppearanceCount int                  `json:"appearanceCount,omitempty"`
	SectionPlanes   []types.SectionPlane `json:"sectionPlanes,omitempty"`
	Camera          *CameraView          `json:"camera,omitempty"`
}

// PositionalInfo is one positional representation: its identity and the count of constraint/
// joint value overrides it carries.
type PositionalInfo struct {
	RepresentationInfo
	OverrideCount int `json:"overrideCount,omitempty"`
}

// LODInfo is one level-of-detail representation: its identity and the count of suppressed
// occurrences.
type LODInfo struct {
	RepresentationInfo
	SuppressedCount int `json:"suppressedCount,omitempty"`
}

// ModelStateInfo is one model state: its identity and the names of the representation it
// selects in each family (empty ⇒ that family is left unchanged on activate).
type ModelStateInfo struct {
	ID            uint64 `json:"id"`
	Name          string `json:"name"`
	DesignView    string `json:"designView,omitempty"`
	Positional    string `json:"positional,omitempty"`
	LevelOfDetail string `json:"levelOfDetail,omitempty"`
	Active        bool   `json:"active,omitempty"`
}

// Result envelopes.
type (
	DesignViewsResult struct {
		Representations []DesignViewInfo `json:"representations"`
	}
	DesignViewResult struct {
		Representation DesignViewInfo `json:"representation"`
	}
	PositionalsResult struct {
		Representations []PositionalInfo `json:"representations"`
	}
	PositionalResult struct {
		Representation PositionalInfo `json:"representation"`
	}
	LODsResult struct {
		Representations []LODInfo `json:"representations"`
	}
	LODResult struct {
		Representation LODInfo `json:"representation"`
	}
	ModelStatesResult struct {
		ModelStates []ModelStateInfo `json:"modelStates"`
	}
	ModelStateResult struct {
		ModelState ModelStateInfo `json:"modelState"`
	}
)

// CaptureRepArgs captures the current scene state into a new representation of its family.
type CaptureRepArgs struct {
	Name string `json:"name"`
}

// RepRef names a representation (or model state) by id — the request of activate/delete.
type RepRef struct {
	ID uint64 `json:"id"`
}

// SetVisibilityArgs hides or shows an occurrence within a design-view representation.
type SetVisibilityArgs struct {
	Rep        uint64 `json:"rep"`
	Occurrence uint64 `json:"occurrence"`
	Visible    bool   `json:"visible"`
}

// SetAppearanceArgs overrides an occurrence's appearance within a design-view representation
// (an empty AppearanceID clears the override).
type SetAppearanceArgs struct {
	Rep          uint64 `json:"rep"`
	Occurrence   uint64 `json:"occurrence"`
	AppearanceID string `json:"appearanceId,omitempty"`
}

// AddSectionArgs adds a section/clipping plane to a design-view representation.
type AddSectionArgs struct {
	Rep   uint64             `json:"rep"`
	Plane types.SectionPlane `json:"plane"`
}

// SetPositionalOverrideArgs overrides a constraint's or joint's value within a positional
// representation (IsJoint selects which relationship id space the id is in).
type SetPositionalOverrideArgs struct {
	Rep          uint64  `json:"rep"`
	Relationship uint64  `json:"relationship"`
	IsJoint      bool    `json:"isJoint,omitempty"`
	Value        float64 `json:"value"`
}

// SetFlexibleArgs sets a subassembly occurrence's flexibility within a positional
// representation.
type SetFlexibleArgs struct {
	Rep        uint64 `json:"rep"`
	Occurrence uint64 `json:"occurrence"`
	Flexible   bool   `json:"flexible"`
}

// SetSuppressedArgs suppresses or restores an occurrence within a level-of-detail
// representation.
type SetSuppressedArgs struct {
	Rep        uint64 `json:"rep"`
	Occurrence uint64 `json:"occurrence"`
	Suppressed bool   `json:"suppressed"`
}

// CreateModelStateArgs creates a model state selecting one representation of each family by
// name (an empty family name leaves that family unchanged on activate).
type CreateModelStateArgs struct {
	Name          string `json:"name"`
	DesignView    string `json:"designView,omitempty"`
	Positional    string `json:"positional,omitempty"`
	LevelOfDetail string `json:"levelOfDetail,omitempty"`
}

// RepresentationEventPayload is the relayed body of a representation/model-state event: the
// kind ("designView"/"positional"/"levelOfDetail"/"modelState"), the id, and the name.
type RepresentationEventPayload struct {
	Kind string `json:"kind"`
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
