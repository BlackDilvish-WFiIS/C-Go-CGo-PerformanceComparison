package main

import (
	"fmt"
	"time"
)

func test(nStrings int) {
	fmt.Printf("Number of appends: %d\n", nStrings)
	start := time.Now()

	str1 := ""
	str2 := "test"
	for i := 0; i < nStrings; i++ {
		str1 += str2
	}

	elapsed := time.Since(start)
	fmt.Printf("Result length: %d\n", len(str1))
	fmt.Printf("Execution time: %.6f seconds\n\n", elapsed.Seconds())
}

func main() {
	test(10000)
	test(50000)
	test(100000)
}
