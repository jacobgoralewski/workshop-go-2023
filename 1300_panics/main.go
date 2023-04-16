package main

import "fmt"

func Div(a, b float32) float32 {
	defer func() {
		r := recover()
		fmt.Println(r)
	}()
	if b == 0 {
		panic("Zero division")
	}
	return a / b
}

func main() {
	fmt.Println(Div(2, 2))
}
