package main

import (
	"bufio"
	"fmt"
	"os"
	"rubiks-solver/internal/solver"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Kociemba Rubik's Solver")
	fmt.Println("-----------------------")
	fmt.Println("Enter 54-char string (Order: U, R, F, D, L, B):")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToUpper(input))

	// 1. Length Validation
	if len(input) != 54 {
		fmt.Printf("\033[31mError: Expected 54 characters, got %d\033[0m\n", len(input))
		return
	}

	// 2. Strict Facelet Count Validation
	counts := make(map[rune]int)
	for _, char := range input {
		counts[char]++
	}

	if len(counts) != 6 {
		fmt.Printf("\033[31mError: Found %d unique characters, expected exactly 6 (e.g., U, R, F, D, L, B).\033[0m\n", len(counts))
		return
	}

	for char, count := range counts {
		if count != 9 {
			fmt.Printf("\033[31mValidation Error: Facelet '%c' appears %d times. Each of the 6 colors MUST appear exactly 9 times.\033[0m\n", char, count)
			return
		}
	}

	// 3. Solve using the internal/solver package
	solution, err := solver.Solve(input)
	if err != nil {
		fmt.Printf("\033[31mSolver Error: %v\033[0m\n", err)
		return
	}

	fmt.Printf("\n\033[32mSolution found (%d moves):\033[0m\n%s\n", len(strings.Fields(solution)), solution)
}