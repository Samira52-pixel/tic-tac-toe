package main

import (
	"slices"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    []int
		want []int
	}{
		{"test1", []int{5, 1, 4, 2}, []int{1, 2, 4, 5}},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			bubbleSort(tt.a)
			if !slices.Equal(tt.a, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", tt.a, tt.want)
			}
		})
	}
}
