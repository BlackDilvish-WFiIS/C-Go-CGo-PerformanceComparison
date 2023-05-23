package main

import (
	"fmt"
	"time"
)

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"

func main() {
	start := time.Now()

	var buff [50000]C.char
	str := C.CString("test")

	for i := 0; i < 10000; i++ {
		C.strcat(&buff[0], str)
	}

	elapsed := time.Since(start)

	result := C.GoString(&buff[0])
	fmt.Printf("Result length: %d\n", len(result))
	fmt.Printf("Execution time: %.6f seconds\n", elapsed.Seconds())
}
