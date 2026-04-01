package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// rotFace rotates the 9 stickers of a specific face clockwise
func rotFace(s []byte, face int) {
	b := face * 9
	// Rotate corners
	s[b+0], s[b+2], s[b+8], s[b+6] = s[b+6], s[b+0], s[b+2], s[b+8]
	// Rotate edges
	s[b+1], s[b+5], s[b+7], s[b+3] = s[b+3], s[b+1], s[b+5], s[b+7]
}

func rotU(s []byte) {
	rotFace(s, 0)
	old := [3]byte{s[45], s[46], s[47]}
	s[45], s[46], s[47] = s[36], s[37], s[38] // L -> B
	s[36], s[37], s[38] = s[18], s[19], s[20] // F -> L
	s[18], s[19], s[20] = s[9], s[10], s[11]  // R -> F
	s[9], s[10], s[11] = old[0], old[1], old[2] // B -> R
}

func rotD(s []byte) {
	rotFace(s, 3)
	old := [3]byte{s[24], s[25], s[26]}
	s[24], s[25], s[26] = s[42], s[43], s[44] // L -> F
	s[42], s[43], s[44] = s[51], s[52], s[53] // B -> L
	s[51], s[52], s[53] = s[15], s[16], s[17] // R -> B
	s[15], s[16], s[17] = old[0], old[1], old[2] // F -> R
}

func rotF(s []byte) {
	rotFace(s, 2)
	old := [3]byte{s[6], s[7], s[8]}
	s[6], s[7], s[8] = s[44], s[41], s[38]    // L -> U
	s[44], s[41], s[38] = s[29], s[28], s[27] // D -> L
	s[29], s[28], s[27] = s[15], s[12], s[9]  // R -> D
	s[15], s[12], s[9] = old[2], old[1], old[0] // U -> R
}

func rotB(s []byte) {
	rotFace(s, 5)
	old := [3]byte{s[0], s[1], s[2]}
	s[0], s[1], s[2] = s[42], s[39], s[36]    // L -> U
	s[42], s[39], s[36] = s[35], s[34], s[33] // D -> L
	s[35], s[34], s[33] = s[17], s[14], s[11] // R -> D
	s[17], s[14], s[11] = old[2], old[1], old[0] // U -> R
}

func rotL(s []byte) {
	rotFace(s, 4)
	old := [3]byte{s[0], s[3], s[6]}
	s[0], s[3], s[6] = s[53], s[50], s[47]    // B -> U
	s[53], s[50], s[47] = s[27], s[30], s[33] // D -> B
	s[27], s[30], s[33] = s[18], s[21], s[24] // F -> D
	s[18], s[21], s[24] = old[0], old[1], old[2] // U -> F
}

func rotR(s []byte) {
	rotFace(s, 1)
	old := [3]byte{s[2], s[5], s[8]}
	s[2], s[5], s[8] = s[20], s[23], s[26]    // F -> U
	s[20], s[23], s[26] = s[29], s[32], s[35] // D -> F
	s[29], s[32], s[35] = s[51], s[48], s[45] // B -> D
	s[51], s[48], s[45] = old[0], old[1], old[2] // U -> B
}

func main() {
	// Initialize a solved string byte array
	state := []byte("UUUUUUUUURRRRRRRRRFFFFFFFFFDDDDDDDDDLLLLLLLLLBBBBBBBBB")
	
	// Create a randomizer (modern Go way)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	moves := []func([]byte){rotU, rotD, rotF, rotB, rotL, rotR}
	moveNames := []string{"U", "D", "F", "B", "L", "R"}

	var sequence []string

	// Apply 30 random moves
	for i := 0; i < 30; i++ {
		idx := rng.Intn(len(moves))
		moves[idx](state)
		sequence = append(sequence, moveNames[idx])
	}

	fmt.Printf("Generated Scramble Sequence:\n%s\n\n", strings.Join(sequence, " "))
	fmt.Println("Scrambled String (Copy this into main.go):")
	fmt.Printf("\033[36m%s\033[0m\n", string(state)) // Prints the string in Cyan!
}