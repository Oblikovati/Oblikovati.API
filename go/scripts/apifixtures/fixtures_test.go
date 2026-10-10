// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

type failingFixtureWriter struct{}

func (failingFixtureWriter) Write([]byte) (int, error) { return 0, errors.New("write refused") }

func TestEveryRegisteredTypeProducesInspectableFixtures(t *testing.T) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	expected := 0
	for _, group := range fixtureRegistry {
		if err := writeFixtureGroup(encoder, group); err != nil {
			t.Fatal(err)
		}
		expected += len(group) * 3
	}
	records := bytes.Split(bytes.TrimSpace(buffer.Bytes()), []byte("\n"))
	if len(records) != expected {
		t.Fatalf("records=%d want=%d", len(records), expected)
	}
	assertFixtureRecords(t, records)
}

func assertFixtureRecords(t *testing.T, records [][]byte) {
	t.Helper()
	for _, encoded := range records {
		var record fixtureRecord
		if err := json.Unmarshal(encoded, &record); err != nil {
			t.Fatal(err)
		}
		if record.Name == "" || !json.Valid(record.Payload) {
			t.Fatalf("invalid record %s", encoded)
		}
		if record.Error == "" && !json.Valid(record.Canonical) {
			t.Fatalf("missing canonical JSON: %s", encoded)
		}
	}
}

func TestFixtureFailuresPropagate(t *testing.T) {
	encoder := json.NewEncoder(failingFixtureWriter{})
	if err := writeFixtureGroup(encoder, fixtureRegistry[0]); err == nil {
		t.Fatal("expected write error")
	}
	unsupported := fixtureType{"channel", reflect.TypeFor[chan int]()}
	if err := writeFixture(encoder, unsupported, "zero"); err == nil {
		t.Fatal("expected marshal error")
	}
}

func TestEdgeFixturesUseTheGoCodecAsTheOracle(t *testing.T) {
	var buffer bytes.Buffer
	if err := writeEdgeFixtures(json.NewEncoder(&buffer)); err != nil {
		t.Fatal(err)
	}
	assertFixtureRecords(t, bytes.Split(bytes.TrimSpace(buffer.Bytes()), []byte("\n")))
	if err := writeEdgeFixtures(json.NewEncoder(failingFixtureWriter{})); err == nil {
		t.Fatal("expected write error")
	}
}

func TestPopulationDepthAndInaccessibleValuesStayUnchanged(t *testing.T) {
	number := 0
	populateValue(reflect.ValueOf(&number).Elem(), "filled", 6)
	if number != 0 {
		t.Fatalf("depth limit ignored: %d", number)
	}
	populateValue(reflect.ValueOf(number), "filled", 0)
}
