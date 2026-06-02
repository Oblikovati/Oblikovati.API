// SPDX-License-Identifier: Apache-2.0

// Package contract is the in-process Go interface surface of the Oblikovati API:
// Application/Session, Document, PartDocument, Parameters, Parameter, Sketch, and
// the collections that hang off them. The GPL implementation (/source) satisfies
// these interfaces with compile-time assertions, so the contract stays honest
// without /api ever importing the implementation.
//
// These interfaces are for FIRST-PARTY, in-process use (one Go runtime). They are
// NOT how out-of-process or C-ABI add-ins reach the host — a live Go interface
// value cannot cross the two-runtime boundary of ADR-0016. Those add-ins use the
// transport-backed [github.com/Oblikovati/api/client] over the
// [github.com/Oblikovati/api/wire] JSON contract instead.
package contract
