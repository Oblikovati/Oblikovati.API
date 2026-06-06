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
| `_Sidebar.md`, `_Footer.md` | wiki chrome |
| _generated_ | **API Docs** — produced from Go doc comments by `scripts/gen-api-docs.sh` |

GitHub maps a filename to a page title by replacing `-` with spaces, so
`First-Steps.md` becomes the page **First Steps** (and `[[First Steps]]` links to it).

## Publishing

A GitHub Actions workflow (`.github/workflows/wiki.yml`) republishes the wiki on every
push to `develop` (i.e. when a PR merges). It runs:

```sh
scripts/publish-wiki.sh   # assembles static pages + generated API Docs, pushes to the wiki
```

To preview the assembled output locally without pushing:

```sh
go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
scripts/gen-api-docs.sh build/wiki/API-Docs.md   # just the generated page
# or run the full assembly (set WIKI_REMOTE to a writable URL to actually push):
WIKI_REMOTE=/tmp/fake.git scripts/publish-wiki.sh
```

## Copyright note

This repository and its documentation must **never** name third-party CAD products. The
generator (`gen-api-docs.sh`) enforces this as a hard gate: it fails if the rendered docs
contain a forbidden product name, so a stray mention in a source doc comment is caught
before it can reach the public wiki.
