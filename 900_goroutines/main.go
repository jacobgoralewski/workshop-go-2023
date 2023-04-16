package main

import (
	"time"
)

func say(s string, delay int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		println(s)
	}
}

func _main() {
	say("Hello", 3000)
	say("Hola", 1000)
}

//func main() {
//	go say("Hello", 3000)
//	go say("Hola", 1000)
//
//	time.Sleep(30 * time.Second)
//}
