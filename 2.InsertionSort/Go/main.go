package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func insertionSortStrings(arr []string) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func insertionSortInts(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func loadData(n_lines int) ([]string, []int, error) {
	stringsFile, err := os.Open("../strings.txt")
	if err != nil {
		return nil, nil, err
	}
	defer stringsFile.Close()

	intsFile, err := os.Open("../ints.txt")
	if err != nil {
		return nil, nil, err
	}
	defer intsFile.Close()

	var strings []string
	var ints []int

	scanner := bufio.NewScanner(stringsFile)
	for scanner.Scan() && len(strings) < n_lines {
		strings = append(strings, scanner.Text())
	}

	scanner = bufio.NewScanner(intsFile)
	for scanner.Scan() && len(ints) < n_lines {
		num, _ := strconv.Atoi(scanner.Text())
		ints = append(ints, num)
	}

	return strings, ints, nil
}

func test(n_lines int, strings []string, ints []int) {
	start := time.Now()
	insertionSortStrings(strings)
	elapsed := time.Since(start)
	fmt.Printf("String sorting execution time: %.6f seconds\n\n", elapsed.Seconds())

	start = time.Now()
	insertionSortInts(ints)
	elapsed = time.Since(start)
	fmt.Printf("Integer sorting execution time: %.6f seconds\n\n", elapsed.Seconds())
}

func main() {
	n_lines := 10000
	strings, ints, _ := loadData(n_lines)
	test(n_lines, strings, ints)
}
