// SPDX-License-Identifier: Apache-2.0

package types

// MessageSeverity grades a message-center entry (M05-F09, #616). The zero value is
// informational so a bare message is never accidentally an error.
type MessageSeverity uint8

const (
	// SeverityInfo is a neutral notice.
	SeverityInfo MessageSeverity = 0
	// SeverityWarning flags a recoverable problem the user should know about.
	SeverityWarning MessageSeverity = 1
	// SeverityError flags a failure; the message center's error state turns on.
	SeverityError MessageSeverity = 2
)

var messageSeverityNames = map[MessageSeverity]string{
	SeverityInfo: "info", SeverityWarning: "warning", SeverityError: "error",
}

// String returns the severity's stable name ("info", "warning", "error").
func (m MessageSeverity) String() string {
	if name, ok := messageSeverityNames[m]; ok {
		return name
	}
	return "messageSeverity(?)"
}
