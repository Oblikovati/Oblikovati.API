// SPDX-License-Identifier: Apache-2.0

// Command apifixtures emits compiler-backed JSON examples for the Rust parity tests.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type fixtureType struct {
	name string
	typ  reflect.Type
}
type fixtureRecord struct {
	Name      string
	Case      string
	Payload   json.RawMessage
	Canonical json.RawMessage
	Error     string
}

func main() {
	encoder := json.NewEncoder(os.Stdout)
	for _, group := range fixtureRegistry {
		if err := writeFixtureGroup(encoder, group); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	if err := writeEdgeFixtures(encoder); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func writeFixtureGroup(encoder *json.Encoder, group []fixtureType) error {
	for _, item := range group {
		for _, mode := range []string{"zero", "filled", "empty"} {
			if err := writeFixture(encoder, item, mode); err != nil {
				return err
			}
		}
	}
	return nil
}

func writeFixture(encoder *json.Encoder, item fixtureType, mode string) error {
	value := reflect.New(item.typ).Elem()
	if mode != "zero" {
		populateValue(value, mode, 0)
	}
	payload, err := json.Marshal(value.Interface())
	if err != nil {
		return fmt.Errorf("fixture %s/%s: %w", item.name, mode, err)
	}
	return writeEncodedFixture(encoder, item, mode, payload)
}

func writeEncodedFixture(encoder *json.Encoder, item fixtureType, mode string, payload []byte) error {
	record := fixtureRecord{Name: item.name, Case: mode, Payload: payload}
	decoded := reflect.New(item.typ)
	if err := json.Unmarshal(payload, decoded.Interface()); err != nil {
		record.Error = err.Error()
		return encoder.Encode(record)
	}
	canonical, err := json.Marshal(decoded.Interface())
	if err != nil {
		return err
	}
	record.Canonical = canonical
	return encoder.Encode(record)
}
