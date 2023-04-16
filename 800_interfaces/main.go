package main

import (
	"fmt"
)

type Greeter interface {
	Hello() string
}

type EnglishGreeter struct{}

func (eg EnglishGreeter) Hello() string {
	return "Hello world!"
}

type SpanishGreeter struct{}

func (sg SpanishGreeter) Hello() string {
	return "Hola Mundo!"
}

func Greet(greeter Greeter) {
	fmt.Println(greeter.Hello())
}

func main() {
	g := SpanishGreeter{}
	Greet(g)
}
