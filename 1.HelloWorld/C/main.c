#include <stdio.h>
#include <time.h>

int main() {
    clock_t start, end;
    double cpu_time_used;
    start = clock();

    printf("Hello, world!");
    end = clock();
    cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
    printf("Execution time: %.6f seconds.\n\n", cpu_time_used);
    return 0;
}
