package solver

import (
	"strings"
	"rubiks-solver/internal/cube"
	"github.com/daosyn/kociemba" 
	"fmt"
)

func Solve(startCube *cube.Cube) ([]string, error) {
	stateStr := string(startCube.State[:])

	// The library returns a single string of moves (e.g., "R U R' U'")
	solution := kociemba.Solve(stateStr) 

	fmt.Printf("Raw library output: '%s'\n", solution)

	if solution == "" {
		return []string{}, nil
	}

	// Split the solution string into individual moves for the terminal output
	return strings.Split(strings.TrimSpace(solution), " "), nil
}