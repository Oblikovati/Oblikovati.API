// SPDX-License-Identifier: Apache-2.0

package types

// WindowState is a view frame's window state — the WindowsSizeEnum equivalent
// (M05-F10, #617).
type WindowState uint8

const (
	// WindowNormal is a regular floating window (the zero value).
	WindowNormal WindowState = 0
	// WindowMaximized fills the screen.
	WindowMaximized WindowState = 1
	// WindowMinimized is iconified.
	WindowMinimized WindowState = 2
)

var windowStateNames = map[WindowState]string{
	WindowNormal: "normal", WindowMaximized: "maximized", WindowMinimized: "minimized",
}

// String returns the state's stable name.
func (w WindowState) String() string {
	return enumName(windowStateNames, w, "windowState(?)")
}
