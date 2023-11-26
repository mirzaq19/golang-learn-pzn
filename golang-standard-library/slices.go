package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Mirzaq", "Richard", "John", "Budi"}
	values := []int{12, 3, 4, 23, 5, 2, 7}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Richard"))
	fmt.Println(slices.Index(names, "John"))
	fmt.Println(slices.Index(names, "Budi"))
}
