//go:build libpact_cgo
// +build libpact_cgo

package main

import (
	"testing"

	"github.com/pact-foundation/pact-go/v2/utils"
)

func TestLibPactFfiProvider(t *testing.T) {
	var port, _ = utils.GetFreePort()
	go startProvider(port)

	var res = test_provider(port)
	if res != 0 {
		t.Fatalf("%v", res)
	}

}
