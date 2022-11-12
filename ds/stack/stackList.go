package stack

type Node struct {
	Val  any
	Next *Node
}

type Stack struct {
	top *Node
	len int
}

func (s *Stack) Push(val any) {
	node := &Node{val, s.top}
	s.top = node
	s.len++
}

func (s *Stack) Pop() any {
	top := s.top
	s.top = top.Next
	return top.Val
}

func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) Peek() any {
	top := s.top
	return top.Val
}
