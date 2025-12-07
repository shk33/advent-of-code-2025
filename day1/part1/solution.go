package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(rotations []string) int {
	currentPosition := 50
	zeroCount := 0

	for _, line := range rotations {
		if len(line) < 2 {
			continue
		}

		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			// Skip lines that don't have a valid number
			continue
		}

		if direction == 'R' {
			currentPosition += distance
		} else if direction == 'L' {
			currentPosition -= distance
		}

		// Go's % operator can return negative results, so we ensure it wraps correctly
		currentPosition = (currentPosition%100 + 100) % 100

		if currentPosition == 0 {
			zeroCount++
		}
	}
	return zeroCount
}

func main() {
	inputFilePath := "day1/part1/input1.txt"
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file '%s': %v\n", inputFilePath, err)
		os.Exit(1)
	}
	defer file.Close()

	var rotations []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rotations = append(rotations, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	password := solve(rotations)
	fmt.Printf("The password is: %d\n", password)
}
