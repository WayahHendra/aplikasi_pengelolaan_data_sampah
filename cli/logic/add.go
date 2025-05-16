package logic

import (
	"fmt"
	"sampah-app/cli/data"
	"sampah-app/cli/utils"
)

func AddData() {
	utils.ClearConsole()

	var (
		wasteType       string
		recyclingMethod string
		quantity        float64
		location        string
		status          string
	)

	ShowRecycleTypeTable() // Memunculkan jenis-jenis sampah daur ulang

	fmt.Println() // Spacing

	fmt.Print("Masukkan jenis sampah: ")
	fmt.Scan(&wasteType)
	wasteType = GarbageTypes(wasteType)

	fmt.Print("Masukkan metode daur ulang sampah: ")
	fmt.Scan(&recyclingMethod)
	recyclingMethod = RecyclingMethods(recyclingMethod, wasteType)

	fmt.Print("Masukkan jumlah sampah (kg): ")
	fmt.Scan(&quantity)

	fmt.Print("Masukkan lokasi pengumpulan: ")
	fmt.Scan(&location)

	fmt.Print("Masukkan status daur ulang (y/n): ")
	fmt.Scan(&status)
	if utils.StrToLower(status) == "y" {
		status = "Sudah"
	} else if utils.StrToLower(status) == "n" {
		status = "Belum"
	} else {
		status = "nil"
	}

	data.WasteData = append(data.WasteData, data.Waste{ // Menambahkan data baru ke dalam DataSampah
		WasteType:       wasteType,
		RecyclingMethod: recyclingMethod,
		Quantity:        quantity,
		Location:        location,
		Status:          status,
	})

	fmt.Println("Data berhasil ditambahkan.")

	fmt.Println() // Spacing

	utils.PressToContinue()
	utils.ClearConsole()
}
