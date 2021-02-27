package okdoc

type Stack [][2]int

func (s *Stack) Push(a, b int) {
	*s = append(*s, [2]int{a, b})
}

func (s *Stack) Pop() (int, int) {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val[0], val[1]
}
