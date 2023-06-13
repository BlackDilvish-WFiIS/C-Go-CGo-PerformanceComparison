package main

/*
#include <stdio.h>
#include <stdlib.h>

void c_print(const char *buffer) {
    printf("%s\n", buffer);
}
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	start := time.Now()

	cs := C.CString("Hello, world!")
    C.c_print(cs)
	C.free(unsafe.Pointer(cs))

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %.6f seconds\n\n", elapsed.Seconds())
}
