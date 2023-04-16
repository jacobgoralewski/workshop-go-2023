package main

import "fmt"

type Node[T any] struct {
	value       T
	left, right *Node[T]
}

func NewNode[T any](value T, left, right *Node[T]) *Node[T] {
	return &Node[T]{
		value: value,
		left:  left,
		right: right,
	}
}

func (s Node[T]) String() string {
	return fmt.Sprintf("node<value:%v|left:%v|right:%v>", s.value, s.left, s.right)
}

func main() {
	tree := NewNode[string]("foo", nil, nil)
	tree.left = NewNode[string]("bar", nil, nil)
	tree.right = NewNode[string]("baz", nil, nil)

	fmt.Printf("%+v", tree)
}
