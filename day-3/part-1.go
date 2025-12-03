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
		var matchingFirstNumberFound bool
		var secondHighestNumber int

		// Loop through the bank a first time
		for index, char := range bank {
			numberFound, _ := strconv.Atoi(string(char))

			if numberFound > firstHighestNumber && (index+1) != len(bank) {
				firstHighestNumber = numberFound
			}
		}

		// Now do it again to get the second highest number
		for _, char := range bank {
			numberFound, _ := strconv.Atoi(string(char))

			// did you find the first highest number and it was not found yet? Then set the value to true and continue with the next one
			if numberFound == firstHighestNumber && !matchingFirstNumberFound {
				// match found, stop the process
				matchingFirstNumberFound = true

				// go to the next number
				continue
			}

			// now just check the highest number in the rest of the numbers
			if matchingFirstNumberFound && numberFound > secondHighestNumber {
				secondHighestNumber = numberFound
			}
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
