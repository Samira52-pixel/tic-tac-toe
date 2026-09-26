package main

import "fmt"

func main() {
	var a1 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a1)
	a1[3] = 100
	fmt.Println(a1)
	//a1 = append(a1, 10)

	s1 := []int{1, 23, 10}
	fmt.Println(s1, len(s1))
	s1 = append(s1, 14)
	fmt.Println(s1, len(s1))
}
