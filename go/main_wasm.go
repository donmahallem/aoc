//go:build wasip1

package main

import (
	"strings"
	"unsafe"
)

// main is required by the compiler
func main() {}

//go:wasmexport malloc
func malloc(size uint32) unsafe.Pointer {
	buf := make([]byte, size)
	return unsafe.Pointer(&buf[0])
}

//go:wasmexport free
func free(ptr unsafe.Pointer, size uint32) {
	// was recommended to include free, but Go has gc
}

//go:wasmexport solveAOC
func solveAOC(year, day, part uint32, inputPtr unsafe.Pointer, inputLen uint32) uint64 {
	// Reconstruct string from WASM memory
	inputData := unsafe.String((*byte)(inputPtr), inputLen)

	res := RunSolver(int(year), int(day), int(part), strings.NewReader(inputData))

	var output string
	if res.Error != nil {
		output = "ERROR: " + res.Error.Error()
	} else {
		output = res.Result
	}

	// Convert result to bytes to get a pointer
	outBytes := []byte(output)
	if len(outBytes) == 0 {
		return 0
	}

	ptr := unsafe.Pointer(&outBytes[0])
	length := uint32(len(outBytes))

	// Pack pointer (high 32) and length (low 32) into a uint64
	return (uint64(uintptr(ptr)) << 32) | uint64(length)
}
