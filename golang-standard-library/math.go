package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.20))
	fmt.Println(math.Floor(1.60))
	fmt.Println(math.Round(1.60))
	fmt.Println(math.Round(1.20))
	fmt.Println(math.Round(1.50))
	fmt.Println(math.Min(14, 42))
	fmt.Println(math.Max(14, 42))
}
