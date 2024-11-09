package main

import (
	"errors"
	"golang.org/x/exp/constraints"
	"log"
)

type Queuer[T constraints.Ordered] interface {
	Enqueue(T)
	Dequeue() (T, error)
	IsEmpty() bool
}

type Queue[T constraints.Ordered] struct {
	data []T
}

func NewQueue[T constraints.Ordered]() Queuer[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

func (s *Queue[T]) Enqueue(item T) {
	s.data = append(s.data, item)
}
func (s *Queue[T]) Dequeue() (T, error) {
	var t T
	if s.IsEmpty() {
		return t, errors.New("the queue is empty")
	}
	item := s.data[0]
	s.data = s.data[1:]

	return item, nil
}
func (s *Queue[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func main() {
	queue := NewQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	log.Println(queue.Dequeue())
	log.Println(queue.Dequeue())
	log.Println(queue.Dequeue())
	log.Println(queue.Dequeue())
}
