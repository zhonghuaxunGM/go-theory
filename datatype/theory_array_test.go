package main

import (
	"testing"
)

func Test_cap1(t *testing.T) {
	cap1()
}

func Test_length(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lengthNotCalculate()
		})
	}
}

func Test_sliceLength(t *testing.T) {
	// tests := []struct {
	// 	name string
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	sliceLength()
	// 	})
	// }
}

func Test_sliceLength2(t *testing.T) {
	// tests := []struct {
	// 	name string
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	sliceLength2()
	// })
	// }
}
