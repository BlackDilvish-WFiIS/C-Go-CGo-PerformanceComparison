#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define LIMIT 40

long long fib(int n) {
    if (n <= 1)
        return n;
    else
        return(fib(n-1) + fib(n-2));
}

void test() {
    clock_t start, end;
    double cpu_time_used;
    long long dummy = 0;

    start = clock();
    for (int i = 0; i < LIMIT; i++) {
        dummy += fib(i);
    }
    end = clock();
    cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
    printf("Fibonacci calculation execution time: %.6f seconds.\n", cpu_time_used);
}

int main() {
    test();
    return 0;
}
