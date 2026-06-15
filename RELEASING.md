# Releasing `oblikovati.org/api`

This module follows [Semantic Versioning 2.0.0](https://semver.org). The single
source of truth for the version is the `Version` constant in
[`version.go`](version.go); the git tag for a release is `v<Version>` (the Go
module tag convention — the `v` prefix is not part of the semver string).

## What is the public surface

Every exported identifier in `types`, `contract`, `wire`, and `client`, and the
JSON shape of every `wire` DTO and method-name constant (the on-the-wire
protocol). A change is **breaking** if it removes or renames any of these, or
alters a wire DTO's JSON shape such that an existing peer would misread it.

## Versioning policy

- **0.y.z (current).** Initial development. The public surface MAY change in any
  minor release. Bump the **MINOR** for a breaking or additive change, the
  **PATCH** for a fix that keeps the surface. No backward-compatibility guarantee
  yet.
- **1.0.0 and later.** The API is declared stable. A backward-incompatible change
  bumps the **MAJOR** — and, per Go's rules, moves the module path to `/v2`,
  `/v3`, … . A backward-compatible addition bumps the **MINOR**; a fix bumps the
  **PATCH**.

## Cutting a release (automatic)

Releases are **not** hand-cut. On every merge to `develop` the release workflow
([`.github/workflows/release.yml`](.github/workflows/release.yml)) derives the next
version from the **scope of the commits merged since the last tag** and does the rest:

1. [`scripts/nextver`](scripts/nextver) reads the conventional-commit messages since the
   last `vX.Y.Z` tag, picks the strongest scope, and applies the policy above to compute
   the next version.
2. It rewrites `Version` in `version.go` and rolls `CHANGELOG.md` (moves `## [Unreleased]`
   over a new `## [x.y.z] - DATE` section and refreshes the compare links).
3. The workflow commits that back to `develop` as `chore(release): vX.Y.Z [skip ci]`
   (the `[skip ci]` stops the bump commit from triggering another release), tags
   `vx.y.z`, and publishes the GitHub release with notes from the changelog section.

A merge whose commits are all `docs`/`chore`/`ci`/`test`/`refactor` releases nothing.

### Telling the workflow the scope: conventional commits

The bump is only as correct as the commit subjects, so use
[Conventional Commits](https://www.conventionalcommits.org):

| Commit subject | Scope | 0.x bump | ≥1.0 bump |
| --- | --- | --- | --- |
| `feat: …` / `feat(scope): …` | additive | MINOR | MINOR |
| `fix: …` / `perf:` / `revert:` | fix | PATCH | PATCH |
| `feat!: …` or a `BREAKING CHANGE:` footer | breaking | MINOR | MAJOR |
| `docs:`/`chore:`/`ci:`/`test:`/`style:`/`build:`/`refactor:` | none | — | — |

Any **unrecognized** subject (no conventional type, e.g. `wire: add …`) is treated as a
conservative **PATCH** so the version never silently stalls — but an additive surface
change deserves a MINOR, so prefix it with `feat:`. The logic lives in `scripts/nextver`
and is covered by `go test ./scripts/nextver`.

## Consuming a release

Pin a release with `require oblikovati.org/api vX.Y.Z`. Local development (the
`Oblikovati` application and the bridge add-in) resolves this module through a
`go.work` workspace over the sibling checkout, so the `require` version is
overridden there and tracks the working tree, not a published tag.
