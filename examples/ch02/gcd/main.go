package main

import (
	"fmt"
)

func abs(h int) int {
	if h < 0 {
		return -h
	}
	return h
}
func gcd(a, b int) int {
	for b != 0 {
		r := a % b
		a = b
		b = r
	}
	return a

}
func main() {
	var a, b int
	fmt.Println("Введите a и b")
	fmt.Scan(&a, &b)
	a = abs(a)
	b = abs(b)
	result := gcd(a, b)

	fmt.Printf("%d и %d = %d\n", a, b, result)

}
