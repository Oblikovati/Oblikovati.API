// SPDX-License-Identifier: Apache-2.0

// Package main builds the small Go add-in used by the shared C ABI test.
package main

/*
#cgo CFLAGS: -I${SRCDIR}/../../include -DOBK_BUILDING_ADDIN
#include <stdlib.h>
#include "oblikovati_addin.h"
static char *example_id(void) { return "com.oblikovati.example.go"; }
static char *example_manifest(void) { return "{\"id\":\"com.oblikovati.example.go\",\"displayName\":\"Go API example\",\"version\":\"1.0.0\",\"capabilities\":[]}"; }
*/
import "C"

import (
	"log/slog"
	"oblikovati.org/api"
	"oblikovati.org/api/client"
	"os"
	"sync"
)

// C exports have no instance parameter. Keep only the worker's lifecycle at this boundary.
var activationMu sync.Mutex
var versionDone chan int

//export ObkAddInId
func ObkAddInId() *C.char { return C.example_id() }

//export ObkAddInManifest
func ObkAddInManifest() *C.char { return C.example_manifest() }

//export ObkAddInApiMajor
func ObkAddInApiMajor() C.int { return C.int(api.Major()) }

//export ObkAddInApiMinor
func ObkAddInApiMinor() C.int { return C.int(api.Minor()) }

//export ObkAddInActivate
func ObkAddInActivate(call C.ObkHostCall, free C.ObkHostFree) C.int {
	if call == nil || free == nil {
		return C.OBK_ERR
	}
	return C.int(activateVersionProbe(cgoCaller{call, free}))
}

func activateVersionProbe(host client.Caller) int {
	activationMu.Lock()
	defer activationMu.Unlock()
	if versionDone != nil {
		return int(C.OBK_ERR)
	}
	versionDone = make(chan int, 1)
	// Return to the host session thread so it can drain queued calls.
	go probeVersion(host, versionDone)
	return int(C.OBK_OK)
}

func probeVersion(host client.Caller, done chan<- int) {
	version, err := client.New(host).Application().ApiVersion()
	if err != nil || version.Major != api.Major() || version.Minor < api.Minor() {
		done <- int(C.OBK_ERR)
		return
	}
	slog.New(slog.NewJSONHandler(os.Stderr, nil)).Info("api-version", "language", "go", "version", version.Version)
	done <- int(C.OBK_OK)
}

//export ObkAddInDeactivate
func ObkAddInDeactivate() C.int {
	activationMu.Lock()
	done := versionDone
	versionDone = nil
	activationMu.Unlock()
	if done == nil {
		return C.OBK_OK
	}
	return C.int(<-done)
}

//export ObkAddInNotify
func ObkAddInNotify(event *C.uint8_t, length C.int) C.int { return C.OBK_OK }

// No add-in-owned buffers are returned by this example.
//
//export ObkFree
func ObkFree(_ *C.uint8_t) {}

func main() {}
