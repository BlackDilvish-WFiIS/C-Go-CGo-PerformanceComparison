#include <stdio.h>
#include <time.h>
#include <stdlib.h>

void multiplyMatrices(int** A, int** B, int** result, int size) {
    for (int i = 0; i < size; i++) {
        for (int j = 0; j < size; j++) {
            result[i][j] = 0;
            for (int k = 0; k < size; k++) {
                result[i][j] += A[i][k] * B[k][j];
            }
        }
    }
}

void test(int matrixSize)
{
    printf("Matrix size: %d x %d\n", matrixSize, matrixSize);
    int **matrixA = (int **)malloc(matrixSize * sizeof(int *));
    int **matrixB = (int **)malloc(matrixSize * sizeof(int *));
    int **resultMatrix = (int **)malloc(matrixSize * sizeof(int *));
    for (int i = 0; i < matrixSize; i++) {
        matrixA[i] = (int *)malloc(matrixSize * sizeof(int));
        matrixB[i] = (int *)malloc(matrixSize * sizeof(int));
        resultMatrix[i] = (int *)malloc(matrixSize * sizeof(int));
    }

    for (int i = 0; i < matrixSize; i++) {
        for (int j = 0; j < matrixSize; j++) {
            matrixA[i][j] = i + j;
            matrixB[i][j] = i - j;
        }
    }

    clock_t start = clock();

    multiplyMatrices(matrixA, matrixB, resultMatrix, matrixSize);

    clock_t end = clock();
    double elapsed = (double)(end - start) / CLOCKS_PER_SEC;
    printf("Execution time: %f seconds\n\n", elapsed);

    for (int i = 0; i < matrixSize; i++) {
        free(matrixA[i]);
        free(matrixB[i]);
        free(resultMatrix[i]);
    }
    free(matrixA);
    free(matrixB);
    free(resultMatrix);
}

int main() {
    test(100);
    test(500);
    test(1000);

    return 0;
}
