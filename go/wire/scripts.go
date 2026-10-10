// SPDX-License-Identifier: Apache-2.0

package wire

// ScriptRunArgs is the request of [MethodScriptRun]: a whole Lua program to run against
// the live model in one call, instead of issuing N separate method calls. The host runs
// it in the same sandboxed runtime the GUI Script Console and the CLI use (ADR-0028), so
// the script reaches the model only through this same wire surface.
//
// An MCP/LLM client uses this to submit a complete automation program (build a sketch,
// extrude it, set parameters) atomically from its side, reading the script's print()
// output and any error back in [ScriptRunResult].
//
// WallMs is an optional per-run wall-clock budget in milliseconds (0 ⇒ the host default);
// the host caps it so a runaway script can never hang the session.
type ScriptRunArgs struct {
	Source string `json:"source"`
	WallMs int    `json:"wallMs,omitempty"`
}

// ScriptRunResult is the response of [MethodScriptRun]. The method succeeds at the
// transport level even when the *script* fails: a syntax/runtime error, quota breach, or
// cancellation is reported in Error (empty on success) alongside whatever Output the
// script printed before it stopped — so a caller always gets both the output and the
// failure reason rather than an opaque transport error.
type ScriptRunResult struct {
	Output     string `json:"output"`          // captured print() output
	Error      string `json:"error,omitempty"` // script failure message; empty on success
	DurationMs int64  `json:"durationMs"`      // wall-clock spent running
	Ops        uint64 `json:"ops,omitempty"`   // opcodes executed (best-effort; metrics)
}
