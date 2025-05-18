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
		fmt.Println("Keluar dari program!")
		*breakLoop = true // Keluar dari program
	default:
		fmt.Println("Pilihan tidak valid, coba lagi.")

		fmt.Println()

		utils.PressToContinue()
		utils.ClearConsole()
	}
}
