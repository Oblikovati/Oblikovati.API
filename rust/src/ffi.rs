// SPDX-License-Identifier: Apache-2.0

//! Exact signatures from `include/oblikovati_addin.h`. Only C primitives cross this ABI.

use std::ffi::{c_char, c_int};

pub const OBK_OK: c_int = 0;
pub const OBK_ERR: c_int = 1;

pub type ObkHostCall =
    unsafe extern "C" fn(*const c_char, *const u8, c_int, *mut *mut u8, *mut c_int) -> c_int;
pub type ObkHostFree = unsafe extern "C" fn(*mut u8);
pub type ObkAddInId = unsafe extern "C" fn() -> *const c_char;
pub type ObkAddInManifest = unsafe extern "C" fn() -> *const c_char;
pub type ObkAddInApiMajor = unsafe extern "C" fn() -> c_int;
pub type ObkAddInApiMinor = unsafe extern "C" fn() -> c_int;
pub type ObkAddInActivate = unsafe extern "C" fn(Option<ObkHostCall>, Option<ObkHostFree>) -> c_int;
pub type ObkAddInDeactivate = unsafe extern "C" fn() -> c_int;
pub type ObkAddInNotify = unsafe extern "C" fn(*const u8, c_int) -> c_int;
pub type ObkFree = unsafe extern "C" fn(*mut u8);
pub type ObkAddInAutomation = ObkHostCall;
