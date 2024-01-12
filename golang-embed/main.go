package main

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
)

//go:embed test/version.txt
var version string

//go:embed test/logo.png
var logo []byte

//go:embed test/files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	file, _ := os.OpenFile("logo_new.png", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	file.Write(logo)

	dir, _ := path.ReadDir("test/files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("test/files/" + entry.Name())
			fmt.Println(string(content))
		}
	}
}
