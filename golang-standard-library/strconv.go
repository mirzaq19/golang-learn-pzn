package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println("result:", result)
	}
	resultInt, err := strconv.Atoi("2343")
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println("result:", resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	var stringInt string = strconv.Itoa(999)
	fmt.Println(stringInt)
}
