// Package main solves Advent of Code 2025 Day 3 puzzle part 1.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// findNumbers recursively finds the highest digits starting from startIndex
// We need to select exactly 'count' digits, so we need at least (count-1) digits remaining after each selection
func findNumbers(bank string, startIndex int, count int) []int {
	if count == 0 {
		return []int{}
	}

	// We need at least 'count' digits remaining total
	if startIndex >= len(bank) || len(bank)-startIndex < count {
		return []int{}
	}

	var highestDigit int
	highestDigitIndex := -1

	// Find the highest digit we can select while ensuring we can still select (count-1) more digits
	// We can only search up to len(bank) - (count-1) to ensure enough digits remain
	maxIndex := len(bank) - (count - 1)
	for index := startIndex; index < maxIndex; index++ {
		digit, _ := strconv.Atoi(string(bank[index]))

		// Select the highest digit available
		if digit > highestDigit {
			highestDigit = digit
			highestDigitIndex = index
		}
	}

	// If we found a digit, recursively find the remaining digits
	if highestDigitIndex >= 0 {
		remainingDigits := findNumbers(bank, highestDigitIndex+1, count-1)
		return append([]int{highestDigit}, remainingDigits...)
	}

	return []int{}
}

func main() {
	// Read the input.txt file and parse the product ranges
	banks, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input.txt: %v", err)
	}

	var totalAmount int

	// Loop through the banks, there is one instruction per line.
	for _, bank := range strings.Split(string(banks), "\n") {
		// Find 12 numbers recursively
		numbers := findNumbers(bank, 0, 12)

		// Paste all numbers together
		var combinedNumberString string
		for _, num := range numbers {
			combinedNumberString += strconv.Itoa(num)
		}

		// Convert the pasted numbers to an int
		parsedCombinedNumber, _ := strconv.Atoi(combinedNumberString)

		// Add the int amount to total
		totalAmount += parsedCombinedNumber
	}

	fmt.Println(totalAmount)
}
