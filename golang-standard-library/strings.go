package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("M. Auliya Mirzaq Romdloni", "Mirzaq"))
	fmt.Println(strings.Split("M. Auliya Mirzaq Romdloni", " "))
	fmt.Println(strings.ToLower("M. Auliya Mirzaq Romdloni"))
	fmt.Println(strings.ToUpper("M. Auliya Mirzaq Romdloni"))
	fmt.Println(strings.Trim("--------M. Auliya Mirzaq Romdloni???????????", "-"))
	fmt.Println(strings.ReplaceAll("M. Auliya Mirzaq Romdloni", "M.", "Muhammad"))
}
