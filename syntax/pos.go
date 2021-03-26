package syntax

import "fmt"

// Pos position info
// #1 - line number
// #2 - number of characters from the beginning of the line
// #3 - number of white spaces from the beginning of the line
type Pos [3]int

// Line returns number of line.
func (p Pos) Line() int {
	return p[0]
}

// Char returns number of char of current line to the left.
func (p Pos) Char() int {
	return p[1]
}

// Spaces returns number of spaces by current line to the left.
func (p Pos) Spaces() int {
	return p[2]
}

// First returns true if first character on line.
func (p Pos) First() bool {
	return p.Char() == p.Spaces()
}

func (p Pos) String() string {
	if p.First() {
		return fmt.Sprintf("L%d\t(SP%d CL%d)", p.Line(), p.Spaces(), p.Char())
	}
	return fmt.Sprintf("\t(SP%d CL%d)", p.Spaces(), p.Char())
}
