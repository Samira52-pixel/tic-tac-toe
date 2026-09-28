package main

import (
	"fmt"
	"tic-tac-toe/examples/ch02/scope/p1"
)

func myfunc() {
	fmt.Println("MyFunc")
	// fmt.Println(a)
	b := 10
	fmt.Println(b)
}

func main() {
	a := 10
	myfunc()
	fmt.Println(a)

	myf1 := myfunc
	myf1()

	myf3 := func() {
		c := 10
		fmt.Println(c)
		fmt.Println(a)
	}

	myf3()

	func() {
		c := 10
		fmt.Println(c)
		fmt.Println(a)
	}()

	//fmt.Println(c)

	{
		c2 := 10
		fmt.Println(c2)
		fmt.Println(a)
	}

	p1.MyFuncPublic()
	//p1.myFuncPrivate()
	var v1 p1.MyTypePublic
	_ = v1
	// var v2 p1.myType

	fmt.Println(p1.MyDataP)
}
