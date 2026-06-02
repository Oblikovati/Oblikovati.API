// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"fmt"
)

// Transport is the one dependency a client has on the host: send a JSON method
// request and get the JSON reply (or an error). An add-in backs it with the host's
// C-ABI ObkHostCall callback (see add-in/include/oblikovati_addin.h); tests back it
// with a fake. The method strings are the [github.com/Oblikovati/api/wire]
// constants.
type Transport interface {
	Call(method string, req []byte) ([]byte, error)
}

// Client is a typed façade over a [Transport]: each method marshals a wire request,
// calls the host, and unmarshals the wire reply, so add-ins program against Go types
// instead of hand-rolling JSON. Reach the operation groups via [Client.Documents],
// [Client.Parameters], [Client.Model], [Client.Sketch], [Client.Features],
// [Client.Commands], [Client.Theme].
type Client struct {
	t Transport
}

// New wraps a transport. Calls on a Client with a nil transport fail with a clear
// error rather than panicking, so a half-wired add-in is diagnosable.
func New(t Transport) *Client { return &Client{t: t} }

// call marshals req (nil → no body), invokes method, and unmarshals the reply into
// out (nil → reply ignored). Errors name the offending method.
func (c *Client) call(method string, req, out any) error {
	if c.t == nil {
		return fmt.Errorf("client: no transport configured for method %q", method)
	}
	var body []byte
	if req != nil {
		b, err := json.Marshal(req)
		if err != nil {
			return fmt.Errorf("client: marshal %q request: %w", method, err)
		}
		body = b
	}
	resp, err := c.t.Call(method, body)
	if err != nil {
		return err
	}
	if out != nil {
		if err := json.Unmarshal(resp, out); err != nil {
			return fmt.Errorf("client: unmarshal %q response %q: %w", method, string(resp), err)
		}
	}
	return nil
}
