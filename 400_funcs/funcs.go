package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mut(x *int) {
	y := *x
	y = y + 1
}

func main() {
	//res := add(2, 3)

	a := 1
	mut(&a)

	fmt.Println(a)
}
