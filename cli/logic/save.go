package logic

import (
	"fmt"
	"sampah-app/core"
	"sampah-app/utils"
)

func SaveWaste(data []core.Waste) {
	utils.ClearConsole()

	var filename, confirmSave string

	fmt.Print("Masukkan nama file untuk menyimpan data: ")
	fmt.Scan(&filename)

	filename = utils.EnsureJSONExtension(filename)

	fmt.Print("Yakin ingin save file? (y/n): ")
	fmt.Scan(&confirmSave)

	if utils.StrToLower(confirmSave) == "y" {
		err := core.SaveWasteToFile(filename, data)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Data berhasil disimpan ke", filename)
		}
	} else if utils.StrToLower(confirmSave) == "n" {
		fmt.Println("Penyimpanan dibatalkan.")
	} else {
		fmt.Println("Input tidak valid!")
	}

	fmt.Println() // Spacing

	utils.PressToContinue()
	utils.ClearConsole()
}
