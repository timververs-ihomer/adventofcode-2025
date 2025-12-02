// Package main solves Advent of Code 2025 Day 1 puzzle.
package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// Keep count of the number of times the dial is left pointing at 0 after any rotation in the sequence.
var count int

func main() {
	// We need to enter a door with a password.
	// There is a safe with a dial on it. ranging from 0 to 99.

	// The safe is actually a decoy.
	// The actual password is the number of times the dial is left pointing at 0 after any rotation in the sequence.

	// Our starting position is 50
	currentPosition := 50

	// Read the input.txt file and parse the instructions.
	instructions, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input.txt: %v", err)
	}

	// Loop through the instructions, there is one instruction per line.
	for _, instruction := range strings.Split(string(instructions), "\n") {
		// parse the instruction and extract the return values from the function in seperate vars
		direction, steps := parseInstruction(instruction)

		// Rotate the dial based on the direction and steps.
		currentPosition = rotateDial(currentPosition, direction, steps)

		// if the current position is 0, increment the count.
		if currentPosition == 0 {
			count++
		}
	}

	// log the count.
	log.Printf("The password is: %d", count)
}

// Function to parse the instruction.
// The instruction is a string of the form "R10" or "L5".
func parseInstruction(instruction string) (byte, int) {
	// The direction is the first character of the instruction.
	direction := instruction[0]

	// Check if the direction is valid.
	if direction != 'R' && direction != 'L' {
		log.Fatalf("Invalid direction: %s", direction)
	}

	// The steps is the number after the direction.
	// strconv.Atoi converts the string to an integer.
	steps, err := strconv.Atoi(instruction[1:])
	if err != nil {
		log.Fatalf("Failed to parse instruction: %v", err)
	}

	// Return the direction and the steps.
	return direction, steps
}

// rotateDial rotates the dial based on the direction and steps.
// If you reach 99 and go right, you wrap around to 0.
// If you reach 0 and go left, you wrap around to 99.
// Counts how many times the dial points at 0 during the rotation.
func rotateDial(currentPosition int, direction byte, steps int) int {
	// Get full rotations (each full rotation passes through 0 once)
	fullRotations := steps / 100
	count += fullRotations

	// Get remainder after full rotations
	remainder := steps % 100

	var newPosition int
	switch direction {
	case 'R':
		// Calculate new position with wrapping
		newPosition = (currentPosition + remainder) % 100
		// Check if we cross the 0 boundary (from 99->0 or any position that wraps)
		// If currentPosition + remainder >= 100, we wrapped and passed through 0
		// But if we end exactly at 0, don't count here (main() will count it)
		if currentPosition+remainder >= 100 && newPosition != 0 {
			count++
		}
	case 'L':
		// Calculate new position with wrapping
		newPosition = (currentPosition - remainder + 100) % 100
		// Check if we cross the 0 boundary (from a positive position to negative, which wraps)
		// If currentPosition - remainder < 0, we wrapped and passed through 0
		// But don't count if we start at 0 (we're already there, not passing through)
		// And if we end exactly at 0, don't count here (main() will count it)
		if currentPosition-remainder < 0 && currentPosition != 0 && newPosition != 0 {
			count++
		}
	default:
		newPosition = currentPosition
	}

	return newPosition
}
