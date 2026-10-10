// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// The status / progress / balloon-tip / prompt / message-center surfaces of
// M05-F09 (#616): the shell's feedback channels, opened to add-ins.

// SetStatusTextArgs is the request of [MethodStatusSetText]: put a transient message
// in the status bar (the same notice slot failed commits use). Empty clears it.
type SetStatusTextArgs struct {
	Text string `json:"text"`
}

// StatusTextResult is the response of [MethodStatusGetText].
type StatusTextResult struct {
	Text string `json:"text"`
}

// BeginProgressArgs is the request of [MethodProgressBegin]: start a progress bar of
// Steps steps with an initial message. Bars nest safely — each begin returns its own
// id and the status bar shows the innermost live bar.
type BeginProgressArgs struct {
	Steps   int    `json:"steps"`
	Message string `json:"message,omitempty"`
}

// BeginProgressResult is the response of [MethodProgressBegin].
type BeginProgressResult struct {
	ID int `json:"id"`
}

// UpdateProgressArgs is the request of [MethodProgressUpdate]: advance a bar to Step
// (of its Steps), optionally replacing its message.
type UpdateProgressArgs struct {
	ID      int    `json:"id"`
	Step    int    `json:"step"`
	Message string `json:"message,omitempty"`
}

// UpdateProgressResult is the response of [MethodProgressUpdate]. Cancelled reports
// the user pressed the bar's cancel control — the polling complement of the
// [ProgressCancelledEvent] push, so a step loop can stop without event plumbing.
type UpdateProgressResult struct {
	OK        bool `json:"ok"`
	Cancelled bool `json:"cancelled,omitempty"`
}

// EndProgressArgs is the request of [MethodProgressEnd].
type EndProgressArgs struct {
	ID int `json:"id"`
}

// ProgressCancelledEvent is the push event (type [EventProgressCancelled]) fired
// when the user cancels a progress bar.
type ProgressCancelledEvent struct {
	Type string `json:"type"` // always EventProgressCancelled
	ID   int    `json:"id"`
}

// RegisterBalloonTipArgs is the request of [MethodBalloonTipRegister]: declare a
// named notification balloon. Registration is separate from showing so the per-user
// "don't show again" suppression has a stable id to remember.
type RegisterBalloonTipArgs struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Icon  string `json:"icon,omitempty"`
}

// ShowBalloonTipArgs is the request of [MethodBalloonTipShow].
type ShowBalloonTipArgs struct {
	ID string `json:"id"`
}

// ShowBalloonTipResult is the response of [MethodBalloonTipShow]: Shown is false
// when the user suppressed this tip ("don't show again").
type ShowBalloonTipResult struct {
	Shown bool `json:"shown"`
}

// BalloonTipClickedEvent is the push event (type [EventBalloonTipClicked]) fired
// when the user clicks a balloon's body (not its close control).
type BalloonTipClickedEvent struct {
	Type string `json:"type"` // always EventBalloonTipClicked
	ID   string `json:"id"`
}

// ShowPromptArgs is the request of [MethodPromptsShow]: queue a declarative prompt.
// ID keys the remembered answer (when Restriction allows remembering); Buttons are
// the answer labels in display order; Default indexes the button focused first.
// The reply is asynchronous: a remembered prompt resolves instantly, otherwise the
// head shows it and the answer arrives as a [PromptAnsweredEvent].
type ShowPromptArgs struct {
	ID          string                  `json:"id"`
	Message     string                  `json:"message"`
	Buttons     []string                `json:"buttons"`
	Default     int                     `json:"default,omitempty"`
	Restriction types.PromptRestriction `json:"restriction,omitempty"`
}

// ShowPromptResult is the response of [MethodPromptsShow]: when Resolved is true the
// remembered Answer is final and no event follows; otherwise the prompt is pending
// and the answer arrives as a [PromptAnsweredEvent].
type ShowPromptResult struct {
	Resolved bool   `json:"resolved"`
	Answer   string `json:"answer,omitempty"`
}

// PromptAnsweredEvent is the push event (type [EventPromptAnswered]) fired when the
// user answers a pending prompt; Remembered reports they chose to keep the answer.
type PromptAnsweredEvent struct {
	Type       string `json:"type"` // always EventPromptAnswered
	ID         string `json:"id"`
	Answer     string `json:"answer"`
	Remembered bool   `json:"remembered,omitempty"`
}

// AddErrorMessageArgs is the request of [MethodErrorsAddMessage]: report into the
// message center, under the innermost open section (or top-level).
type AddErrorMessageArgs struct {
	Text     string                `json:"text"`
	Severity types.MessageSeverity `json:"severity,omitempty"`
}

// BeginMessageSectionArgs is the request of [MethodErrorsBeginSection]: group the
// following messages under a titled section (sections nest).
type BeginMessageSectionArgs struct {
	Title string `json:"title"`
}

// BeginMessageSectionResult is the response of [MethodErrorsBeginSection].
type BeginMessageSectionResult struct {
	Section int `json:"section"`
}

// EndMessageSectionArgs is the request of [MethodErrorsEndSection].
type EndMessageSectionArgs struct {
	Section int `json:"section"`
}

// MessageEntry is one message-center entry.
type MessageEntry struct {
	Text     string                `json:"text"`
	Severity types.MessageSeverity `json:"severity,omitempty"`
}

// MessageSectionView is one section of [MethodErrorsList]: its own messages plus
// nested sections.
type MessageSectionView struct {
	Title    string               `json:"title,omitempty"`
	Messages []MessageEntry       `json:"messages,omitempty"`
	Sections []MessageSectionView `json:"sections,omitempty"`
}

// ListErrorsResult is the response of [MethodErrorsList]: the message tree and the
// aggregate flags (the ErrorManager HasErrors/HasWarnings equivalents).
type ListErrorsResult struct {
	Root        MessageSectionView `json:"root"`
	HasErrors   bool               `json:"hasErrors"`
	HasWarnings bool               `json:"hasWarnings"`
	LastMessage string             `json:"lastMessage,omitempty"`
}
