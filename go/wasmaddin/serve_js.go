//go:build js && wasm

// SPDX-License-Identifier: Apache-2.0

package wasmaddin

import (
	"fmt"
	"syscall/js"
)

const hostGlobalName = "__obkHost"

// Host returns the browser worker host bridge object (`globalThis.__obkHost`).
//
//	host := wasmaddin.Host()
//	if !host.IsUndefined() { /* read optional browser channel */ }
func Host() js.Value {
	return js.Global().Get(hostGlobalName)
}

// Serve binds handler to `globalThis.__obkHost` and keeps the Go WASM runtime
// alive for the lifetime of the Web Worker.
//
//	func main() { wasmaddin.Serve(myAddIn) }
func Serve(handler Handler) {
	host := Host()
	if host.IsUndefined() || host.IsNull() {
		panic(fmt.Sprintf("wasmaddin: globalThis.%s is missing, expected Oblikovati add-in worker host", hostGlobalName))
	}
	disp, err := NewDispatcher(handler, jsBridge{host: host})
	if err != nil {
		panic(err.Error())
	}
	exports := buildExports(disp)
	host.Call("register", exports)
	select {}
}

func buildExports(disp *Dispatcher) js.Value {
	obj := js.Global().Get("Object").New()
	obj.Set("id", js.FuncOf(func(js.Value, []js.Value) any { return disp.ID() }))
	obj.Set("manifest", js.FuncOf(func(js.Value, []js.Value) any { return disp.Manifest() }))
	obj.Set("apiMajor", js.FuncOf(func(js.Value, []js.Value) any { return disp.APIMajor() }))
	obj.Set("apiMinor", js.FuncOf(func(js.Value, []js.Value) any { return disp.APIMinor() }))
	obj.Set("hasAutomation", js.FuncOf(func(js.Value, []js.Value) any { return disp.HasAutomation() }))
	// Lifecycle hooks run in a goroutine and return a JS Promise so blocking Go channels
	// or synchronous host RPCs never deadlock Go's single-threaded js/wasm event loop.
	obj.Set("activate", js.FuncOf(func(js.Value, []js.Value) any {
		return runAsync(func() any { return disp.Activate() })
	}))
	obj.Set("deactivate", js.FuncOf(func(js.Value, []js.Value) any {
		return runAsync(func() any { return disp.Deactivate() })
	}))
	obj.Set("notify", js.FuncOf(func(_ js.Value, args []js.Value) any {
		var payload []byte
		if len(args) > 0 {
			payload = bytesFromJS(args[0])
		}
		return runAsync(func() any { return disp.Notify(payload) })
	}))
	obj.Set("automation", js.FuncOf(func(_ js.Value, args []js.Value) any {
		method := ""
		var req []byte
		if len(args) > 0 {
			method = args[0].String()
		}
		if len(args) > 1 {
			req = bytesFromJS(args[1])
		}
		return runAsync(func() any { return invokeAutomation(disp, method, req) })
	}))
	return obj
}

func runAsync(fn func() any) js.Value {
	promiseCtor := js.Global().Get("Promise")
	var executor js.Func
	executor = js.FuncOf(func(_ js.Value, args []js.Value) any {
		resolve := args[0]
		go func() {
			defer executor.Release()
			resolve.Invoke(fn())
		}()
		return nil
	})
	return promiseCtor.New(executor)
}

func invokeAutomation(disp *Dispatcher, method string, req []byte) any {
	resp, err := disp.Automation(method, req)
	result := js.Global().Get("Object").New()
	if err != nil {
		result.Set("ok", false)
		result.Set("data", bytesToJS([]byte(err.Error())))
		return result
	}
	result.Set("ok", true)
	result.Set("data", bytesToJS(resp))
	return result
}

type jsBridge struct {
	host js.Value
}

func (b jsBridge) Call(method string, req []byte) ([]byte, bool) {
	reply := b.host.Call("call", method, bytesToJS(req))
	if reply.IsUndefined() || reply.IsNull() {
		return []byte("wasmaddin: host returned empty reply"), false
	}
	ok := reply.Get("ok").Bool()
	return bytesFromJS(reply.Get("data")), ok
}

func bytesToJS(data []byte) js.Value {
	arr := js.Global().Get("Uint8Array").New(len(data))
	if len(data) > 0 {
		js.CopyBytesToJS(arr, data)
	}
	return arr
}

func bytesFromJS(val js.Value) []byte {
	if val.IsUndefined() || val.IsNull() {
		return nil
	}
	n := val.Get("byteLength").Int()
	if n <= 0 {
		return []byte{}
	}
	out := make([]byte, n)
	js.CopyBytesToGo(out, val)
	return out
}
