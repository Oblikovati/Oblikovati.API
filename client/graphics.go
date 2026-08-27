// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// Graphics is the client/interaction-graphics group: it lets an add-in draw its own
// geometry — meshes, heatmaps, lines, point markers, and labels — into the 3D view.
// Set/List/Delete/SetVisible manage persistent (document-owned) graphics; the typed
// helpers (AddMesh, AddHeatmap, AddLines, AddPoints, AddLabel) wrap the common cases;
// Interaction reaches the transient command-preview lanes.
//
// Geometry travels as flat arrays: coords/normals are xyz triples, colors rgba quads in
// 0..1, indices are 0-based into the primitive's own vertices.
type Graphics struct{ c *Client }

// Graphics returns the client-graphics group.
func (c *Client) Graphics() Graphics { return Graphics{c} }

// Set submits or replaces the whole named graphics group (idempotent by ClientId).
//
// mcp:tool set_client_graphics
// mcp:summary Create or replace a named client-graphics overlay (declarative nodes/primitives drawn in the viewport, e.g. sim results).
func (g Graphics) Set(args wire.SetClientGraphicsArgs) (wire.SetClientGraphicsResult, error) {
	return call[wire.SetClientGraphicsResult](g.c, wire.MethodClientGraphicsSet, args)
}

// List enumerates the live graphics groups across all lanes.
//
// mcp:tool list_client_graphics
// mcp:summary List the add-in's client-graphics overlays (id, visibility).
func (g Graphics) List() (wire.ListClientGraphicsResult, error) {
	return call[wire.ListClientGraphicsResult](g.c, wire.MethodClientGraphicsList, nil)
}

// Delete removes the named graphics group.
//
// mcp:tool delete_client_graphics
// mcp:summary Delete a client-graphics overlay by id.
func (g Graphics) Delete(clientID string) error {
	return g.c.invoke(wire.MethodClientGraphicsDelete, wire.DeleteClientGraphicsArgs{ClientId: clientID}, nil)
}

// SetVisible toggles a group's visibility without resubmitting its geometry.
//
// mcp:tool set_client_graphics_visible
// mcp:summary Show or hide a client-graphics overlay by id.
func (g Graphics) SetVisible(clientID string, visible bool) error {
	return g.c.invoke(wire.MethodClientGraphicsSetVisible, wire.SetClientGraphicsVisibleArgs{ClientId: clientID, Visible: visible}, nil)
}

// AddMesh submits a single-color triangle mesh as a persistent group (coords xyz triples,
// indices 0-based triangle corners, color rgba 0..1).
func (g Graphics) AddMesh(clientID string, coords []float64, indices []int, color []float32) (wire.SetClientGraphicsResult, error) {
	return g.Set(oneShot(clientID, wire.GraphicsPrimitive{
		Kind: string(types.GraphicsTriangles), Coordinates: coords, Indices: indices,
		Color: color, ColorBinding: string(types.GraphicsColorOverall),
	}))
}

// AddHeatmap submits a triangle mesh colored per vertex from scalar values through a
// color mapper — the FEA/simulation-result case (coords xyz triples, indices triangle
// corners, scalars one per vertex).
func (g Graphics) AddHeatmap(clientID string, coords []float64, indices []int, scalars []float64, mapper wire.GraphicsColorMapper) (wire.SetClientGraphicsResult, error) {
	return g.Set(oneShot(clientID, wire.GraphicsPrimitive{
		Kind: string(types.GraphicsTriangles), Coordinates: coords, Indices: indices,
		Scalars: scalars, ColorMapper: &mapper, ColorBinding: string(types.GraphicsColorPerVertex),
	}))
}

// AddFloodPlot submits a scalar-mapped triangle mesh as an FEA-style flood plot drawn ON
// TOP of the model (depth test disabled) at the given opacity, so the field projects over
// the analyzed geometry instead of being occluded by it — the canonical way an FEA add-in
// shows a result field over its part/assembly. opacity is 0..1 (e.g. 0.6 lets the part
// edges read through the field); coords are xyz triples, scalars one per vertex.
//
// Unlike AddHeatmap (a persistent, depth-tested overlay for coloring a 3D surface that has
// its own normals), a flood plot is a flat data layer with no surface normals, so it is
// rendered unlit and over the geometry.
func (g Graphics) AddFloodPlot(clientID string, coords []float64, indices []int, scalars []float64, mapper wire.GraphicsColorMapper, opacity float32) (wire.SetClientGraphicsResult, error) {
	return g.Set(oneShot(clientID, wire.GraphicsPrimitive{
		Kind: string(types.GraphicsTriangles), Coordinates: coords, Indices: indices,
		Scalars: scalars, ColorMapper: &mapper, ColorBinding: string(types.GraphicsColorPerVertex),
		OnTop: true, Opacity: opacity,
	}))
}

// AddLines submits an indexed line list (coords xyz triples, indices segment-endpoint
// pairs) in one color as a persistent group.
func (g Graphics) AddLines(clientID string, coords []float64, indices []int, color []float32) (wire.SetClientGraphicsResult, error) {
	return g.Set(oneShot(clientID, wire.GraphicsPrimitive{
		Kind: string(types.GraphicsLines), Coordinates: coords, Indices: indices, Color: color,
	}))
}

// AddPoints submits point markers drawn with the given glyph style (coords xyz triples).
func (g Graphics) AddPoints(clientID string, coords []float64, style types.GraphicsPointStyle, color []float32) (wire.SetClientGraphicsResult, error) {
	return g.Set(oneShot(clientID, wire.GraphicsPrimitive{
		Kind: string(types.GraphicsPoints), Coordinates: coords, PointStyle: string(style), Color: color,
	}))
}

// AddLabel submits a single world-anchored text label (anchor is the xyz world point).
func (g Graphics) AddLabel(clientID, text string, anchor []float64, color []float32) (wire.SetClientGraphicsResult, error) {
	return g.Set(oneShot(clientID, wire.GraphicsPrimitive{
		Kind: string(types.GraphicsText), Text: text, Anchor: anchor, Color: color,
	}))
}

// AddBodyOverlay renders an existing B-rep body/face (by its persistent reference key) as a
// persistent overlay in an override color — the geometry stays host-side (no mesh shipped).
//
// mcp:tool add_body_overlay
// mcp:summary Overlay an existing B-rep body/face (by reference key) in an override color, rendered host-side.
func (g Graphics) AddBodyOverlay(clientID, bodyKey string, color []float32) (wire.SetClientGraphicsResult, error) {
	return g.Set(oneShot(clientID, wire.GraphicsPrimitive{
		Kind: string(types.GraphicsSurface), BodyKey: bodyKey, Color: color,
	}))
}

// AddImage submits a world-anchored image billboard (a textured quad) sized in model units.
//
// mcp:tool add_image_overlay
// mcp:summary Place an image billboard (textured quad) at a world point in the viewport.
func (g Graphics) AddImage(clientID, imagePath string, anchor []float64, width, height float64) (wire.SetClientGraphicsResult, error) {
	return g.Set(oneShot(clientID, wire.GraphicsPrimitive{
		Kind: string(types.GraphicsImage), ImagePath: imagePath, Anchor: anchor,
		ImageWidth: width, ImageHeight: height, Behavior: types.FrontFacingBehavior,
	}))
}

// AddStripMesh submits a triangle-strip mesh in one color (coords xyz triples in strip order)
// — the compact encoding for terrain/ribbon overlays.
//
// mcp:tool add_strip_mesh
// mcp:summary Draw a triangle-strip mesh overlay in one color (vertices in strip order).
func (g Graphics) AddStripMesh(clientID string, coords []float64, color []float32) (wire.SetClientGraphicsResult, error) {
	return g.Set(oneShot(clientID, wire.GraphicsPrimitive{
		Kind: string(types.GraphicsTriangleStrip), Coordinates: coords, Color: color,
		ColorBinding: string(types.GraphicsColorOverall),
	}))
}

// RegisterColorMapper stores a named, reusable color mapper that heatmap primitives reference
// by name (via the MapperName field) instead of carrying an inline legend.
//
// mcp:tool register_color_mapper
// mcp:summary Register a named, reusable heatmap color mapper shared across overlays.
func (g Graphics) RegisterColorMapper(name string, mapper wire.GraphicsColorMapper) error {
	return g.c.invoke(wire.MethodClientGraphicsRegisterMapper, wire.RegisterColorMapperArgs{Name: name, Mapper: mapper}, nil)
}

// ColorMappers lists the registered named color mappers.
//
// mcp:tool list_color_mappers
// mcp:summary List the registered named color mappers.
func (g Graphics) ColorMappers() (wire.ColorMappersResult, error) {
	return call[wire.ColorMappersResult](g.c, wire.MethodClientGraphicsListMappers, nil)
}

// SetNodeTransform moves one node within a group without resubmitting its geometry (transform
// is a 16-element row-major matrix; empty resets to identity).
//
// mcp:tool set_graphics_node_transform
// mcp:summary Move one client-graphics node (by id) without resending its mesh.
func (g Graphics) SetNodeTransform(clientID, nodeID string, transform []float64) error {
	return g.c.invoke(wire.MethodGraphicsNodeSetTransform, wire.SetNodeTransformArgs{ClientId: clientID, NodeId: nodeID, Transform: transform}, nil)
}

// SetNodeVisible toggles one node's visibility within a group without resubmitting geometry.
//
// mcp:tool set_graphics_node_visible
// mcp:summary Show or hide one client-graphics node (by id) without resending its mesh.
func (g Graphics) SetNodeVisible(clientID, nodeID string, visible bool) error {
	return g.c.invoke(wire.MethodGraphicsNodeSetVisible, wire.SetNodeVisibleArgs{ClientId: clientID, NodeId: nodeID, Visible: visible}, nil)
}

// SetNodeSelectable toggles whether one node's primitives participate in picking.
//
// mcp:tool set_graphics_node_selectable
// mcp:summary Toggle whether one client-graphics node (by id) is pickable.
func (g Graphics) SetNodeSelectable(clientID, nodeID string, selectable bool) error {
	return g.c.invoke(wire.MethodGraphicsNodeSetSelectable, wire.SetNodeSelectableArgs{ClientId: clientID, NodeId: nodeID, Selectable: selectable}, nil)
}

// oneShot wraps one primitive into a single-node persistent group request.
func oneShot(clientID string, p wire.GraphicsPrimitive) wire.SetClientGraphicsArgs {
	return wire.SetClientGraphicsArgs{
		ClientId: clientID,
		Lane:     string(types.GraphicsLanePersistent),
		Nodes:    []wire.GraphicsNode{{Primitives: []wire.GraphicsPrimitive{p}}},
	}
}

// Interaction returns the transient command-preview graphics group.
func (g Graphics) Interaction() InteractionGraphics { return InteractionGraphics(g) }

// InteractionGraphics is the command-scoped preview/overlay surface:
// Update replaces a transient lane's nodes (rubber-band/manipulator
// feedback on mouse move) and Clear removes them — both vanish when the command ends.
type InteractionGraphics struct{ c *Client }

// Update replaces the nodes of one interaction lane (overlay draws on top of the scene;
// preview draws depth-tested with it).
//
// mcp:tool update_interaction_graphics
// mcp:summary Set the transient interaction-graphics overlay (a short-lived preview/highlight pass, replaced each call).
func (i InteractionGraphics) Update(lane types.GraphicsLane, nodes []wire.GraphicsNode) error {
	return i.c.invoke(wire.MethodInteractionGraphicsUpdate, wire.UpdateInteractionGraphicsArgs{Lane: string(lane), Nodes: nodes}, nil)
}

// Clear removes all transient interaction graphics (both lanes).
//
// mcp:tool clear_interaction_graphics
// mcp:summary Clear the transient interaction-graphics overlay.
func (i InteractionGraphics) Clear() error {
	return i.c.invoke(wire.MethodInteractionGraphicsClear, nil, nil)
}
