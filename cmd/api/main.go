package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"rubiks-solver/internal/cube"
	"rubiks-solver/internal/solver"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the 54-character scrambled cube state:")

	// Read the scrambled string from the terminal
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if len(input) != 54 {
		fmt.Printf("Error: Expected 54 characters, got %d\n", len(input))
		return
	}

	// Initialize the cube state
	myCube := &cube.Cube{}
	copy(myCube.State[:], input)

	// Solve and output the result
	solution, err := solver.Solve(myCube)
	if err != nil {
		fmt.Println("Error solving cube:", err)
		return
	}

	fmt.Println("\nMinimum rotations required:")
	fmt.Println(strings.Join(solution, " "))
}