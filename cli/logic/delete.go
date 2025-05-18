package logic

import (
	"fmt"
	"sampah-app/core"
	"sampah-app/utils"
)

func DeleteWaste() {
	utils.ClearConsole()

	var (
		index         int
		confirmDelete string
	)

	if len(core.WasteData) == 0 { // Panjang data sampah == 0
		fmt.Println("Tidak ada data yang tersedia untuk dihapus.")
		fmt.Println() // Spacing

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	core.TriggerShowData = false
	ReadWaste() // Memunculkan tabel data
	core.TriggerShowData = true

	fmt.Print("Masukkan nomor data yang ingin dihapus: ")
	fmt.Scan(&index)

	index-- // Ubah ke index array (mulai dari 0)

	if index < 0 || index >= len(core.WasteData) { // Input < 0 / input >= panjang data sampah
		fmt.Println("Nomor data tidak valid.")

		fmt.Println() // Spacing

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	fmt.Print("Yakin ingin menghapus data? (y/n): ")
	fmt.Scan(&confirmDelete)

	if utils.StrToLower(confirmDelete) == "y" {
		core.DeleteData(index) // Hapus data

		utils.ClearConsole()

		fmt.Println("Data berhasil dihapus.")
	} else if utils.StrToLower(confirmDelete) == "n" {
		fmt.Println("Gagal menghapus data.")
	} else {
		fmt.Println("Input tidak valid!")
	}

	fmt.Println() // Spacing

	utils.PressToContinue()
	utils.ClearConsole()
}
