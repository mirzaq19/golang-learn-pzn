package main

import "fmt"

func factorialLoop(value int) int {
	total := 1

	for i := value; i >= 1; i-- {
		total *= i
	}

	return total
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}
