package stack

import (
	"sync/atomic"
)

type Stack[T any] struct {
	head atomic.Pointer[Node[T]]
}

type Node[T any] struct {
	element *T
	next    *Node[T]
}

func NewTreiberStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func newNode[T any](element *T) *Node[T] {
	return &Node[T]{
		element: element,
		next:    nil,
	}
}

func (s *Stack[T]) Push(element T) {
	newElement := newNode[T](&element)

	for {
		currentHead := s.head.Load()
		newElement.next = currentHead

		if s.head.CompareAndSwap(currentHead, newElement) {
			return
		}
		// retry
	}
}

func (s *Stack[T]) Pop() *T {
	for {
		currentHead := s.head.Load()
		if currentHead == nil {
			return nil
		}

		next := currentHead.next
		if s.head.CompareAndSwap(currentHead, next) {
			return currentHead.element
		}

		// retry
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.head.Load() == nil
}
