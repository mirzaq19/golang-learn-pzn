package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	goodBye := getGoodBye
	fmt.Println(goodBye("Mirzaq"))
}
