package main

import "github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"

type streamContext struct {
	// we must embed the default context so that you need not to reimplement all the methods by yourself
	proxywasm.DefaultStreamContext
	contextID uint32
}

//NewStreamContext Override DefaultRootContext.
func (ctx *rootContext) NewStreamContext(contextID uint32) proxywasm.StreamContext {
	return &streamContext{contextID: contextID}
}
