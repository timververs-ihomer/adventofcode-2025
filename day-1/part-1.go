// Package main solves Advent of Code 2025 Day 1 puzzle.
package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// We need to enter a door with a password.
	// There is a safe with a dial on it. ranging from 0 to 99.

	// The safe is actually a decoy.
	// The actual password is the number of times the dial is left pointing at 0 after any rotation in the sequence.

	// Our starting position is 50
	currentPosition := 50

	// Keep count of the number of times the dial is left pointing at 0 after any rotation in the sequence.
	count := 0

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
func rotateDial(currentPosition int, direction byte, steps int) int {
	switch direction {
	case 'R':
		// Modulo wraps around to the correct position.
		return (currentPosition + steps) % 100
	case 'L':
		// For left rotation, first add 100 to be able to use Modulo
		return (currentPosition - steps + 100) % 100
	default:
		return currentPosition
	}
}
