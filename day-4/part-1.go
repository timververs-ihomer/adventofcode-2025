// Package main solves Advent of Code 2025 Day 4 puzzle part 1.
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var rows []string
var maxColumnIndex int

func main() {
	// Read the input.txt file and parse the product ranges
	inputData, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input.txt: %v", err)
	}

	rows = strings.Split(string(inputData), "\n")

	// Paper roll is accessable
	var paperRollAccessible int

	// first we determine the max column length
	maxColumnIndex = len(rows[0]) - 1

	fmt.Println(maxColumnIndex)

	// Loop through all the rows and then all the columns
	for rowIndex, row := range rows {
		for columnIndex, column := range row {
			// if the column is a '.' skip it.
			if column == '@' {
				if checkPaperRoll(rowIndex, columnIndex) {
					paperRollAccessible++
				}
			}
		}
	}

	fmt.Println(paperRollAccessible)
}

func checkPaperRoll(rowIndex int, columnIndex int) bool {
	paperRollsAbove := 0
	paperRollsCurrent := 0
	paperRollsBelow := 0

	// Validate the row above you only if there is a row
	if rowIndex != 0 {
		paperRollsAbove = checkRowAbove(rowIndex, columnIndex)
	}

	//validate the current row only if there is a row
	paperRollsCurrent = checkCurrentRow(rowIndex, columnIndex)

	//validate the row below you only if there is a row
	if rowIndex != len(rows)-1 {
		paperRollsBelow = checkRowBelow(rowIndex, columnIndex)
	}

	return (paperRollsAbove + paperRollsBelow + paperRollsCurrent) < 4
}

func checkRowAbove(rowIndex int, columnIndex int) int {
	foundPaperRolls := 0

	if columnIndex != 0 && rows[rowIndex-1][columnIndex-1] == '@' {
		foundPaperRolls++
	}

	if rows[rowIndex-1][columnIndex] == '@' {
		foundPaperRolls++
	}

	if columnIndex+1 <= maxColumnIndex && rows[rowIndex-1][columnIndex+1] == '@' {
		foundPaperRolls++
	}

	return foundPaperRolls
}

func checkCurrentRow(rowIndex int, columnIndex int) int {
	foundPaperRolls := 0

	if columnIndex != 0 && rows[rowIndex][columnIndex-1] == '@' {
		foundPaperRolls++
	}

	if columnIndex+1 <= maxColumnIndex && rows[rowIndex][columnIndex+1] == '@' {
		foundPaperRolls++
	}

	return foundPaperRolls
}

func checkRowBelow(rowIndex int, columnIndex int) int {
	foundPaperRolls := 0

	if columnIndex != 0 && rows[rowIndex+1][columnIndex-1] == '@' {
		foundPaperRolls++
	}

	if rows[rowIndex+1][columnIndex] == '@' {
		foundPaperRolls++

	}
	if columnIndex+1 <= maxColumnIndex && rows[rowIndex+1][columnIndex+1] == '@' {
		foundPaperRolls++
	}

	return foundPaperRolls
}
