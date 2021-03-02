package syntax

type IntStack []int

func (s *IntStack) Push(in int) {
	*s = append(*s, in)
}

func (s IntStack) Top() int {
	idx := len(s) - 1
	if idx < 0 {
		return -1
	}
	return s[idx]
}

func (s *IntStack) Pop() int {
	idx := len(*s) - 1
	if idx < 0 {
		return -1
	}

	top := (*s)[idx]
	*s = (*s)[:idx]

	return top
}
