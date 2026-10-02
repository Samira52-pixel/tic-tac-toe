package main

import (
	"math"
	"testing"
)

func Test_simpson(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    float64
		b    float64
		n    int
		f    func(float64) float64
		want float64
	}{
		{"Тест1", 0, 2, 4, func(x float64) float64 { return x * x }, 8.0 / 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := simpson(tt.a, tt.b, tt.n, tt.f)
			// TODO: update the condition below to compare got with tt.want.
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("simpson() = %v, want %v", got, tt.want)
			}
		})
	}
}
