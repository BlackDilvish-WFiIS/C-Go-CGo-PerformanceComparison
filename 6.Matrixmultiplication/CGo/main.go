package main

import (
	"fmt"
	"time"
	"unsafe"
)

/*
#include <stdio.h>

void multiplyMatrices(int* A, int* B, int* result, int rowsA, int colsA, int colsB) {
    for (int i = 0; i < rowsA; i++) {
        for (int j = 0; j < colsB; j++) {
            result[i * colsB + j] = 0;
            for (int k = 0; k < colsA; k++) {
                result[i * colsB + j] += A[i * colsA + k] * B[k * colsB + j];
            }
        }
    }
}
*/
import "C"

func multiplyMatrices(A, B, result [][]int) {
	rowsA := len(A)
	colsA := len(A[0])
	colsB := len(B[0])

	// Convert 2D Go slices to 1D C arrays
	flatten := func(matrix [][]int) []int {
		arr := make([]int, len(matrix)*len(matrix[0]))
		for i, row := range matrix {
			for j, val := range row {
				arr[i*len(row)+j] = val
			}
		}
		return arr
	}

	cArrA := (*C.int)(unsafe.Pointer(&flatten(A)[0]))
	cArrB := (*C.int)(unsafe.Pointer(&flatten(B)[0]))
	cArrResult := (*C.int)(unsafe.Pointer(&flatten(result)[0]))
	C.multiplyMatrices(cArrA, cArrB, cArrResult, C.int(rowsA), C.int(colsA), C.int(colsB))
}

func test(matrixSize int) {
	fmt.Printf("Matrix size: %d x %d\n", matrixSize, matrixSize)
	matrixA := make([][]int, matrixSize)
	matrixB := make([][]int, matrixSize)
	resultMatrix := make([][]int, matrixSize)

	for i := 0; i < matrixSize; i++ {
		matrixA[i] = make([]int, matrixSize)
		matrixB[i] = make([]int, matrixSize)
		resultMatrix[i] = make([]int, matrixSize)
		for j := 0; j < matrixSize; j++ {
			matrixA[i][j] = i + j
			matrixB[i][j] = i - j
		}
	}

	start := time.Now()

	multiplyMatrices(matrixA, matrixB, resultMatrix)

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %.6f seconds\n\n", elapsed.Seconds())
}

func main() {
	test(100)
	test(500)
	test(1000)
}
