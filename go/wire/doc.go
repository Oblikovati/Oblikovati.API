// SPDX-License-Identifier: Apache-2.0

// Package wire is the host↔add-in transport contract: one method-name constant and a pair of
// JSON request/response DTOs per operation. It is the single place either side of the C ABI
// declares a method's name or payload shape — the GPL application's router serves these
// methods keyed on the constants, and add-ins call them through the typed client package
// rather than raw JSON (ADR-0016/ADR-0018). DTOs are plain data with json tags only; wire has
// no behavior and imports nothing but types.
package wire
