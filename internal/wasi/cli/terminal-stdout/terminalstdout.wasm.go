// Code generated by wit-bindgen-go. DO NOT EDIT.

package terminalstdout

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasi:cli@0.2.0".

//go:wasmimport wasi:cli/terminal-stdout@0.2.0 get-terminal-stdout
//go:noescape
func wasmimport_GetTerminalStdout(result *cm.Option[TerminalOutput])
