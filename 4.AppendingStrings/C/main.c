#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

void test(int nStrings)
{
    printf("Number of appends: %d\n", nStrings);
    clock_t start = clock();

    char* buff = malloc(sizeof(char) * (nStrings*4 + 10));
    buff[0] = '\n';
    const char* str = "test";

    for (int i = 0; i < nStrings; i++) {
        strcat(buff, str);
    }

    clock_t end = clock();
    double elapsed = (double)(end - start) / CLOCKS_PER_SEC;

    printf("Result length: %zu\n", strlen(buff));
    printf("Execution time: %f seconds\n\n", elapsed);

    free(buff);
}

int main() {
    test(10000);
    test(50000);
    test(100000);

    return 0;
}
