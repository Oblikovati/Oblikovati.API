# SPDX-License-Identifier: Apache-2.0

import unittest
from unittest.mock import patch
from typing import NoReturn
from urllib.error import HTTPError
from release_checks import EXPECTED, validate_metadata, main


class MissingImportEndpoint:
    def __call__(self, url: str, timeout: int) -> NoReturn:
        raise HTTPError(url, 404, "Not Found", {}, None)


class ReleaseMetadataTests(unittest.TestCase):
    def test_unpublished_endpoint_has_an_actionable_failure(self) -> None:
        with patch("release_checks.urlopen", MissingImportEndpoint()):
            with self.assertRaisesRegex(
                SystemExit, "oblikovati.org/api.*expected a published.*404"
            ):
                main()

    def test_subdirectory_metadata_accepts_whitespace(self) -> None:
        validate_metadata(
            f'<html><meta name="go-import" content="  {EXPECTED}  "></html>'
        )

    def test_missing_legacy_and_ambiguous_metadata_are_refused(self) -> None:
        legacy = EXPECTED.removesuffix(" go")
        pages = [
            "",
            '<meta name="go-import">',
            f'<meta name="go-import" content="{legacy}">',
            f'<meta name="go-import" content="{EXPECTED}"><meta name="go-import" content="{legacy}">',
        ]
        for page in pages:
            with self.assertRaisesRegex(
                ValueError, "expected exactly.*deploy the vanity page"
            ):
                validate_metadata(page)


if __name__ == "__main__":
    unittest.main()
