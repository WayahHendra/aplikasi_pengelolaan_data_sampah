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

	if core.SwitchLanguage {
		fmt.Print("Enter the name of the file to save data: ")
	} else {
		fmt.Print("Masukkan nama file untuk menyimpan data: ")
	}
	fmt.Scan(&filename)

	filename = utils.EnsureJSONExtension(filename) // Pastikan nama file memiliki ekstensi .json

	if core.SwitchLanguage {
		fmt.Print("Are you sure you want to save the file? (y/n): ")
	} else {
		fmt.Print("Yakin ingin save file? (y/n): ")
	}
	fmt.Scan(&confirmSave)

	// Proses penyimpanan berdasarkan konfirmasi
	if utils.StrToLower(confirmSave) == "yes" || utils.StrToLower(confirmSave) == "ye" || utils.StrToLower(confirmSave) == "y" {
		// Simpan data ke file
		err := core.SaveWasteToFile(filename, data)
		if err != nil {
			fmt.Println(err)
		} else {
			if core.SwitchLanguage {
				fmt.Print("Data successfully saved to ", filename)
			} else {
				fmt.Println("Data berhasil disimpan ke", filename)
			}
		}
	} else if utils.StrToLower(confirmSave) == "no" || utils.StrToLower(confirmSave) == "n" {
		if core.SwitchLanguage {
			fmt.Print("File save canceled.")
		} else {
			fmt.Println("Penyimpanan dibatalkan.")
		}
	} else {
		if core.SwitchLanguage {
			fmt.Print("Invalid input!")
		} else {
			fmt.Println("Input tidak valid!")
		}
	}

	utils.PressToContinue()
	utils.ClearConsole()
}
