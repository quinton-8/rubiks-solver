package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rubiks-solver/internal/cube"
	"rubiks-solver/internal/solver"
	
)

type SolveRequest struct {
	Scramble string `json:"scramble"`
}

type SolveResponse struct {
	Moves []string `json:"moves"`
}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure it's a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. Decode the incoming JSON request
	var req SolveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 2. Validate the scramble string length (must be exactly 54 characters)
	if len(req.Scramble) != 54 {
		http.Error(w, "Scramble string must be exactly 54 characters long", http.StatusBadRequest)
		return
	}

	// 3. Initialize the cube with the user's scramble string
	myCube := &cube.Cube{}
	copy(myCube.State[:], req.Scramble)

	// 4. Pass it to the solver
	solutionMoves, err := solver.Solve(myCube)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 5. Send back the solution as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SolveResponse{Moves: solutionMoves})
}

func main() {
	myCube := cube.NewSolvedCube()

	// Perform a sequence of moves (The Scramble)
	myCube.RotateRight()
	myCube.RotateUp()
	myCube.RotateFront()
	myCube.RotateLeft()

	// Pass this scrambled cube to your library-based solver
	solution, _ := solver.Solve(myCube)

	fmt.Println("Scrambled State:", string(myCube.State[:]))
	fmt.Println("Solution to get back to solved:", solution)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/solve", solveHandler)

	fmt.Println("Rubik's Solver API running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
