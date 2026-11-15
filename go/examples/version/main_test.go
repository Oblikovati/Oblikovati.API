// SPDX-License-Identifier: Apache-2.0

package main

import (
	"io"
	"oblikovati.org/api"
	"oblikovati.org/api/client"
	"os"
	"testing"
)

func TestVersionExampleUsesSharedContract(t *testing.T) {
	got, err := client.New(exampleHost{}).Application().ApiVersion()
	if err != nil || got.Version != api.Version {
		t.Fatalf("version = %+v, %v; expected %s", got, err, api.Version)
	}
	if _, err := (exampleHost{}).Call("bad", nil); err == nil {
		t.Fatal("bad method should fail")
	}
}

func TestMainPrintsContractVersion(t *testing.T) {
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	original := os.Stdout
	os.Stdout = writer
	defer func() { os.Stdout = original }()
	main()
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	if err := reader.Close(); err != nil {
		t.Fatal(err)
	}
	if string(output) != api.Version+"\n" {
		t.Fatalf("output = %q; expected version %q", output, api.Version)
	}
}
