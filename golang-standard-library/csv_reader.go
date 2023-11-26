package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "M.,Auliya,Mirzaq,Romdloni\n" +
		"Richard,Asmarakandi,\n" +
		"Ahdan,Amanullah,Muzaffar\n" +
		"Zydhan,Linnar,Putra\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
