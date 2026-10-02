package main

func bubbleSort(a []int) {
	n := len(a)
	i := 0
	for i < n-1 {
		j := 0
		for j < n-1-i {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
			j = j + 1
		}
		i = i + 1
	}
}
