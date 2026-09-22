package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	value := rand.IntN(2) + 1
	fmt.Println(value)
}
