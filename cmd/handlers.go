package main

import "fmt"

func addData() {
	clearConsole()

	var (
		wasteType       string
		recyclingMethod string
		quantity        int
		location        string
		status          string
	)

	showRecycleTypeTable() // Memunculkan jenis-jenis sampah daur ulang

	fmt.Println() // Spacing

	fmt.Print("Masukkan jenis sampah: ")
	fmt.Scan(&wasteType)
	wasteType = garbageTypes(wasteType) // Update data

	fmt.Print("Masukkan metode daur ulang sampah: ")
	fmt.Scan(&recyclingMethod)
	recyclingMethod = recyclingMethods(recyclingMethod, wasteType) // Update data

	fmt.Print("Masukkan jumlah sampah: ")
	fmt.Scan(&quantity)

	fmt.Print("Masukkan lokasi pengumpulan: ")
	fmt.Scan(&location)

	fmt.Print("Masukkan status daur ulang (y/n): ")
	fmt.Scan(&status)
	if strToLower(status) == "y" {
		status = "Sudah"
	} else if strToLower(status) == "n" {
		status = "Belum"
	} else {
		status = "nil"
	}

	WasteData = append(WasteData, Waste{ // Menambahkan data baru ke dalam DataSampah
		WasteType:       wasteType,
		RecyclingMethod: recyclingMethod,
		Quantity:        quantity,
		Location:        location,
		Status:          status,
	})

	fmt.Println("Data berhasil ditambahkan.")

	fmt.Println() // Spacing

	pressToContinue()
	clearConsole()
}

func showAllData() {
	clearConsole()

	var data Waste

	if len(WasteData) == 0 { // Panjang DataSampah == 0
		fmt.Println("Belum ada data sampah yang tersedia.")
		fmt.Println() // Spacing
		return
	}

	fmt.Println("========================================================================================")
	fmt.Printf("|| %-3s || %-15s|| %-20s || %-6s  || %-10s || %-8s ||\n", "No", "jenis", "Metode Daur Ulang", "Jumlah", "Lokasi", "Status")
	fmt.Println("||====================================================================================||")
	for i := 0; i < len(WasteData); i++ { // Loop data
		data = WasteData[i]
		fmt.Printf("|| %-3d || %-15s || %-20s || %-6d || %-10s || %-8s ||\n", i+1, data.WasteType, data.RecyclingMethod, data.Quantity, data.Location, data.Status)
	}
	fmt.Println("========================================================================================")

	fmt.Println() // Spacing

	if TriggerShowData {
		pressToContinue()
		clearConsole()
	}
}

func updateData() {
	clearConsole()

	var (
		index           int
		wasteType       string
		recyclingMethod string
		quantity        int
		location        string
		status          string
	)

	TriggerShowData = false
	showAllData() // Memunculkan tabel data
	TriggerShowData = true

	fmt.Print("Masukkan nomor data yang ingin diubah: ")
	fmt.Scan(&index)

	index-- // Ubah ke index array (mulai dari 0)

	if index < 0 || index >= len(WasteData) { // Input < 0 / input >= panjang DataSampah
		fmt.Println("Nomor data tidak valid.")

		fmt.Println() // Spacing

		pressToContinue()
		clearConsole()

		return
	}

	clearConsole()

	showRecycleTypeTable() // Memunculkan jenis-jenis sampah daur ulang

	fmt.Println() // Spacing

	fmt.Print("Masukkan jenis sampah: ")
	fmt.Scan(&wasteType)
	wasteType = garbageTypes(wasteType) // Update data

	fmt.Print("Masukkan metode daur ulang sampah: ")
	fmt.Scan(&recyclingMethod)
	recyclingMethod = recyclingMethods(recyclingMethod, wasteType) // Update data

	fmt.Print("Masukkan jumlah sampah: ")
	fmt.Scan(&quantity)

	fmt.Print("Masukkan lokasi pengumpulan: ")
	fmt.Scan(&location)

	fmt.Print("Masukkan status daur ulang (y/n): ")
	fmt.Scan(&status)
	if strToLower(status) == "y" {
		status = "Sudah"
	} else if strToLower(status) == "n" {
		status = "Belum"
	} else {
		status = "nil"
	}

	WasteData[index] = Waste{ // Update DataSampah
		WasteType:       wasteType,
		RecyclingMethod: recyclingMethod,
		Quantity:        quantity,
		Location:        location,
		Status:          status,
	}

	fmt.Println("Data berhasil diubah.")

	fmt.Println() // Spacing

	pressToContinue()
	clearConsole()
}

func deleteData() {
	clearConsole()

	var (
		index         int
		confirmDelete string
	)

	if len(WasteData) == 0 { // Panjang DataSampah == 0
		fmt.Println("Tidak ada data yang tersedia untuk dihapus.")
		fmt.Println() // Spacing
		return
	}

	TriggerShowData = false
	showAllData() // Memunculkan tabel data
	TriggerShowData = true

	fmt.Print("Masukkan nomor data yang ingin dihapus: ")
	fmt.Scan(&index)

	index-- // Ubah ke index array (mulai dari 0)

	if index < 0 || index >= len(WasteData) { // Input < 0 / input >= panjang WasteData
		fmt.Println("Nomor data tidak valid.")

		fmt.Println() // Spacing

		pressToContinue()
		clearConsole()

		return
	}

	fmt.Print("Yakin ingin menghapus data? (y/n): ")
	fmt.Scan(&confirmDelete)

	if strToLower(confirmDelete) == "y" {
		for i := index; i < len(WasteData)-1; i++ { // Geser ke kiri
			WasteData[i] = WasteData[i+1]
		}
		WasteData = WasteData[:len(WasteData)-1] // Pangkas slice

		clearConsole()

		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Gagal menghapus data.")
	}

	fmt.Println() // Spacing

	pressToContinue()
	clearConsole()
}

func searchData() {
	// TODO:
}

func sortData() {
	// TODO:
}

func recordProcess() {
	// TODO:
}

func showStatistics() {
	// TODO:
}
