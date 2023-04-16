package main

import (
	"time"
)

func saySync(c chan int, s string, delay int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		println(s)
	}

	c <- 1
}

func main() {
	c := make(chan int)

	go saySync(c, "Hello", 3000)
	go saySync(c, "Hola", 1000)

	<-c
	<-c
}

//func main() {
//	go say("Hello", 3000)
//	go say("Hola", 1000)
//
//	time.Sleep(30 * time.Second)
//}
