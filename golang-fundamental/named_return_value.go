package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Muhammad"
	middleName = "Auliya Mirzaq"
	lastName = "Romdloni"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()

	fmt.Println(firstName, middleName, lastName)
}
