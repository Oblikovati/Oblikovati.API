// SPDX-License-Identifier: Apache-2.0

use oblikovati_api::{
    Error,
    client::{Caller, Client},
    types::{self, StartupSketchPlane},
    wire::{METHOD_OPTIONS_SET_GROUP, OptionGroupView, PartOptionsView},
};
use std::{cell::RefCell, rc::Rc};

#[derive(Clone, Default)]
struct RecordingPartOptionsHost(Rc<RefCell<Vec<Vec<u8>>>>);

impl Caller for RecordingPartOptionsHost {
    fn call(&self, method: &str, request: &[u8]) -> Result<Vec<u8>, Error> {
        assert_eq!(method, METHOD_OPTIONS_SET_GROUP);
        self.0.borrow_mut().push(request.to_vec());
        Ok(br#"{"ok":true}"#.to_vec())
    }
}

fn options(plane: Option<StartupSketchPlane>, facing: Option<bool>) -> OptionGroupView {
    OptionGroupView {
        group: "part".into(),
        part: Some(Box::new(PartOptionsView {
            startup_sketch_plane: plane.map(Box::new),
            look_at_sketch_plane: facing.map(Box::new),
            ..Default::default()
        })),
        ..Default::default()
    }
}

#[test]
fn invalid_startup_plane_is_refused_before_host_call() {
    let host = RecordingPartOptionsHost::default();
    let client = Client::new(host.clone());
    let invalid = StartupSketchPlane("diagonal".into());
    let error = client
        .options()
        .set_part(&options(Some(invalid), None))
        .unwrap_err();
    assert!(error.0.contains("diagonal"));
    assert!(host.0.borrow().is_empty());
}

#[test]
fn accepted_planes_and_optional_facing_reach_host_without_value_loss() {
    let host = RecordingPartOptionsHost::default();
    let client = Client::new(host.clone());
    for plane in [
        types::STARTUP_SKETCH_NONE,
        types::STARTUP_SKETCH_XY,
        types::STARTUP_SKETCH_YZ,
        types::STARTUP_SKETCH_XZ,
    ] {
        client
            .options()
            .set_part(&options(Some(plane), Some(false)))
            .unwrap();
    }
    client.options().set_part(&options(None, None)).unwrap();
    let calls = host.0.borrow();
    assert_eq!(calls.len(), 5);
    for bytes in &calls[..4] {
        let json: serde_json::Value = serde_json::from_slice(bytes).unwrap();
        assert_eq!(json["part"]["lookAtSketchPlane"], false);
    }
    let omitted: serde_json::Value = serde_json::from_slice(&calls[4]).unwrap();
    assert!(omitted["part"].get("lookAtSketchPlane").is_none());
    assert!(omitted["part"].get("startupSketchPlane").is_none());
}
