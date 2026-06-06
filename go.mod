// Module oblikovati/api is the public, Apache-2.0 automation contract
// for Oblikovati: the value types, enums, Go interfaces, JSON wire DTOs, and a
// transport-backed client that both the GPL application (/source) and third-party
// (possibly closed-source) add-ins depend on. It is the single source of truth for
// the API surface and MUST NOT import the GPL implementation module
// oblikovati — the dependency only ever flows the other way.
// See architecture/decisions/ADR-0018.
module oblikovati/api

go 1.22
