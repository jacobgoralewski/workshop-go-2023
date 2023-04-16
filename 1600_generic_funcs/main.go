package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strconv"
)

type X = int

type Foo struct {
	bar int
}

func (f Foo) String() string {
	return strconv.Itoa(f.bar)
}

func Add[T constraints.Ordered](a, b T) T {
	return a + b
}

func ToString[T fmt.Stringer](a T) string {
	return a.String()
}

func Baz[T int](x T) {
	fmt.Println(x)
}

func main() {
	fmt.Println(Add[int](2, 3))
	fmt.Printf("%v\n", ToString(Foo{bar: 3}))

	Baz(X(1))
}
