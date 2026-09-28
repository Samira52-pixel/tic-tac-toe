package main

import "fmt"

func fib(n int) int {
	var value int
	if n <= 1 {
		value = n
	} else {
		prev := 0
		curr := 1
		i := 2
		for i <= n {
			next := prev + curr
			prev = curr
			curr = next
			i = i + 1
		}
		value = curr
	}
	return value
}
func main() {
	var n int
	fmt.Println("Введите n")
	fmt.Scan(&n)
	if n < 0 {
		return
	}
	result := fib(n)
	fmt.Printf("%d = %d\n", n, result)

}
