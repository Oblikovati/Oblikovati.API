// SPDX-License-Identifier: Apache-2.0

use oblikovati_api::{
    client::{Caller, HostTransport},
    ffi::{OBK_ERR, OBK_OK},
};
use std::{
    cell::RefCell,
    ffi::{CStr, c_char, c_int},
};

#[derive(Default)]
struct CallbackHost {
    status: c_int,
    length: c_int,
    allocate: bool,
    frees: usize,
}

// The ABI has no context argument. Thread-local state keeps these callbacks independent.
thread_local! { static HOST: RefCell<CallbackHost> = RefCell::default(); }

unsafe extern "C" fn call_host(
    method: *const c_char,
    request: *const u8,
    length: c_int,
    response: *mut *mut u8,
    response_length: *mut c_int,
) -> c_int {
    // SAFETY: HostTransport supplies live method/request buffers for the duration of this call.
    unsafe {
        assert_eq!(CStr::from_ptr(method).to_str().unwrap(), "example.method");
        assert_eq!(std::slice::from_raw_parts(request, length as usize), b"{}");
    }
    HOST.with_borrow(|host| {
        // SAFETY: HostTransport supplies exclusive output pointers; free_host owns this allocation.
        unsafe {
            *response = if host.allocate {
                Box::into_raw(Box::new(*b"true")).cast()
            } else {
                std::ptr::null_mut()
            };
            *response_length = host.length;
        }
        host.status
    })
}

unsafe extern "C" fn free_host(pointer: *mut u8) {
    // SAFETY: only call_host allocations reach this callback, once through HostBuffer::drop.
    unsafe {
        drop(Box::from_raw(pointer.cast::<[u8; 4]>()));
    }
    HOST.with_borrow_mut(|host| host.frees += 1);
}

fn transport(status: c_int, length: c_int, allocate: bool) -> HostTransport<'static> {
    HOST.set(CallbackHost {
        status,
        length,
        allocate,
        frees: 0,
    });
    // SAFETY: test callbacks have process lifetime and share the allocation contract above.
    unsafe { HostTransport::new(Some(call_host), Some(free_host)).unwrap() }
}

#[test]
fn callback_validation_requires_both_functions() {
    // SAFETY: constructing with an absent callback returns before it can be called.
    unsafe {
        assert!(HostTransport::new(None, None).is_err());
        assert!(HostTransport::new(Some(call_host), None).is_err());
        assert!(HostTransport::new(None, Some(free_host)).is_err());
    }
}

#[test]
fn successful_and_failed_calls_release_the_host_allocation() {
    for status in [OBK_OK, OBK_ERR, 99] {
        let reply = transport(status, 4, true).call("example.method", b"{}");
        assert_eq!(reply.is_ok(), status == OBK_OK);
        if let Ok(bytes) = reply {
            assert_eq!(bytes, b"true");
        }
        HOST.with_borrow(|host| assert_eq!(host.frees, 1));
    }
}

#[test]
fn invalid_buffer_shapes_are_rejected_and_freed() {
    for (length, allocate) in [(-1, true), (1, false), (-1, false)] {
        let failure = transport(OBK_OK, length, allocate)
            .call("example.method", b"{}")
            .unwrap_err();
        assert!(failure.0.contains("expected a nonnegative length"));
        HOST.with_borrow(|host| assert_eq!(host.frees, usize::from(allocate)));
    }
}

#[test]
fn empty_buffers_are_valid_with_either_pointer_shape() {
    for allocate in [false, true] {
        assert_eq!(
            transport(OBK_OK, 0, allocate)
                .call("example.method", b"{}")
                .unwrap(),
            b""
        );
        HOST.with_borrow(|host| assert_eq!(host.frees, usize::from(allocate)));
    }
}

#[test]
fn embedded_nul_is_rejected_before_calling_c() {
    let failure = transport(OBK_OK, 4, true)
        .call("bad\0method", b"{}")
        .unwrap_err();
    assert!(failure.0.contains("expected no NUL"));
    HOST.with_borrow(|host| assert_eq!(host.frees, 0));
}
