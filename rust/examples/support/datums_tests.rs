// SPDX-License-Identifier: Apache-2.0
use super::*;

struct DatumListHost {
    method: &'static str,
    response: &'static str,
}
impl Caller for DatumListHost {
    fn call(&self, method: &str, request: &[u8]) -> Result<Vec<u8>, Error> {
        assert_eq!(method, self.method);
        assert_eq!(request, br#"{"includeConstruction":true}"#);
        Ok(self.response.as_bytes().to_vec())
    }
}
#[test]
fn axes_example_includes_construction_rows() {
    let host = DatumListHost {
        method: wire::METHOD_WORK_AXES_LIST,
        response: r#"{"axes":[{"index":4,"construction":true}]}"#,
    };
    let rows = list_all_axes(host).unwrap().axes.unwrap();
    assert!(rows[0].construction);
    assert_eq!(rows[0].index, 4);
}
#[test]
fn points_example_includes_construction_rows() {
    let host = DatumListHost {
        method: wire::METHOD_WORK_POINTS_LIST,
        response: r#"{"points":[{"index":2,"construction":true}]}"#,
    };
    let rows = list_all_points(host).unwrap().points.unwrap();
    assert!(rows[0].construction);
    assert_eq!(rows[0].index, 2);
}

#[test]
fn planes_example_includes_construction_rows() {
    let host = DatumListHost {
        method: wire::METHOD_WORK_PLANES_LIST,
        response: r#"{"planes":[{"index":4,"construction":true}]}"#,
    };
    let rows = list_all_planes(host).unwrap().planes.unwrap();
    assert!(rows[0].construction);
    assert_eq!(rows[0].index, 4);
}
