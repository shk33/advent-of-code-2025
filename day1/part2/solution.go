package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// A helper for true mathematical floor division
func floorDiv(a, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}

func solve(rotations []string) int {
	currentPosition := 50
	totalZeroCount := 0

	for _, line := range rotations {
		if len(line) < 2 {
			continue
		}

		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			continue
		}

		zerosThisTurn := 0
		if direction == 'R' {
			zerosThisTurn = floorDiv(currentPosition+distance, 100) - floorDiv(currentPosition, 100)
		} else if direction == 'L' {
			zerosThisTurn = floorDiv(currentPosition-1, 100) - floorDiv(currentPosition-distance-1, 100)
		}
		totalZeroCount += zerosThisTurn

		if direction == 'R' {
			currentPosition += distance
		} else if direction == 'L' {
			currentPosition -= distance
		}
	}
	return totalZeroCount
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
