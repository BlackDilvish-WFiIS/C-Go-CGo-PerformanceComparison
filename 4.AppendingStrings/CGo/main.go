package main

import (
	"fmt"
	"log"
	"time"
	"unsafe"
)

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

double measureExecutionTime(const char* filename) {
	clock_t start = clock();
	char buff[50000];
    buff[0] = '\n';
    const char* str = "test";

    for (int i = 0; i < 10000; i++) {
        strcat(buff, str);
    }

	clock_t end = clock();
	printf("Result length: %zu\n", strlen(buff));
	return (double)(end - start) / CLOCKS_PER_SEC;
}
*/
import "C"

func main() {
	filename := C.CString("input2.txt")
	defer C.free(unsafe.Pointer(filename))

	start := time.Now()
	executionTime := C.measureExecutionTime(filename)
	if executionTime < 0 {
		log.Fatal("Error opening file.")
	}
	elapsed := time.Since(start)

	// Print the execution time
	fmt.Printf("Execution time: %.6f seconds.\n", executionTime)

	fmt.Printf("Execution time: %.6f seconds\n", elapsed.Seconds())
}
