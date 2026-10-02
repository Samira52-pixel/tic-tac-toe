package main

import "fmt"

func reverse(a []int) {
	i := 0
	j := len(a) - 1
	for i < j {
		px := &a[i]
		py := &a[j]
		tmp := *px
		*px = *py
		*py = tmp
		i = i + 1
		j = j - 1
	}

}

/*
func main() {
	fmt.Println("Введите числа:")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	var a []int
	for _, word := range strings.Fields(sc.Text()) {
		x, err := strconv.Atoi(word)
		if err != nil {
			break
		}
		a = append(a, x)
	}
	reverse(a)
	fmt.Println(a)
}
*/

func main() {
	var size int
	fmt.Println("Введите количество элементов:")
	fmt.Scan(&size)
	a := make([]int, size)
	fmt.Printf("Введите %d чисел \n", size)
	for i := 0; i < size; i++ {
		fmt.Scan(&a[i])
	}
	reverse(a)
	fmt.Printf("a = %v\n", a)
}
