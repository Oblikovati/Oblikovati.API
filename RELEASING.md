# Releasing `oblikovati.org/api`

This module follows [Semantic Versioning 2.0.0](https://semver.org). The single
source of truth for the version is the `Version` constant in
[`go/version.go`](go/version.go). The release tags are `v<Version>` and
`go/v<Version>` on the same commit. The second tag selects the Go submodule.

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

Before versioning, the workflow checks every Go wire method for a client-side
`mcp:tool` and `mcp:summary`, or an intentional `mcp:skip` with a reason. The
candidate CI gate runs the same test before publication.

1. [`go/scripts/nextver`](go/scripts/nextver) reads the conventional-commit messages since the
   last `vX.Y.Z` tag, picks the strongest scope, and applies the policy above to compute
   the next version.
2. It rewrites `Version` in `go/version.go` and rolls `CHANGELOG.md` (moves `## [Unreleased]`
   over a new `## [x.y.z] - DATE` section and refreshes the compare links).
3. `make generate` updates the Rust package version and C ABI handshake constants.
   The workflow commits the synchronized changes as a unique release candidate.
   `cargo package` checks that exact commit, and the workflow checks the public Go
   import metadata selects `go/`.
4. The workflow pushes a temporary candidate branch and dispatches `CI` for it.
   It accepts only the returned run on the candidate SHA, with the Linux gate and
   both Windows and macOS jobs completed successfully. Missing, skipped, failed,
   or mismatched jobs prevent publication. `workflow_dispatch` can start a run
   with the workflow token even though ordinary token-generated pushes cannot.
5. Only after CI succeeds does the workflow create annotated `vX.Y.Z` and
   `go/vX.Y.Z` tags. It atomically pushes the tested commit and both tags, with a
   lease that rejects a changed `develop`. The GitHub release contains the notes
   and verified Rust `.crate` archive. The workflow deletes its candidate branch
   only if that branch still points to its candidate commit.

If release creation succeeds but its archive upload fails, the workflow retries
and checks the existing release's author, notes, tag targets, and archive bytes.
It never replaces an existing tag or asset. If retries fail after the atomic
push, investigate and publish the verified archive for the existing tag after
checking its SHA and CI result. A new release must use a new version.

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
change deserves a MINOR, so prefix it with `feat:`. The logic lives in `go/scripts/nextver`
and is covered by `go test ./scripts/nextver`.

When an already published commit has a nonconventional subject that understates
an additive API change, add a correction commit that explains the affected
changes and includes `Release-Scope: feature` on a separate line in its body.
The trailer sets a minimum feature scope for the release. A breaking-change
marker still takes precedence. This preserves the original commit identities.

## Consuming a release

Pin a release with `require oblikovati.org/api vX.Y.Z`. Local development (the
`Oblikovati` application and the bridge add-in) resolves this module through a
`go.work` workspace over the sibling checkout, so the `require` version is
overridden there and tracks the working tree, not a published tag.

## Language directories

A release has two tags on the same commit: `vX.Y.Z` for the shared contract and
`go/vX.Y.Z` for the Go module in the `go/` subdirectory. `make generate` copies
the Go version into `rust/Cargo.toml` and the Rust handshake constants.

Before the first release from this layout, update the import metadata served
at `https://oblikovati.org/api?go-get=1` to:

```html
<meta name="go-import" content="oblikovati.org/api git https://github.com/Oblikovati/Oblikovati.API go">
```

The fourth field selects the module subdirectory (Go 1.25 or later). See the
[Go module reference](https://go.dev/ref/mod#vcs-find). Update `go-source` links
to `tree/develop/go{/dir}` and `blob/develop/go{/dir}/{file}#L{line}` as well.
The website source is `website-oblikovati.org/public/api/index.html`.
Existing local workspaces must point to `Oblikovati.API/go`.

The release checker identifies itself as `Oblikovati.API-release-check/1.0`.
It requests the canonical HTTPS endpoint, accepts only HTTP 200 at `/api` or
`/api/` with the `go-get=1` query, and requires the exact import metadata above.
A forbidden response, or a final response at another origin or over HTTP, remains
a failure.

The release workflow refuses to tag or push if this metadata is not live.
The directory move is an alpha API migration. Older source tags use the former
root layout; keep old checkouts on that layout when reproducing old releases.

Rust consumers can use `oblikovati-api = { git =
"https://github.com/Oblikovati/Oblikovati.API", tag = "vX.Y.Z" }` or the local
`rust/` path. Publishing to crates.io is a separate distribution step; this
workflow does not claim a registry name or require a registry token.

## Workflow dependency provenance

Remote workflow actions use full source commit SHAs. The comments retain their
selected major versions; Rust steps also set `toolchain: 1.95.0` explicitly so
pinning the action implementation does not change the compiler version.
`make gate` checks these rules through `scripts/workflow_dependencies.py`.

The documentation tool remains gomarkdoc v1.1.0. Workflows install its official
source commit `e62c5abf78916697dbc4b25f590deb28f26f5184`, which Go resolves to
v1.1.0. Go version suffixes select exact module versions, rather than version
ranges; normal public-module checksum verification still applies. For this
revision, `go mod download -json` reports:

- Module zip: `h1:xtl7mESKQWVuGiFdd1AO3dFA6OenWG86bZu97IqBNPE=`.
- Module definition: `h1:HI3w0Zv8H03ecak/IqVAcPFTuPt7sn7Top6xbgCs1Qk=`.

The full source revision also makes the tool pin explicit to workflow scanners.
No tool or language version is upgraded by this pinning change.
