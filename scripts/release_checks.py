# SPDX-License-Identifier: Apache-2.0
"""Refuse a Go subdirectory release until its public import metadata is ready."""

from html.parser import HTMLParser
from urllib.request import urlopen

ENDPOINT = "https://oblikovati.org/api?go-get=1"
EXPECTED = "oblikovati.org/api git https://github.com/Oblikovati/Oblikovati.API go"


class ImportMetadata(HTMLParser):
    def __init__(self) -> None:
        super().__init__()
        self.values: list[str] = []

    def handle_starttag(self, tag: str, attrs: list[tuple[str, str | None]]) -> None:
        attributes = dict(attrs)
        if tag == "meta" and attributes.get("name") == "go-import":
            self.values.append(attributes.get("content") or "")


def validate_metadata(page: str) -> None:
    parser = ImportMetadata()
    parser.feed(page)
    values = [" ".join(value.split()) for value in parser.values]
    if values != [EXPECTED]:
        raise ValueError(
            f"{ENDPOINT} metadata {values!r}: expected exactly {EXPECTED!r}; deploy the vanity page before releasing"
        )


def main() -> None:
    try:
        with urlopen(ENDPOINT, timeout=20) as response:
            page = response.read().decode("utf-8")
    except (OSError, UnicodeError) as error:
        raise SystemExit(
            f"{ENDPOINT}: expected a published UTF-8 Go import page; request failed: {error}"
        ) from error
    validate_metadata(page)
    print("Go import metadata selects the released go/ module")


if __name__ == "__main__":
    main()
