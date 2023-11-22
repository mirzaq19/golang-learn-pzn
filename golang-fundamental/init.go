package main

import (
	"fmt"
	"golang-fundamental/database"
	_ "golang-fundamental/internal"
)

func main() {
	fmt.Println(database.GetDatabaseConnection())
}
