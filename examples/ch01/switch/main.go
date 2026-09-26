package main

import "fmt"

/*
func main() {
	score := 61
	if score < 60 {
		fmt.Println("незачёт")
	} else if score >= 60 && score <= 89 {
		fmt.Println("зачёт")
	} else if score >= 90 {
		fmt.Println("отлично")
	}
}

*/

func main() {

	score := 61
	switch {
	case score < 60:
		fmt.Println("незачёт")
	case score >= 60 && score <= 89:
		fmt.Println("зачёт")
	case score >= 90:
		fmt.Println("отлично")

	}
}
