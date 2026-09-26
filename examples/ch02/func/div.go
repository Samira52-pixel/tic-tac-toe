package main

import "fmt"

func div(a, b int) (c int, err error) {
	if b == 0 {
		err = fmt.Errorf("На ноль делить нельзя")
		return
	}

	c = a / b
	return
}

func div2(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("На ноль делить нельзя")
	}
	return a / b, nil
}
