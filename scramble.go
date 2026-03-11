package main

import (
	"fmt"
	"math/rand"
	"time"

	"rubiks-solver/internal/cube"
)

func main() {
	// Initialize a solved cube
	myCube := cube.NewSolvedCube()
	rand.Seed(time.Now().UnixNano())

	// List of all available rotation methods
	moves := []func(){
		myCube.RotateRight, myCube.RotateLeft,
		myCube.RotateUp, myCube.RotateDown,
		myCube.RotateFront, myCube.RotateBack,
	}

	// Apply 20 random moves to ensure a complex, valid scramble
	for i := 0; i < 20; i++ {
		moveIndex := rand.Intn(len(moves))
		moves[moveIndex]()
	}

	fmt.Println("Copy this scrambled string into your solver:")
	fmt.Println(string(myCube.State[:]))
}
