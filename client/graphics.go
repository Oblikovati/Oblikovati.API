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
	var r wire.SetClientGraphicsResult
	return r, g.c.call(wire.MethodClientGraphicsSet, args, &r)
}

// List enumerates the live graphics groups across all lanes.
//
// mcp:tool list_client_graphics
// mcp:summary List the add-in's client-graphics overlays (id, visibility).
func (g Graphics) List() (wire.ListClientGraphicsResult, error) {
	var r wire.ListClientGraphicsResult
	return r, g.c.call(wire.MethodClientGraphicsList, nil, &r)
}

// Delete removes the named graphics group.
//
// mcp:tool delete_client_graphics
// mcp:summary Delete a client-graphics overlay by id.
func (g Graphics) Delete(clientID string) error {
	return g.c.call(wire.MethodClientGraphicsDelete, wire.DeleteClientGraphicsArgs{ClientId: clientID}, nil)
}

// SetVisible toggles a group's visibility without resubmitting its geometry.
//
// mcp:tool set_client_graphics_visible
// mcp:summary Show or hide a client-graphics overlay by id.
func (g Graphics) SetVisible(clientID string, visible bool) error {
	return g.c.call(wire.MethodClientGraphicsSetVisible, wire.SetClientGraphicsVisibleArgs{ClientId: clientID, Visible: visible}, nil)
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
	return i.c.call(wire.MethodInteractionGraphicsUpdate, wire.UpdateInteractionGraphicsArgs{Lane: string(lane), Nodes: nodes}, nil)
}

// Clear removes all transient interaction graphics (both lanes).
//
// mcp:tool clear_interaction_graphics
// mcp:summary Clear the transient interaction-graphics overlay.
func (i InteractionGraphics) Clear() error {
	return i.c.call(wire.MethodInteractionGraphicsClear, nil, nil)
}
