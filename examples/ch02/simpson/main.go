package main

import "fmt"

func simpson(a, b float64, n int, f func(float64) float64) float64 {
	h := (b - a) / float64(n)
	sum := f(a) + f(b)
	i := 1
	for i < n {
		x := a + float64(i)*h
		if i%2 != 0 {
			sum = sum + 4*f(x)
		} else {
			sum = sum + 2*f(x)
		}
		i = i + 1
	}
	area := sum * h / 3
	return area

}
func main() {
	var a, b float64
	var n int
	fmt.Println("Введи a, b и чётное n")
	fmt.Scan(&a, &b, &n)
	if n < 2 || n%2 != 0 {
		fmt.Println("Ошибка ввода")
		return
	}
	f := func(x float64) float64 {
		return x * x
	}
	area := simpson(a, b, n, f)
	fmt.Printf("area = %v\n", area)

}
