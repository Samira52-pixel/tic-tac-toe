package main

import "fmt"

func main() {
	var value int8 = 127
	fmt.Println(value)
	value += 5
	fmt.Println(value)

	var value2 uint8 = 255
	fmt.Println(value2)
	value2 += 5
	fmt.Printf("value2 = %d\n", value2)

	fmt.Println("Max int8:", int8(^uint8(0)>>1))
	fmt.Println("Min int8:", -int8(^uint8(0)>>1)-1)
	fmt.Println("Max uint8:", uint8(^uint8(0)))
	fmt.Println("Min uint8:", uint8(0))

}
