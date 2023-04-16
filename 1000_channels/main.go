package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	c := make(chan int)
	go func(c chan int) {
		log.Println("Start")
		c <- 1
		log.Println("Step")
		c <- 2
		log.Println("Stop")
	}(c)

	log.Println("Sleep 1")
	time.Sleep(1 * time.Second)
	fmt.Println(<-c)
	log.Println("Sleep 2")
	time.Sleep(1 * time.Second)
	fmt.Println(<-c)
	log.Println("Sleep 3")
	time.Sleep(1 * time.Second)

}
