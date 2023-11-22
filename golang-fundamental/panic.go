package main

import "fmt"

func endApp() {
	fmt.Println("End app")

	message := recover()
	fmt.Println("Terjadi Error", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups error")
	}

}

func main() {
	runApp(true)

	fmt.Println("M. Auliya Mirzaq Romdloni")
}
