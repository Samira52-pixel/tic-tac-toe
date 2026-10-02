package main

import "testing"

func Test_factorial(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want int
	}{
		{"test1", 0, 1},
		{"test2", 1, 1},
		{"test3", 5, 120},
		{"test4", 3, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := factorial(tt.n)
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_factorial_negative(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test1", 1, 100},
		{"test2", 2, 21},
		{"test3", 0, 0},
		{"test4", 5, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := factorial(tt.n)
			if tt.want == got {
				t.Errorf("factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
