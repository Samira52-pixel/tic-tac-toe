package main

import "fmt"

func main() {

	arguments := []string{"arg1", "arg2", "arg3"}
	index := 0

	// Пример цикла с предусловием (while-do)
	for {
		if index >= len(arguments) {
			break
		}

		fmt.Println(arguments[index])
		index++
	}

}
