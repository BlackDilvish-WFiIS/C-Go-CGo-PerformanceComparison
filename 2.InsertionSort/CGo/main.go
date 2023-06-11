package main

/*
#include <stdlib.h>
#include <string.h>

int compareStrings(const void* a, const void* b) {
    char *strA = *(char**)a;
    char *strB = *(char**)b;
    return strcmp(strA, strB);
}

void insertionSortInts(int* arr, int n) {
    int i, key, j;
    for (i = 1; i < n; i++) {
        key = arr[i];
        j = i - 1;
        while (j >= 0 && arr[j] > key) {
            arr[j + 1] = arr[j];
            j = j - 1;
        }
        arr[j + 1] = key;
    }
}

void insertionSortStrings(char** arr, int n) {
    int i, j;
    char* key;
    for (i = 1; i < n; i++) {
        key = arr[i];
        j = i - 1;
        while (j >= 0 && strcmp(arr[j], key) > 0) {
            arr[j + 1] = arr[j];
            j = j - 1;
        }
        arr[j + 1] = key;
    }
}
*/
import "C"
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"unsafe"
)

func loadData(numOfLines int) ([]string, []int, error) {
	stringsFile, err := os.Open("..//strings.txt")
	if err != nil {
		return nil, nil, err
	}
	defer stringsFile.Close()

	intsFile, err := os.Open("../ints.txt")
	if err != nil {
		return nil, nil, err
	}
	defer intsFile.Close()

	var strings []string
	var ints []int

	scanner := bufio.NewScanner(stringsFile)
	for scanner.Scan() && len(strings) < numOfLines {
		strings = append(strings, scanner.Text())
	}

	scanner = bufio.NewScanner(intsFile)
	for scanner.Scan() && len(ints) < numOfLines {
		num, _ := strconv.Atoi(scanner.Text())
		ints = append(ints, num)
	}

	return strings, ints, nil
}

func test(numOfLines int, strings []string, ints []int) {
	cStrings := make([]*C.char, len(strings))
	for i, s := range strings {
		cStrings[i] = C.CString(s)
		defer C.free(unsafe.Pointer(cStrings[i]))
	}

	cInts := make([]C.int, len(ints))
	for i, n := range ints {
		cInts[i] = C.int(n)
	}

	start := time.Now()
	C.insertionSortStrings((**C.char)(unsafe.Pointer(&cStrings[0])), C.int(numOfLines))
	elapsed := time.Since(start)
	fmt.Printf("String sorting execution time: %.6f seconds\n\n", elapsed.Seconds())

	start = time.Now()
	C.insertionSortInts((*C.int)(unsafe.Pointer(&cInts[0])), C.int(numOfLines))
	elapsed = time.Since(start)
	fmt.Printf("Integer sorting execution time: %.6f seconds\n\n", elapsed.Seconds())
}

func main() {
	numOfLines := 10000
	strings, ints, _ := loadData(numOfLines)
	test(numOfLines, strings, ints)
}
