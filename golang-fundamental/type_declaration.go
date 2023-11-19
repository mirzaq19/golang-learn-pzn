package main

import "fmt"

func main() {
	type NoKTP string

	var ktpMirzaq NoKTP = "123123123123"
	var contoh string = "111111111111"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpMirzaq)
	fmt.Println(contoh)
	fmt.Println(contohKTP)

}
