package main

import "fmt"

func getFullName() (string, string) {
	return "M. Auliya", "Mirzaq Romdloni"
}

func main() {
	//firstName, lastName := getFullName()
	//fmt.Println(firstName, lastName)

	firstName, _ := getFullName()
	fmt.Println(firstName)
}
