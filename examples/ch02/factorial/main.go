package main

import "fmt"

func factorial(n int) int {
	acc := 1
	for n > 1 {
		acc = acc * n
		n = n - 1
	}
	return acc
}
func main() {
	var n int
	fmt.Println("Введите n")
	fmt.Scan(&n)

	if n < 0 {
		return
	}
	result := factorial(n)
	fmt.Printf("%d! = %d\n", n, result)
}
