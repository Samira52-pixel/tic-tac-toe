package main

// add функция с аргументами и возвращаемым значением
func add(a int, b int) int {
	c := a + b
	return c
}

// add2 функция с именованным возвращаемым значением
func add2(a, b int) (c int) {
	c = a + b
	return
}

func add3(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}
