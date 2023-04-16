package main

import "fmt"

const (
	A int64 = iota + 2
	B
	_
	C
)

func main() {
	fmt.Printf("%d, %d, %d\n", A, B, C)
}
