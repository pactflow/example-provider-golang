//go:build (darwin || linux) && !cgo
// +build darwin linux
// +build !cgo

package main

import "github.com/ebitengine/purego"

func openLibrary(name string) (uintptr, error) {
	return purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}
