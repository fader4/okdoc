package starlark

type stackChars []int

func (s *stackChars) push(in int) {
	*s = append(*s, in)
}

func (s stackChars) top() int {
	idx := len(s) - 1
	if idx < 0 {
		return -1
	}
	return s[idx]
}

func (s *stackChars) pop() int {
	idx := len(*s) - 1
	if idx < 0 {
		return -1
	}

	top := (*s)[idx]
	*s = (*s)[:idx]

	return top
}

// file position - file line
// #1 - line number
// #2 - number of characters from the beginning of the line
// #3 - number of white spaces from the beginning of the line
type pos [3]int

func (p pos) Line() int {
	return p[0]
}

func (p pos) Char() int {
	return p[1]
}

func (p pos) Spaces() int {
	return p[2]
}

// First returns true if first character on line.
func (p pos) First() bool {
	return p.Char() == p.Spaces()
}
