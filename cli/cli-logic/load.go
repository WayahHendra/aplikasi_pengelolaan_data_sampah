package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

func LoadWaste(waste *[]core.Waste) {
	utils.ClearConsole()

	var filename string

	fmt.Print("Masukkan nama file untuk load data: ")
	fmt.Scan(&filename)

	filename = utils.EnsureJSONExtension(filename)

	data, err := core.LoadWasteFromFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		*waste = data
		fmt.Println("Data berhasil dimuat dari", filename)
	}

	fmt.Println() // Spacing

	utils.PressToContinue()
	utils.ClearConsole()
}
