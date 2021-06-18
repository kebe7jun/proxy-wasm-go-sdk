// Copyright 2020-2021 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"math/rand"
	"time"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func main() {
	proxywasm.SetNewRootContextFn(newRootContext)
}

type rootContext struct {
	// Embed the default root context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultRootContext
}

func newRootContext(uint32) types.RootContext { return &rootContext{} }

// Override DefaultRootContext.
func (*rootContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpRouting{}
}

type httpRouting struct {
	// Embed the default http context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultHttpContext
}

// Unittest purpose.
var now = func() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}

// Override DefaultHttpContext.
func (ctx *httpRouting) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	// Randomly routing to the canary cluster.
	dice := now()
	proxywasm.LogInfof("dice: %d\n", dice)
	if dice%2 == 0 {
		const authorityKey = ":authority"
		value, err := proxywasm.GetHttpRequestHeader(authorityKey)
		if err != nil {
			proxywasm.LogCritical("failed to get request header: ':authority'")
			return types.ActionPause
		}
		// Append "-canary" suffix to route this request to the canary cluster.
		value += "-canary"
		if err := proxywasm.ReplaceHttpRequestHeader(":authority", value); err != nil {
			proxywasm.LogCritical("failed to set request header: test")
			return types.ActionPause
		}
	}
	return types.ActionContinue
}