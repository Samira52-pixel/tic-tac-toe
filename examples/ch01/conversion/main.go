package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i1 int = 10
	var f1 float64 = float64(i1)

	fmt.Printf("i1 = %d, f1 = %f \n", i1, f1)
	var f2 float64 = 10.12
	var i2 int = int(f2)

	fmt.Printf("i2 = %d, f2 = %f \n", i2, f2)

	var i3 int = 123
	// var s1 string = fmt.Sprintf("%d", i3)
	var s1 string = strconv.Itoa(i3)
	fmt.Printf("i3 = %d, s1 = %s \n", i3, s1)

	var s2 string = "123"

	// i4, err := strconv.Atoi(s2)
	// if err != nil {
	// 	fmt.Println("Error converting string to int:", err)
	// }
	// fmt.Printf("s2 = %s, i4 = %d \n", s2, i4)

	if i4, err := strconv.Atoi(s2); err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Printf("s2 = %s, i4 = %d \n", s2, i4)
	}

}
