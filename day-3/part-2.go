// Package main solves Advent of Code 2025 Day 3 puzzle part 1.
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
	banks, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input.txt: %v", err)
	}

	var totalAmount int

	// Loop through the banks, there is one instruction per line.
	for _, bank := range strings.Split(string(banks), "\n") {
		// Find the highest two numbers in the bank

		// Loop through all the characters of the string and store the value found
		// If the value is (parsed to int) greater than the value that is found, update the value

		var firstHighestNumber int
		var firstHighestNumberIndex int

		// Loop through the bank a first time
		for index, char := range bank {
			numberFound, _ := strconv.Atoi(string(char))

			// The current position should have atleast 11 chars after it
			if numberFound > firstHighestNumber && len(bank)-(index+1) < 12 {
				firstHighestNumber = numberFound
				firstHighestNumberIndex = index
			}
		}

		// Now do it again to get the remaining 11 numbers
		for _, char := range bank[firstHighestNumberIndex:] {
			numberFound, _ := strconv.Atoi(string(char))

			// Loop throught the remaing

		}

		// Paste the two numbers together
		combinedNumberString := strconv.Itoa(firstHighestNumber) + strconv.Itoa(secondHighestNumber)

		// Convert the pasted numbers to a int
		parsedCombinedNumber, _ := strconv.Atoi(combinedNumberString)

		fmt.Println((parsedCombinedNumber))

		// Add the int amount to total
		totalAmount += parsedCombinedNumber
	}

	fmt.Println((totalAmount))

}
