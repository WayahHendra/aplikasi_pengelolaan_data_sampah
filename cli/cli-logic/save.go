package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

// SaveWaste memungkinkan pengguna untuk menyimpan data sampah ke dalam file JSON.
func SaveWaste(data []core.Waste) {
	utils.ClearConsole()

	var confirmSave string // Variabel untuk konfirmasi penyimpanan

	if core.SwitchLanguage {
		fmt.Print("Are you sure you want to save the file? [1] Yes, [2] No: ")
	} else {
		fmt.Print("Yakin ingin save file? [1] Ya, [2] Tidak: ")
	}
	fmt.Scan(&confirmSave)

	// Proses penyimpanan berdasarkan konfirmasi
	if confirmSave == "1" {
		var (
			filename string
			err      error
		)

		// Gunakan FormatingSaveData untuk membuat nama file
		filename = utils.FormatingSaveData()

		// Simpan data ke file
		err = core.SaveWasteToFile(filename, data)
		if err != nil {
			fmt.Println(err)
		} else {
			if core.SwitchLanguage {
				fmt.Println("Data successfully saved to", filename)
			} else {
				fmt.Println("Data berhasil disimpan ke", filename)
			}
		}
	} else if confirmSave == "2" {
		fmt.Println()

		if core.SwitchLanguage {
			fmt.Println("File save canceled.")
		} else {
			fmt.Println("Penyimpanan dibatalkan.")
		}
	} else {
		fmt.Println()

		if core.SwitchLanguage {
			fmt.Println("Invalid input!")
		} else {
			fmt.Println("Input tidak valid!")
		}
	}

	if core.SwitchLanguage {
		utils.PressToContinue("Press Enter to continue...")
	} else {
		utils.PressToContinue("Tekan Enter untuk melanjutkan...")
	}

	utils.ClearConsole()
}
