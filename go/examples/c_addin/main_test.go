// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"oblikovati.org/api"
	"oblikovati.org/api/examples/c_addin/testdata/host"
	"oblikovati.org/api/wire"
)

func TestExportsReportCompatibleVersionAndAcceptEmptyNotifications(t *testing.T) {
	if ObkAddInId() == nil || ObkAddInManifest() == nil {
		t.Fatal("expected non-nil add-in metadata")
	}
	if int(ObkAddInApiMajor()) != api.Major() || int(ObkAddInApiMinor()) != api.Minor() {
		t.Fatal("expected the contract major and minor versions")
	}
	if ObkAddInNotify(nil, 0) != 0 || ObkAddInDeactivate() != 0 {
		t.Fatal("expected empty notification and inactive shutdown to succeed")
	}
	ObkFree(nil)
}

func TestActivationRejectsMissingHostCallbacks(t *testing.T) {
	if ObkAddInActivate(nil, nil) != 1 {
		t.Fatal("expected missing host callbacks to be rejected")
	}
}

func TestCallerRejectsMethodWithNUL(t *testing.T) {
	_, err := (cgoCaller{}).Call("invalid\x00method", nil)
	if err == nil || !strings.Contains(err.Error(), "NUL-free") {
		t.Fatalf("error = %v; expected NUL-free method diagnostic", err)
	}
}

type fakeVersionHost struct {
	reply []byte
	err   error
}

func (host fakeVersionHost) Call(method string, request []byte) ([]byte, error) {
	if method != wire.MethodApplicationApiVersion || len(request) != 0 {
		return nil, fmt.Errorf("unexpected host call %q with %q", method, request)
	}
	return host.reply, host.err
}

func TestActivationChecksHostVersionAndLifecycle(t *testing.T) {
	valid := fmt.Sprintf(`{"version":%q,"major":%d,"minor":%d}`, api.Version, api.Major(), api.Minor())
	incompatible := fmt.Sprintf(`{"version":"incompatible","major":%d,"minor":0}`, api.Major()+1)
	cases := []struct {
		name   string
		host   fakeVersionHost
		status int
	}{
		{"compatible", fakeVersionHost{reply: []byte(valid)}, 0},
		{"incompatible", fakeVersionHost{reply: []byte(incompatible)}, 1},
		{"invalid reply", fakeVersionHost{reply: []byte("invalid")}, 1},
		{"host failure", fakeVersionHost{err: errors.New("host failed")}, 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := activateVersionProbe(tc.host); got != 0 {
				t.Fatalf("activate = %d, want 0", got)
			}
			if got := activateVersionProbe(tc.host); got != 1 {
				t.Fatalf("duplicate activate = %d, want 1", got)
			}
			if got := int(ObkAddInDeactivate()); got != tc.status {
				t.Fatalf("deactivate = %d, want %d", got, tc.status)
			}
		})
	}
}

func TestActivationCallsNativeHost(t *testing.T) {
	if got := host.Activate(api.Major(), api.Minor()); got != 0 {
		t.Fatalf("native activation = %d, want 0", got)
	}
	if got := ObkAddInDeactivate(); got != 0 {
		t.Fatalf("native deactivation = %d, want 0", got)
	}
}
