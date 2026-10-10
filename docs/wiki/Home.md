# Oblikovati Add-in Developer Wiki

Build add-ins with Go or Rust through the same public API. Both bindings are
Apache-2.0 and share one C interface, one JSON protocol and one version.
The application owns the model; the add-in sends typed requests through a
host callback.

This wiki is built from `Oblikovati.API/docs/wiki` and the checked-in API
sources. Edit those sources. The workflow replaces manual wiki edits.

## Start here

| Page | Purpose |
| --- | --- |
| [[First Steps]] | Install either binding and run a complete typed-client example. |
| [[C ABI Add-ins]] | Build Go and Rust shared libraries with the same host interface. |
| [[Typed Operations]] | Use matching Go and Rust requests, values and events. |
| [[Oblikovati API Architecture]] | Understand the contract modules and compatibility rules. |
| [[Testing Automation]] | Test clients, JSON compatibility and C buffer ownership. |
| [[API Docs]] | Open the generated Go and Rust reference pages. |
| [[Oblikovati Architecture]] | Understand the application's model and geometry layers. |

## License boundary

The Go module `oblikovati.org/api` and Rust crate `oblikovati-api` are
Apache-2.0. The application is GPL-2.0-only. Neither API binding imports or
links the application. Add-ins communicate through C functions and byte
buffers. They do not pass Go interfaces or Rust trait objects across C.
