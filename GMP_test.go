package main

import (
	"testing"
)

func top(max int) int {
	total := 0
	for i := 0; i < max; i++ {
		total += i
	}
	return total
}
func demo1() {
	defer func() {
		top(10)
	}()
	top(100)
}
func demo2() {
	top(100)
	top(10)
}
func BenchmarkDemo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		demo1()
	}
}
func BenchmarkDemo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		demo2()
	}
}
