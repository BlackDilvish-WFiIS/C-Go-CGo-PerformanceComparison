package main

/*
#include <stdio.h>

long long fib(int n) {
    if (n <= 1)
        return n;
    else
        return(fib(n-1) + fib(n-2));
}
*/
import "C"
import (
	"fmt"
	"time"
)

const LIMIT = 40

func test() {
	start := time.Now()
	dummy := 0
	for i := 0; i < LIMIT; i++ {
		dummy += int(C.fib(C.int(i)))
	}
	elapsed := time.Since(start)
	fmt.Printf("Fibonacci calculation execution time: %.6f seconds\n\n", elapsed.Seconds())
}

func main() {
	test()
}
