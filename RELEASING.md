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

## Cutting a release

1. Determine the new version from the changes since the last tag, per the policy
   above.
2. Bump `Version` in `version.go`.
3. In `CHANGELOG.md`, move the `## [Unreleased]` entries under a new
   `## [x.y.z] - YYYY-MM-DD` heading and refresh the compare links at the bottom.
4. Open a PR; merge to `develop` once CI is green.
5. On the merge, the release workflow
   ([`.github/workflows/release.yml`](.github/workflows/release.yml)) reads
   `version.go`, creates the `vx.y.z` tag, and publishes a GitHub release. It is
   idempotent: a push that does not change `Version` is a no-op because the tag
   already exists, so unrelated merges to `develop` never re-tag.

## Consuming a release

Pin a release with `require oblikovati.org/api vX.Y.Z`. Local development (the
`Oblikovati` application and the bridge add-in) resolves this module through a
`go.work` workspace over the sibling checkout, so the `require` version is
overridden there and tracks the working tree, not a published tag.
