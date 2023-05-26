package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	str1 := ""
	str2 := "test"
	for i := 0; i < 50000; i++ {
		str1 += str2
	}

	elapsed := time.Since(start)
	fmt.Printf("Result length: %d\n", len(str1))
	fmt.Printf("Execution time: %.6f seconds\n", elapsed.Seconds())
}
