package main

import "testing"

type S struct {
	a int
}

func FibR(n uint) uint {
	if n < 2 {
		return n
	}
	return FibR(n-1) + FibR(n-2)
}

func FibI(n uint) uint {
	if n < 2 {
		return n
	}

	var a, b uint
	b = 1
	for n--; n > 0; n-- {
		a += b
		a, b = b, a
	}

	return b
}

// from fib_test.go
func BenchmarkFibR10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibR(uint(n) % 10)
	}
}

// from fib_test.go
func BenchmarkFibI10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibI(uint(n) % 10)
	}
}
