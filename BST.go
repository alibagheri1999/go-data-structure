package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"log"
)

type TNode[T constraints.Ordered] struct {
	Value T
	Left  *TNode[T]
	Right *TNode[T]
}

type BST[T constraints.Ordered] struct {
	root *TNode[T]
}

func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{
		root: nil,
	}
}

func (b *BST[T]) Insert(t T) {
	var insert func(n **TNode[T], node *TNode[T])
	insert = func(n **TNode[T], node *TNode[T]) {
		if *n == nil {
			*n = node
			return
		}
		if node.Value < (*n).Value {
			insert(&(*n).Left, node)
		} else {
			insert(&(*n).Right, node)
		}
	}
	insert(&b.root, &TNode[T]{Value: t})
}

func (b *BST[T]) String() string {
	var result string
	var order func(n *TNode[T])
	order = func(n *TNode[T]) {
		if n == nil {
			return
		}
		if n.Left != nil {
			order(n.Left)
		}
		result += fmt.Sprint(n.Value, " ")
		if n.Right != nil {
			order(n.Right)
		}
	}
	order(b.root)
	return result
}

func (b *BST[T]) Find(t T) *TNode[T] {
	var find func(n *TNode[T], t T) *TNode[T]
	find = func(n *TNode[T], t T) *TNode[T] {
		if n == nil {
			return nil
		}
		if n.Value == t {
			return n
		}
		if n.Value > t {
			return find(n.Left, t)
		} else {
			return find(n.Right, t)
		}
	}
	return find(b.root, t)
}

func main() {
	bst := NewBST[int]()
	bst.Insert(5)
	bst.Insert(2)
	bst.Insert(8)
	bst.Insert(1)
	bst.Insert(13)
	bst.Insert(7)
	bst.Insert(-1)
	bst.Insert(11)
	log.Println(bst)
}
