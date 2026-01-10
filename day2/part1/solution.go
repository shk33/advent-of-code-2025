package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
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
	base := int64(1)
	sBasePrev := "0"

	for {
		sBase := fmt.Sprintf("%d", base)
		sInvalid := sBase + sBase
		nInvalid, _ := new(big.Int).SetString(sInvalid, 10)

		if nInvalid.Cmp(maxID) > 0 {
			if len(sBase) > len(sBasePrev) {
				break
			}
		} else {
			invalidIDs = append(invalidIDs, nInvalid)
		}
		sBasePrev = sBase
		base++
	}
	return invalidIDs
}

func solve(inputStr string) *big.Int {
	ranges, maxID := parseRanges(inputStr)
	potentialIDs := generateInvalidIDs(maxID)

	foundInvalidIDs := make(map[string]*big.Int)

	for _, invalidID := range potentialIDs {
		for _, r := range ranges {
			if r.Start.Cmp(invalidID) <= 0 && invalidID.Cmp(r.End) <= 0 {
				foundInvalidIDs[invalidID.String()] = invalidID
				break
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
	inputFilePath := "day2/part1/input.txt"
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputStr := scanner.Text()

	result := solve(inputStr)
	fmt.Printf("The sum of all invalid IDs is: %s\n", result.String())
}

