package golang_embed

import (
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
