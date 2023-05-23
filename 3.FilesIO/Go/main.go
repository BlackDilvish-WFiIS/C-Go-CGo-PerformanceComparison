package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	filename := "../input.txt"

	// Start measuring time
	start := time.Now()

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf(scanner.Text() + "\n")
	}

	elapsed := time.Since(start)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Printf("Execution time: %.6f seconds\n", elapsed.Seconds())
}
