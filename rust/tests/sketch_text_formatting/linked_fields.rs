// SPDX-License-Identifier: Apache-2.0

use super::{Client, RefCell, TextFormattingHost, json, wire};
use oblikovati_api::types::{
    SKETCH_TEXT_FIELD_FILE_NAME, SKETCH_TEXT_FIELD_MODEL_PARAMETER, SketchTextField,
    SketchTextFraction,
};

const CONTENT: &str = "A🦀\u{fffc}\n\u{fffc}";
const REPLY: &str = r#"{"entityId":42,"style":{"content":"A🦀\ufffc\n\ufffc","height":0.5,"linkedFields":[{"at":2,"kind":"modelParameter","parameter":"d0","precision":0},{"at":4,"kind":"fileName"}]}}"#;

#[test]
fn typed_linked_fields_round_trip() {
    let requests = RefCell::new(Vec::new());
    let sketch = Client::new(TextFormattingHost {
        requests: &requests,
        resolved: REPLY,
    });
    let fields = sample_fields();
    let options = wire::AddTextArgs {
        text: CONTENT.into(),
        linked_fields: Some(fields.clone()),
        ..Default::default()
    };
    assert_eq!(
        sketch.sketch().add_text_with(&options).unwrap().entity_id,
        42
    );
    let edit = wire::EditTextArgs {
        entity_id: 42,
        linked_fields: Some(Box::new(Some(fields.clone()))),
        ..Default::default()
    };
    let edited = sketch.sketch().edit_text(&edit).unwrap();
    assert_eq!(edited.style.content, CONTENT);
    assert_eq!(edited.style.linked_fields, Some(fields));
    let get = wire::GetTextArgs {
        entity_id: 42,
        ..Default::default()
    };
    assert_eq!(sketch.sketch().get_text(&get).unwrap(), edited);
    assert_field_requests(&requests.borrow());
}

fn sample_fields() -> Vec<SketchTextField> {
    vec![
        SketchTextField {
            at: 2,
            kind: SKETCH_TEXT_FIELD_MODEL_PARAMETER,
            parameter: "d0".into(),
            precision: Some(Box::new(0)),
        },
        SketchTextField {
            at: 4,
            kind: SKETCH_TEXT_FIELD_FILE_NAME,
            ..Default::default()
        },
    ]
}

fn assert_field_requests(requests: &[serde_json::Value]) {
    let fields = json!([
        {"at":2,"kind":"modelParameter","parameter":"d0","precision":0},
        {"at":4,"kind":"fileName"}
    ]);
    for request in &requests[..2] {
        assert_eq!(request["formattingRuns"], serde_json::Value::Null);
        assert_eq!(request["linkedFields"], fields);
    }
}

#[test]
fn absent_edit_keeps_fields_and_explicit_empty_clears() {
    let requests = RefCell::new(Vec::new());
    let resolved = r#"{"entityId":42,"style":{"content":"plain","height":0.5}}"#;
    let client = Client::new(TextFormattingHost {
        requests: &requests,
        resolved,
    });
    for linked_fields in [None, Some(Box::new(Some(Vec::new())))] {
        let edit = wire::EditTextArgs {
            entity_id: 42,
            linked_fields,
            ..Default::default()
        };
        assert!(
            client
                .sketch()
                .edit_text(&edit)
                .unwrap()
                .style
                .linked_fields
                .is_none()
        );
    }
    assert!(requests.borrow()[0].get("linkedFields").is_none());
    assert_eq!(requests.borrow()[1]["linkedFields"], json!([]));
}

#[test]
fn linked_text_rejects_invalid_request_and_response() {
    let requests = RefCell::new(Vec::new());
    let bad_reply = REPLY.replace("\"precision\":0", "\"precision\":9");
    let client = Client::new(TextFormattingHost {
        requests: &requests,
        resolved: &bad_reply,
    });
    let bad = SketchTextField {
        at: 1,
        kind: SKETCH_TEXT_FIELD_FILE_NAME,
        ..Default::default()
    };
    let add = wire::AddTextArgs {
        text: CONTENT.into(),
        linked_fields: Some(vec![bad]),
        ..Default::default()
    };
    assert!(
        client
            .sketch()
            .add_text_with(&add)
            .unwrap_err()
            .0
            .contains("U+FFFC")
    );
    assert!(requests.borrow().is_empty());
    let get = wire::GetTextArgs {
        entity_id: 42,
        ..Default::default()
    };
    assert!(
        client
            .sketch()
            .get_text(&get)
            .unwrap_err()
            .0
            .contains("zero through eight")
    );
}

#[test]
fn linked_text_rejects_duplicate_offsets_and_fraction_overlap() {
    let fields = sample_fields();
    assert!(SketchTextField::validate(Some(CONTENT), &fields, &[]).is_ok());
    let duplicates = [fields[0].clone(), fields[0].clone()];
    assert!(
        SketchTextField::validate(Some(CONTENT), &duplicates, &[])
            .unwrap_err()
            .0
            .contains("greater than")
    );
    let fraction = SketchTextFraction { start: 0, end: 3 };
    assert!(
        SketchTextField::validate(Some("1/\u{fffc}"), &fields[..1], &[fraction])
            .unwrap_err()
            .0
            .contains("outside all stacked fractions")
    );
}
