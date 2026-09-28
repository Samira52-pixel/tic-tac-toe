package main

import "fmt"

func maxGap(n int) (int, string, error) {
	if n < 2 {
		fmt.Println("ошибка ввода")
		return 0, "", fmt.Errorf("ошибка ввода")
	}
	var firstTime int
	fmt.Println("Введите время первого сеанса")
	fmt.Scan(&firstTime)

	if firstTime < 0 || firstTime > 600 {
		return 0, "", fmt.Errorf("ошибка ввода")
	}
	prevTime := firstTime
	maxGapValue := 0

	for i := 2; i <= n; i++ {
		var currentTime int
		fmt.Println("Введите время очередного сеанса")
		fmt.Scan(&currentTime)

		if !(currentTime > 0 && currentTime <= 600 && currentTime > prevTime) {
			return 0, "", fmt.Errorf("Ошибка ввода")
		}

		gap := currentTime - prevTime
		if gap > maxGapValue {
			maxGapValue = gap
		}
		prevTime = currentTime

	}
	var rating string
	if maxGapValue >= 180 {
		rating = "редкие сеансы"
	} else if maxGapValue >= 60 {
		rating = "обычное расписание"
	} else {
		rating = "плотное расписание"
	}
	return maxGapValue, rating, nil

}

func main() {
	var n int
	fmt.Println("Введи n")
	fmt.Scan(&n)
	if gap, rating, err := maxGap(n); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(gap, rating)
	}
}
