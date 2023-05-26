package main

import (
	"fmt"
	"time"
)

const (
	Rows = 1000
	Cols = 1000
)

func multiplyMatrices(A, B, result [][]int) {
	for i := 0; i < Rows; i++ {
		for j := 0; j < Cols; j++ {
			result[i][j] = 0
			for k := 0; k < Cols; k++ {
				result[i][j] += A[i][k] * B[k][j]
			}
		}
	}
}

func main() {
	matrixA := make([][]int, Rows)
	matrixB := make([][]int, Rows)
	resultMatrix := make([][]int, Rows)

	for i := 0; i < Rows; i++ {
		matrixA[i] = make([]int, Cols)
		matrixB[i] = make([]int, Cols)
		resultMatrix[i] = make([]int, Cols)
		for j := 0; j < Cols; j++ {
			matrixA[i][j] = i + j
			matrixB[i][j] = i - j
		}
	}

	start := time.Now()
	multiplyMatrices(matrixA, matrixB, resultMatrix)

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %.6f seconds\n", elapsed.Seconds())
}
