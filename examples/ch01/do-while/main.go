package main

import "fmt"

func main() {

	arguments := []string{"arg1", "arg2", "arg3"}
	index := 0
	// Цикл с постусловием
	for {
		fmt.Println(arguments[index])
		index++
		if index >= len(arguments) {
			break
		}
	}

}
