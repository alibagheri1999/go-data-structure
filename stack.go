package main

import (
	"errors"
	"golang.org/x/exp/constraints"
	"log"
)

type Stacker[T constraints.Ordered] interface {
	Push(t T)
	Pop() (T, error)
	IsEmpty() bool
}

type Stack[T constraints.Ordered] struct {
	data []T
}

func NewStack[T constraints.Ordered]() Stacker[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}
func (s *Stack[T]) Pop() (T, error) {
	var t T
	if s.IsEmpty() {
		return t, errors.New("the stack is empty")
	}
	index := len(s.data) - 1
	item := s.data[index]
	s.data = s.data[:index]

	return item, nil
}
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func main() {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	log.Println(stack.Pop())
	log.Println(stack.Pop())
	log.Println(stack.Pop())
	log.Println(stack.Pop())
}
