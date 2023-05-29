package main

import (
	"fmt"
	"time"
	"unsafe"
)

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

func main() {
	filename := C.CString("../input.txt")
	defer C.free(unsafe.Pointer(filename))

	// Call the C function to measure execution time
	start := time.Now()
	file := C.fopen(filename, C.CString("r"))
	if file == nil {
		fmt.Printf("Failed to open file: %s\n", filename)
		return
	}
	defer C.fclose(file)

	var buff [255]C.char
	for {
		result := C.fgets(&buff[0], C.int(255), file)
		if result == nil {
			break // Reached end of file or encountered an error
		}
		fmt.Print(C.GoString(&buff[0]))
	}

	elapsed := time.Since(start)

	fmt.Printf("Execution time: %.6f seconds\n\n", elapsed.Seconds())
}
