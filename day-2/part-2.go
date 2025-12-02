// Package main solves Advent of Code 2025 Day 2 puzzle part 2.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// var invalidIdCount int
var inValidProductIdCount int
var invalidProductIdTotal int

func main() {
	// Read the input.txt file and parse the product ranges
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input.txt: %v", err)
	}

	// split the ranges comma seperated
	productRanges := strings.Split(string(input), ",")

	// Loop through all ranges
	for _, productRange := range productRanges {
		// element is the element from someSlice for where we are
		analyseRange(productRange)
	}

	// print the invalid productId count and total
	fmt.Printf(
		"We found a total of %d invalid product id's that add up to %d\n",
		inValidProductIdCount,
		invalidProductIdTotal,
	)
}

func analyseRange(productRange string) {
	// Split the product range into a start and end number, parse to number aswell
	numbers := strings.Split(productRange, "-")

	// Convert string to int
	firstNumber, _ := strconv.Atoi(numbers[0])
	lastNumber, _ := strconv.Atoi(numbers[1])

	// Loop through all the numbers (productId's)
	for i := firstNumber; i <= lastNumber; i++ {
		// Now we can validate each product id
		if !validateProductId(i) {
			// If the productId is valid, increment the valid productId count
			inValidProductIdCount++

			// Also add the productId (the number) to the total amount.
			invalidProductIdTotal += i
		}
	}
}

// Now we need to analyse the productId to see if there is a pattern that has duplicate numbers in it
func validateProductId(productId int) bool {
	// Convert the productId to a string
	productIdString := strconv.Itoa(productId)

	// split the string through the middle
	leftSide := productIdString[:len(productIdString)/2]
	rightSide := productIdString[len(productIdString)/2:]

	// Check if the left side is equal to the right side (exact match)
	if leftSide == rightSide {
		return false
	}

	return true
}
