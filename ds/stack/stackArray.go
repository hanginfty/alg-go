package stack

type StackArray struct {
	elements []any
}

func (s *StackArray) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *StackArray) Push(val any) {
	s.elements = append([]any{val}, s.elements)
}

func (s *StackArray) Pop() any {
	top := s.elements[0]
	s.elements = s.elements[1:]
	return top
}

func (s *StackArray) Peek() any {
	return s.elements[0]
}

func (s *StackArray) Len() int {
	return len(s.elements)
}
