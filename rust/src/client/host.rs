// SPDX-License-Identifier: Apache-2.0

use crate::{
    Error,
    ffi::{OBK_ERR, OBK_OK, ObkHostCall, ObkHostFree},
};
use std::{
    ffi::{CString, c_int},
    marker::PhantomData,
    rc::Rc,
};

/// A host callback pair whose buffers are always released by the owning host.
/// Construct it on the worker that makes calls. The host serializes them on its session.
pub struct HostTransport<'host> {
    call: ObkHostCall,
    free: ObkHostFree,
    lifetime: PhantomData<(&'host (), Rc<()>)>,
}

impl<'host> HostTransport<'host> {
    /// Bind callbacks passed to `ObkAddInActivate`.
    /// Example: `unsafe { HostTransport::new(call, free) }?`.
    ///
    /// # Safety
    /// The callbacks and their library must remain valid for `'host`. Calls must
    /// follow `oblikovati_addin.h`: each returned non-null pointer names an allocation
    /// of at least the returned length and is accepted exactly once by `free`.
    pub unsafe fn new(call: Option<ObkHostCall>, free: Option<ObkHostFree>) -> Result<Self, Error> {
        match (call, free) {
            (Some(call), Some(free)) => Ok(Self {
                call,
                free,
                lifetime: PhantomData,
            }),
            _ => Err(Error(format!(
                "host callbacks call={}, free={}: expected two non-null callbacks",
                call.is_some(),
                free.is_some()
            ))),
        }
    }
}

impl super::Caller for HostTransport<'_> {
    fn call(&self, method: &str, request: &[u8]) -> Result<Vec<u8>, Error> {
        let name = CString::new(method)
            .map_err(|_| Error(format!("method {method:?}: expected no NUL bytes")))?;
        let length = c_int::try_from(request.len()).map_err(|_| {
            Error(format!(
                "request length {}: expected at most {}",
                request.len(),
                c_int::MAX
            ))
        })?;
        let mut response = HostBuffer {
            pointer: std::ptr::null_mut(),
            length: 0,
            free: self.free,
        };
        // SAFETY: construction establishes callback validity; the input slices live through the call.
        let status = unsafe {
            (self.call)(
                name.as_ptr(),
                request.as_ptr(),
                length,
                &mut response.pointer,
                &mut response.length,
            )
        };
        let bytes = response.copy(method)?;
        match status {
            OBK_OK => Ok(bytes),
            OBK_ERR => Err(Error(format!(
                "host method {method:?} failed: {}",
                String::from_utf8_lossy(&bytes)
            ))),
            _ => Err(Error(format!(
                "host method {method:?} status {status}: expected OBK_OK(0) or OBK_ERR(1)"
            ))),
        }
    }
}

struct HostBuffer {
    pointer: *mut u8,
    length: c_int,
    free: ObkHostFree,
}

impl HostBuffer {
    fn copy(&self, method: &str) -> Result<Vec<u8>, Error> {
        if self.length < 0 || (self.length > 0 && self.pointer.is_null()) {
            return Err(Error(format!(
                "host method {method:?} buffer {:?}, length {}: expected a nonnegative length and an allocated buffer for nonzero length",
                self.pointer, self.length
            )));
        }
        if self.length == 0 {
            return Ok(Vec::new());
        }
        // SAFETY: the constructor contract guarantees the reported allocation length; null and sign were checked.
        Ok(unsafe { std::slice::from_raw_parts(self.pointer, self.length as usize) }.to_vec())
    }
}

impl Drop for HostBuffer {
    fn drop(&mut self) {
        if self.pointer.is_null() {
            return;
        }
        // SAFETY: this unique guard releases only the allocation returned by its paired callback.
        unsafe { (self.free)(self.pointer) };
    }
}
