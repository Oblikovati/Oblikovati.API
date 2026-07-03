// SPDX-License-Identifier: Apache-2.0

package types

// ObjectKind identifies the kind of document object a metadata-mutation event concerns
// (Oblikovati/Oblikovati#1644): a body, sketch, feature, occurrence, or the document itself. It is
// the discriminator carried by object.renamed and property.changed so ONE generic event serves
// every rename / property change instead of a bespoke event per mutation — an add-in matches on
// Kind to route the event to the right mirror of the document structure.
type ObjectKind string

const (
	// ObjectKindBody is a solid/surface body in a part.
	ObjectKindBody ObjectKind = "body"
	// ObjectKindSketch is a 2D or 3D sketch.
	ObjectKindSketch ObjectKind = "sketch"
	// ObjectKindFeature is a history feature (extrude, fillet, work plane, …).
	ObjectKindFeature ObjectKind = "feature"
	// ObjectKindOccurrence is an assembly component occurrence.
	ObjectKindOccurrence ObjectKind = "occurrence"
	// ObjectKindDocument is the document itself.
	ObjectKindDocument ObjectKind = "document"
)
