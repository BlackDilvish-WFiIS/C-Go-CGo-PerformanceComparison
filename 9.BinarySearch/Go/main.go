package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func loadData(n_lines int, strings *[]string, ints *[]int) {
    file, _ := os.Open("../strings.txt")
    scanner := bufio.NewScanner(file)
    for i := 0; scanner.Scan() && i < n_lines; i++ {
        *strings = append(*strings, scanner.Text())
    }
    file.Close()

    file, _ = os.Open("../ints.txt")
    scanner = bufio.NewScanner(file)
    for i := 0; scanner.Scan() && i < n_lines; i++ {
        num, _ := strconv.Atoi(scanner.Text())
        *ints = append(*ints, num)
    }
    file.Close()
}

func main() {
    n_lines := 5000
    var strings []string
    var ints []int
    loadData(n_lines, &strings, &ints)

    searchString := strings[0]
    searchInt := ints[0]

	start := time.Now()
    sort.Strings(strings)
    resultString := sort.SearchStrings(strings, searchString)
	elapsed := time.Since(start)
    fmt.Printf("Target string %s found: %s\n", searchString, strings[resultString]);
    fmt.Printf("String lookup time: %.6f seconds.\n", elapsed.Seconds());

	start = time.Now()
    sort.Ints(ints)
    resultInt := sort.SearchInts(ints, searchInt)
	elapsed = time.Since(start)
    fmt.Printf("Target int %d found: %d\n", searchInt, ints[resultInt]);
    fmt.Printf("Int lookup time: %.6f seconds.\n", elapsed.Seconds());
}
