/* SPDX-License-Identifier: Apache-2.0 */

#include "oblikovati_addin.h"
#include <assert.h>
#include <dlfcn.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdatomic.h>

static atomic_int freed;
static atomic_int status;
static atomic_int malformed;
static int expected_major;
static int expected_minor;
static char version_response[256];

static int host_call(const char *method, const uint8_t *request, int length,
                     uint8_t **response, int *response_length) {
    assert(strcmp(method, "application.apiVersion") == 0);
    assert(length == 0);
    (void)request;
    const char *text = status ? "host failure" : version_response;
    if (malformed) text = "invalid JSON";
    *response_length = (int)strlen(text);
    *response = malloc((size_t)*response_length);
    assert(*response != NULL);
    memcpy(*response, text, (size_t)*response_length);
    return status;
}

static void host_free(uint8_t *pointer) {
    assert(pointer != NULL);
    ++freed;
    free(pointer);
}

static void *symbol(void *library, const char *name) {
    void *address = dlsym(library, name);
    if (address == NULL) { fprintf(stderr, "missing export %s: %s\n", name, dlerror()); exit(1); }
    return address;
}

static void check_callbacks(void *library) {
    int (*activate)(ObkHostCall, ObkHostFree) = symbol(library, "ObkAddInActivate");
    int (*deactivate)(void) = symbol(library, "ObkAddInDeactivate");
    assert(activate(NULL, NULL) == OBK_ERR);
    freed = 0; status = OBK_OK; malformed = 0;
    assert(activate(host_call, host_free) == OBK_OK);
    assert(activate(host_call, host_free) == OBK_ERR);
    assert(deactivate() == OBK_OK);
    assert(freed == 1);
    status = OBK_ERR;
    assert(activate(host_call, host_free) == OBK_OK);
    assert(deactivate() == OBK_ERR);
    assert(freed == 2);
    status = OBK_OK; malformed = 1;
    assert(activate(host_call, host_free) == OBK_OK);
    assert(deactivate() == OBK_ERR);
    assert(freed == 3);
}

static void check_exports(void *library) {
    const char *(*identity)(void) = symbol(library, "ObkAddInId");
    const char *(*manifest)(void) = symbol(library, "ObkAddInManifest");
    int (*major)(void) = symbol(library, "ObkAddInApiMajor");
    int (*minor)(void) = symbol(library, "ObkAddInApiMinor");
    int (*deactivate)(void) = symbol(library, "ObkAddInDeactivate");
    int (*notify)(const uint8_t *, int) = symbol(library, "ObkAddInNotify");
    assert(strlen(identity()) > 0 && strlen(manifest()) > 0);
    assert(major() == expected_major && minor() == expected_minor);
    assert(notify(NULL, 0) == OBK_OK);
    assert(deactivate() == OBK_OK);
    assert(symbol(library, "ObkFree") != NULL);
}

int main(int argc, char **argv) {
    assert(argc == 3);
    int patch;
    assert(sscanf(argv[2], "%d.%d.%d", &expected_major, &expected_minor, &patch) == 3);
    int written = snprintf(version_response, sizeof(version_response),
        "{\"version\":\"%s\",\"major\":%d,\"minor\":%d}", argv[2], expected_major, expected_minor);
    assert(written > 0 && (size_t)written < sizeof(version_response));
    void *library = dlopen(argv[1], RTLD_NOW | RTLD_LOCAL);
    if (library == NULL) { fprintf(stderr, "%s\n", dlerror()); return 1; }
    check_exports(library);
    check_callbacks(library);
    /* Go's runtime owns worker threads; leave each library loaded until process exit. */
    puts("C ABI: exports, typed call, error propagation and buffer ownership passed");
    return 0;
}
