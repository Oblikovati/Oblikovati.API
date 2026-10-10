// SPDX-License-Identifier: Apache-2.0

package client_test

import (
	"encoding/json"
	"errors"
	"oblikovati.org/api/client"
	"reflect"
	"testing"
)

var errParityHost = errors.New("parity host reached")

type parityCaller struct {
	test   *testing.T
	method string
}

func (host parityCaller) Call(method string, request []byte) ([]byte, error) {
	host.test.Helper()
	if method != host.method {
		host.test.Fatalf("method=%q expected=%q", method, host.method)
	}
	if len(request) > 0 && !json.Valid(request) {
		host.test.Fatalf("request=%q expected JSON", request)
	}
	return nil, errParityHost
}

func checkGoMethod(t *testing.T, method reflect.Value) {
	t.Helper()
	results := invokeReflected(method)
	failure, ok := results[len(results)-1].Interface().(error)
	if !ok || !errors.Is(failure, errParityHost) {
		t.Fatalf("client did not reach the expected host: %v", failure)
	}
}

func zeroArguments(method reflect.Value) []reflect.Value {
	args := make([]reflect.Value, method.Type().NumIn())
	for index := range args {
		args[index] = reflect.Zero(method.Type().In(index))
	}
	return args
}

func invokeReflected(method reflect.Value) []reflect.Value {
	if method.Type().IsVariadic() {
		return method.CallSlice(zeroArguments(method))
	}
	return method.Call(zeroArguments(method))
}

func checkGoGroup(t *testing.T, c *client.Client, group, method string) {
	t.Helper()
	if group == "Client" {
		checkGoMethod(t, reflect.ValueOf(c).MethodByName(method))
		return
	}
	queue := []reflect.Value{reflect.ValueOf(c)}
	seen := make(map[reflect.Type]bool)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if seen[current.Type()] {
			continue
		}
		seen[current.Type()] = true
		if current.Type().Name() == group {
			checkGoMethod(t, current.MethodByName(method))
			return
		}
		queue = append(queue, clientGroups(current)...)
	}
	t.Fatalf("group %q: expected a public client group", group)
}

func clientGroups(current reflect.Value) []reflect.Value {
	var groups []reflect.Value
	for index := 0; index < current.NumMethod(); index++ {
		method := current.Method(index)
		if method.Type().NumOut() != 1 {
			continue
		}
		result := method.Type().Out(0)
		if result.Kind() != reflect.Struct || result.PkgPath() != "oblikovati.org/api/client" {
			continue
		}
		groups = append(groups, invokeReflected(method)[0])
	}
	return groups
}
