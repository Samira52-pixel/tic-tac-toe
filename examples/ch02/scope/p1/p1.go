package p1

import "fmt"

func myFuncPrivate() {
	fmt.Println("myFuncPrivate")
}

func MyFuncPublic() {
	myFuncPrivate()
	fmt.Println("MyFuncPublic")
}

type myType int
type MyTypePublic int

const myData int = 10
const MyDataP int = 100
