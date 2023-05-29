package main

import (
	"fmt"
	"time"
)

func swap(str []rune, i, j int) {
	str[i], str[j] = str[j], str[i]
}

func permute(str []rune, l, r int, count *int) {
	if l == r {
		(*count)++
	} else {
		for i := l; i <= r; i++ {
			swap(str, l, i)
			permute(str, l+1, r, count)
			swap(str, l, i)
		}
	}
}

func test(str string) {
	fmt.Printf("String to permute: %s, len: %d\n", str, len(str))
	runes := []rune(str)
	count := 0

	startTime := time.Now()
	permute(runes, 0, len(runes)-1, &count)
	endTime := time.Now()

	executionTime := endTime.Sub(startTime)

	fmt.Println("Permutations:", count)
	fmt.Printf("Execution time: %.6f seconds\n\n", executionTime.Seconds())
}

func main() {
	test("abcdefghij")
	test("abcdefghijk")
	test("abcdefghijkl")
}
