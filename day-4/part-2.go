// Package main solves Advent of Code 2025 Day 4 puzzle part 2.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var rows []string
var maxColumnIndex int
var paperRollAbleToRemove int

func main() {
	// Read the input.txt file and parse the product ranges
	inputData, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input.txt: %v", err)
	}

	rows = strings.Split(string(inputData), "\n")

	// first we determine the max column length
	maxColumnIndex = len(rows[0]) - 1

	cleanPaperRolls()

	fmt.Println(paperRollAbleToRemove)
}

func cleanPaperRolls() {
	var paperRollsToRemove []string

	// Loop through all the rows and then all the columns
	for rowIndex, row := range rows {
		for columnIndex, column := range row {
			// if the column is a '.' skip it.
			if column == '@' {
				if checkPaperRoll(rowIndex, columnIndex) {
					paperRollAbleToRemove++

					//mark rowIndex + columnIndex for removal by adding it to the paperRollsToRemove slice
					paperRollsToRemove = append(paperRollsToRemove, fmt.Sprintf("%d,%d", rowIndex, columnIndex))
				}
			}
		}
	}

	if len(paperRollsToRemove) > 0 {
		// Remove the paper Towels
		removePaperRolls(paperRollsToRemove)

		// Recursivly recall the cleanPaperRolls with the new rows variable
		cleanPaperRolls()
	}
}

func removePaperRolls(paperRollsToRemove []string) {

	for _, paperRoll := range paperRollsToRemove {
		rowIndex, _ := strconv.Atoi(strings.Split(paperRoll, ",")[0])
		columnIndex, _ := strconv.Atoi(strings.Split(paperRoll, ",")[1])

		// Rebuild the string by slicing and concatenating (strings are immutable in Go :( sad face)
		row := rows[rowIndex]
		rows[rowIndex] = row[:columnIndex] + "X" + row[columnIndex+1:]
	}
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
