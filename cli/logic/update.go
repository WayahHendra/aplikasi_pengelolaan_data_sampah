package logic

import (
	"fmt"
	"sampah-app/cli/data"
	"sampah-app/cli/utils"
)

func UpdateData() {
	utils.ClearConsole()

	var (
		index           int
		wasteType       string
		recyclingMethod string
		quantity        float64
		location        string
		status          string
	)

	data.TriggerShowData = false
	ShowAllData() // Memunculkan tabel data
	data.TriggerShowData = true

	if len(data.WasteData) == 0 { // Panjang data sampah == 0
		return
	}

	fmt.Print("Masukkan nomor data yang ingin diubah: ")
	fmt.Scan(&index)

	index-- // Ubah ke index array (mulai dari 0)

	if index < 0 || index >= len(data.WasteData) { // Input < 0 / input >= panjang data sampah
		fmt.Println("Nomor data tidak valid.")

		fmt.Println() // Spacing

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	utils.ClearConsole()

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

	data.WasteData[index] = data.Waste{ // Update DataSampah
		WasteType:       wasteType,
		RecyclingMethod: recyclingMethod,
		Quantity:        quantity,
		Location:        location,
		Status:          status,
	}

	fmt.Println("Data berhasil diubah.")

	fmt.Println() // Spacing

	utils.PressToContinue()
	utils.ClearConsole()
}
