# Oblikovati Architecture

A high-level tour of how Oblikovati is built, aimed at add-in developers who want to
understand what they are driving. You do **not** need to know the host internals to write an
add-in — but knowing the shape of the system makes the API obvious.

## The layers

Oblikovati is a parametric, **feature-based** modeler. From the bottom up:

```
            ┌──────────────────────────────────────────────────────┐
  UI head   │  ribbon, browser, viewport, dialogs, interaction      │
            ├──────────────────────────────────────────────────────┤
  Model     │  documents · parameters · sketches · features         │
            │  (the recipe: an ordered, re-computable feature tree)  │
            ├──────────────────────────────────────────────────────┤
  Kernel    │  geom (curves/surfaces) · topo (B-rep) · ops          │
            │  (boolean, fillet, shell, draft, tessellate, mass)     │
            └──────────────────────────────────────────────────────┘
```

### Geometry kernel

The kernel is the geometry engine. It has three cooperating parts:

- **`geom`** — analytic geometry: transient primitives and curves/surfaces (lines, arcs,
  circles, ellipses, planes, cylinders, cones, spheres, tori, helices, B-splines) plus the
  transforms that move them. These are the building blocks a feature turns into solids.
- **`topo`** — the **boundary representation** (B-rep): the connected vertices, edges, loops,
  faces, shells and bodies that make a solid watertight and queryable. Every face/edge carries
  a **lineage** so a reference to it survives edits (persistent topological naming).
- **`ops`** — operations over bodies: booleans (union / difference / intersection), fillet,
  chamfer, shell, draft, plus *tessellation* (turning exact geometry into the triangle mesh the
  viewport draws) and *mass properties* (volume, area, centroid via the divergence theorem).

Booleans run on an **exact planar B-rep** path for planar-faceted solids and fall back to a
welded-mesh CSG path for curved cases. The kernel validates results (manifold, closed,
orientation) so a bad operation is caught rather than silently shipped.

### Model layer (the recipe)

A document is not a static mesh — it is a **recipe**: an ordered tree of *features* that
recompute from *parameters* and *sketches*.

- **Parameters** are named, unit-bearing expressions (`length = 40 mm`, `width = length/2`).
  They form a dependency graph; changing one recomputes everything downstream.
- **Sketches** are 2D (on a plane) or 3D profiles built from constrained geometry. A sketch is
  solved by a constraint engine and ideally **fully constrained** (zero degrees of freedom)
  before a feature consumes it.
- **Features** are the modeling operations — extrude, revolve, sweep, loft, coil, hole, fillet,
  chamfer, shell, draft, rib, patterns, mirror, and more. Each takes sketches/parameters and
  earlier geometry as input and produces or modifies the body. Editing a feature's parameters or
  its geometry references and recomputing is the heart of parametric modeling.

This recipe is what an add-in manipulates: you add parameters, author sketches, and apply
features, then read back the resulting geometry and physical properties.

### UI head

The head renders the model (Vulkan viewport), presents the **ribbon** (tabs → panels →
buttons), the model **browser** (the feature tree), and the interactive tools. An add-in's
ribbon button lives here; clicking it routes back to the add-in (see [[First Steps]]).

## How an add-in fits in

The host owns the model and runs it on a single session goroutine. An add-in is a separate
shared library with its **own runtime**; it never touches host memory directly. Instead it
sends JSON *method calls* (`documents.create`, `sketch.rectangle`, `features.add`, …) across the
C ABI; the host executes each against the live model and returns JSON. The full method
vocabulary and the typed client that wraps it are described in
[[Oblikovati API Architecture]].

The practical consequences for you:

- **Everything is a method call.** There is no shared object graph; you drive the model the same
  way the UI does, through named operations.
- **Units are explicit.** Lengths are expressions like `"40 mm"` or `"5 cm"`; the host's database
  unit is centimetres, so volumes come back in cm³.
- **References are strings.** When a feature needs an earlier face/edge/plane, you pass a
  reference key that the host resolves and tracks across recompute.
