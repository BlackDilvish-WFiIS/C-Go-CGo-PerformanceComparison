#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#define MAX_STR_LEN 9
#define MAX_STR_COUNT 10000
#define MAX_INT_COUNT 10000

void loadData(int n_lines, char strings[][MAX_STR_LEN], int ints[]) {
    FILE *f;
    int i = 0, j = 0;
    char line[100];

    f = fopen("../strings.txt", "r");
    while (i < n_lines && fgets(line, sizeof(line), f)) {
        line[strcspn(line, "\n")] = 0;  // remove newline
        strncpy(strings[i], line, MAX_STR_LEN);
        strings[i++][MAX_STR_LEN] = '\0';
    }
    fclose(f);

    f = fopen("../ints.txt", "r");
    while (j < n_lines && fgets(line, sizeof(line), f)) {
        ints[j++] = atoi(line);
    }
    fclose(f);
}

int cmpstr(const void* a, const void* b) {
    return strcmp((char*)a, (char*)b);
}

int cmpint(const void* a, const void* b) {
    return (*(int*)a - *(int*)b);
}

int main() {
    int n_lines = 5000;
    char strings[n_lines][MAX_STR_LEN];
    int ints[n_lines];
    loadData(n_lines, strings, ints);

    char* searchString = strings[0];
    int searchInt = ints[0];

    clock_t start, end;
    double cpu_time_used;

    start = clock();
    qsort(strings, n_lines, MAX_STR_LEN, cmpstr);
    char* resultString = (char*)bsearch(searchString, strings, n_lines, MAX_STR_LEN, cmpstr);
    end = clock();
    cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
    printf("Target string %s found: %s\n", searchString, resultString);
    printf("String lookup time: %.6f seconds.\n", cpu_time_used);

    start = clock();
    qsort(ints, n_lines, sizeof(int), cmpint);
    int* resultInt = (int*)bsearch(&searchInt, ints, n_lines, sizeof(int), cmpint);
    end = clock();
    cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
    printf("Target integer %d found: %d\n", searchInt, *resultInt);
    printf("Integer lookup time: %.6f seconds.\n", cpu_time_used);

    return 0;
}
