# SPDX-License-Identifier: Apache-2.0

.PHONY: gate build test lint generate check-generated docs docs-lint abi-test
export GOWORK := off
# Keep rustc, rustdoc and rustfmt on the pinned toolchain when rustup is installed.
export PATH := $(HOME)/.cargo/bin:$(PATH)

gate: check-generated build test lint docs-lint docs abi-test

build:
	cd go && go build ./...
	cargo build --manifest-path rust/Cargo.toml --locked --all-targets

test:
	cd go && go test ./...
	cargo test --manifest-path rust/Cargo.toml --locked
	cargo test --manifest-path rust/Cargo.toml --locked --examples
	python3 -m unittest discover -s scripts/rustgen -p 'test_*.py'
	python3 -m unittest discover -s scripts -p 'test_*.py'

lint:
	cd go && go vet ./...
	test -z "$$(gofmt -l go)"
	cargo fmt --manifest-path rust/Cargo.toml -- --check
	cargo clippy --manifest-path rust/Cargo.toml --locked --all-targets -- -D warnings
	python3 scripts/check-spdx.py --check
	python3 scripts/check-license-boundary.py

generate:
	mkdir -p build
	cd go && go run ./scripts/apischema > ../build/api-schema.json
	python3 scripts/rustgen/main.py < build/api-schema.json
	gofmt -w go/scripts/apifixtures go/scripts/apischema go/client go/examples
	cd go && go run ./scripts/apifixtures > ../build/go-wire.jsonl
	python3 scripts/split-fixtures.py
	cargo fmt --manifest-path rust/Cargo.toml
	python3 scripts/wiki-examples.py
	cargo update --manifest-path rust/Cargo.toml --workspace

check-generated:
	python3 scripts/check-generated.py

docs:
	bash scripts/assemble-wiki.sh
	cargo doc --manifest-path rust/Cargo.toml --locked --no-deps

docs-lint:
	npx --yes markdownlint-cli2@0.18.1 README.md RELEASING.md 'docs/wiki/*.md'
	python3 scripts/wiki-examples.py --check

abi-test:
	bash scripts/test-c-abi.sh
