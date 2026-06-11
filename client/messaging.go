// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Status is the status-bar text operation group (M05-F09).
type Status struct{ c *Client }

// Status returns the status-bar operation group.
func (c *Client) Status() Status { return Status{c} }

// SetText puts a transient message in the status bar; empty clears it.
func (s Status) SetText(text string) (wire.OKResult, error) {
	var r wire.OKResult
	return r, s.c.call(wire.MethodStatusSetText, wire.SetStatusTextArgs{Text: text}, &r)
}

// Text returns the current status-bar message.
func (s Status) Text() (wire.StatusTextResult, error) {
	var r wire.StatusTextResult
	return r, s.c.call(wire.MethodStatusGetText, nil, &r)
}

// Progress is the progress-bar operation group (M05-F09): begin a bar for a long
// operation, advance it step by step, and end it. Cancellation arrives both in each
// update's reply and as a progress.cancelled push event.
type Progress struct{ c *Client }

// Progress returns the progress-bar operation group.
func (c *Client) Progress() Progress { return Progress{c} }

// Begin starts a bar of steps steps and returns its id.
//
//	bar, _ := client.Progress().Begin(100, "Meshing…")
//	for i := 1; i <= 100; i++ {
//	    if r, _ := client.Progress().Update(bar.ID, i, ""); r.Cancelled { break }
//	}
//	client.Progress().End(bar.ID)
func (p Progress) Begin(steps int, message string) (wire.BeginProgressResult, error) {
	var r wire.BeginProgressResult
	return r, p.c.call(wire.MethodProgressBegin, wire.BeginProgressArgs{Steps: steps, Message: message}, &r)
}

// Update advances the bar to step, optionally replacing its message; the reply's
// Cancelled reports the user pressed cancel.
func (p Progress) Update(id, step int, message string) (wire.UpdateProgressResult, error) {
	var r wire.UpdateProgressResult
	args := wire.UpdateProgressArgs{ID: id, Step: step, Message: message}
	return r, p.c.call(wire.MethodProgressUpdate, args, &r)
}

// End removes the bar.
func (p Progress) End(id int) (wire.OKResult, error) {
	var r wire.OKResult
	return r, p.c.call(wire.MethodProgressEnd, wire.EndProgressArgs{ID: id}, &r)
}

// Messages is the notification + prompt + message-center operation group (M05-F09):
// balloon tips, declarative prompts with remembered answers, and the error manager's
// sectioned message tree.
type Messages struct{ c *Client }

// Messages returns the messaging operation group.
func (c *Client) Messages() Messages { return Messages{c} }

// RegisterBalloonTip declares a named notification balloon (stable id ⇒ the user's
// "don't show again" suppression survives sessions).
func (m Messages) RegisterBalloonTip(args wire.RegisterBalloonTipArgs) (wire.OKResult, error) {
	var r wire.OKResult
	return r, m.c.call(wire.MethodBalloonTipRegister, args, &r)
}

// ShowBalloonTip displays a registered balloon; Shown is false when suppressed.
func (m Messages) ShowBalloonTip(id string) (wire.ShowBalloonTipResult, error) {
	var r wire.ShowBalloonTipResult
	return r, m.c.call(wire.MethodBalloonTipShow, wire.ShowBalloonTipArgs{ID: id}, &r)
}

// ShowPrompt queues a declarative prompt: a remembered prompt resolves in the reply,
// otherwise the user's answer arrives as a prompt.answered push event.
func (m Messages) ShowPrompt(args wire.ShowPromptArgs) (wire.ShowPromptResult, error) {
	var r wire.ShowPromptResult
	return r, m.c.call(wire.MethodPromptsShow, args, &r)
}

// AddMessage reports into the message center under the innermost open section.
func (m Messages) AddMessage(args wire.AddErrorMessageArgs) (wire.OKResult, error) {
	var r wire.OKResult
	return r, m.c.call(wire.MethodErrorsAddMessage, args, &r)
}

// BeginSection groups the following messages under a titled, nestable section.
func (m Messages) BeginSection(title string) (wire.BeginMessageSectionResult, error) {
	var r wire.BeginMessageSectionResult
	return r, m.c.call(wire.MethodErrorsBeginSection, wire.BeginMessageSectionArgs{Title: title}, &r)
}

// EndSection closes a section opened by BeginSection.
func (m Messages) EndSection(section int) (wire.OKResult, error) {
	var r wire.OKResult
	return r, m.c.call(wire.MethodErrorsEndSection, wire.EndMessageSectionArgs{Section: section}, &r)
}

// List returns the message tree and the aggregate error/warning flags.
func (m Messages) List() (wire.ListErrorsResult, error) {
	var r wire.ListErrorsResult
	return r, m.c.call(wire.MethodErrorsList, nil, &r)
}

// Clear empties the message center.
func (m Messages) Clear() (wire.OKResult, error) {
	var r wire.OKResult
	return r, m.c.call(wire.MethodErrorsClear, nil, &r)
}

// Show opens the host's message-center panel.
func (m Messages) Show() (wire.OKResult, error) {
	var r wire.OKResult
	return r, m.c.call(wire.MethodErrorsShow, nil, &r)
}
