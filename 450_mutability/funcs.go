package main

import "fmt"

func mut(x *int) {
	y := *x
	y = y + 1
}

func main() {
	a := 1
	mut(&a)

	fmt.Println(a)
}
