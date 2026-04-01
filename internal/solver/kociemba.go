package solver

import (
	"fmt"
	"os/exec"
	"strings"
)

// Solve bridges your Go CLI to the official Python Kociemba engine
func Solve(cubeString string) (string, error) {
	if len(cubeString) != 54 {
		return "", fmt.Errorf("invalid cube string length: expected 54, got %d", len(cubeString))
	}

	// Executes the python 'kociemba' CLI tool in the background
	cmd := exec.Command("kociemba", cubeString)
	output, err := cmd.CombinedOutput()
	solution := strings.TrimSpace(string(output))
	
	if err != nil {
		// If the engine throws a ValueError, the cube state is physically impossible
		if strings.Contains(solution, "ValueError") {
			return "", fmt.Errorf("the cube is mathematically unsolvable (invalid parity or flipped edge)")
		}
		// If the command fails entirely
		return "", fmt.Errorf("engine failure (check pip installation): %v - %s", err, solution)
	}

	// Fix: The Python engine outputs nothing if the cube is ALREADY solved
	if solution == "" {
		return "Cube is already solved! (0 moves)", nil
	}

	return solution, nil
}