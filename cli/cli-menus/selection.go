package cli_menus

import (
	"fmt"
	logic "trash-app/cli/cli-logic"
	"trash-app/core"
	"trash-app/utils"
)

// HandleSubSelection menangani pilihan sub menu berdasarkan input pengguna.
func HandleSubSelection(breakLoop *bool) {
	utils.ClearConsole()

	ShowTableSubMenu()

	var subSelection int
	fmt.Scan(&subSelection)

	switch subSelection {
	case 1:
		logic.ReadWaste("") // Tampilkan semua data
	case 2:
		var date string = logic.LoadWasteByDate(&core.WasteData)
		logic.ReadWaste(date) // Tampilkan berdasarkan tanggal
	case -1:
		return // Kembali ke menu utama
	default:
		utils.ClearConsole()

		if core.SwitchLanguage {
			fmt.Println("Invalid input! Please select a valid menu.")
			utils.PressToContinue("Press Enter to continue...")
		} else {
			fmt.Println("Input tidak valid! Silakan pilih menu yang benar.")
			utils.PressToContinue("Tekan Enter untuk melanjutkan...")
		}

		utils.ClearConsole()
	}
}

// HandleSelection menangani pilihan menu berdasarkan input pengguna.
func HandleSelection(value int, breakLoop *bool) {
	switch value {
	case 1:
		logic.CreateWaste() // Tambah data sampah
	case 2:
		HandleSubSelection(breakLoop) // Tampilkan data sampah
	case 3:
		logic.UpdateWaste() // Ubah data sampah
	case 4:
		logic.DeleteWaste() // Hapus data sampah
	case 5:
		logic.SearchWaste() // Cari data sampah
	case 6:
		logic.SortWaste() // Urutkan data sampah
	case 7:
		logic.ShowStatistics() // Tampilkan statistik data sampah
	case 8:
		logic.SaveWaste(core.WasteData) // Simpan data sampah
	case 9:
		logic.LoadAllWaste(&core.WasteData) // Muat semua data sampah
	case 10:
		logic.LoadWasteByDate(&core.WasteData) // Muat data sampah berdasarkan tanggal
	case 11:
		logic.SwitchLanguage() // Ganti bahasa
	case -1:
		logic.ExitProgram(breakLoop) // Keluar dari program
	default:
		utils.ClearConsole()

		if core.SwitchLanguage {
			fmt.Println("Invalid input! Please select a valid menu.")
			utils.PressToContinue("Press Enter to continue...")
		} else {
			fmt.Println("Input tidak valid! Silakan pilih menu yang benar.")
			utils.PressToContinue("Tekan Enter untuk melanjutkan...")
		}

		utils.ClearConsole()
	}
}
