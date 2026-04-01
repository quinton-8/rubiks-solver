package cube

import (
	"strings"
)
// Cube describes a cube state using Corner/Edge Permutation and Orientation (CEPO).
type Cube struct {
	CP    [8]int8  // Corner Permutation (0-7)
	CO    [8]int8  // Corner Orientation (0-2)
	EP    [12]int8 // Edge Permutation (0-11)
	EO    [12]int8 // Edge Orientation (0-1)
	Move  string   // Last move performed
	Move2 string   // Move before last
}

// NewSolvedCube returns a new solved cube.
func NewSolvedCube() *Cube {
	c := &Cube{}
	for i := range c.CP {
		c.CP[i] = int8(i)
	}
	for i := range c.EP {
		c.EP[i] = int8(i)
	}
	return c
}

// modulo3 handles orientation math: 0 = good, 1 = twisted clockwise, 2 = twisted anti-clockwise.
func modulo3(n int8) int8 {
	if n == -1 {
		return 2
	}
	return n % 3
}

// RotateUp (U) rotates the top face clockwise.
func (c *Cube) RotateUp() {
	// corner permutation
	tmpC := c.CP[0]
	c.CP[0] = c.CP[4]
	c.CP[4] = c.CP[3]
	c.CP[3] = c.CP[7]
	c.CP[7] = tmpC
	// edge permutation
	tmpE := c.EP[0]
	c.EP[0] = c.EP[8]
	c.EP[8] = c.EP[3]
	c.EP[3] = c.EP[11]
	c.EP[11] = tmpE
	// corner orientation
	c.CO[c.CP[0]] = modulo3(c.CO[c.CP[0]] - 1)
	c.CO[c.CP[3]] = modulo3(c.CO[c.CP[3]] - 1)
	c.CO[c.CP[4]] = modulo3(c.CO[c.CP[4]] + 1)
	c.CO[c.CP[7]] = modulo3(c.CO[c.CP[7]] + 1)
	// edge orientation
	c.EO[c.EP[0]] = (c.EO[c.EP[0]] + 1) % 2
	c.EO[c.EP[3]] = (c.EO[c.EP[3]] + 1) % 2
	c.EO[c.EP[8]] = (c.EO[c.EP[8]] + 1) % 2
	c.EO[c.EP[11]] = (c.EO[c.EP[11]] + 1) % 2
}

// RotateDown (D) rotates the bottom face clockwise.
func (c *Cube) RotateDown() {
	tmpC := c.CP[1]
	c.CP[1] = c.CP[5]
	c.CP[5] = c.CP[2]
	c.CP[2] = c.CP[6]
	c.CP[6] = tmpC
	tmpE := c.EP[1]
	c.EP[1] = c.EP[10]
	c.EP[10] = c.EP[2]
	c.EP[2] = c.EP[9]
	c.EP[9] = tmpE
	c.CO[c.CP[1]] = modulo3(c.CO[c.CP[1]] - 1)
	c.CO[c.CP[2]] = modulo3(c.CO[c.CP[2]] - 1)
	c.CO[c.CP[5]] = modulo3(c.CO[c.CP[5]] + 1)
	c.CO[c.CP[6]] = modulo3(c.CO[c.CP[6]] + 1)
	c.EO[c.EP[1]] = (c.EO[c.EP[1]] + 1) % 2
	c.EO[c.EP[10]] = (c.EO[c.EP[10]] + 1) % 2
	c.EO[c.EP[2]] = (c.EO[c.EP[2]] + 1) % 2
	c.EO[c.EP[9]] = (c.EO[c.EP[9]] + 1) % 2
}

// RotateFront (F) rotates the front face clockwise.
func (c *Cube) RotateFront() {
	tmpC := c.CP[4]
	c.CP[4] = c.CP[1]
	c.CP[1] = c.CP[6]
	c.CP[6] = c.CP[3]
	c.CP[3] = tmpC
	tmpE := c.EP[5]
	c.EP[5] = c.EP[9]
	c.EP[9] = c.EP[6]
	c.EP[6] = c.EP[8]
	c.EP[8] = tmpE
	c.CO[c.CP[1]] = modulo3(c.CO[c.CP[1]] + 1)
	c.CO[c.CP[3]] = modulo3(c.CO[c.CP[3]] + 1)
	c.CO[c.CP[4]] = modulo3(c.CO[c.CP[4]] - 1)
	c.CO[c.CP[6]] = modulo3(c.CO[c.CP[6]] - 1)
}

// RotateBack (B) rotates the back face clockwise.
func (c *Cube) RotateBack() {
	tmpC := c.CP[7]
	c.CP[7] = c.CP[2]
	c.CP[2] = c.CP[5]
	c.CP[5] = c.CP[0]
	c.CP[0] = tmpC
	tmpE := c.EP[7]
	c.EP[7] = c.EP[10]
	c.EP[10] = c.EP[4]
	c.EP[4] = c.EP[11]
	c.EP[11] = tmpE
	c.CO[c.CP[0]] = modulo3(c.CO[c.CP[0]] + 1)
	c.CO[c.CP[2]] = modulo3(c.CO[c.CP[2]] + 1)
	c.CO[c.CP[5]] = modulo3(c.CO[c.CP[5]] - 1)
	c.CO[c.CP[7]] = modulo3(c.CO[c.CP[7]] - 1)
}

// RotateLeft (L) rotates the left face clockwise.
func (c *Cube) RotateLeft() {
	tmpC := c.CP[0]
	c.CP[0] = c.CP[5]
	c.CP[5] = c.CP[1]
	c.CP[1] = c.CP[4]
	c.CP[4] = tmpC
	tmpE := c.EP[4]
	c.EP[4] = c.EP[1]
	c.EP[1] = c.EP[5]
	c.EP[5] = c.EP[0]
	c.EP[0] = tmpE
}

// RotateRight (R) rotates the right face clockwise.
func (c *Cube) RotateRight() {
	tmpC := c.CP[3]
	c.CP[3] = c.CP[6]
	c.CP[6] = c.CP[2]
	c.CP[2] = c.CP[7]
	c.CP[7] = tmpC
	tmpE := c.EP[6]
	c.EP[6] = c.EP[2]
	c.EP[2] = c.EP[7]
	c.EP[7] = c.EP[3]
	c.EP[3] = tmpE
}

func (c *Cube) ApplyMoves(movesStr string) {
	sequence := strings.Fields(movesStr)
	for _, move := range sequence {
		switch move {
		case "U": c.RotateUp()
		case "U2": c.RotateUp(); c.RotateUp()
		case "U'": c.RotateUp(); c.RotateUp(); c.RotateUp()
		// Repeat for D, L, R, F, B...
		}
		c.Move2 = c.Move
		c.Move = move
	}
}

// ListMoves returns valid moves for the current subgroup
func ListMoves(c *Cube, subgroup int8) []string {
	var moves []string
	if subgroup == 0 {
		moves = []string{"U", "U'", "U2", "D", "D'", "D2", "R", "R'", "R2", "L", "L'", "L2", "F", "F'", "F2", "B", "B'", "B2"}
	} else if subgroup == 3 {
		moves = []string{"U2", "D2", "R2", "L2", "F2", "B2"} // Only double turns
	}
	// ... logic for subgroups 1 and 2
	return moves
}