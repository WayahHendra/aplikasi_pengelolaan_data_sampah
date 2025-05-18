package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

func UpdateWaste() {
	utils.ClearConsole()

	var (
		wasteType       string
		recyclingMethod string
		quantity        float64
		location        string
		status          string
	)

	var (
		confirmUpdate string
		index         int
	)

	core.TriggerShowData = false
	ReadWaste() // Menampilkan tabel data
	core.TriggerShowData = true

	if len(core.WasteData) == 0 { // Panjang data sampah == 0
		return
	}

	fmt.Print("Masukkan nomor data yang ingin diubah: ")
	fmt.Scan(&index)

	index-- // Ubah ke index 0-based

	if index < 0 || index >= len(core.WasteData) { // Input < 0 / input >= panjang data sampah
		fmt.Println("Nomor data tidak valid.")

		fmt.Println() // Spacing

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	fmt.Print("Yakin ingin mengubah data? (y/n): ")
	fmt.Scan(&confirmUpdate)

	if utils.StrToLower(confirmUpdate) == "y" {
		utils.ClearConsole()

		core.ShowRecycleTypeTable() // Menampilkan jenis-jenis sampah daur ulang

		fmt.Println() // Spacing

		fmt.Println("Data yang akan diubah:")
		fmt.Println("========================================================================================================")
		fmt.Printf("|| %-3s || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", "No", "Jenis", "Metode Daur Ulang", "Jumlah", "Lokasi", "Status")
		fmt.Println("||====================================================================================================||")
		fmt.Printf("|| %-3d || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", index+1, core.WasteData[index].WasteType, core.WasteData[index].RecyclingMethod, fmt.Sprint(core.WasteData[index].Quantity, " kg"), core.WasteData[index].Location, core.WasteData[index].Status)
		fmt.Println("========================================================================================================")

		fmt.Println() // Spacing

		fmt.Print("Masukkan jenis sampah: ")
		fmt.Scan(&wasteType)

		fmt.Print("Masukkan metode daur ulang sampah: ")
		fmt.Scan(&recyclingMethod)

		fmt.Print("Masukkan jumlah sampah (kg): ")
		fmt.Scan(&quantity)

		fmt.Print("Masukkan lokasi pengumpulan: ")
		fmt.Scan(&location)

		fmt.Print("Masukkan status daur ulang (y/n): ")
		fmt.Scan(&status)

		// Pakai core.Create untuk normalisasi dan validasi
		core.WasteData[index] = core.CreateData(wasteType, recyclingMethod, quantity, location, status)

		fmt.Println("Data berhasil diubah.")
	} else if utils.StrToLower(confirmUpdate) == "n" {
		fmt.Println("Gagal mengubah data.")
	} else {
		fmt.Println("Input tidak valid!")
	}

	fmt.Println() // Spacing

	utils.PressToContinue()
	utils.ClearConsole()
}
