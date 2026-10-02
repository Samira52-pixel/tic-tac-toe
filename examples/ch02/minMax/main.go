package main

import "fmt"

func minMax(a []int, min, max *int) {
	*min = a[0]
	*max = a[0]
	i := 1

	for i < len(a) {
		if a[i] < *min {
			*min = a[i]
		}
		if a[i] > *max {
			*max = a[i]
		}
		i = i + 1
	}

}
func main() {
	var size int
	fmt.Println("Введите количество элементов:")
	fmt.Scan(&size)
	a := make([]int, size)
	fmt.Printf("Введите %d чисел \n", size)
	for i := 0; i < size; i++ {
		fmt.Scan(&a[i])
	}
	min := 0
	max := 0
	minMax(a, &min, &max)
	fmt.Printf("%d %d \n", min, max)
}
