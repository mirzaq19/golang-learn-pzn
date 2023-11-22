package main

import "fmt"

func random() any {
	//return "Ok"
	//return 234
	return true
}

func main() {
	result := random()
	//resultString := result.(string)
	//fmt.Println(resultString)

	//resultInt := result.(int) //error
	//fmt.Println(resultInt)
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
