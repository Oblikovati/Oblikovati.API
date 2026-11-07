# SPDX-License-Identifier: Apache-2.0
"""Verify the public release-check request without external network I/O."""
from io import BytesIO
import unittest
from unittest.mock import patch
from urllib.error import HTTPError
from urllib.request import Request
from release_checks import ENDPOINT, EXPECTED, USER_AGENT, main, validate_endpoint


class FakeImportResponse(BytesIO):
    status: int = 200

    def geturl(self) -> str:
        return "https://oblikovati.org/api/?go-get=1"


class RecordingImportEndpoint:
    def __init__(self) -> None:
        self.request: Request | None = None
        self.timeout: int | None = None

    def __call__(self, request: Request, timeout: int) -> FakeImportResponse:
        self.request = request
        self.timeout = timeout
        return FakeImportResponse(f'<meta name="go-import" content="{EXPECTED}">'.encode())


class ForbiddenImportEndpoint:
    def __call__(self, request: Request, timeout: int) -> FakeImportResponse:
        raise HTTPError(request.full_url, 403, "Forbidden", {}, None)


class ReleaseRequestTests(unittest.TestCase):
    def test_truthful_header_preserves_the_canonical_endpoint_and_timeout(self) -> None:
        endpoint = RecordingImportEndpoint()
        with patch("release_checks.urlopen", endpoint):
            main()
        self.assertIsNotNone(endpoint.request)
        self.assertEqual(endpoint.request.full_url, ENDPOINT)
        self.assertEqual(endpoint.request.get_header("User-agent"), USER_AGENT)
        self.assertEqual(USER_AGENT, "Oblikovati.API-release-check/1.0")
        self.assertEqual(endpoint.timeout, 20)

    def test_forbidden_response_is_not_ignored(self) -> None:
        with patch("release_checks.urlopen", ForbiddenImportEndpoint()):
            with self.assertRaisesRegex(SystemExit, "expected a published.*403"):
                main()

    def test_only_successful_https_canonical_responses_are_accepted(self) -> None:
        for endpoint in (ENDPOINT, "https://oblikovati.org/api/?go-get=1"):
            validate_endpoint(endpoint, 200)
        refused = [(ENDPOINT, 204), (ENDPOINT, 301),
                   ("http://oblikovati.org:8080/api/?go-get=1", 200),
                   ("https://example.org/api?go-get=1", 200),
                   ("https://oblikovati.org/api/?other=1", 200)]
        for endpoint, status in refused:
            with self.assertRaisesRegex(ValueError, "expected HTTP 200"):
                validate_endpoint(endpoint, status)
