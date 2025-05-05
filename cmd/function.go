package main

import "fmt"

// showMenu function for showing menu
func showMenu() {
	fmt.Println("======================================================")
	fmt.Println("||   SELAMAT DATANG DI APLIKASI DAUR ULANG SAMPAH   ||")
	fmt.Println("======================================================")
	fmt.Println("||  1. Tambah data sampah                           ||")
	fmt.Println("||  2. Ubah data sampah                             ||")
	fmt.Println("||  3. Hapus data sampah                            ||")
	fmt.Println("||  4. Catat proses daur ulang                      ||")
	fmt.Println("||  5. Cari data sampah                             ||")
	fmt.Println("||  6. Urutkan data sampah                          ||")
	fmt.Println("||  7. Tampilkan statistik                          ||")
	fmt.Println("||  8. Tampilkan semua data                         ||")
	fmt.Println("||  9. Clear Console                                ||")
	fmt.Println("||  10. Keluar program                              ||")
	fmt.Println("======================================================")
	fmt.Print("Pilih menu (1-9): ")
}

func addData() {
	clear() // Clear console

	var (
		dataBaru        Sampah
		kategori        string
		jumlah          int
		metodeDaurUlang string
		lokasi          string
		status          string
	)

	fmt.Println() // Spacing

	fmt.Print("Masukkan kategori sampah (Organik/Anorganik/B3): ")
	fmt.Scan(&kategori)

	fmt.Print("Masukkan jumlah sampah: ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan metode daur ulang sampah: ")
	fmt.Scan(&metodeDaurUlang)

	fmt.Print("Masukkan lokasi pengumpulan: ")
	fmt.Scan(&lokasi)

	fmt.Print("Masukkan status daur ulang (Sudah/Belum): ")
	fmt.Scan(&status)

	dataBaru = Sampah{
		Kategori:        kategori,
		Jumlah:          jumlah,
		MetodeDaurUlang: metodeDaurUlang,
		Lokasi:          lokasi,
		Status:          status,
	}

	DataSampah = append(DataSampah, dataBaru) // Tambah dataBaru ke DataSampah
	fmt.Println("Data berhasil ditambahkan.")

	fmt.Println() // Spacing
}

func updateData() {
	// TODO:
}

func deleteData() {
	// TODO:
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
	clear() // Clear console

	if len(DataSampah) == 0 { // Panjang DataSampah == 0
		fmt.Println("Belum ada data sampah yang tersedia.")
		fmt.Println() // Spacing
		return
	}

	fmt.Println() // Spacing

	fmt.Println("========================================================================================")
	fmt.Printf("|| %-3s || %-15s || %-6s || %-20s || %-10s || %-8s ||\n", "No", "Kategori", "Jumlah", "Metode Daur Ulang", "Lokasi", "Status")
	fmt.Println("========================================================================================")
	for i := 0; i < len(DataSampah); i++ {
		data := DataSampah[i]
		fmt.Printf("|| %-3d || %-15s || %-6d || %-20s || %-10s || %-8s ||\n", i+1, data.Kategori, data.Jumlah, data.MetodeDaurUlang, data.Lokasi, data.Status)
	}
	fmt.Println("========================================================================================")

	fmt.Println() // Spacing
}
