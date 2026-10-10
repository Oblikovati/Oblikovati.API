// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"oblikovati.org/api/types"
	"reflect"
	"time"
)

type edgeFixture struct {
	value    fixtureType
	payloads []string
}

func writeEdgeFixtures(encoder *json.Encoder) error {
	cases := []edgeFixture{
		{fixtureType{"std.Bytes", reflect.TypeFor[[]byte]()}, []string{`"AB=="`, `"AA\r\nH/"`, `[0,255]`, `null`, `""`, `"?"`, `[256]`}},
		{fixtureType{"std.Time", reflect.TypeFor[time.Time]()}, []string{`"2016-12-31T23:59:60Z"`, `"2026-10-10T12:30:00.1000Z"`, `"2026-10-10T12:30:00+02:00"`, `null`, `"bad"`}},
		{fixtureType{"types.Point", reflect.TypeFor[types.Point]()}, []string{`[null,1,2]`, `[1,2]`, `[1,2,3,4]`, `null`}},
		{fixtureType{"types.UnitVector", reflect.TypeFor[types.UnitVector]()}, []string{`[null,3,4]`, `[0,0,0]`, `[1e-12,0,0]`, `[1e-13,0,0]`}},
		{fixtureType{"types.Variant", reflect.TypeFor[types.Variant]()}, []string{`{"type":"integer","value":9223372036854775807}`, `{"type":"integer","value":null}`, `{"type":"integer"}`, `{"type":"bytes","value":"AB=="}`, `{"type":"unknown","value":1}`}},
	}
	for _, item := range cases {
		for index, payload := range item.payloads {
			if err := writeEncodedFixture(encoder, item.value, fmt.Sprintf("edge-%d", index), []byte(payload)); err != nil {
				return err
			}
		}
	}
	return nil
}
