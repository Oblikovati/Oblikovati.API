// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Diagnostics is the operation-trace/log group: a real-time, pollable tail of the host's
// activity (every router call with timing + outcome, plus captured slog records). Tail with
// the previous result's NextSeq to fetch only new records.
type Diagnostics struct{ c *Client }

// Diagnostics returns the log/trace operation group.
func (c *Client) Diagnostics() Diagnostics { return Diagnostics{c} }

// Tail returns trace records newer than sinceSeq (0 ⇒ from the oldest retained), optionally
// filtered to a minimum level ("debug"|"info"|"warn"|"error", empty ⇒ all) and capped to max
// (0 ⇒ server default). The result's NextSeq is the cursor for the next poll.
//
// mcp:tool tail_logs
// mcp:summary Tail the host operation trace in real time: records of every command (method, durationMicros, ok/error, and panic+stack for caught kernel bugs) plus structured logs. Poll with sinceSeq=<previous result's nextSeq> to get only new records; filter with level=debug|info|warn|error; cap with max. Use after operations to verify timing and surface errors/panics while stress-testing.
func (d Diagnostics) Tail(sinceSeq uint64, level string, max int) (wire.LogsResult, error) {
	var r wire.LogsResult
	args := wire.LogsTailArgs{SinceSeq: sinceSeq, Level: level, Max: max}
	return r, d.c.call(wire.MethodLogsTail, args, &r)
}
