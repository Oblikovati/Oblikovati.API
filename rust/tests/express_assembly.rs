// SPDX-License-Identifier: Apache-2.0
use oblikovati_api::wire::{DocumentInfo, LoadFullDocumentArgs, OpenDocumentArgs};

#[test]
fn old_open_requests_and_responses_default_to_full_mode() {
    let open: OpenDocumentArgs =
        serde_json::from_str(r#"{"fullDocumentName":"a.oad","visible":true}"#).unwrap();
    let info: DocumentInfo = serde_json::from_str(
        r#"{"id":7,"name":"a","type":"assembly","dirty":false,"visible":true,"active":true}"#,
    )
    .unwrap();
    assert!(!open.express && !info.express);
    assert!(serde_json::to_value(open).unwrap().get("express").is_none());
}

#[test]
fn express_option_and_load_full_identity_round_trip() {
    let open = OpenDocumentArgs {
        full_document_name: "a.oad".into(),
        visible: true,
        express: true,
        ..Default::default()
    };
    let decoded: OpenDocumentArgs =
        serde_json::from_value(serde_json::to_value(open.clone()).unwrap()).unwrap();
    assert_eq!(decoded, open);
    let args = LoadFullDocumentArgs { id: 7 };
    assert_eq!(serde_json::to_value(args).unwrap()["id"], 7);
}
