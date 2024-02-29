package graham

type Stack[T any] struct {
	stack []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{stack: make([]T, 0)}
}

func (s *Stack[T]) Push(x T) {
	s.stack = append(s.stack, x)
}

func (s *Stack[T]) Pop() T {
	x := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return x
}

func (s *Stack[T]) Top() T {
	return s.stack[len(s.stack)-1]
}

func (s *Stack[T]) SemiTop() T {
	if len(s.stack) == 1 {
		panic("stack length is 1")
	}
	return s.stack[len(s.stack)-2]
}

func (s *Stack[T]) Len() int {
	return len(s.stack)
}

func (s *Stack[T]) Empty() bool {
	return len(s.stack) == 0
}

func (s *Stack[T]) Slice() []T {
	return s.stack
}
