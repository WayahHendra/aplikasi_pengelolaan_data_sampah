package main

import "fmt"

func addData() {
	clearConsole() // Clear console

	var (
		jenis           string
		metodeDaurUlang string
		jumlah          int
		lokasi          string
		status          string
	)

	showTypes() // Memunculkan jenis sampah

	fmt.Println() // Spacing

	fmt.Print("Masukkan jenis sampah: ")
	fmt.Scan(&jenis)
	jenis = garbageTypes(jenis) // Update data

	fmt.Print("Masukkan metode daur ulang sampah: ")
	fmt.Scan(&metodeDaurUlang)
	metodeDaurUlang = recyclingMethods(metodeDaurUlang, jenis) // Update data

	fmt.Print("Masukkan jumlah sampah: ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan lokasi pengumpulan: ")
	fmt.Scan(&lokasi)

	fmt.Print("Masukkan status daur ulang (y/n): ")
	fmt.Scan(&status)
	if strToLower(status) == "y" {
		status = "Sudah"
	} else if strToLower(status) == "n" {
		status = "Belum"
	} else {
		status = "nil"
	}

	DataSampah = append(DataSampah, Sampah{ // Menambahkan data baru ke dalam DataSampah
		Jenis:           jenis,
		MetodeDaurUlang: metodeDaurUlang,
		Jumlah:          jumlah,
		Lokasi:          lokasi,
		Status:          status,
	})

	fmt.Println("Data berhasil ditambahkan.")

	fmt.Println() // Spacing
}

func updateData() {
	clearConsole() // Clear console

	var (
		index           int
		jenis           string
		metodeDaurUlang string
		jumlah          int
		lokasi          string
		status          string
	)

	showAllData() // Memunculkan tabel data

	fmt.Print("Masukkan nomor data yang ingin diubah: ")
	fmt.Scan(&index)

	index-- // Ubah ke index array (mulai dari 0)

	if index < 0 || index >= len(DataSampah) { // Input < 0 / input >= panjang DataSampah
		fmt.Println("Nomor data tidak valid.")
		fmt.Println() // Spacing
		return
	}

	clearConsole() // Clear console

	showTypes() // Memunculkan jenis sampah

	fmt.Println() // Spacing

	fmt.Print("Masukkan jenis sampah: ")
	fmt.Scan(&jenis)
	jenis = garbageTypes(jenis) // Update data

	fmt.Print("Masukkan metode daur ulang sampah: ")
	fmt.Scan(&metodeDaurUlang)
	metodeDaurUlang = recyclingMethods(metodeDaurUlang, jenis) // Update data

	fmt.Print("Masukkan jumlah sampah: ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan lokasi pengumpulan: ")
	fmt.Scan(&lokasi)

	fmt.Print("Masukkan status daur ulang (y/n): ")
	fmt.Scan(&status)
	if strToLower(status) == "y" {
		status = "Sudah"
	} else if strToLower(status) == "n" {
		status = "Belum"
	} else {
		status = "nil"
	}

	DataSampah[index] = Sampah{ // Update DataSampah
		Jenis:           jenis,
		MetodeDaurUlang: metodeDaurUlang,
		Jumlah:          jumlah,
		Lokasi:          lokasi,
		Status:          status,
	}

	fmt.Println("Data berhasil diubah.")

	fmt.Println() // Spacing
}

func deleteData() {
	clearConsole() // Clear console

	var (
		index         int
		confirmDelete string
	)

	if len(DataSampah) == 0 { // Panjang DataSampah == 0
		fmt.Println("Tidak ada data yang tersedia untuk dihapus.")
		fmt.Println() // Spacing
		return
	}

	showAllData() // Memunculkan tabel data

	fmt.Print("Masukkan nomor data yang ingin dihapus: ")
	fmt.Scan(&index)

	index-- // Ubah ke index array (mulai dari 0)

	if index < 0 || index >= len(DataSampah) { // Input < 0 / input >= panjang DataSampah
		fmt.Println("Nomor data tidak valid.")
		fmt.Println() // Spacing
		return
	}

	fmt.Print("Yakin ingin menghapus data? (y/n): ")
	fmt.Scan(&confirmDelete)

	if strToLower(confirmDelete) == "y" {
		for i := index; i < len(DataSampah)-1; i++ { // Geser ke kiri
			DataSampah[i] = DataSampah[i+1]
		}
		DataSampah = DataSampah[:len(DataSampah)-1] // Pangkas slice

		clearConsole() // Clear console

		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Gagal menghapus data.")
	}

	fmt.Println() // Spacing
}

func recordProcess() {
	// TODO:
}

func searchData() {
	// TODO:
}

func sortData() {
	// TODO:
}

func showStatistics() {
	// TODO:
}

func showAllData() {
	clearConsole() // Clear console

	var data Sampah

	if len(DataSampah) == 0 { // Panjang DataSampah == 0
		fmt.Println("Belum ada data sampah yang tersedia.")
		fmt.Println() // Spacing
		return
	}

	fmt.Println("========================================================================================")
	fmt.Printf("|| %-3s || %-15s || %-6s || %-20s || %-10s || %-8s ||\n", "No", "jenis", "Jumlah", "Metode Daur Ulang", "Lokasi", "Status")
	fmt.Println("||====================================================================================||")
	for i := 0; i < len(DataSampah); i++ { // Loop data
		data = DataSampah[i]
		fmt.Printf("|| %-3d || %-15s || %-6d || %-20s || %-10s || %-8s ||\n", i+1, data.Jenis, data.Jumlah, data.MetodeDaurUlang, data.Lokasi, data.Status)
	}
	fmt.Println("========================================================================================")

	fmt.Println() // Spacing
}
