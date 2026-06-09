// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Scripts is the operation group for running a Lua program against the live model in one
// call (the sandboxed runtime of ADR-0028). An MCP/LLM client uses it to submit a whole
// automation program instead of issuing many individual method calls.
type Scripts struct{ c *Client }

// Scripts returns the scripting operation group.
func (c *Client) Scripts() Scripts { return Scripts{c} }

// Run executes a Lua source against the host and returns its captured output, any script
// error (reported in the result, not as a transport error), and timing. wallMs is an
// optional per-run wall-clock budget in milliseconds (0 ⇒ the host default).
//
//	res, _ := client.Scripts().Run(`oblikovati.documents.create{ type="part" }`, 0)
//	if res.Error != "" { log.Print(res.Error) }
func (s Scripts) Run(source string, wallMs int) (wire.ScriptRunResult, error) {
	var r wire.ScriptRunResult
	return r, s.c.call(wire.MethodScriptRun, wire.ScriptRunArgs{Source: source, WallMs: wallMs}, &r)
}
