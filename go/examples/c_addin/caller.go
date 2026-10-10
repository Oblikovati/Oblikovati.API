// SPDX-License-Identifier: Apache-2.0

package main

/*
#cgo CFLAGS: -I${SRCDIR}/../../../include -DOBK_BUILDING_ADDIN
#include <stdlib.h>
#include <limits.h>
#include "oblikovati_addin.h"
static int invoke_host(ObkHostCall call, const char *method, const uint8_t *request, int length, uint8_t **response, int *size) { return call(method, request, length, response, size); }
static void release_host(ObkHostFree free, uint8_t *response) { free(response); }
*/
import "C"

import (
	"fmt"
	"strings"
	"unsafe"
)

type cgoCaller struct {
	call C.ObkHostCall
	free C.ObkHostFree
}

func (host cgoCaller) Call(method string, request []byte) ([]byte, error) {
	if strings.ContainsRune(method, 0) || len(request) > int(C.INT_MAX) {
		return nil, fmt.Errorf("request %q length %d: expected a NUL-free method and C int length", method, len(request))
	}
	name := C.CString(method)
	defer C.free(unsafe.Pointer(name))
	var response *C.uint8_t
	var length C.int
	status := C.invoke_host(host.call, name, (*C.uint8_t)(unsafe.Pointer(unsafe.SliceData(request))), C.int(len(request)), &response, &length)
	if response != nil {
		defer C.release_host(host.free, response)
	}
	if length < 0 || (length > 0 && response == nil) {
		return nil, fmt.Errorf("response %p length %d: expected allocated bytes and a nonnegative length", response, length)
	}
	bytes := C.GoBytes(unsafe.Pointer(response), length)
	if status != C.OBK_OK {
		return nil, fmt.Errorf("host method %q status %d: %s", method, status, bytes)
	}
	return bytes, nil
}
