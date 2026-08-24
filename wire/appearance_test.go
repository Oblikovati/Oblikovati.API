// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/types"
)

// openPBRSurfaceParams bundles one instance of every OpenPBR parameter group at its
// spec default (plus a non-nil Geometry.Normal, to exercise the pointer field), for
// building a fully-populated test fixture DTO.
type openPBRSurfaceParams struct {
	Base         types.OpenPBRBase
	Specular     types.OpenPBRSpecular
	Transmission types.OpenPBRTransmission
	Subsurface   types.OpenPBRSubsurface
	Coat         types.OpenPBRCoat
	Fuzz         types.OpenPBRFuzz
	ThinFilm     types.OpenPBRThinFilm
	Emission     types.OpenPBREmission
	Geometry     types.OpenPBRGeometry
}

func fullOpenPBRSurfaceParams() openPBRSurfaceParams {
	n := types.NewVector(0, 0, 1)
	p := openPBRSurfaceParams{
		Base: types.DefaultOpenPBRBase(), Specular: types.DefaultOpenPBRSpecular(),
		Transmission: types.DefaultOpenPBRTransmission(), Subsurface: types.DefaultOpenPBRSubsurface(),
		Coat: types.DefaultOpenPBRCoat(), Fuzz: types.DefaultOpenPBRFuzz(),
		ThinFilm: types.DefaultOpenPBRThinFilm(), Emission: types.DefaultOpenPBREmission(),
		Geometry: types.DefaultOpenPBRGeometry(),
	}
	p.Geometry.Normal = &n
	return p
}

// TestAppearanceInfoRoundTrip checks every group of an AppearanceInfo — one
// non-default value per field, including a non-nil Geometry.Normal — survives a JSON
// round-trip intact.
func TestAppearanceInfoRoundTrip(t *testing.T) {
	p := fullOpenPBRSurfaceParams()
	want := AppearanceInfo{
		ID: "brushed-steel", DisplayName: "Brushed Steel", Source: "builtin",
		Base: p.Base, Specular: p.Specular, Transmission: p.Transmission,
		Subsurface: p.Subsurface, Coat: p.Coat, Fuzz: p.Fuzz, ThinFilm: p.ThinFilm,
		Emission: p.Emission, Geometry: p.Geometry,
	}

	b, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	var got AppearanceInfo
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}

	if got.ID != want.ID || got.DisplayName != want.DisplayName || got.Source != want.Source {
		t.Errorf("identity round-trip = %+v, want %+v", got, want)
	}
	if got.Base != want.Base {
		t.Errorf("Base round-trip = %+v, want %+v", got.Base, want.Base)
	}
	if got.Specular != want.Specular {
		t.Errorf("Specular round-trip = %+v, want %+v", got.Specular, want.Specular)
	}
	if got.Transmission != want.Transmission {
		t.Errorf("Transmission round-trip = %+v, want %+v", got.Transmission, want.Transmission)
	}
	if got.Subsurface != want.Subsurface {
		t.Errorf("Subsurface round-trip = %+v, want %+v", got.Subsurface, want.Subsurface)
	}
	if got.Coat != want.Coat {
		t.Errorf("Coat round-trip = %+v, want %+v", got.Coat, want.Coat)
	}
	if got.Fuzz != want.Fuzz {
		t.Errorf("Fuzz round-trip = %+v, want %+v", got.Fuzz, want.Fuzz)
	}
	if got.ThinFilm != want.ThinFilm {
		t.Errorf("ThinFilm round-trip = %+v, want %+v", got.ThinFilm, want.ThinFilm)
	}
	if got.Emission != want.Emission {
		t.Errorf("Emission round-trip = %+v, want %+v", got.Emission, want.Emission)
	}
	if got.Geometry.Opacity != want.Geometry.Opacity || got.Geometry.ThinWalled != want.Geometry.ThinWalled {
		t.Errorf("Geometry scalar round-trip = %+v, want %+v", got.Geometry, want.Geometry)
	}
	if got.Geometry.Normal == nil || *got.Geometry.Normal != *want.Geometry.Normal {
		t.Errorf("Geometry.Normal round-trip = %v, want %v", got.Geometry.Normal, want.Geometry.Normal)
	}
}

// TestAppearanceArgsRoundTrip checks the create/update/assign request shapes
// survive a JSON round-trip.
func TestAppearanceArgsRoundTrip(t *testing.T) {
	create := CreateAppearanceArgs{BaseID: "brushed-steel", Name: "My Steel"}
	var gotCreate CreateAppearanceArgs
	b, _ := json.Marshal(create)
	if err := json.Unmarshal(b, &gotCreate); err != nil || gotCreate != create {
		t.Errorf("CreateAppearanceArgs round-trip = %+v, want %+v (err %v)", gotCreate, create, err)
	}

	p := fullOpenPBRSurfaceParams()
	update := UpdateAppearanceArgs{
		ID: "my-steel", DisplayName: "My Steel", Base: p.Base, Specular: p.Specular,
		Transmission: p.Transmission, Subsurface: p.Subsurface, Coat: p.Coat, Fuzz: p.Fuzz,
		ThinFilm: p.ThinFilm, Emission: p.Emission, Geometry: p.Geometry,
	}
	var gotUpdate UpdateAppearanceArgs
	b, _ = json.Marshal(update)
	if err := json.Unmarshal(b, &gotUpdate); err != nil {
		t.Fatal(err)
	}
	if gotUpdate.ID != update.ID || gotUpdate.DisplayName != update.DisplayName ||
		gotUpdate.Base != update.Base || gotUpdate.Coat != update.Coat {
		t.Errorf("UpdateAppearanceArgs round-trip = %+v, want %+v", gotUpdate, update)
	}

	assign := AssignAppearanceArgs{Scope: "body", Key: "body/0", AppearanceID: "my-steel"}
	var gotAssign AssignAppearanceArgs
	b, _ = json.Marshal(assign)
	if err := json.Unmarshal(b, &gotAssign); err != nil || gotAssign != assign {
		t.Errorf("AssignAppearanceArgs round-trip = %+v, want %+v (err %v)", gotAssign, assign, err)
	}
}
