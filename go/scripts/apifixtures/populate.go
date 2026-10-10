// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"oblikovati.org/api/types"
	"reflect"
	"time"
)

func populateValue(value reflect.Value, mode string, depth int) {
	if !value.CanSet() || depth > 5 {
		return
	}
	if populateCustom(value) {
		return
	}
	switch value.Kind() {
	case reflect.Bool:
		value.SetBool(true)
	case reflect.String:
		value.SetString("example <&> 日本語")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value.SetInt(7)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value.SetUint(7)
	case reflect.Float32, reflect.Float64:
		value.SetFloat(1.25)
	case reflect.Struct:
		populateFields(value, mode, depth)
	default:
		populateCollection(value, mode, depth)
	}
}

func populateCustom(value reflect.Value) bool {
	switch value.Type() {
	case reflect.TypeFor[types.Variant]():
		value.Set(reflect.ValueOf(types.UnitVariant(2.5, "mm")))
	case reflect.TypeFor[types.UnitVector]():
		value.Set(reflect.ValueOf(types.UnitVector{X: 1}))
	case reflect.TypeFor[types.UnitVector2d]():
		value.Set(reflect.ValueOf(types.UnitVector2d{X: 1}))
	case reflect.TypeFor[json.RawMessage]():
		value.SetBytes([]byte(`{"opaque":18446744073709551615}`))
	case reflect.TypeFor[time.Time]():
		value.Set(reflect.ValueOf(time.Date(2026, 10, 10, 12, 30, 0, 123456789, time.UTC)))
	default:
		return false
	}
	return true
}

func populateFields(value reflect.Value, mode string, depth int) {
	for index := 0; index < value.NumField(); index++ {
		populateValue(value.Field(index), mode, depth+1)
	}
}

func populateCollection(value reflect.Value, mode string, depth int) {
	switch value.Kind() {
	case reflect.Pointer:
		value.Set(reflect.New(value.Type().Elem()))
		populateValue(value.Elem(), mode, depth+1)
	case reflect.Array:
		for index := 0; index < value.Len(); index++ {
			populateValue(value.Index(index), mode, depth+1)
		}
	case reflect.Slice:
		length := 1
		if mode == "empty" {
			length = 0
		}
		value.Set(reflect.MakeSlice(value.Type(), length, length))
		for index := 0; index < length; index++ {
			populateValue(value.Index(index), mode, depth+1)
		}
	case reflect.Map:
		populateMap(value, mode, depth)
	}
}

func populateMap(value reflect.Value, mode string, depth int) {
	value.Set(reflect.MakeMap(value.Type()))
	if mode == "empty" {
		return
	}
	key := reflect.New(value.Type().Key()).Elem()
	item := reflect.New(value.Type().Elem()).Elem()
	populateValue(key, mode, depth+1)
	populateValue(item, mode, depth+1)
	value.SetMapIndex(key, item)
}
