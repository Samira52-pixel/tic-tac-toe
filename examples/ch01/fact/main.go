package main

import "fmt"

// !5 = 1 * 2 * 3 * 4 *5
// Расчет факториала с помощью рекурсии
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func main() {
	f1 := Factorial(5)
	fmt.Println(f1)
}
