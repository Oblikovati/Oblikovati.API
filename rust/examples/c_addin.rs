// SPDX-License-Identifier: Apache-2.0

//! Minimal shared-library add-in: query the host from a worker after activation returns.

use oblikovati_api::{
    API_MAJOR, API_MINOR,
    client::{Client, HostTransport},
    ffi::*,
};
use std::ffi::{c_char, c_int};
use std::sync::Mutex;
use std::thread::JoinHandle;

// C exports have no instance parameter. The transport itself stays inside the worker.
static VERSION_WORKER: Mutex<Option<JoinHandle<c_int>>> = Mutex::new(None);

#[unsafe(no_mangle)]
pub extern "C" fn ObkAddInId() -> *const c_char {
    c"com.oblikovati.example.rust".as_ptr()
}

#[unsafe(no_mangle)]
pub extern "C" fn ObkAddInManifest() -> *const c_char {
    c"{\"id\":\"com.oblikovati.example.rust\",\"displayName\":\"Rust API example\",\"version\":\"1.0.0\",\"capabilities\":[]}".as_ptr()
}

#[unsafe(no_mangle)]
pub extern "C" fn ObkAddInApiMajor() -> c_int {
    API_MAJOR
}

#[unsafe(no_mangle)]
pub extern "C" fn ObkAddInApiMinor() -> c_int {
    API_MINOR
}

/// Start a version query without blocking the host session. Example: load this library in the host.
///
/// # Safety
/// Both callback pointers must remain valid until `ObkAddInDeactivate` returns.
#[unsafe(no_mangle)]
pub unsafe extern "C" fn ObkAddInActivate(
    call: Option<ObkHostCall>,
    free: Option<ObkHostFree>,
) -> c_int {
    if call.is_none() || free.is_none() {
        return OBK_ERR;
    }
    let Ok(mut worker) = VERSION_WORKER.lock() else {
        return OBK_ERR;
    };
    if worker.is_some() {
        return OBK_ERR;
    }
    // Activation runs on the session thread; it must return before the host drains calls.
    match std::thread::Builder::new()
        .name("api-version".into())
        .spawn(move || probe_version(call, free))
    {
        Ok(started) => {
            *worker = Some(started);
            OBK_OK
        }
        Err(_) => OBK_ERR,
    }
}

fn probe_version(call: Option<ObkHostCall>, free: Option<ObkHostFree>) -> c_int {
    // SAFETY: activation's contract keeps callbacks alive until deactivation joins this worker.
    let Ok(transport) = (unsafe { HostTransport::new(call, free) }) else {
        return OBK_ERR;
    };
    let client = Client::new(transport);
    match client.application().api_version() {
        Ok(version)
            if version.major == i64::from(API_MAJOR) && version.minor >= i64::from(API_MINOR) =>
        {
            eprintln!(
                "{}",
                serde_json::json!({"msg":"api-version", "language":"rust", "version":version.version})
            );
            OBK_OK
        }
        _ => OBK_ERR,
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn ObkAddInDeactivate() -> c_int {
    let Ok(mut worker) = VERSION_WORKER.lock() else {
        return OBK_ERR;
    };
    worker
        .take()
        .map_or(OBK_OK, |started| started.join().unwrap_or(OBK_ERR))
}

#[unsafe(no_mangle)]
pub extern "C" fn ObkAddInNotify(_event: *const u8, _length: c_int) -> c_int {
    OBK_OK
}

// This example never returns an add-in-owned buffer, so there is nothing to release.
#[unsafe(no_mangle)]
pub extern "C" fn ObkFree(_pointer: *mut u8) {}
