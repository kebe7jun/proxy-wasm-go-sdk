package main

import (
	"strconv"

	"github.com/mathetake/proxy-wasm-go/runtime"
	"github.com/mathetake/proxy-wasm-go/runtime/types"
)

func main() {
	runtime.SetNewHttpContext(newContext)
}

type httpHeaders struct {
	// you must embed the default context so that you need not to reimplement all the methods by yourself
	runtime.DefaultContext
	contextID uint32
}

func newContext(contextID uint32) runtime.HttpContext {
	return &httpHeaders{contextID: contextID}
}

// override
func (ctx *httpHeaders) OnHttpRequestHeaders(_ int, _ bool) types.Action {
	hs, err := ctx.GetHttpRequestHeaders()
	if err != nil {
		runtime.LogCritical("failed to get request headers: " + err.Error())
	}

	for _, h := range hs {
		runtime.LogInfo("request header: " + h[0] + ": " + h[1])
	}
	return types.ActionContinue
}

// override
func (ctx *httpHeaders) OnHttpResponseHeaders(_ int, _ bool) types.Action {
	hs, err := ctx.GetHttpResponseHeaders()
	if err != nil {
		runtime.LogCritical("failed to get request headers: " + err.Error())
	}

	for _, h := range hs {
		runtime.LogInfo("response header: " + h[0] + ": " + h[1])
	}
	return types.ActionContinue
}

// override
func (ctx *httpHeaders) OnLog() {
	runtime.LogInfo(strconv.FormatUint(uint64(ctx.contextID), 10) + " finished")
}