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
#include <time.h>

void swap(char* a, char* b) {
    char temp = *a;
    *a = *b;
    *b = temp;
}

void permute(char* str, int l, int r, int* count) {
    if (l == r) {
        (*count)++;
    } else {
        for (int i = l; i <= r; i++) {
            swap((str + l), (str + i));
            permute(str, l + 1, r, count);
            swap((str + l), (str + i)); // backtrack
        }
    }
}

void calculate_permutations(char* str, int* count) {
    int len = strlen(str);
    permute(str, 0, len - 1, count);
}
*/
import "C"

func calculatePermutations(str string) int {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	var count C.int
	C.calculate_permutations(cStr, &count)

	return int(count)
}

func test(str string) {
	fmt.Printf("String to permute: %s, len: %d\n", str, len(str))
	count := 0

	startTime := time.Now()
	count = calculatePermutations(str)
	endTime := time.Now()

	executionTime := endTime.Sub(startTime)

	fmt.Println("Permutations:", count)
	fmt.Printf("Execution time: %.6f seconds\n\n", executionTime.Seconds())
}

func main() {
	test("abcdefghij")
	test("abcdefghijk")
	test("abcdefghijkl")
}
