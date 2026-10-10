// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// The assembly contact & interference surface (M12-F05, Oblikovati/Oblikovati#362/#368):
// manage contact sets (occurrences that resist interpenetration when dragged), toggle the
// contact solver, and run a static interference analysis that reports the overlapping volumes
// between occurrences. Members and occurrences are addressed by occurrence id.

// ContactSetInfo is one contact set: its id, name, and member occurrence ids.
type ContactSetInfo struct {
	ID      uint64   `json:"id"`
	Name    string   `json:"name"`
	Members []uint64 `json:"members,omitempty"`
}

// ContactSolverInfo is the contact solver's state: whether it is enabled and how many sets it
// governs.
type ContactSolverInfo struct {
	Enabled  bool `json:"enabled"`
	SetCount int  `json:"setCount,omitempty"`
}

// InterferenceResultInfo is one overlapping pair: the two occurrence ids, the overlap volume
// (cm³), and a representative point inside the overlap.
type InterferenceResultInfo struct {
	OccurrenceA uint64      `json:"occurrenceA"`
	OccurrenceB uint64      `json:"occurrenceB"`
	Volume      float64     `json:"volume"`
	Center      types.Point `json:"center"`
}

// Result envelopes.
type (
	ContactSetsResult struct {
		ContactSets []ContactSetInfo `json:"contactSets"`
	}
	ContactSetResult struct {
		ContactSet ContactSetInfo `json:"contactSet"`
	}
	ContactSolverResult struct {
		Solver ContactSolverInfo `json:"solver"`
	}
	InterferenceResultsResult struct {
		Results     []InterferenceResultInfo `json:"results"`
		TotalVolume float64                  `json:"totalVolume,omitempty"`
	}
)

// CreateContactSetArgs creates a new contact set named Name.
type CreateContactSetArgs struct {
	Name string `json:"name"`
}

// ContactSetRef names a contact set by id (the request of delete).
type ContactSetRef struct {
	ID uint64 `json:"id"`
}

// ContactMemberArgs adds or removes an occurrence to/from a contact set.
type ContactMemberArgs struct {
	Set        uint64 `json:"set"`
	Occurrence uint64 `json:"occurrence"`
}

// ContactSolverEnableArgs enables or disables the contact solver.
type ContactSolverEnableArgs struct {
	Enabled bool `json:"enabled"`
}

// AnalyzeInterferenceArgs requests a static interference analysis. Occurrences optionally
// restricts the analysis to a subset (by id); empty analyses every pair in the active assembly.
type AnalyzeInterferenceArgs struct {
	Occurrences []uint64 `json:"occurrences,omitempty"`
}
