package main

import "fmt"

func myf1(x *int) {
	fmt.Println("myf1:", *x)
	*x = 100
}
func myf2(x int) {
	fmt.Println("myf2:", x)
	x = 200
}

func main() {
	x := 10
	myf1(&x)
	fmt.Println("main:", x)
	myf2(x)
	fmt.Println("main:", x)
}
