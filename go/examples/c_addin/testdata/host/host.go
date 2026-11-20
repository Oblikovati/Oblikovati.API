// SPDX-License-Identifier: Apache-2.0

// Package host supplies C callbacks for the add-in ABI tests.
package host

/*
#cgo CFLAGS: -I${SRCDIR}/../../../../include
#include "oblikovati_addin.h"
#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
static int major_version;
static int minor_version;
static int call_host(const char *method, const uint8_t *request, int length, uint8_t **response, int *size) {
    (void)method; (void)request; (void)length;
    char json[128];
    int n = snprintf(json, sizeof(json), "{\"version\":\"test\",\"major\":%d,\"minor\":%d}", major_version, minor_version);
    *response = malloc((size_t)n);
    if (!*response) return 1;
    memcpy(*response, json, (size_t)n);
    *size = n;
    return 0;
}
static void free_host(uint8_t *response) { free(response); }
static void configure(int major, int minor) { major_version = major; minor_version = minor; }
static int activate_host(void) { return ObkAddInActivate(call_host, free_host); }
*/
import "C"

// Activate starts the add-in with native host callbacks for the requested version.
// For example, Activate(0, 159) starts a host that reports API version 0.159.
func Activate(major, minor int) int {
	C.configure(C.int(major), C.int(minor))
	return int(C.activate_host())
}
