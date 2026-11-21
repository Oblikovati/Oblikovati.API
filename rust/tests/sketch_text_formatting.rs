// SPDX-License-Identifier: Apache-2.0

use oblikovati_api::{
    Error,
    client::{Caller, Client},
    types::{
        Color, OVERRIDE_COLOR_SOURCE, SketchTextFraction, SketchTextLineSpacing, SketchTextRun,
        TextLineSpacingMode,
    },
    wire,
};
use serde_json::{Value, json};
use std::cell::RefCell;

const MIXED: &str = "A🦀e\u{0301}界";
const STYLE: &str = r#"{"entityId":42,"style":{"content":"A🦀界","height":0.5,"formattingRuns":[{"start":1,"end":2,"bold":true,"strikeThrough":true}]}}"#;

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
    let run = SketchTextRun {
        strike_through: true,
        ..bold_run()
    };
    let requests = RefCell::new(Vec::new());
    let client = Client::new(TextFormattingHost {
        requests: &requests,
        resolved: STYLE,
    });
    let sketch = client.sketch();
    let add = wire::AddTextArgs {
        text: "A🦀界".into(),
        formatting_runs: Some(vec![run.clone()]),
        ..Default::default()
    };
    assert_eq!(sketch.add_text_with(&add).unwrap().entity_id, 42);
    let edit = wire::EditTextArgs {
        entity_id: 42,
        formatting_runs: Some(Box::new(Some(vec![run.clone()]))),
        ..Default::default()
    };
    let edited = sketch.edit_text(&edit).unwrap();
    assert_eq!(edited.style.formatting_runs, Some(vec![run.clone()]));
    let args = wire::GetTextArgs {
        entity_id: 42,
        ..Default::default()
    };
    assert_eq!(sketch.get_text(&args).unwrap(), edited);
    assert_eq!(
        requests.borrow()[0]["formattingRuns"],
        json!([{"start":1,"end":2,"bold":true,"strikeThrough":true}])
    );
    assert_eq!(
        requests.borrow()[1]["formattingRuns"],
        json!([{"start":1,"end":2,"bold":true,"strikeThrough":true}])
    );
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

fn text_spacing(mode: &str, value: f64) -> SketchTextLineSpacing {
    SketchTextLineSpacing {
        mode: TextLineSpacingMode(mode.to_owned().into()),
        value,
    }
}

#[test]
fn line_spacing_validates_modes_and_values() {
    assert!(SketchTextLineSpacing::validate(None).is_ok());
    for (mode, value) in [
        ("single", 0.),
        ("oneAndHalf", 0.),
        ("double", 0.),
        ("multiple", 1.75),
        ("exactly", 0.61),
    ] {
        assert!(SketchTextLineSpacing::validate(Some(&text_spacing(mode, value))).is_ok());
    }
    for (mode, value) in [
        ("unknown", 0.),
        ("double", 2.),
        ("multiple", 0.),
        ("exactly", -1.),
        ("multiple", f64::INFINITY),
        ("exactly", f64::NAN),
    ] {
        assert!(SketchTextLineSpacing::validate(Some(&text_spacing(mode, value))).is_err());
    }
}

#[test]
fn typed_spacing_edits_distinguish_preserve_and_single() {
    let requests = RefCell::new(Vec::new());
    let client = Client::new(TextFormattingHost {
        requests: &requests,
        resolved: r#"{"entityId":42,"style":{"content":"a\nb","height":0.5,"lineSpacing":{"mode":"double"}}}"#,
    });
    let sketch = client.sketch();
    let mut edit = wire::EditTextArgs {
        entity_id: 42,
        ..Default::default()
    };
    let prior = sketch.edit_text(&edit).unwrap();
    assert_eq!(
        prior.style.line_spacing.as_deref(),
        Some(&text_spacing("double", 0.))
    );
    edit.line_spacing = Some(Box::new(text_spacing("single", 0.)));
    sketch.edit_text(&edit).unwrap();
    assert!(requests.borrow()[0].get("lineSpacing").is_none());
    assert_eq!(
        requests.borrow()[1]["lineSpacing"],
        json!({"mode":"single"})
    );
}

#[test]
fn invalid_typed_spacing_is_rejected_on_both_sides() {
    let requests = RefCell::new(Vec::new());
    let client = Client::new(TextFormattingHost {
        requests: &requests,
        resolved: r#"{"entityId":42,"style":{"content":"a","height":0.5,"lineSpacing":{"mode":"exactly"}}}"#,
    });
    let sketch = client.sketch();
    let add = wire::AddTextArgs {
        text: "a".into(),
        line_spacing: Some(Box::new(text_spacing("multiple", 0.))),
        ..Default::default()
    };
    assert!(
        sketch
            .add_text(&add)
            .unwrap_err()
            .0
            .contains("positive finite")
    );
    assert!(requests.borrow().is_empty());
    let read = wire::GetTextArgs {
        entity_id: 42,
        ..Default::default()
    };
    assert!(
        sketch
            .get_text(&read)
            .unwrap_err()
            .0
            .contains("positive finite")
    );
}

const FRACTION: SketchTextFraction = SketchTextFraction { start: 1, end: 4 };
const FRACTION_STYLE: &str = r#"{"entityId":42,"style":{"content":"⌀1/2","height":0.5,"stackedFractions":[{"start":1,"end":4}]}}"#;

#[test]
fn stacked_fraction_checks_bounds_and_both_terms() {
    for text in ["⌀1/2", "⌀界/二"] {
        assert!(SketchTextFraction::validate(Some(text), &[FRACTION]).is_ok());
    }
    assert!(SketchTextFraction::validate(None, &[FRACTION]).is_ok());
    assert!(SketchTextFraction::validate(None, &[FRACTION, FRACTION]).is_err());
    for text in ["1/2", "⌀ /2", "⌀123", "⌀1//", "⌀1/\n"] {
        assert!(
            SketchTextFraction::validate(Some(text), &[FRACTION]).is_err(),
            "{text:?}"
        );
    }
}

#[test]
fn typed_stacked_fraction_round_trips_and_distinguishes_empty_edit() {
    let requests = RefCell::new(Vec::new());
    let client = Client::new(TextFormattingHost {
        requests: &requests,
        resolved: FRACTION_STYLE,
    });
    let sketch = client.sketch();
    let add = wire::AddTextArgs {
        text: "⌀1/2".into(),
        stacked_fractions: Some(vec![FRACTION]),
        ..Default::default()
    };
    assert_eq!(sketch.add_text_with(&add).unwrap().entity_id, 42);
    let mut edit = wire::EditTextArgs {
        entity_id: 42,
        ..Default::default()
    };
    assert_eq!(
        sketch.edit_text(&edit).unwrap().style.stacked_fractions,
        Some(vec![FRACTION])
    );
    edit.stacked_fractions = Some(Box::new(Some(Vec::new())));
    sketch.edit_text(&edit).unwrap();
    assert_eq!(requests.borrow()[0]["stackedFractions"], json!([FRACTION]));
    assert!(requests.borrow()[1].get("stackedFractions").is_none());
    assert_eq!(requests.borrow()[2]["stackedFractions"], json!([]));
}

#[test]
fn typed_fraction_validates_request_and_response() {
    let requests = RefCell::new(Vec::new());
    let bad_style = FRACTION_STYLE.replace("\"end\":4", "\"end\":6");
    let client = Client::new(TextFormattingHost {
        requests: &requests,
        resolved: &bad_style,
    });
    let sketch = client.sketch();
    let add = wire::AddTextArgs {
        text: "⌀123".into(),
        stacked_fractions: Some(vec![FRACTION]),
        ..Default::default()
    };
    assert!(
        sketch
            .add_text_with(&add)
            .unwrap_err()
            .0
            .contains("one slash")
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
            .contains("4 Unicode scalars")
    );
}
