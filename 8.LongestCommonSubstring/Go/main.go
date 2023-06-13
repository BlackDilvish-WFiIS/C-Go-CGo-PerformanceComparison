package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func loadData() ([]string, error) {
	file, err := os.Open("../strings.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func lcs(str1, str2 string) string {
    m := len(str1)
    n := len(str2)
    maxLength := 0
    endingIndex := m
    length := make([][]int, 2)
    currRow := 0

    for i := range length {
        length[i] = make([]int, n+1)
    }

    for i := 0; i <= m; i++ {
        for j := 0; j <= n; j++ {
            if i == 0 || j == 0 {
                length[currRow][j] = 0
            } else if str1[i-1] == str2[j-1] {
                length[currRow][j] = length[1-currRow][j-1] + 1
                if length[currRow][j] > maxLength {
                    maxLength = length[currRow][j]
                    endingIndex = i - 1
                }
            } else {
                length[currRow][j] = 0
            }
        }
        currRow = 1 - currRow
    }

    if maxLength == 0 {
        return ""
    }

    return str1[endingIndex-maxLength+1 : endingIndex+1]
}


func main() {
	strings, _ := loadData()
	start := time.Now()
	result := lcs(strings[0], strings[1])
	elapsed := time.Since(start)
	if result != "" {
		fmt.Println(result)
	} else {
		fmt.Println("No common substring")
	}
	fmt.Printf("LCS execution time: %.6f seconds\n\n", elapsed.Seconds())
}
