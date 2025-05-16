package logic

import (
	"fmt"
	"sampah-app/cli/data"
	"sampah-app/cli/utils"
)

func DeleteData() {
	utils.ClearConsole()

	var (
		index         int
		confirmDelete string
	)

	if len(data.WasteData) == 0 { // Panjang data sampah == 0
		fmt.Println("Tidak ada data yang tersedia untuk dihapus.")
		fmt.Println() // Spacing

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	data.TriggerShowData = false
	ShowAllData() // Memunculkan tabel data
	data.TriggerShowData = true

	fmt.Print("Masukkan nomor data yang ingin dihapus: ")
	fmt.Scan(&index)

	index-- // Ubah ke index array (mulai dari 0)

	if index < 0 || index >= len(data.WasteData) { // Input < 0 / input >= panjang data sampah
		fmt.Println("Nomor data tidak valid.")

		fmt.Println() // Spacing

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	fmt.Print("Yakin ingin menghapus data? (y/n): ")
	fmt.Scan(&confirmDelete)

	if utils.StrToLower(confirmDelete) == "y" {
		for i := index; i < len(data.WasteData)-1; i++ { // Geser ke kiri
			data.WasteData[i] = data.WasteData[i+1]
		}
		data.WasteData = data.WasteData[:len(data.WasteData)-1] // Pangkas slice

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
