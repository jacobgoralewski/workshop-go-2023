package main

import "fmt"

func foo(a *int) {
	*a++
	fmt.Println(*a)
}


func main() {
	a := 1
	fmt.Println(a)
	foo(&a)
	fmt.Println(a)
}
