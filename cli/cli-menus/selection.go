package cli_menus

import (
	"fmt"
	logic "trash-app/cli/cli-logic"
	"trash-app/core"
	"trash-app/utils"
)

// HandleSelection menangani pilihan menu berdasarkan input pengguna.
func HandleSelection(value int, breakLoop *bool) {
	switch value {
	case 1:
		logic.CreateWaste() // Tambah data sampah
	case 2:
		logic.ReadWaste() // Tampilkan semua data sampah
	case 3:
		logic.UpdateWaste() // Ubah data sampah
	case 4:
		logic.DeleteWaste() // Hapus data sampah
	case 5:
		logic.SearchWaste() // Cari data sampah
	case 6:
		logic.SortWaste() // Urutkan data sampah
	case 7:
		logic.RecordProcess() // Catat proses daur ulang
	case 8:
		logic.ShowStatistics() // Tampilkan statistik data
	case 9:
		logic.SaveWaste(core.WasteData) // Simpan data ke file
	case 10:
		logic.LoadWaste(&core.WasteData) // Muat data dari file
	case 11:
		logic.SwitchLanguage() // Ganti bahasa
	case 12:
		var exitConfirmation string

		if core.SwitchLanguage {
			fmt.Print("Do you want to exit the program? (y/n): ")
		} else {
			fmt.Print("Apakah Anda ingin keluar dari program? (y/n): ")
		}
		fmt.Scan(&exitConfirmation)

		if utils.StrToLower(exitConfirmation) == "yes" || utils.StrToLower(exitConfirmation) == "ye" || utils.StrToLower(exitConfirmation) == "y" {
			fmt.Println("Exiting the program...")
			*breakLoop = true // Set breakLoop menjadi true untuk keluar dari loop
			return
		} else if utils.StrToLower(exitConfirmation) == "no" || utils.StrToLower(exitConfirmation) == "n" {
			fmt.Println("Keluar dari program...")
			*breakLoop = false // Set breakLoop menjadi false untuk tetap di dalam loop
		} else {
			if core.SwitchLanguage {
				fmt.Println("Invalid input! Please select a valid option.")
			} else {
				fmt.Println("Input tidak valid! Silakan pilih opsi yang benar.")
			}
		}
	default:
		if core.SwitchLanguage {
			fmt.Println("Selection not valid, please try again.")
		} else {
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}

		fmt.Println()

		utils.PressToContinue()
		utils.ClearConsole()
	}
}
