// SPDX-License-Identifier: Apache-2.0

// Command version demonstrates a typed client with a named fake host.
package main

import (
	"encoding/json"
	"fmt"
	"oblikovati.org/api"
	"oblikovati.org/api/client"
	"oblikovati.org/api/wire"
)

type exampleHost struct{}

func (exampleHost) Call(method string, request []byte) ([]byte, error) {
	if method != wire.MethodApplicationApiVersion || len(request) != 0 {
		return nil, fmt.Errorf("method %q payload %q: expected application.apiVersion without a body", method, request)
	}
	return json.Marshal(wire.ApplicationApiVersionResult{Version: api.Version, Major: api.Major(), Minor: api.Minor()})
}

func main() {
	version, err := client.New(exampleHost{}).Application().ApiVersion()
	if err != nil {
		panic(err)
	}
	fmt.Println(version.Version)
}
