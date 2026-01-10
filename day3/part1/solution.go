package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func solve() {
	// Correctly resolve the input file path relative to the executable
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("failed to get current directory: %v", err)
	}
	inputPath := filepath.Join(dir, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		// Fallback for when running with "go run"
		inputPath = filepath.Join(".", "day3", "part1", "input.txt")
	}

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalJoltage := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		maxLineJoltage := 0
		for i := 0; i < len(line); i++ {
			for j := i + 1; j < len(line); j++ {
				joltageStr := string(line[i]) + string(line[j])
				joltage, _ := strconv.Atoi(joltageStr)
				if joltage > maxLineJoltage {
					maxLineJoltage = joltage
				}
			}
		}
		totalJoltage += maxLineJoltage
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input file: %v", err)
	}

	fmt.Printf("The total output joltage is: %d\n", totalJoltage)
}

func main() {
	solve()
}
