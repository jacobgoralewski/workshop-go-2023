package main

import (
	"fmt"
)

func forLikeInPrimarySchool() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forButMoreLikeWhileButStillFor() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func forRange() {
	r := []int{5, 4, 3, 2, 1}
	for _, v := range r {
		fmt.Printf("k: %d, v: %d\n", v)
	}
}

func main() {

}