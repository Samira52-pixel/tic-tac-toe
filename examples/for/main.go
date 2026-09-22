package main

import "fmt"

func main() {

	names := make(map[string]string)
	names["Karina"] = "Work"
	names["Samir"] = "Учиться"
	names["Artur"] = "В отпуске"
	for name, status := range names {
		fmt.Println(name, " ", status)
	}
}
