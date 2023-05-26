#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

int main() {
    clock_t start = clock();

    char* buff = malloc(sizeof(char) * 200010);
    buff[0] = '\n';
    const char* str = "test";

    for (int i = 0; i < 50000; i++) {
        strcat(buff, str);
    }

    clock_t end = clock();
    double elapsed = (double)(end - start) / CLOCKS_PER_SEC;

    printf("Result length: %zu\n", strlen(buff));
    printf("Execution time: %f seconds\n", elapsed);

    free(buff);

    return 0;
}
