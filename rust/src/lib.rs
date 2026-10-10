// SPDX-License-Identifier: Apache-2.0

//! Oblikovati's language-neutral automation contract.
//! Both language bindings use the same method names, JSON shapes and C ABI.
//!
//! ```
//! use oblikovati_api::{types, wire};
//! let args = wire::CreateWorkPointArgs { ..Default::default() };
//! assert_eq!(oblikovati_api::API_MAJOR, 0);
//! ```

pub mod client;
pub mod contract;
mod error;
pub mod ffi;
pub mod json;
pub mod types;
pub mod wire;
pub use error::Error;

mod version;
pub use version::{API_MAJOR, API_MINOR, VERSION};
