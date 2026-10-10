// SPDX-License-Identifier: Apache-2.0

use oblikovati_api::{
    Error,
    client::{Caller, Client},
    wire,
    wire::featureargs,
};
use serde_json::json;

mod feature_kinds;

struct ReplyHost {
    method: &'static str,
    request: serde_json::Value,
    reply: &'static str,
}

impl Caller for ReplyHost {
    fn call(&self, method: &str, request: &[u8]) -> Result<Vec<u8>, Error> {
        assert_eq!(method, self.method);
        let actual = if request.is_empty() {
            serde_json::Value::Null
        } else {
            serde_json::from_slice(request).unwrap()
        };
        assert_eq!(actual, self.request);
        Ok(self.reply.as_bytes().to_vec())
    }
}

#[test]
fn client_decodes_typed_results_and_reports_invalid_replies() {
    let client = Client::new(ReplyHost {
        method: wire::METHOD_APPLICATION_API_VERSION,
        request: json!(null),
        reply: r#"{"version":"1.2.3","major":1,"minor":2}"#,
    });
    assert_eq!(client.application().api_version().unwrap().minor, 2);
    let broken = Client::new(ReplyHost {
        method: wire::METHOD_APPLICATION_API_VERSION,
        request: json!(null),
        reply: "broken",
    });
    let failure = broken.application().api_version().unwrap_err();
    assert!(failure.to_string().contains("broken"));
    assert!(
        failure
            .to_string()
            .contains(wire::METHOD_APPLICATION_API_VERSION)
    );
    let _: &dyn std::error::Error = &failure;
}

#[test]
fn typed_features_build_the_same_envelope_as_go() {
    let client = Client::new(ReplyHost {
        method: wire::METHOD_FEATURES_ADD,
        request: json!({"kind":"extrude", "args":{"sketchIndex":0,"profileIndex":0,"distance":"50 mm","operation":"new"}}),
        reply: "{}",
    });
    let args = featureargs::Extrude {
        distance: "50 mm".into(),
        operation: "new".into(),
        ..Default::default()
    };
    assert!(client.features().add_feature(&args).is_ok());
}

#[test]
fn invalid_requests_fail_before_host_io() {
    let client = Client::new(ReplyHost {
        method: "must not be called",
        request: json!(null),
        reply: "",
    });
    let args = featureargs::Grill {
        draft: f64::NAN,
        ..Default::default()
    };
    assert!(
        client
            .features()
            .add_feature(&args)
            .unwrap_err()
            .0
            .contains("grill")
    );
    let args = wire::CreateWorkPointArgs {
        at: Some(vec![f64::NAN]),
        ..Default::default()
    };
    assert!(client.work_points().create(&args).is_err());
}
