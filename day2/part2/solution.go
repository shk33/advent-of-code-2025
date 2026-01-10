
package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// Define a struct to hold the range properly
type Range struct {
	Start *big.Int
	End   *big.Int
}

func parseRanges(inputStr string) ([]Range, *big.Int) {
	var ranges []Range
	maxID := big.NewInt(0)
	rangeStrs := strings.Split(strings.TrimSpace(inputStr), ",")
	for _, rStr := range rangeStrs {
		if rStr == "" {
			continue
		}
		parts := strings.Split(rStr, "-")
		start, _ := new(big.Int).SetString(parts[0], 10)
		end, _ := new(big.Int).SetString(parts[1], 10)

		ranges = append(ranges, Range{Start: start, End: end})

		if end.Cmp(maxID) > 0 {
			maxID.Set(end)
		}
	}
	return ranges, maxID
}

func generateInvalidIDs(maxID *big.Int) []*big.Int {
	var invalidIDs []*big.Int
	baseNum := int64(1)

	for {
		baseS := strconv.FormatInt(baseNum, 10)
		
		// Optimization: if the smallest possible repetition (baseS + baseS) is too big, stop.
		firstRepetitionVal, _ := new(big.Int).SetString(baseS + baseS, 10)
		if firstRepetitionVal.Cmp(maxID) > 0 {
			break
		}
		
		currentRepeatedS := baseS
		
		for {
			currentRepeatedS += baseS // Append baseS again
			nInvalid, _ := new(big.Int).SetString(currentRepeatedS, 10)

			if nInvalid.Cmp(maxID) > 0 {
				break // This baseS repetition is too long
			}
			
			invalidIDs = append(invalidIDs, nInvalid)
		}
		
		baseNum++
	}
	return invalidIDs
}

func solve(inputStr string) *big.Int {
	ranges, maxID := parseRanges(inputStr)
	potentialIDs := generateInvalidIDs(maxID)

	foundInvalidIDs := make(map[string]*big.Int) // Using a map as a Set

	for _, invalidID := range potentialIDs {
		for _, r := range ranges {
			if r.Start.Cmp(invalidID) <= 0 && invalidID.Cmp(r.End) <= 0 {
				foundInvalidIDs[invalidID.String()] = invalidID
				break // Found in a range, break from inner loop
			}
		}
	}

	totalSum := big.NewInt(0)
	for _, v := range foundInvalidIDs {
		totalSum.Add(totalSum, v)
	}

	return totalSum
}

func main() {
	inputFilePath := "day2/part2/input.txt" // Relative path to input.txt in the same directory
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file '%s': %v\n", inputFilePath, err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputStr := scanner.Text()

	result := solve(inputStr)
	fmt.Printf("The sum of all invalid IDs is: %s\n", result.String())
}
