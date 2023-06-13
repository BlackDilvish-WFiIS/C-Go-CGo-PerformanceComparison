package main

// #include <stdlib.h>
// #include <string.h>
//
// int cmpstr(const void* a, const void* b) {
//     return strcmp(*(const char**)a, *(const char**)b);
// }
//
// int cmpint(const void* a, const void* b) {
//     return (*(int*)a - *(int*)b);
// }
//
// void qsortStrings(char* strings[], size_t n) {
//     qsort(strings, n, sizeof(char*), cmpstr);
// }
//
// void qsortInts(int* ints, size_t n) {
//     qsort(ints, n, sizeof(int), cmpint);
// }
//
// char* bsearchString(char* key, char* strings[], size_t n) {
//     return *(char**)bsearch(&key, strings, n, sizeof(char*), cmpstr);
// }
//
// int* bsearchInt(int key, int* ints, size_t n) {
//     return (int*)bsearch(&key, ints, n, sizeof(int), cmpint);
// }
import "C"

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func loadData(n_lines int, strings *[]*C.char, ints *[]C.int) {
	file, _ := os.Open("../strings.txt")
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan() && i < n_lines; i++ {
		*strings = append(*strings, C.CString(scanner.Text()))
	}
	file.Close()

	file, _ = os.Open("../ints.txt")
	scanner = bufio.NewScanner(file)
	for i := 0; scanner.Scan() && i < n_lines; i++ {
		num, _ := strconv.Atoi(scanner.Text())
		*ints = append(*ints, C.int(num))
	}
	file.Close()
}

func main() {
	n_lines := 5000
	var strings []*C.char
	var ints []C.int
	loadData(n_lines, &strings, &ints)

	searchString := strings[0]
	searchInt := ints[0]

	start := time.Now()
	C.qsortStrings(&strings[0], C.size_t(len(strings)))
	resultString := C.bsearchString(searchString, &strings[0], C.size_t(len(strings)))
	elapsed := time.Since(start)
	fmt.Printf("Target string %s found: %s\n", C.GoString(searchString), C.GoString(resultString))
	fmt.Printf("String lookup time: %.6f seconds.\n", elapsed.Seconds())

	start = time.Now()
	C.qsortInts(&ints[0], C.size_t(len(ints)))
	resultInt := C.bsearchInt(C.int(searchInt), &ints[0], C.size_t(len(ints)))
	elapsed = time.Since(start)
	fmt.Printf("Target integer %d found: %d\n", searchInt, int(*resultInt))
	fmt.Printf("Integer lookup time: %.6f seconds.\n", elapsed.Seconds())
}
