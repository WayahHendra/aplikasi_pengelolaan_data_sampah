package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

// CreateWaste memungkinkan pengguna untuk menambahkan data sampah baru.
func CreateWaste() {
	utils.ClearConsole()

	var (
		wasteType       string  // Jenis sampah
		recyclingMethod string  // Metode daur ulang
		quantity        float64 // Jumlah sampah (dalam kg)
		location        string  // Lokasi pengumpulan
		status          string  // Status daur ulang (sudah/belum)
	)

	var (
		count    int          // Jumlah data yang akan ditambahkan
		tempData []core.Waste // Slice sementara untuk menyimpan data baru
		waste    core.Waste   // Variabel sementara untuk data sampah
	)

	fmt.Println("Berapa banyak data sampah yang ingin ditambahkan?")
	fmt.Print("Masukkan jumlah (bilangan bulat positif): ")
	fmt.Scan(&count)

	// Validasi jumlah data
	if count <= 0 {
		fmt.Println("Jumlah data tidak valid. Minimal 1.")

		fmt.Println()

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	utils.ClearConsole()

	// Loop untuk menambahkan data sebanyak jumlah yang dimasukkan
	for i := 0; i < count; i++ {
		core.ShowRecycleTypeTable() // Tampilkan tabel data daur ulang yang tersedia
		fmt.Println()

		fmt.Print("========= Data ke-", i+1, " =========\n")
		fmt.Print("Masukkan jenis sampah: ")
		fmt.Scan(&wasteType)

		fmt.Print("Masukkan metode daur ulang: ")
		fmt.Scan(&recyclingMethod)

		fmt.Print("Masukkan jumlah sampah (kg): ")
		fmt.Scan(&quantity)

		fmt.Print("Masukkan lokasi pengumpulan: ")
		fmt.Scan(&location)

		fmt.Print("Masukkan status daur ulang (y/n): ")
		fmt.Scan(&status)

		// Tambahkan data ke slice sementara
		tempData = append(tempData, core.CreateData(wasteType, recyclingMethod, quantity, location, status))

		fmt.Print("Data ke-", i+1, " berhasil ditambahkan.\n\n")

		utils.PressToContinue()
		utils.ClearConsole()
	}

	// Tampilkan data yang baru ditambahkan
	fmt.Println("Berikut adalah data yang baru ditambahkan:")
	fmt.Println("========================================================================================================")
	fmt.Printf("|| %-3s || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", "No", "Jenis", "Metode Daur Ulang", "Jumlah", "Lokasi", "Status")
	fmt.Println("||====================================================================================================||")
	for i := 0; i < len(tempData); i++ {
		waste = tempData[i] // Ambil data dari slice sementara
		fmt.Printf("|| %-3d || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", i+1, waste.WasteType, waste.RecyclingMethod, fmt.Sprint(waste.Quantity, " kg"), waste.Location, waste.Status)
	}
	fmt.Println("========================================================================================================")

	core.WasteData = append(core.WasteData, tempData...) // Tambahkan data baru ke dalam slice utama

	utils.PressToContinue()
	utils.ClearConsole()
}
