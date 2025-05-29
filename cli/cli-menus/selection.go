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
		logic.LoadAllWaste(&core.WasteData)
		logic.ReadWaste("") // Tampilkan semua data
	case 2:
		date := logic.LoadWasteByDate(&core.WasteData)
		logic.ReadWaste(date) // Tampilkan berdasarkan tanggal
	case 3:
		return // Kembali ke menu utama
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

// HandleSelection menangani pilihan menu berdasarkan input pengguna.
func HandleSelection(value int, breakLoop *bool) {
	switch value {
	case 1:
		logic.LoadWaste(&core.WasteData) // Muat data sampah
		logic.CreateWaste() // Tambah data sampah
		logic.SaveWaste(core.WasteData) // Simpan data sampah
	case 2: 
		HandleSubSelection(breakLoop)
	case 3:
		logic.LoadWaste(&core.WasteData) // Muat data sampah
		logic.UpdateWaste() // Ubah data sampah
		logic.SaveWaste(core.WasteData) // Simpan data sampah
	case 4:
		logic.LoadWaste(&core.WasteData) // Muat data sampah
		logic.DeleteWaste() // Hapus data sampah
		logic.SaveWaste(core.WasteData) // Simpan data sampah
	case 5:
		logic.SearchWaste() // Cari data sampah
	case 6:
		logic.SortWaste() // Urutkan data sampah
	case 7:
		logic.RecordProcess() // Catat proses daur ulang
	case 8:
		logic.ShowStatistics() // Tampilkan statistik data sampah
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
