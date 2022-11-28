package structures

type Queue[T any] struct {
	Data []T
}

// Enqueue to end of queue
func (s *Queue[T]) Enqueue(value T) {
	s.Data = append(s.Data, value)
}

// Dequeue from from of queue
func (s *Queue[T]) Dequeue() T {
	toRemove := s.Data[0]
	s.Data = s.Data[1:]
	return toRemove
}
