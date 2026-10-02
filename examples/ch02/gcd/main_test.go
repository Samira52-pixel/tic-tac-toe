package main

import "testing"

func Test_gcd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"обычный", 48, 18, 6},
		{"аргументы наоборот", 18, 48, 6},
		{"равные", 7, 7, 7},
		{"один делит другой", 100, 25, 25},
		{"взаимно простые", 17, 13, 1},
		{"единица", 1, 100, 1},
		{"ноль вторым", 5, 0, 5},
		{"ноль первым", 0, 5, 5},
		{"оба нуля", 0, 0, 0},
		{"степени двойки", 64, 24, 8},
		{"классический пример", 1071, 462, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gcd(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			} else {
				t.Logf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}
