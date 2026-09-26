package main

import (
	"fmt"
	"math"
)

// mymin
// min(1,2,4,5) = 1
func mymin(values ...int) int {
	result := math.MaxInt64
	for _, num := range values {
		if num < result {
			result = num
		}
	}
	return result
}
func mymin1(values []int) int {
	result := math.MaxInt64
	for _, num := range values {
		if num < result {
			result = num
		}
	}
	return result
}

// mymax
// max(1,2,4,5) = 5
func mymax(values ...int) int {
	result := math.MinInt64
	for _, num := range values {
		if num > result {
			result = num
		}
	}
	return result
}
func mymax1(values []int) int {
	result := math.MinInt64
	for _, num := range values {
		if num > result {
			result = num
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	m5 := mymin1(slice)
	m4 := mymax1(slice)
	m1 := mymax(1, 2)
	m2 := mymax(1, -2, 3)
	m3 := mymax(1, 2, 3, -4)
	fmt.Println(m1, m2, m3, m4, m5)
}
