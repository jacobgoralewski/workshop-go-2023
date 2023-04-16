package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func div(a, b float64) (float64, error) {
	if b == 0.0 {
		return math.Inf(1), errors.New("zero division")
	}
	return a / b, nil
}

func main() {
	res, err := div(2, 0)
	if err != nil {
		log.Printf("Err: %v\n", err)
	}

	fmt.Println(res)
}
