package main

import (
	"fmt"
	"time"
	"unsafe"
)

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"

func test(nStrings int) {
	fmt.Printf("Number of appends: %d\n", nStrings)
	start := time.Now()

	buff := (*C.char)(C.malloc(C.size_t(nStrings*4 + 10)))
	str := C.CString("test")

	for i := 0; i < nStrings; i++ {
		C.strcat(buff, str)
	}

	elapsed := time.Since(start)

	result := C.GoString(buff)
	fmt.Printf("Result length: %d\n", len(result))
	fmt.Printf("Execution time: %.6f seconds\n\n", elapsed.Seconds())

	C.free(unsafe.Pointer(buff))
}

func main() {
	test(10000)
	test(50000)
	test(100000)
}
