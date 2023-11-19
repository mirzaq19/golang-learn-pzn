package main

import "fmt"

func main() {
	const nickname = "Mirzaq"
	const (
		firstname string = "M. Auliya"
		lastname         = "Mirzaq Romdloni"
	)

	fmt.Println(nickname)
	fmt.Println(firstname)
	fmt.Println(lastname)
	//error
	//firstname = "Auliya"
	//lastname = "Mirzaq"
}
