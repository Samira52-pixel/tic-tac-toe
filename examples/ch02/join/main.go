// join("a", "Привет", "b") => "aПриветb"
// join("a", "Привет", "b") => "a Привет b"
// fmt.Sprint("%s%s", "a", "b") => "ab"
// "a" + "b" => "ab" // Конкатенация строк

package main

import "fmt"

func join(values ...string) string {
	result := ""
	for index, value := range values {
		if index == len(values)-1 {
			result = result + value
		} else {
			result = result + value + " "
		}
	}
	return result
}

func join2(values []string) string {
	result := ""
	for index, value := range values {
		if index == len(values)-1 {
			result = result + value
		} else {
			result = result + value + " "
		}
	}
	return result
}

func main() {

	slice := []string{"a", "Привет", "b"}
	a1 := join2(slice)

	a2 := join("a1", "a2", "a3", "a4")
	fmt.Println(a1, a2)
}
