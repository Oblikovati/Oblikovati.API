# API duplication check for issue 301

`make gate` runs `make duplication`. The check invokes the exact npm package
`jscpd@5.4.1` and requires Node.js and npm. It scans `go`, `rust/src`,
`rust/tests`, `rust/examples`, `scripts`, `include`, `tests`, and `.github`,
plus the root build and configuration files and `rust/Cargo.toml`, including
generated source, tests, and ignored paths. It uses mild detection
with a five-line and 50-token minimum, so shorter fragments do not contribute
to clone counts. Supported formats include Go, Rust, Python, shell, C, headers,
YAML, and TOML. No source directories or patterns are excluded. The check refuses
links and files larger than its explicit 1 GB scanner limit, rather than
silently skipping them.

The scanner uses an isolated empty configuration and writes a fresh JSON
report outside the tree. The gate validates the scanner exit status, report,
and nonempty counts for every listed source format. It compares integer duplicated-line and
total-line counts and fails when the ratio is 3% or higher, even if a rounded
percentage would appear to pass. A missing scanner or invalid report fails the
gate. To run this check alone, use `make duplication`.

On 2026-11-16 the widened candidate based on `b48b923` reported 1,082 duplicated
lines of 104,150 total lines (1.0389%). This includes 115 duplicated C header
lines. Earlier, `GOWORK=off go test ./... -cover` at `276a401`
reported 80.2% to 100% statement coverage for each Go package with executable
statements; `contract` and `wire` report no statements. These coverage values
are separate from duplication and do not establish coverage for another tree.
`cargo llvm-cov --manifest-path rust/Cargo.toml --locked --summary-only`
reported 97.86% total Rust line coverage on the same API sources.

The Rust generator shares the common scalar representation and DTO field
serialization rules through macros in `rust/src`. Each named Go scalar still
produces a distinct public Rust tuple struct. Each Go DTO still produces a
separate public Rust struct with its own fields. The field macro retains the
same JSON names, null handling, omission rules, finite checks and flattening.
The generated Go documentation remains on the corresponding types and fields.

On the same machine with jscpd 5.4.1 and its default thresholds, the issue's
`go rust/src` scan changed from 5,974 of 96,515 lines (6.1897%) at
`9c95fa09508c9e95d5b1e939f39582056fc8c194` to 833 of 81,118 lines
(1.0269%). The broader `go rust/src rust/tests` scan with five-line and
50-token thresholds changed from 6,066 of 113,633 lines (5.3382%) to 925 of
98,268 lines (0.9413%). No files were excluded from either comparison.

A source comparison checked all 203 generated scalar names, underlying types,
documents and derives, and all 3,938 DTO fields, field documents and serde
rules against the prior generator. It found no differences. The full API gate
also runs the generated Go/Rust JSON parity tests and regeneration check.
Focused Rust tests cover scalar type identity, layout and wire values, plus DTO
null defaults, omission and non-finite rejection. The Rust library checks also
passed as cross-builds for Windows GNU x86-64 and macOS ARM64; those are not
native runtime tests. Go aggregate statement coverage was 84.6% and Rust line
coverage was 97.86%; aggregate results do not prove that each example or file
meets 80%.
