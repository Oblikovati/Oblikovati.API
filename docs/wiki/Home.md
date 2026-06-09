# Oblikovati Add-in Developer Wiki

Welcome. This wiki is for developers building **add-ins** for Oblikovati — the
parametric, feature-based CAD application. Add-ins extend the host with new ribbon
commands and drive the live model (sketches, parameters, features, view, materials)
through a stable, permissively-licensed contract.

> **This wiki is generated.** It is rebuilt automatically from
> [`Oblikovati.API`](https://github.com/Oblikovati/Oblikovati.API) every time a pull
> request merges to `develop`. Edit the sources under `docs/wiki/` in that repo — manual
> edits made here will be overwritten on the next merge. The **API Docs** page is produced
> directly from the contract's Go source and doc comments.

## What an add-in is

An add-in is a shared library (`.so` / `.dll` / `.dylib`) that the host loads in-process at
runtime. It links **only** the Apache-2.0 contract module (`oblikovati.org/api`) and talks to the
host across a small C ABI. Because the boundary is a byte protocol (JSON over a single host
callback), an add-in can be **closed-source** while the contract it depends on stays open.

```
┌────────────────────┐        C ABI (JSON method calls)        ┌──────────────────────┐
│   Your add-in      │  ────────────────────────────────────▶  │   Oblikovati host    │
│  (.so / .dll)      │           oblikovati.org/api/client          │  (kernel + UI head)  │
│  links only        │  ◀────────────────────────────────────  │  owns the live model │
│  oblikovati.org/api    │            replies / events              │                      │
└────────────────────┘                                          └──────────────────────┘
```

## Start here

| Page | What it covers |
|---|---|
| **[[Oblikovati Architecture]]** | How the application and its geometry kernel are put together. |
| **[[First Steps]]** | Build an add-in that adds a ribbon button which sketches a rectangle and extrudes a cube. |
| **[[Oblikovati API Architecture]]** | The four contract packages, the two consumption paths, and how to add to the surface. |
| **[[Testing Automation]]** | How to unit- and integration-test an add-in while respecting the license boundary. |
| **[[API Docs]]** | Auto-generated reference for every exported type, function, and method in the contract. |

## License at a glance

- The **contract** you build against (`oblikovati.org/api`) is **Apache-2.0**.
- The **host application** is GPL-2.0 and is never linked by an add-in.
- Your add-in may carry **any license you choose** — see [[Testing Automation]] for how the
  build keeps that boundary honest.
