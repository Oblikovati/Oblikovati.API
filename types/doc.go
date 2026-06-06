// SPDX-License-Identifier: Apache-2.0

// Package types holds the shared vocabulary of the Oblikovati API: enums,
// stable identifiers, and value/option structs with no behavior. Both the GPL
// implementation and add-ins use these as a common, Apache-2.0 type currency so
// the same concepts (a document kind, a parameter kind, a 2D point) are named once.
//
// Types here are pure data — no dependency on the implementation, no methods that
// touch live model state. The richer behavioral surface lives in
// [oblikovati/api/contract] (in-proc Go interfaces) and
// [oblikovati/api/wire] (the JSON contract).
package types
