package set

import "fmt"

type Set struct {
	data map[int]bool
}

func NewSet() *Set {
	return &Set{
		data: make(map[int]bool),
	}
}

func (s *Set) Add(item int) {
	s.data[item] = true
}

func (s *Set) Remove(item int) {
	delete(s.data, item)
}

func (s *Set) Contains(item int) bool {
	return s.data[item]
}

func (s *Set) Len() int {
	return len(s.data)
}

func (s *Set) Clear() {
	s.data = make(map[int]bool)
}

func (s *Set) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Set) PrintSet() {
	for item := range s.data {
		fmt.Println(item)
	}
}

func (s *Set) Union(another *Set) *Set {
	union := NewSet()
	for item := range s.data {
		union.Add(item)
	}

	for item := range another.data {
		union.Add(item)
	}
	return union
}
