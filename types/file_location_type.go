// SPDX-License-Identifier: Apache-2.0

package types

// FileLocationType classifies where a referenced file lives relative to the
// design project's search roots — the editable workspace, a shared workgroup,
// a read-only library, or somewhere unmanaged (M03-F07,
// Oblikovati/Oblikovati#608).
//
// The values are a frozen block matching the reference API's location-type
// enum; never renumber them.
type FileLocationType int32

const (
	// LocationWorkspace is the project's primary editable root.
	LocationWorkspace FileLocationType = 45057
	// LocationLocal is an auxiliary local workspace root.
	LocationLocal FileLocationType = 45058
	// LocationWorkgroup is a shared root of files being concurrently worked on.
	LocationWorkgroup FileLocationType = 45059
	// LocationLibrary is a shared, typically read-only catalog root.
	LocationLibrary FileLocationType = 45060
	// LocationUnknown is a path outside every configured root.
	LocationUnknown FileLocationType = 45061
	// LocationOwnerDirectory is the directory of the file holding the reference
	// (sibling-relative resolution).
	LocationOwnerDirectory FileLocationType = 45062
	// LocationCloud is a connected/synchronized network location.
	LocationCloud FileLocationType = 45063
)

// fileLocationTypeNames are the frozen wire spellings.
var fileLocationTypeNames = map[FileLocationType]string{
	LocationWorkspace:      "workspace",
	LocationLocal:          "local",
	LocationWorkgroup:      "workgroup",
	LocationLibrary:        "library",
	LocationUnknown:        "unknown",
	LocationOwnerDirectory: "ownerDirectory",
	LocationCloud:          "cloud",
}

// String returns the location type's wire spelling.
func (l FileLocationType) String() string { return enumName(fileLocationTypeNames, l, "enum(?)") }

// ParseFileLocationType resolves a wire spelling back to its FileLocationType.
func ParseFileLocationType(s string) (FileLocationType, bool) {
	return enumFromName(fileLocationTypeNames, s)
}
