package stack

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Peek() string {
	return (*s)[len(*s)-1]
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(*s) - 1
	lastElement := (*s)[index]
	*s = (*s)[:index]

	return lastElement, true
}
