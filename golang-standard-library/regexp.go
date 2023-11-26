package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`m([a-z]*)q`)

	fmt.Println(regex.MatchString("mirzaq"))
	fmt.Println(regex.MatchString("miRZaq"))
	fmt.Println(regex.MatchString("miRzaq"))
	fmt.Println(regex.MatchString("mq"))

	fmt.Println(regex.FindAllString("mirzaq mirzaq irzaq irza mirzaq mizaq mzq", 10))
}
