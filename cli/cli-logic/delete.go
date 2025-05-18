package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

// DeleteWaste memungkinkan pengguna untuk menghapus data sampah.
func DeleteWaste() {
	utils.ClearConsole()

	var (
		index         int    // Indeks data yang akan dihapus
		confirmDelete string // Konfirmasi penghapusan
	)

	// Periksa apakah ada data yang tersedia
	if len(core.WasteData) == 0 {
		fmt.Println("Tidak ada data yang tersedia untuk dihapus.")
		fmt.Println()

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	// Tampilkan data sampah
	core.TriggerShowData = false // Nonaktifkan trigger untuk membersihkan konsol
	ReadWaste()                  // Tampilkan data
	core.TriggerShowData = true  // Aktifkan kembali trigger

	fmt.Print("Masukkan nomor data yang ingin dihapus: ")
	fmt.Scan(&index)

	index-- // Konversi nomor ke indeks (dimulai dari 0)

	// Validasi indeks
	if index < 0 || index >= len(core.WasteData) {
		fmt.Println("Nomor data tidak valid.")

		fmt.Println()

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	fmt.Print("Yakin ingin menghapus data? (y/n): ")
	fmt.Scan(&confirmDelete)

	// Proses penghapusan berdasarkan konfirmasi
	if utils.StrToLower(confirmDelete) == "y" {
		core.DeleteData(index) // Hapus data berdasarkan indeks

		utils.ClearConsole()

		fmt.Println("Data berhasil dihapus.")
	} else if utils.StrToLower(confirmDelete) == "n" {
		fmt.Println("Penghapusan data dibatalkan.")
	} else {
		fmt.Println("Input tidak valid!")
	}

	fmt.Println()

	utils.PressToContinue()
	utils.ClearConsole()
}
