// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Status is the status-bar text operation group (M05-F09).
type Status struct{ c *Client }

// Status returns the status-bar operation group.
func (c *Client) Status() Status { return Status{c} }

// SetText puts a transient message in the status bar; empty clears it.
//
// mcp:tool status_set_text
// mcp:summary Puts a transient message in the status bar; empty clears it.
func (s Status) SetText(text string) (wire.OKResult, error) {
	return call[wire.OKResult](s.c, wire.MethodStatusSetText, wire.SetStatusTextArgs{Text: text})
}

// Text returns the current status-bar message.
//
// mcp:tool status_get_text
// mcp:summary Returns the current status-bar message.
func (s Status) Text() (wire.StatusTextResult, error) {
	return call[wire.StatusTextResult](s.c, wire.MethodStatusGetText, nil)
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
//
// mcp:tool progress_begin
// mcp:summary Starts a bar of steps steps and returns its id.
func (p Progress) Begin(steps int, message string) (wire.BeginProgressResult, error) {
	return call[wire.BeginProgressResult](p.c, wire.MethodProgressBegin, wire.BeginProgressArgs{Steps: steps, Message: message})
}

// Update advances the bar to step, optionally replacing its message; the reply's
// Cancelled reports the user pressed cancel.
//
// mcp:tool progress_update
// mcp:summary Advances the bar to step, optionally replacing its message; the reply's Cancelled reports the user pressed cancel.
func (p Progress) Update(id, step int, message string) (wire.UpdateProgressResult, error) {
	args := wire.UpdateProgressArgs{ID: id, Step: step, Message: message}
	return call[wire.UpdateProgressResult](p.c, wire.MethodProgressUpdate, args)
}

// End removes the bar.
//
// mcp:tool progress_end
// mcp:summary Removes the bar.
func (p Progress) End(id int) (wire.OKResult, error) {
	return call[wire.OKResult](p.c, wire.MethodProgressEnd, wire.EndProgressArgs{ID: id})
}

// Messages is the notification + prompt + message-center operation group (M05-F09):
// balloon tips, declarative prompts with remembered answers, and the error manager's
// sectioned message tree.
type Messages struct{ c *Client }

// Messages returns the messaging operation group.
func (c *Client) Messages() Messages { return Messages{c} }

// RegisterBalloonTip declares a named notification balloon (stable id ⇒ the user's
// "don't show again" suppression survives sessions).
//
// mcp:tool balloon_tip_register
// mcp:summary Declares a named notification balloon (stable id ⇒ the user's "don't show again" suppression survives sessions).
func (m Messages) RegisterBalloonTip(args wire.RegisterBalloonTipArgs) (wire.OKResult, error) {
	return call[wire.OKResult](m.c, wire.MethodBalloonTipRegister, args)
}

// ShowBalloonTip displays a registered balloon; Shown is false when suppressed.
//
// mcp:tool balloon_tip_show
// mcp:summary Displays a registered balloon; Shown is false when suppressed.
func (m Messages) ShowBalloonTip(id string) (wire.ShowBalloonTipResult, error) {
	return call[wire.ShowBalloonTipResult](m.c, wire.MethodBalloonTipShow, wire.ShowBalloonTipArgs{ID: id})
}

// ShowPrompt queues a declarative prompt: a remembered prompt resolves in the reply,
// otherwise the user's answer arrives as a prompt.answered push event.
//
// mcp:tool prompts_show
// mcp:summary Queues a declarative prompt: a remembered prompt resolves in the reply, otherwise the user's answer arrives as a prompt.answered push event.
func (m Messages) ShowPrompt(args wire.ShowPromptArgs) (wire.ShowPromptResult, error) {
	return call[wire.ShowPromptResult](m.c, wire.MethodPromptsShow, args)
}

// AddMessage reports into the message center under the innermost open section.
//
// mcp:tool errors_add_message
// mcp:summary Reports into the message center under the innermost open section.
func (m Messages) AddMessage(args wire.AddErrorMessageArgs) (wire.OKResult, error) {
	return call[wire.OKResult](m.c, wire.MethodErrorsAddMessage, args)
}

// BeginSection groups the following messages under a titled, nestable section.
//
// mcp:tool errors_begin_section
// mcp:summary Groups the following messages under a titled, nestable section.
func (m Messages) BeginSection(title string) (wire.BeginMessageSectionResult, error) {
	return call[wire.BeginMessageSectionResult](m.c, wire.MethodErrorsBeginSection, wire.BeginMessageSectionArgs{Title: title})
}

// EndSection closes a section opened by BeginSection.
//
// mcp:tool errors_end_section
// mcp:summary Closes a section opened by BeginSection.
func (m Messages) EndSection(section int) (wire.OKResult, error) {
	return call[wire.OKResult](m.c, wire.MethodErrorsEndSection, wire.EndMessageSectionArgs{Section: section})
}

// List returns the message tree and the aggregate error/warning flags.
//
// mcp:tool errors_list
// mcp:summary Returns the message tree and the aggregate error/warning flags.
func (m Messages) List() (wire.ListErrorsResult, error) {
	return call[wire.ListErrorsResult](m.c, wire.MethodErrorsList, nil)
}

// Clear empties the message center.
//
// mcp:tool errors_clear
// mcp:summary Empties the message center.
func (m Messages) Clear() (wire.OKResult, error) {
	return call[wire.OKResult](m.c, wire.MethodErrorsClear, nil)
}

// Show opens the host's message-center panel.
//
// mcp:tool errors_show
// mcp:summary Opens the host's message-center panel.
func (m Messages) Show() (wire.OKResult, error) {
	return call[wire.OKResult](m.c, wire.MethodErrorsShow, nil)
}
