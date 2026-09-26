package main

import "fmt"

// Задание.
// Объяви переменную score типа int и задай ей любое число от 0 до 100.
// С помощью if / else if / else выведи ровно одну строку:
//   score < 60            -> "незачёт"
//   score от 60 до 89     -> "зачёт"
//   score >= 90           -> "отлично"
// Если score меньше 0 или больше 100, выведи "неверная оценка" и ничего больше.

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
