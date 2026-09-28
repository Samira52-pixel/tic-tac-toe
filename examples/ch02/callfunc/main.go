package main

import "fmt"

func inc(a int) int {
	return a + 1
}

func dec(a int) int {
	return a - 1
}

func square(a int) int {
	return a * a
}

func callFunc(caller func(int) int, b int) {
	c := caller(b)
	fmt.Println(c)
}

func StringerF(operation func(a string, b string) string, del string, strs []string) string {
	var result string
	if len(strs) == 0 {
		return result
	}

	if len(strs) == 1 {
		result = strs[0]
		return result
	}

	var i, j int = 0, 1
	for i < len(strs)-1 && j < len(strs) {
		result = result + del + operation(strs[i], strs[j])
		i++
		j++
	}

	return result

}

func main() {
	// callFunc(inc, 10)
	// callFunc(dec, 112)
	// callFunc(func(a int) int { return 2 * a }, 10)
	// callFunc(square, 5)
	// callFunc(func(d int) int { return d * 100 }, 12)

	result := StringerF(func(a, b string) string { return a + ", " + b }, " ", []string{"a1", "a2", "a3"})
	fmt.Println(result)
}
