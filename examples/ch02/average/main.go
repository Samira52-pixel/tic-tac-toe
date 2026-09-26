package main

import (
	"fmt"
)

// average - она принимает набор целых чисел и возвращает их среднее значение.

func average(values ...float64) float64 {
	sum := 0.0
	for _, n := range values {
		sum += n
	}

	f := sum / float64(len(values))

	return f
}

// average2 - она принимает два целых числа и возвращает их среднее значение.
func average2(a, b float64) float64 {
	// (1 + 2) / 2 = 1,5
	c := (a + b) / 2
	return c
}

// average3
func average3(a, b, c float64) float64 {
	f := (a + b + c) / 3
	return f
}

// average4
func average4(a, b, c, d float64) float64 {
	f := (a + b + c + d) / 4
	return f
}

func average5(values []float64) float64 {
	sum := 0.0
	for _, n := range values {
		sum += n
	}

	f := sum / float64(len(values))

	return f
}

func main() {
	slices := []float64{1, 2, 3, 4, 5}
	a5 := average5(slices)

	a4 := average(1, 2, 3, 4)
	a3 := average(1, 2, 3)
	a2 := average(1, 2)
	a1 := average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(a2, a3, a4, a1, a5)
}
