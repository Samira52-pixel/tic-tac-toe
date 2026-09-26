package main

import "fmt"

func main() {
	var b1 bool = false
	var b2 bool = false
	b3 := b1 || b2
	fmt.Println(b3)
	b4 := b1 && b2
	fmt.Println(b4)
}
