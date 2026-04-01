package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rubiks-solver/internal/solver"
	"strings"
)

type SolveRequest struct {
	CubeString string `json:"cubeString"`
}

type SolveResponse struct {
	Solution string `json:"solution,omitempty"`
	Error    string `json:"error,omitempty"`
}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SolveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(SolveResponse{Error: "Invalid request payload"})
		return
	}

	input := strings.TrimSpace(strings.ToUpper(req.CubeString))

	// 1. Length Validation
	if len(input) != 54 {
		json.NewEncoder(w).Encode(SolveResponse{Error: fmt.Sprintf("Expected 54 characters, got %d", len(input))})
		return
	}

	// 2. Strict Facelet Count Validation
	counts := make(map[rune]int)
	for _, char := range input {
		counts[char]++
	}

	if len(counts) != 6 {
		json.NewEncoder(w).Encode(SolveResponse{Error: fmt.Sprintf("Found %d unique characters, expected exactly 6.", len(counts))})
		return
	}

	for char, count := range counts {
		if count != 9 {
			json.NewEncoder(w).Encode(SolveResponse{Error: fmt.Sprintf("Facelet '%c' appears %d times. Each must appear exactly 9 times.", char, count)})
			return
		}
	}

	// 3. Solve
	solution, err := solver.Solve(input)
	if err != nil {
		json.NewEncoder(w).Encode(SolveResponse{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(SolveResponse{Solution: solution})
}

func main() {
	// Serve the static HTML/CSS/JS files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// API Endpoint for the solver
	http.HandleFunc("/api/solve", solveHandler)

	port := "8080"
	fmt.Printf("🌐 Interactive Server running on http://localhost:%s\n", port)
	fmt.Println("Press Ctrl+C to stop.")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}