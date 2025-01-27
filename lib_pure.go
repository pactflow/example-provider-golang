//go:build !cgo
// +build !cgo

// Package native contains the c bindings into the Pact Reference types.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ebitengine/purego"
)

func getSystemLibrary() string {
	switch runtime.GOOS {
	case "darwin":
		return "libpact_ffi.dylib"
	case "linux":
		return "libpact_ffi.so"
	case "windows":
		return "pact_ffi.dll"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

type (
	size_t uintptr
)

var pactffi_verifier_new func() uintptr
var pactffi_verifier_shutdown func(uintptr)
var pactffi_verifier_set_provider_info func(uintptr, string, string, string, uint16, string)
var pactffi_verifier_add_file_source func(uintptr, string)
var pactffi_verifier_execute func(uintptr) int32

func init() {
	libpact_ffi, err := openLibrary(filepath.Join(os.Getenv("PACT_DOWNLOAD_DIR"), getSystemLibrary()))
	if err != nil {
		panic(err)
	}
	purego.RegisterLibFunc(&pactffi_verifier_new, libpact_ffi, "pactffi_verifier_new")
	purego.RegisterLibFunc(&pactffi_verifier_set_provider_info, libpact_ffi, "pactffi_verifier_set_provider_info")
	purego.RegisterLibFunc(&pactffi_verifier_add_file_source, libpact_ffi, "pactffi_verifier_add_file_source")
	purego.RegisterLibFunc(&pactffi_verifier_execute, libpact_ffi, "pactffi_verifier_execute")
	purego.RegisterLibFunc(&pactffi_verifier_shutdown, libpact_ffi, "pactffi_verifier_shutdown")
}

func test_provider(port int) int {
	verifier := pactffi_verifier_new()
	pactffi_verifier_set_provider_info(verifier, "pactflow-example-provider-golang", "http", "localhost", uint16(port), "/")
	pactFile := os.Getenv("PACT_FILE")
	if pactFile == "" {
		pactFile = "pact.json"
	}
	pactffi_verifier_add_file_source(verifier, pactFile)
	result := pactffi_verifier_execute(verifier)
	pactffi_verifier_shutdown(verifier)
	if result != 0 {
		fmt.Printf("Result is not 0: %d", result)
	} else {
		fmt.Print("Result success")
	}
	return int(result)
}