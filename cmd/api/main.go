package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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

	if len(input) != 54 {
		json.NewEncoder(w).Encode(SolveResponse{Error: fmt.Sprintf("Expected 54 characters, got %d", len(input))})
		return
	}

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

	solution, err := solver.Solve(input)
	if err != nil {
		json.NewEncoder(w).Encode(SolveResponse{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(SolveResponse{Solution: solution})
}

func main() {
	// SANITY CHECK: Does the static folder actually exist where Go is looking?
	if _, err := os.Stat("static/index.html"); os.IsNotExist(err) {
		fmt.Println("\033[31m🚨 CRITICAL ERROR: Cannot find static/index.html\033[0m")
		fmt.Println("Make sure you are in the root 'rubiks-solver' directory, NOT inside 'cmd' or 'api'!")
		fmt.Println("Your current folder structure MUST look like this:")
		fmt.Println("  rubiks-solver/")
		fmt.Println("  ├── cmd/")
		fmt.Println("  └── static/")
		os.Exit(1)
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/api/solve", solveHandler)

	port := "8080"
	fmt.Printf("\n🟢 SUCCESS: Found static files!\n")
	fmt.Printf("🌐 Interactive Server running on http://localhost:%s\n", port)
	fmt.Println("Press Ctrl+C to stop.")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}