package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

// LoadWaste memungkinkan pengguna untuk memuat data sampah dari file JSON.
func LoadWaste(waste *[]core.Waste) {
	utils.ClearConsole()

	var filename string // Nama file yang akan dimuat

	if core.SwitchLanguage {
		fmt.Print("Enter the filename to load data: ")
	} else {
		fmt.Print("Masukkan nama file untuk load data: ")
	}
	fmt.Scan(&filename)

	filename = utils.EnsureJSONExtension(filename) // Pastikan nama file memiliki ekstensi .json

	// Memuat data dari file
	data, err := core.LoadWasteFromFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		// Jika berhasil, data dimasukkan ke dalam slice
		*waste = data

		fmt.Println()

		if core.SwitchLanguage {
			fmt.Println("Data successfully loaded from", filename)
		} else {
			fmt.Println("Data berhasil dimuat dari", filename)
		}
	}

	if core.SwitchLanguage {
		utils.PressToContinue("Press Enter to return to the main menu...")
	} else {
		utils.PressToContinue("Tekan Enter untuk kembali ke menu utama...")
	}

	utils.ClearConsole()
}
