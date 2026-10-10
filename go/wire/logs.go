// SPDX-License-Identifier: Apache-2.0

package wire

// The diagnostics log/trace surface ([MethodLogsTail]): a real-time, pollable view of the
// host's operation trace — every add-in method call the router served (with timing and
// outcome), plus any structured (slog) records. It exists so a driver (an LLM test harness,
// say) can watch the kernel work, see detailed errors, and catch panics while stress-testing.
// The stream is a monotonic, gap-free sequence: poll with the last NextSeq to get only what
// is new (the cursor pattern), so "real time" is cheap repeated tailing.

// LogsTailArgs is the request of [MethodLogsTail]. SinceSeq returns only records with a
// sequence strictly greater than it (0 ⇒ from the oldest retained record); Level filters to
// records at or above a minimum severity ("debug"|"info"|"warn"|"error", empty ⇒ all); Max
// caps how many records are returned (0 ⇒ a server default).
type LogsTailArgs struct {
	SinceSeq uint64 `json:"sinceSeq,omitempty"`
	Level    string `json:"level,omitempty"`
	Max      int    `json:"max,omitempty"`
}

// LogRecord is one entry of the host operation trace. A record is either an operation entry
// (Method set — a router call, with DurationMicros and OK/Error/Panic) or a structured log
// entry (Method empty — a captured slog record with Message). Panic carries the recovered
// value and Stack the goroutine stack when a handler panicked (the kernel-bug signal).
type LogRecord struct {
	Seq            uint64 `json:"seq"`
	TimeMillis     int64  `json:"timeMillis"` // Unix ms when the record was created
	Level          string `json:"level"`      // debug|info|warn|error
	Method         string `json:"method,omitempty"`
	DurationMicros int64  `json:"durationMicros,omitempty"`
	OK             bool   `json:"ok,omitempty"`
	Message        string `json:"message,omitempty"`
	Error          string `json:"error,omitempty"`
	Panic          string `json:"panic,omitempty"`
	Stack          string `json:"stack,omitempty"`
}

// LogsResult is the response of [MethodLogsTail]: the matching records oldest-first, the
// next cursor to poll with (the highest Seq seen, or SinceSeq when none matched), and
// Dropped — how many records aged out of the ring buffer before this poll (a hint that the
// caller polled too slowly to see everything).
type LogsResult struct {
	Records []LogRecord `json:"records"`
	NextSeq uint64      `json:"nextSeq"`
	Dropped uint64      `json:"dropped,omitempty"`
}
