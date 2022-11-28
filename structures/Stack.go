package structures

type Stack[T any] struct {
	Val []T
}

// Push on top of stack
func (s *Stack[T]) Push(value T) {
	s.Val = append(s.Val, value)
}

// Pop from top of stack
func (s *Stack[T]) Pop() T {
	l := len(s.Val) - 1
	toRemove := s.Val[l]
	s.Val = s.Val[:l]
	return toRemove
}
