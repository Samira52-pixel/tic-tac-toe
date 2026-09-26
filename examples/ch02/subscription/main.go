package main

import "fmt"

func calcSubs(coast float64, month float64) (float64, error) {
	if coast <= 0 && month < 1 {
		fmt.Println("ошибка ввода")
		return 0, fmt.Errorf("Ошибка coast = %f  month = %f", coast, month)
	}

	//
	sum := coast * month
	discount := 0.0
	if month >= 12 {
		discount = 20.0
	} else {
		discount = 0.0
	}
	res := sum * (100.0 - discount) / 100.0

	return res, nil

}

func main() {
	var coast, month float64
	fmt.Scan(&coast, &month)
	if result, err := calcSubs(coast, month); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
