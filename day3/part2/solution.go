package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func findMaxSubsequence(line string, k int) string {
	if len(line) < k || k == 0 {
		return strings.Repeat("0", k)
	}

	var result strings.Builder
	currentStartIndex := 0

	for i := 0; i < k; i++ {
		remainingToFind := k - i
		searchEndIndex := len(line) - remainingToFind

		bestDigit := -1
		bestDigitIndex := -1

		for j := currentStartIndex; j <= searchEndIndex; j++ {
			digit, _ := strconv.Atoi(string(line[j]))
			if digit > bestDigit {
				bestDigit = digit
				bestDigitIndex = j
			}
		}
		
		result.WriteString(strconv.Itoa(bestDigit))
		currentStartIndex = bestDigitIndex + 1
	}

	return result.String()
}

func solve() {
	// Correctly resolve the input file path
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("failed to get current directory: %v", err)
	}
	inputPath := filepath.Join(dir, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		inputPath = filepath.Join(".", "day3", "part2", "input.txt")
	}

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalJoltage := new(big.Int)
	k := 12

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		maxLineJoltageStr := findMaxSubsequence(line, k)
		
		currentLineJoltage := new(big.Int)
		currentLineJoltage.SetString(maxLineJoltageStr, 10)
		totalJoltage.Add(totalJoltage, currentLineJoltage)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input file: %v", err)
	}

	fmt.Printf("The total output joltage is: %s\n", totalJoltage.String())
}

func main() {
	solve()
}
