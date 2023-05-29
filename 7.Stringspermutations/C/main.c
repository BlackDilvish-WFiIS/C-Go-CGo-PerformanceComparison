#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

void swap(char* a, char* b) {
    char temp = *a;
    *a = *b;
    *b = temp;
}

void permute(char* str, int l, int r, int* count) {
    if (l == r) {
        (*count)++;
    } else {
        for (int i = l; i <= r; i++) {
            swap((str + l), (str + i));
            permute(str, l + 1, r, count);
            swap((str + l), (str + i));
        }
    }
}

void test(const char* str)
{
    int len = strlen(str);
    char* str_mut = malloc(sizeof(char) * len);
    strcpy(str_mut, str);
    printf("String to permute: %s, len: %d\n", str, len);
    int count = 0;

    clock_t start_time = clock();
    permute(str_mut, 0, len - 1, &count);
    clock_t end_time = clock();

    double execution_time = (double)(end_time - start_time) / CLOCKS_PER_SEC;

    printf("Permutations: %d\n", count);
    printf("Execution Time: %.5f seconds\n\n", execution_time);
}

int main() {
    test("abcdefghij");
	test("abcdefghijk");
	test("abcdefghijkl");

    return 0;
}
