package main

import (
	"fmt"
)

func average(n int) (float64, string, error) {
	if n < 1 {
		fmt.Println("ошибка ввода")
		return 0, "", fmt.Errorf("Ошибка ввода n = %d", n)
	}
	sum := 0

	for i := 1; i <= n; i++ {
		var grade int
		fmt.Println("Введи grade")
		fmt.Scan(&grade)
		if grade < 1 || grade > 10 {
			fmt.Println("ошибка ввода")
			return 0, "", fmt.Errorf("Ошибка ввода grade = %d", grade)
		}
		sum = sum + grade

	}
	average := sum / n

	var verdict string
	if average >= 8 {
		verdict = "рекомендуем"
	} else if average >= 5 {
		verdict = "можно смотреть"
	} else {
		verdict = "не рекомендуем"
	}
	return float64(average), verdict, nil
}
func main() {
	var n int
	fmt.Println("Введи n")
	fmt.Scan(&n)
	if av, verd, err := average(n); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(av, verd)
	}
}
