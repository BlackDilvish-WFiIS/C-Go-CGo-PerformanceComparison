package main

import (
	"fmt"
	"time"
)

const LIMIT = 40

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func test() {
	start := time.Now()
	dummy := 0
	for i := 0; i < LIMIT; i++ {
		dummy += fib(i)
	}
	elapsed := time.Since(start)
	fmt.Printf("Fibonacci calculation execution time: %.6f seconds\n\n", elapsed.Seconds())
}

func main() {
	test()
}
