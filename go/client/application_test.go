// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"oblikovati.org/api/wire"
)

// TestApplicationApiVersionDecodesReply asserts the version query targets the right
// wire method, sends no body, and decodes the host's reported version + major.
func TestApplicationApiVersionDecodesReply(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"version":"1.4.2","major":1,"minor":4}`)}
	c := New(ft)

	got, err := c.Application().ApiVersion()
	if err != nil {
		t.Fatalf("ApiVersion: %v", err)
	}
	if ft.gotMethod != wire.MethodApplicationApiVersion {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodApplicationApiVersion)
	}
	if ft.gotReq != nil {
		t.Errorf("request body = %q, want nil (no-arg method)", ft.gotReq)
	}
	if got.Version != "1.4.2" || got.Major != 1 || got.Minor != 4 {
		t.Errorf("decoded = %+v, want {Version:1.4.2 Major:1 Minor:4}", got)
	}
}
