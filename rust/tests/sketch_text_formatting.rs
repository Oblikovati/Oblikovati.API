// SPDX-License-Identifier: Apache-2.0

use oblikovati_api::{
    Error,
    client::{Caller, Client},
    types::{Color, OVERRIDE_COLOR_SOURCE, SketchTextRun},
    wire,
};
use serde_json::{Value, json};
use std::cell::RefCell;

const MIXED: &str = "A🦀e\u{0301}界";
const STYLE: &str = r#"{"entityId":42,"style":{"content":"A🦀界","height":0.5,"formattingRuns":[{"start":1,"end":2,"bold":true}]}}"#;

fn bold_run() -> SketchTextRun {
    SketchTextRun {
        start: 1,
        end: 2,
        bold: true,
        ..Default::default()
    }
}

#[test]
fn scalar_positions_count_emoji_and_combining_marks() {
    let last_run = SketchTextRun {
        start: 3,
        end: 5,
        italic: true,
        underline: true,
        width_percent: Some(Box::new(50.)),
        color: Some(Box::new(Color {
            r: 80,
            opacity: 0.5,
            source: OVERRIDE_COLOR_SOURCE,
            ..Default::default()
        })),
        ..Default::default()
    };
    let runs = [bold_run(), last_run];
    assert!(SketchTextRun::validate(Some(MIXED), &runs).is_ok());
    assert!(SketchTextRun::validate(None, &runs).is_ok());
    let encoded = serde_json::to_value(&runs).unwrap();
    let decoded: Vec<SketchTextRun> = serde_json::from_value(encoded).unwrap();
    assert_eq!(decoded, runs);
}

#[test]
fn invalid_sketch_text_ranges_are_rejected() {
    let empty = SketchTextRun {
        start: 1,
        end: 1,
        ..Default::default()
    };
    let long = SketchTextRun {
        start: 3,
        end: 6,
        ..Default::default()
    };
    for (runs, reason) in [
        (vec![empty], "start < end"),
        (vec![bold_run(), bold_run()], "2 <= start"),
        (vec![long], "5 Unicode scalars"),
    ] {
        assert!(
            SketchTextRun::validate(Some(MIXED), &runs)
                .unwrap_err()
                .0
                .contains(reason)
        );
    }
}

#[test]
fn invalid_sketch_text_width_and_color_are_rejected() {
    for width in [0., -2., f64::INFINITY, f64::NAN] {
        let run = SketchTextRun {
            width_percent: Some(Box::new(width)),
            ..bold_run()
        };
        assert!(
            SketchTextRun::validate(Some(MIXED), &[run])
                .unwrap_err()
                .0
                .contains("positive finite percentage")
        );
    }
    for color in [
        Color::default(),
        Color {
            opacity: 1.1,
            source: OVERRIDE_COLOR_SOURCE,
            ..Default::default()
        },
    ] {
        let run = SketchTextRun {
            color: Some(Box::new(color)),
            ..bold_run()
        };
        assert!(
            SketchTextRun::validate(Some(MIXED), &[run])
                .unwrap_err()
                .0
                .contains("override color")
        );
    }
}

struct TextFormattingHost<'a> {
    requests: &'a RefCell<Vec<Value>>,
    resolved: &'a str,
}

impl Caller for TextFormattingHost<'_> {
    fn call(&self, method: &str, request: &[u8]) -> Result<Vec<u8>, Error> {
        self.requests
            .borrow_mut()
            .push(serde_json::from_slice(request).unwrap());
        let response = match method {
            wire::METHOD_SKETCH_ADD_TEXT => r#"{"entityId":42}"#,
            wire::METHOD_SKETCH_EDIT_TEXT | wire::METHOD_SKETCH_GET_TEXT => self.resolved,
            _ => panic!("unknown method {method}"),
        };
        Ok(response.as_bytes().to_vec())
    }
}

#[test]
fn typed_sketch_text_methods_round_trip_character_formatting() {
    let requests = RefCell::new(Vec::new());
    let client = Client::new(TextFormattingHost {
        requests: &requests,
        resolved: STYLE,
    });
    let sketch = client.sketch();
    let add = wire::AddTextArgs {
        text: "A🦀界".into(),
        formatting_runs: Some(vec![bold_run()]),
        ..Default::default()
    };
    assert_eq!(sketch.add_text_with(&add).unwrap().entity_id, 42);
    let edit = wire::EditTextArgs {
        entity_id: 42,
        formatting_runs: Some(Box::new(Some(vec![bold_run()]))),
        ..Default::default()
    };
    let edited = sketch.edit_text(&edit).unwrap();
    assert_eq!(edited.style.formatting_runs, Some(vec![bold_run()]));
    let args = wire::GetTextArgs {
        entity_id: 42,
        ..Default::default()
    };
    assert_eq!(sketch.get_text(&args).unwrap(), edited);
    assert_eq!(requests.borrow()[0]["formattingRuns"], json!([bold_run()]));
    assert_eq!(requests.borrow()[1]["formattingRuns"], json!([bold_run()]));
}

#[test]
fn explicit_empty_runs_clear_and_legacy_payloads_omit_the_key() {
    let requests = RefCell::new(Vec::new());
    let resolved = r#"{"entityId":42,"style":{"content":"abc","height":0.5}}"#;
    let client = Client::new(TextFormattingHost {
        requests: &requests,
        resolved,
    });
    let sketch = client.sketch();
    let mut edit = wire::EditTextArgs {
        entity_id: 42,
        ..Default::default()
    };
    assert!(
        sketch
            .edit_text(&edit)
            .unwrap()
            .style
            .formatting_runs
            .is_none()
    );
    edit.formatting_runs = Some(Box::new(Some(Vec::new())));
    assert!(
        sketch
            .edit_text(&edit)
            .unwrap()
            .style
            .formatting_runs
            .is_none()
    );
    assert!(requests.borrow()[0].get("formattingRuns").is_none());
    assert_eq!(requests.borrow()[1]["formattingRuns"], json!([]));
}

#[test]
fn invalid_typed_requests_and_responses_are_rejected() {
    let requests = RefCell::new(Vec::new());
    let resolved = STYLE.replace("\"end\":2", "\"end\":4");
    let client = Client::new(TextFormattingHost {
        requests: &requests,
        resolved: &resolved,
    });
    let sketch = client.sketch();
    let add = wire::AddTextArgs {
        text: "🦀".into(),
        formatting_runs: Some(vec![bold_run()]),
        ..Default::default()
    };
    assert!(
        sketch
            .add_text(&add)
            .unwrap_err()
            .0
            .contains("1 Unicode scalars")
    );
    assert!(requests.borrow().is_empty());
    let args = wire::GetTextArgs {
        entity_id: 42,
        ..Default::default()
    };
    assert!(
        sketch
            .get_text(&args)
            .unwrap_err()
            .0
            .contains("3 Unicode scalars")
    );
}
