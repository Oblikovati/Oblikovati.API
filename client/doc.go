// SPDX-License-Identifier: Apache-2.0

// Package client is the typed automation client add-ins use to drive a running
// Oblikovati host. It marshals [github.com/Oblikovati/api/wire] DTOs onto a
// caller-supplied [Transport] (the add-in backs it with the host's C-ABI
// ObkHostCall — see add-in/include/oblikovati_addin.h) and unmarshals the replies.
//
// This is the out-of-runtime path: an add-in links only the Apache-2.0 /api module
// and never the GPL implementation, so a closed-source add-in stays decoupled from
// /source both legally and at the ABI (the only thing crossing the boundary is JSON,
// per ADR-0016).
package client
