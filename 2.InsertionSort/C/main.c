#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#define MAX_STR_LEN 8
#define MAX_STR_COUNT 10000
#define MAX_INT_COUNT 10000

void insertionSortString(char arr[][MAX_STR_LEN], int n) {
    int i, j;
    char key[MAX_STR_LEN];
    for (i = 1; i < n; i++) {
        strcpy(key, arr[i]);
        j = i - 1;
        while (j >= 0 && strcmp(arr[j], key) > 0) {
            strcpy(arr[j + 1], arr[j]);
            j = j - 1;
        }
        strcpy(arr[j + 1], key);
    }
}

void insertionSortInt(int arr[], int n) {
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

void loadData(int n_lines, char strings[][MAX_STR_LEN], int ints[]) {
    FILE *f;
    int i = 0, j = 0;
    char line[100];

    f = fopen("../strings.txt", "r");
    while (i < n_lines && fgets(line, sizeof(line), f)) {
        line[strcspn(line, "\n")] = 0;  // remove newline
        strcpy(strings[i++], line);
    }
    fclose(f);

    f = fopen("../ints.txt", "r");
    while (j < n_lines && fgets(line, sizeof(line), f)) {
        ints[j++] = atoi(line);
    }
    fclose(f);
}

void test(int n_lines, char strings[][MAX_STR_LEN], int ints[]) {
    clock_t start, end;
    double cpu_time_used;

    start = clock();
    insertionSortString(strings, n_lines);
    end = clock();
    cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
    printf("String sorting execution time: %.6f seconds.\n", cpu_time_used);

    start = clock();
    insertionSortInt(ints, n_lines);
    end = clock();
    cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
    printf("Integer sorting execution time: %.6f seconds.\n", cpu_time_used);
}

int main() {
    int n_lines = 5000;
    char strings[n_lines][MAX_STR_LEN];
    int ints[n_lines];

    loadData(n_lines, strings, ints);
    test(n_lines, strings, ints);

    return 0;
}
