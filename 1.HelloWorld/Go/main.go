package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Printf("Hello, world!\n")

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %.6f seconds\n\n", elapsed.Seconds())
}
