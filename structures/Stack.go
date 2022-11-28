package structures

type Stack[T any] struct {
	Data []T
}

// Push on top of stack
func (s *Stack[T]) Push(value T) {
	s.Data = append(s.Data, value)
}

// Pop from top of stack
func (s *Stack[T]) Pop() T {
	l := len(s.Data) - 1
	toRemove := s.Data[l]
	s.Data = s.Data[:l]
	return toRemove
}
