package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

func CreateWaste() {
	utils.ClearConsole()

	var (
		wasteType       string
		recyclingMethod string
		quantity        float64
		location        string
		status          string
	)

	var (
		count    int
		tempData []core.Waste
		waste    core.Waste
	)

	fmt.Println("Berapa banyak data sampah yang ingin ditambahkan?")
	fmt.Print("Masukkan jumlah (bilangan bulat positif): ")
	fmt.Scan(&count)

	if count <= 0 {
		fmt.Println("Jumlah data tidak valid. Minimal 1.")

		fmt.Println() // Spacing

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	utils.ClearConsole()

	for i := 0; i < count; i++ {
		core.ShowRecycleTypeTable()
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

		tempData = append(tempData, core.CreateData(wasteType, recyclingMethod, quantity, location, status))

		fmt.Print("Data ke-", i+1, " berhasil ditambahkan.\n\n")

		utils.PressToContinue()
		utils.ClearConsole()
	}

	fmt.Println("Berikut adalah data yang baru ditambahkan:")
	fmt.Println("========================================================================================================")
	fmt.Printf("|| %-3s || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", "No", "Jenis", "Metode Daur Ulang", "Jumlah", "Lokasi", "Status")
	fmt.Println("||====================================================================================================||")
	for i := 0; i < len(tempData); i++ {
		waste = tempData[i]
		fmt.Printf("|| %-3d || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", i+1, waste.WasteType, waste.RecyclingMethod, fmt.Sprint(waste.Quantity, " kg"), waste.Location, waste.Status)
	}
	fmt.Println("========================================================================================================")

	core.WasteData = append(core.WasteData, tempData...) // Menambahkan data baru ke dalam slice WasteData

	utils.PressToContinue()
	utils.ClearConsole()
}
