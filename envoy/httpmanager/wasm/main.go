package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

// Other examples can be found at https://github.com/tetratelabs/proxy-wasm-go-sdk/tree/v0.1.1/examples

func main() {
	proxywasm.SetNewRootContext(newRootContext)
}

type rootContext struct {
	// You'd better embed the default root context
	// so that you don't need to reimplement all the methods by yourself.
	proxywasm.DefaultRootContext
}

func newRootContext(uint32) proxywasm.RootContext { return &rootContext{} }
