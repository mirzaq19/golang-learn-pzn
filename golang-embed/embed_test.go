package golang_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	//err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm) //deprecated
	file, err := os.OpenFile("logo_new.png", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	_, err = file.Write(logo)
	if err != nil {
		return
	}

}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(content))
		}
	}
}
