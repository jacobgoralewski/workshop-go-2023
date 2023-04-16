package main

import "fmt"

func foo() {
	s2 := "hello"
	s := &s2
	defer fmt.Printf("1 %s", *s)
	defer fmt.Printf("2 %s", *s)
	defer fmt.Printf("3 %s", *s)
}

func main() {
	foo()
}
