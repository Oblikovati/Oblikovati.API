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
//
// mcp:tool run_script
// mcp:summary Run a whole sandboxed Lua program against the model in one call (instead of many tool calls). Drive the model with oblikovati.call("method", {args}) or the typed sugar oblikovati.<group>.<method>{args} (e.g. oblikovati.documents.create{type="part"}, oblikovati.parameters.add{name="h",expression="3 cm"}); print(...) for output. Returns {output, error, durationMs}. Optional wallMs bounds the run (default 10s, max 60s).
func (s Scripts) Run(source string, wallMs int) (wire.ScriptRunResult, error) {
	return call[wire.ScriptRunResult](s.c, wire.MethodScriptRun, wire.ScriptRunArgs{Source: source, WallMs: wallMs})
}
