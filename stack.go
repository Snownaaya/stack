package structs

import "cmp"

type MinStack[T cmp.Ordered] struct {
	items []T
}

func Constructor[T cmp.Ordered]() MinStack[T] {
	return MinStack[T]{items: []T{}}
}

func (s *MinStack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *MinStack[T]) Pop() (T, bool) {
	var zero T

	if s.IsEmpty() {
		return zero, false
	}

	var item T = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item, true
}

func (s *MinStack[T]) Top() T {
	return s.items[len(s.items)-1]
}

func (s *MinStack[T]) getMin() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}

	var min T = s.items[0]

	for i := 1; i < len(s.items); i++ {
		if s.items[i] < min {
			min = s.items[i]
		}
	}

	return min
}

func (s *MinStack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
