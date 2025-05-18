package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

// UpdateWaste memungkinkan pengguna untuk mengubah data sampah yang sudah ada.
func UpdateWaste() {
	utils.ClearConsole()

	var (
		wasteType       string  // Jenis sampah
		recyclingMethod string  // Metode daur ulang
		quantity        float64 // Jumlah sampah (dalam kg)
		location        string  // Lokasi pengumpulan
		status          string  // Status daur ulang (sudah/belum)
	)

	var (
		confirmUpdate string // Konfirmasi pengubahan data
		index         int    // Indeks data yang akan diubah
	)

	core.TriggerShowData = false // Nonaktifkan trigger untuk membersihkan konsol
	ReadWaste()                  // Menampilkan tabel data
	core.TriggerShowData = true  // Aktifkan kembali trigger

	// Validasi apakah ada data yang tersedia
	if len(core.WasteData) == 0 {
		return
	}

	fmt.Print("Masukkan nomor data yang ingin diubah: ")
	fmt.Scan(&index)

	index-- // Ubah ke indeks berbasis 0

	// Validasi indeks
	if index < 0 || index >= len(core.WasteData) {
		fmt.Println("Nomor data tidak valid.")

		fmt.Println()

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	fmt.Print("Yakin ingin mengubah data? (y/n): ")
	fmt.Scan(&confirmUpdate)

	if utils.StrToLower(confirmUpdate) == "y" {
		utils.ClearConsole()

		core.ShowRecycleTypeTable() // Menampilkan jenis-jenis sampah daur ulang

		fmt.Println()

		fmt.Println("Data yang akan diubah:")
		fmt.Println("========================================================================================================")
		fmt.Printf("|| %-3s || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", "No", "Jenis", "Metode Daur Ulang", "Jumlah", "Lokasi", "Status")
		fmt.Println("||====================================================================================================||")
		fmt.Printf("|| %-3d || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", index+1, core.WasteData[index].WasteType, core.WasteData[index].RecyclingMethod, fmt.Sprint(core.WasteData[index].Quantity, " kg"), core.WasteData[index].Location, core.WasteData[index].Status)
		fmt.Println("========================================================================================================")

		fmt.Println()

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

		// Gunakan core.CreateData untuk normalisasi dan validasi data baru
		core.WasteData[index] = core.CreateData(wasteType, recyclingMethod, quantity, location, status)

		fmt.Println("Data berhasil diubah.")
	} else if utils.StrToLower(confirmUpdate) == "n" {
		fmt.Println("Gagal mengubah data.")
	} else {
		fmt.Println("Input tidak valid!")
	}

	fmt.Println()

	utils.PressToContinue()
	utils.ClearConsole()
}
