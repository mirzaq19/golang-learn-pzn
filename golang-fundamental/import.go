package main

import (
	"fmt"
	"golang-fundamental/helper"
)

func main() {
	result := helper.SayHello("Mirzaq")
	fmt.Println(result)

	fmt.Println(helper.Application)
	//fmt.Println(helper.version)           // Tidak bisa diakses
	//fmt.Println(helper.sayGoodBye("Eko")) // Tidak bisa diakses
}
