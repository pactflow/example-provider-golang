//go:build cgo
// +build cgo

// Package native contains the c bindings into the Pact Reference types.
package main

/*
#cgo darwin,arm64 LDFLAGS: -L/tmp -L/usr/local/lib -Wl,-rpath -Wl,/tmp -Wl,-rpath -Wl,/usr/local/lib -lpact_ffi
#cgo darwin,amd64 LDFLAGS: -L/tmp -L/usr/local/lib -Wl,-rpath -Wl,/tmp -Wl,-rpath -Wl,/usr/local/lib -lpact_ffi
#cgo windows,amd64 LDFLAGS: -lpact_ffi
#cgo linux,amd64 LDFLAGS: -L/tmp -L/opt/pact/lib -L/usr/local/lib -Wl,-rpath -Wl,/opt/pact/lib -Wl,-rpath -Wl,/tmp -Wl,-rpath -Wl,/usr/local/lib -lpact_ffi
#cgo linux,arm64 LDFLAGS: -L/tmp -L/opt/pact/lib -L/usr/local/lib -Wl,-rpath -Wl,/opt/pact/lib -Wl,-rpath -Wl,/tmp -Wl,-rpath -Wl,/usr/local/lib -lpact_ffi
#include "pact.h"
*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

func test_provider(port int) int {
	verifier := C.pactffi_verifier_new()
	C.pactffi_verifier_set_provider_info(verifier, C.CString("pactflow-example-provider-golang"), C.CString("http"), C.CString("localhost"), C.ushort(port), C.CString("/"))
	C.pactffi_verifier_add_file_source(verifier, C.CString(os.Getenv("PACT_FILE")))

	defer C.pactffi_verifier_shutdown(verifier)
	result := C.pactffi_verifier_execute(verifier)
	if result != 0 {
		fmt.Printf("Result is not 0: %d", result)
	} else {
		fmt.Print("Result success")
	}
	return int(result)
}


type interactionPart int

const (
	INTERACTION_PART_REQUEST interactionPart = iota
	INTERACTION_PART_RESPONSE
)


func free(str *C.char) {
	C.free(unsafe.Pointer(str))
}
