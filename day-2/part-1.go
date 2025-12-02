// Package main solves Advent of Code 2025 Day 2 puzzle.
package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

//var invalidIdCount int

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
		validateProductId(i)
	}
}

func validateProductId(productId int) {
	// Now we need to analyse the productId to see if there is a pattern that has duplicate numbers in it

	// First check if we can find single digit matches

}
