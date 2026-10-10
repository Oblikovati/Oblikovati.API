// SPDX-License-Identifier: Apache-2.0

package wire

// The help and language surface of M05-F14 (#621): add-ins route their F1/help
// requests through the host (so add-in help opens like product help) and read the
// host's locale. Sources resolve to URLs or local files; the host opens them with
// its platform opener (or an in-process interceptor handles them first).

// RegisterHelpContextArgs is the request of [MethodHelpRegisterContext]: declare a
// help source — Base is a URL prefix (https://…) or a local directory; topics
// resolve against it.
type RegisterHelpContextArgs struct {
	Source string `json:"source"`
	Base   string `json:"base"`
}

// DisplayHelpArgs is the request of [MethodHelpDisplay]: open a topic of a
// registered source ("" ⇒ the host's own documentation).
type DisplayHelpArgs struct {
	Source string `json:"source,omitempty"`
	Topic  string `json:"topic,omitempty"`
}

// HelpPathResult is the response of [MethodHelpPath]: a source's registered base.
type HelpPathResult struct {
	Source string `json:"source,omitempty"`
	Base   string `json:"base"`
}

// LanguageInfoResult is the response of [MethodLanguageInfo]: the host's locale as
// a BCP-47 tag (e.g. "en-US"), for add-ins formatting numbers/dates to match.
// Translated-name catalogs do not exist yet; command display names are canonical
// (the #160 ADR records the declined translation surface).
type LanguageInfoResult struct {
	Locale string `json:"locale"`
}
