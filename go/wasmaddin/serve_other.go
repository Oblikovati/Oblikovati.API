//go:build !js || !wasm

// SPDX-License-Identifier: Apache-2.0

package wasmaddin

import (
	"fmt"
	"runtime"
)

// Serve is implemented only on GOOS=js GOARCH=wasm. Native add-ins use the C ABI
// defined in include/oblikovati_addin.h.
func Serve(_ Handler) {
	panic(fmt.Sprintf("wasmaddin.Serve called on %s/%s: expected js/wasm", runtime.GOOS, runtime.GOARCH))
}
