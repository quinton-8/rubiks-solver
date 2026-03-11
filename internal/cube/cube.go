package cube

import "fmt"

// Cube represents the 54 stickers of a 3x3 Rubik's Cube.
type Cube struct {
	State [54]byte
}

// NewSolvedCube initializes a cube in its solved state.
// U=White, R=Blue, F=Red, D=Yellow, L=Green, B=Orange
func NewSolvedCube() *Cube {
	solvedState := "WWWWWWWWWBBBBBBBBBRRRRRRRRRYYYYYYYYYGGGGGGGGGOOOOOOOOO"
	var c Cube
	copy(c.State[:], solvedState)
	return &c
}

// RotateRight performs a 90-degree clockwise turn of the Right (R) face.
func (c *Cube) RotateRight() {
	// A temporary copy to hold the old state during rotation
	old := c.State

	// 1. Rotate the Right face itself (indices 9 to 17)
	c.State[9], c.State[10], c.State[11] = old[15], old[12], old[9]
	c.State[12], c.State[14] = old[16], old[10]
	c.State[15], c.State[16], c.State[17] = old[17], old[14], old[11]

	// 2. Cycle the adjacent edges (Up, Back, Down, Front)
	// Top (Up) gets Front
	c.State[2], c.State[5], c.State[8] = old[20], old[23], old[26]
	// Back gets Top (reversed index mapping due to physical wrapping)
	c.State[45], c.State[48], c.State[51] = old[8], old[5], old[2]
	// Bottom (Down) gets Back
	c.State[29], c.State[32], c.State[35] = old[51], old[48], old[45]
	// Front gets Bottom
	c.State[20], c.State[23], c.State[26] = old[29], old[32], old[35]
}

// RotateLeft (L)
func (c *Cube) RotateLeft() {
	old := c.State
	// Face rotation
	c.State[36], c.State[37], c.State[38] = old[42], old[39], old[36]
	c.State[39], c.State[41] = old[43], old[37]
	c.State[42], c.State[43], c.State[44] = old[44], old[41], old[38]
	// Edges: Up -> Front -> Down -> Back -> Up
	c.State[0], c.State[3], c.State[6] = old[53], old[50], old[47]
	c.State[18], c.State[21], c.State[24] = old[0], old[3], old[6]
	c.State[27], c.State[30], c.State[33] = old[18], old[21], old[24]
	c.State[47], c.State[50], c.State[53] = old[33], old[30], old[27]
}

// RotateUp (U)
func (c *Cube) RotateUp() {
	old := c.State
	// Face rotation
	c.State[0], c.State[1], c.State[2] = old[6], old[3], old[0]
	c.State[3], c.State[5] = old[7], old[1]
	c.State[6], c.State[7], c.State[8] = old[8], old[5], old[2]
	// Edges: Front -> Left -> Back -> Right -> Front
	c.State[18], c.State[19], c.State[20] = old[9], old[10], old[11]
	c.State[36], c.State[37], c.State[38] = old[18], old[19], old[20]
	c.State[45], c.State[46], c.State[47] = old[36], old[37], old[38]
	c.State[9], c.State[10], c.State[11] = old[45], old[46], old[47]
}

// RotateDown (D)
func (c *Cube) RotateDown() {
	old := c.State
	// Face rotation
	c.State[27], c.State[28], c.State[29] = old[33], old[30], old[27]
	c.State[30], c.State[32] = old[34], old[28]
	c.State[33], c.State[34], c.State[35] = old[35], old[32], old[29]
	// Edges: Front -> Right -> Back -> Left -> Front
	c.State[24], c.State[25], c.State[26] = old[42], old[43], old[44]
	c.State[15], c.State[16], c.State[17] = old[24], old[25], old[26]
	c.State[51], c.State[52], c.State[53] = old[15], old[16], old[17]
	c.State[42], c.State[43], c.State[44] = old[51], old[52], old[53]
}

// RotateFront (F)
func (c *Cube) RotateFront() {
	old := c.State
	// Face rotation
	c.State[18], c.State[19], c.State[20] = old[24], old[21], old[18]
	c.State[21], c.State[23] = old[25], old[19]
	c.State[24], c.State[25], c.State[26] = old[26], old[23], old[20]
	// Edges: Up -> Right -> Down -> Left -> Up
	c.State[6], c.State[7], c.State[8] = old[44], old[41], old[38]
	c.State[9], c.State[12], old[15] = old[6], old[7], old[8]
	c.State[27], c.State[28], c.State[29] = old[15], old[12], old[9]
	c.State[38], c.State[41], c.State[44] = old[27], old[28], old[29]
}

// RotateBack (B)
func (c *Cube) RotateBack() {
	old := c.State
	// Face rotation
	c.State[45], c.State[46], c.State[47] = old[51], old[48], old[45]
	c.State[48], c.State[50] = old[52], old[46]
	c.State[51], c.State[52], c.State[53] = old[53], old[50], old[47]
	// Edges: Up -> Left -> Down -> Right -> Up
	c.State[0], c.State[1], c.State[2] = old[11], old[14], old[17]
	c.State[36], c.State[39], c.State[42] = old[0], old[1], old[2]
	c.State[33], c.State[34], c.State[35] = old[36], old[39], old[42]
	c.State[11], c.State[14], c.State[17] = old[33], old[34], old[35]
}

// Print outputs the current string state of the cube.
func (c *Cube) Print() {
	fmt.Println(string(c.State[:]))
}
