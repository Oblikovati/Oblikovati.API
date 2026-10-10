# Wiki source (developer documentation)

These Markdown files are the **source of truth** for the Oblikovati add-in developer
[GitHub Wiki](https://github.com/Oblikovati/Oblikovati.API/wiki). The wiki itself is a
downstream mirror — **edit pages here, not in the wiki UI** (manual wiki edits are
overwritten on the next publish).

## Pages

| File | Wiki page |
|---|---|
| `Home.md` | Home (landing + navigation) |
| `Oblikovati-Architecture.md` | Oblikovati Architecture |
| `First-Steps.md` | First Steps |
| `Oblikovati-API-Architecture.md` | Oblikovati API Architecture |
| `Testing-Automation.md` | Testing Automation |
| `Typed-Operations.md` | Typed Operations (compiled Go and Rust examples) |
| `C-ABI-Add-ins.md` | C ABI Add-ins (compiled shared libraries in both languages) |
| `_Sidebar.md`, `_Footer.md` | wiki chrome |
| _generated_ | **API Docs** — links to both language references |
| _generated_ | **Go API Reference** — Go doc comments and declarations |
| _generated_ | **Rust API Reference** — Rust declarations and client method index |

GitHub maps a filename to a page title by replacing `-` with spaces, so
`First-Steps.md` becomes the page **First Steps** (and `[[First Steps]]` links to it).

## Publishing

A GitHub Actions workflow (`.github/workflows/wiki.yml`) republishes the wiki on every
push to `develop` (i.e. when a PR merges). It runs:

```sh
bash scripts/publish-wiki.sh
```

To preview the assembled output locally without pushing:

```sh
go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@v1.1.0
make generate
make gate
make docs
```

The assembled pages are in `build/wiki/`. Rustdoc is in `rust/target/doc/`.
`make docs` does not publish. Edit example source files under `go/examples/` or
`rust/examples/`, then run `make generate`. CI rejects stale copied code blocks.

## Copyright note

This repository and its documentation must **never** name third-party CAD products. The
generator (`gen-api-docs.sh`) enforces this as a hard gate: it fails if the rendered docs
contain a forbidden product name, so a stray mention in a source doc comment is caught
before it can reach the public wiki.
