// SPDX-License-Identifier: Apache-2.0
use oblikovati_api::wire::{
    BatchSaveItemResult, CloseDocumentsResult, SAVE_WARNING_ASSEMBLY_CACHE_UNAVAILABLE,
    SaveDocumentResult, SaveWarning,
};

#[test]
fn warnings_survive_all_save_result_shapes() {
    let warning = SaveWarning {
        document: 7,
        code: SAVE_WARNING_ASSEMBLY_CACHE_UNAVAILABLE.into(),
        detail: "component is dirty".into(),
    };
    let values = [
        serde_json::to_value(SaveDocumentResult {
            full_document_name: "a.oad".into(),
            warnings: Some(vec![warning.clone()]),
        })
        .unwrap(),
        serde_json::to_value(BatchSaveItemResult {
            document: 7,
            ok: true,
            warnings: Some(vec![warning.clone()]),
            ..Default::default()
        })
        .unwrap(),
        serde_json::to_value(CloseDocumentsResult {
            closed: 1,
            warnings: Some(vec![warning]),
        })
        .unwrap(),
    ];
    for value in values {
        assert_eq!(value["warnings"][0]["document"], 7);
        assert_eq!(value["warnings"][0]["code"], "assemblyCacheUnavailable");
    }
}

#[test]
fn older_save_results_remain_valid_and_omit_empty_warnings() {
    let result: SaveDocumentResult =
        serde_json::from_str(r#"{"fullDocumentName":"a.oad"}"#).unwrap();
    assert!(result.warnings.is_none());
    assert!(
        serde_json::to_value(result)
            .unwrap()
            .get("warnings")
            .is_none()
    );
}
