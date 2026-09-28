package main

import "fmt"

func incProducer(inc float64) func(float64) float64 {
	return func(a float64) float64 {
		return a + inc
	}
}

func main() {
	inc1 := incProducer(1)
	r1 := inc1(10)
	fmt.Println(r1)

	inc10 := incProducer(10)
	r10 := inc10(10)
	fmt.Println(r10)

	fmt.Println(inc10(12))
	fmt.Println(inc1(12))
}
