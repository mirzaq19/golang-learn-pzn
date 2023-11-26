package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"M.", "Auliya", "Mirzaq", "Romdloni"})
	_ = writer.Write([]string{"Richard", "Asmarakandi", "", ""})
	_ = writer.Write([]string{"Zydhan", "Linnar", "Putra", ""})
	_ = writer.Write([]string{"Ahdan", "Amanullah", "Irfan", "Muzaffar"})

	writer.Flush()
}
