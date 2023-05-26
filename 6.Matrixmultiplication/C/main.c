#include <stdio.h>
#include <time.h>
#include <stdlib.h>

#define ROWS 1000
#define COLS 1000

void multiplyMatrices(int** A, int** B, int** result) {
    for (int i = 0; i < ROWS; i++) {
        for (int j = 0; j < COLS; j++) {
            result[i][j] = 0;
            for (int k = 0; k < COLS; k++) {
                result[i][j] += A[i][k] * B[k][j];
            }
        }
    }
}

int main() {
    int **matrixA = (int **)malloc(ROWS * sizeof(int *));
    int **matrixB = (int **)malloc(ROWS * sizeof(int *));
    int **resultMatrix = (int **)malloc(ROWS * sizeof(int *));
    for (int i = 0; i < ROWS; i++) {
        matrixA[i] = (int *)malloc(COLS * sizeof(int));
        matrixB[i] = (int *)malloc(COLS * sizeof(int));
        resultMatrix[i] = (int *)malloc(COLS * sizeof(int));
    }

    for (int i = 0; i < ROWS; i++) {
        for (int j = 0; j < COLS; j++) {
            matrixA[i][j] = i + j;
            matrixB[i][j] = i - j;
        }
    }

    clock_t start = clock();

    multiplyMatrices(matrixA, matrixB, resultMatrix);

    clock_t end = clock();
    double elapsed = (double)(end - start) / CLOCKS_PER_SEC;
    printf("Execution time: %f seconds\n", elapsed);

    for (int i = 0; i < ROWS; i++) {
        free(matrixA[i]);
        free(matrixB[i]);
        free(resultMatrix[i]);
    }
    free(matrixA);
    free(matrixB);
    free(resultMatrix);

    return 0;
}
