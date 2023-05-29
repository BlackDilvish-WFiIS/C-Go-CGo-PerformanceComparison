package main

import (
	"fmt"
	"time"
)

func multiplyMatrices(A, B, result [][]int) {
	size := len(A)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			result[i][j] = 0
			for k := 0; k < size; k++ {
				result[i][j] += A[i][k] * B[k][j]
			}
		}
	}
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
