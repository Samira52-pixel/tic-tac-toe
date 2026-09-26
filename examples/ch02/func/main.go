package main

import "fmt"

func main() {

	// 3 + 5 = 8
	var i1 int = 3
	i2 := 5
	result := add(i1, i2)
	fmt.Println("Result:", result)

	added := add2(10, 20)
	fmt.Println("Added:", added)

	added2 := add3(1, 2, 3, 4, 5)
	fmt.Println("Added2:", added2)
	added3 := add3(1, 2)
	fmt.Println("Added3:", added3)

	values := []int{1, 2, 3, 4, 5}
	added4 := add3(values...)
	fmt.Println("Added4:", added4)

	// Пример анонимной функции
	sum2 := func(a, b int) int {
		return a + b
	}(i1, i2)

	fmt.Println("Sum2:", sum2)

}
