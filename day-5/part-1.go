// Package main solves Advent of Code 2025 Day 5 puzzle part 2.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read the input.txt file and parse the product ranges
	inputData, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input.txt: %v", err)
	}

	lines := strings.Split(string(inputData), "\n")

	// loop through the input data and find the first line that is an empty string
	var indexOfEmptyLine int
	for index, line := range lines {
		if line == "" {
			indexOfEmptyLine = index
			break
		}
	}

	// everything before the empty line is the fresh product ranges
	freshProductRanges := lines[:indexOfEmptyLine]
	// everything after the empty line
	freshProductsAvailable := lines[indexOfEmptyLine:]

	var freshProductsAvailableCount int

	// loop through each available product and check if it's fresh
	for _, availableProduct := range freshProductsAvailable {

		// parse the available productId to an int
		availableProductInt, _ := strconv.Atoi(availableProduct)

		for _, productRange := range freshProductRanges {
			// if the number is higher than the range start and lower than the range end, then it's fresh
			rangeNumbers := strings.Split(productRange, "-")
			rangeStart, _ := strconv.Atoi(rangeNumbers[0])
			rangeEnd, _ := strconv.Atoi(rangeNumbers[1])
			if availableProductInt >= rangeStart && availableProductInt <= rangeEnd {
				// up the count
				freshProductsAvailableCount++

				// go to the next available product (break the for loop)
				break
			}
		}
	}

	// print the count of fresh products available
	fmt.Println(freshProductsAvailableCount)
}

//636 to high
