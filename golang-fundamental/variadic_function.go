package main

import "fmt"

func sumAll(flag bool, numbers ...int) int {
	total := 0

	for i, number := range numbers {
		if flag == true {
			total += number
		} else {
			if i == 0 {
				total = number
			} else {
				total *= number
			}
		}

	}

	return total
}

func main() {
	//result := sumAll(10, 20, 10, 40, 20)
	//fmt.Println(result)

	numbers := []int{10, 1, 2, 3, 4}
	fmt.Println(sumAll(false, numbers...))
}
