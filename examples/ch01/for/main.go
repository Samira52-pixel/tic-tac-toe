package main

import "fmt"

func main() {

	names := make(map[string]string)
	names["Karina"] = "Work"
	names["Samir"] = "Учиться"
	names["Artur"] = "В отпуске"

	// fmt.Println(names["Karina"], " ", "Work")
	// fmt.Println(names["Samir"], " ", "Учиться")
	// fmt.Println(names["Artur"], " ", "В отпуске")
	for name, status := range names {
		fmt.Println(name, " ", status)
	}

	args := []string{"Karina", "Samir", "Artur"}
	for _, name := range args {
		fmt.Println(name)
	}

	for index := 0; index < len(args); index++ {
		fmt.Println(args[index])
	}
}
