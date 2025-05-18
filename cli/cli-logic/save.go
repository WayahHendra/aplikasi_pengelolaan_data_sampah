package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

// SaveWaste memungkinkan pengguna untuk menyimpan data sampah ke dalam file JSON.
func SaveWaste(data []core.Waste) {
	utils.ClearConsole()

	var filename, confirmSave string // Variabel untuk nama file dan konfirmasi penyimpanan

	fmt.Print("Masukkan nama file untuk menyimpan data: ")
	fmt.Scan(&filename)

	filename = utils.EnsureJSONExtension(filename) // Pastikan nama file memiliki ekstensi .json

	fmt.Print("Yakin ingin save file? (y/n): ")
	fmt.Scan(&confirmSave)

	// Proses penyimpanan berdasarkan konfirmasi
	if utils.StrToLower(confirmSave) == "yes" || utils.StrToLower(confirmSave) == "ye" || utils.StrToLower(confirmSave) == "y" {
		// Simpan data ke file
		err := core.SaveWasteToFile(filename, data)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Data berhasil disimpan ke", filename)
		}
	} else if utils.StrToLower(confirmSave) == "no" || utils.StrToLower(confirmSave) == "n" {
		fmt.Println("Penyimpanan dibatalkan.")
	} else {
		fmt.Println("Input tidak valid!")
	}

	utils.PressToContinue()
	utils.ClearConsole()
}
