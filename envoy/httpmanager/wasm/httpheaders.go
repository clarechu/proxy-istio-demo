package main

import (
	"fmt"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

type httpHeaders struct {
	// we must embed the default context so that you need not to reimplement all the methods by yourself
	proxywasm.DefaultHttpContext
	contextID uint32
}

//NewHttpContext Override DefaultRootContext.
func (ctx *rootContext) NewHttpContext(contextID uint32) proxywasm.HttpContext {
	return &httpHeaders{contextID: contextID}
}

//OnHttpRequestHeaders Override DefaultHttpContext.
func (ctx *httpHeaders) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	hs, err := proxywasm.GetHttpRequestHeaders()
	if err != nil {
		proxywasm.LogCriticalf("failed to get request headers: %v", err)
	}

	for _, h := range hs {
		proxywasm.LogInfof("request header: %s: %s", h[0], h[1])
	}
	return types.ActionContinue
}

//OnHttpResponseHeaders Override DefaultHttpContext.
func (ctx *httpHeaders) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	if err := proxywasm.SetHttpResponseHeader("hello1", "world"); err != nil {
		proxywasm.LogCriticalf("failed to set response header: %v", err)
		return types.ActionContinue
	}
	proxywasm.LogCritical("set response http header success ..")
	return types.ActionContinue
}

//OnHttpStreamDone Override DefaultHttpContext.
func (ctx *httpHeaders) OnHttpStreamDone() {
	proxywasm.LogInfof("%d finished", ctx.contextID)
}

//OnHttpRequestBody Override DefaultHttpContext.
func (ctx *httpHeaders) OnHttpRequestBody(numHeaders int, endOfStream bool) types.Action {
	if endOfStream {
		return types.ActionContinue
	}
	start := 0
	maxsize := 10
	body := make([]byte, 0)
	for {
		hs, err := proxywasm.GetHttpRequestBody(start, 10)
		start = start + maxsize
		if err != nil {
			proxywasm.LogCriticalf("failed to get request body: %v", err)
			break
		}
		body = append(body, hs...)
		fmt.Printf("len:%d, start:%d, debug:%+v \n", len(hs), start, string(hs))
		if len(hs) < 10 {
			break
		}
	}
	fmt.Println("body ----> ", string(body))
	return types.ActionContinue
}
