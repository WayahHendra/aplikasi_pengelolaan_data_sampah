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
