package main

import "fmt"

func showMenu() { // showMenu function for showing menu
	fmt.Println("======================================================")
	fmt.Println("||   SELAMAT DATANG DI APLIKASI DAUR ULANG SAMPAH   ||")
	fmt.Println("||==================================================||")
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
	fmt.Print("Pilih menu (1-10): ")
}

func addData() {
	clearConsole() // Clear console

	var (
		dataBaru        Sampah
		kategori        string
		metodeDaurUlang string
		jumlah          int
		lokasi          string
		status          string
	)

	fmt.Println("=============================================================================================================================")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "Jenis Sampah", "Metode Daur Ulang", "Deskripsi Singkat")
	fmt.Println("||=========================================================================================================================||")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "Kompos", "Sampah seperti daun dan sisa makanan diubah menjadi pupuk")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "Organik", "Biogas", "Limbah organik difermentasi menghasilkan gas dan pupuk")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "BSF", "Larva BSF mengurai sampah, hasilnya jadi pakan dan kompos")
	fmt.Println("||=========================================================================================================================||")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "Peleburan", "Logam & plastik dilebur")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "Reuse", "Botol atau kemasan dipakai kembali tanpa diproses ulang")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "Anorganik", "Downcycling", "Plastik diolah menjadi produk kualitas lebih rendah (mis. tekstil)")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "Upcycling", "Sampah dijadikan barang kerajinan atau produk kreatif")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "Ecobricks", "Sampah plastik dimampatkan ke dalam botol jadi bahan bangunan")
	fmt.Println("||=========================================================================================================================||")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "Insinerasi", "Limbah dibakar dalam suhu tinggi agar aman")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "B3 (Berbahaya)", "Stabilisasi", "Limbah diubah jadi bentuk padat agar tak mencemari lingkungan")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "Elektronik", "E-waste didaur ulang untuk ambil logam berharga (emas, tembaga)")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "Kimia", "Limbah cair dinetralkan atau diolah dengan bahan kimia khusus")
	fmt.Println("=============================================================================================================================")

	fmt.Println() // Spacing

	fmt.Print("Masukkan jenis sampah: ")
	fmt.Scan(&kategori)
	k := strToLower(kategori)
	if k == "o" || k == "or" || k == "org" || k == "orga" || k == "organ" || k == "organi" || k == "organik" {
		kategori = "Organik"
	} else if k == "a" || k == "an" || k == "ano" || k == "anor" || k == "anorg" || k == "anorga" || k == "anorgan" || k == "anorgani" || k == "anorganik" {
		kategori = "Anorganik"
	} else if k == "b" || k == "b3" {
		kategori = "B3"
	} else {
		kategori = "nil"
	}

	fmt.Print("Masukkan metode daur ulang sampah: ")
	fmt.Scan(&metodeDaurUlang)
	mdu := strToLower(metodeDaurUlang)
	if kategori == "Organik" {
		if mdu == "k" || mdu == "ko" || mdu == "kom" || mdu == "komp" || mdu == "kompo" || mdu == "kompos" {
			metodeDaurUlang = "Kompos"
		} else if mdu == "b" || mdu == "bi" || mdu == "bio" || mdu == "biog" || mdu == "bioga" || mdu == "biogas" {
			metodeDaurUlang = "Biogas"
		} else if mdu == "b" || mdu == "bs" || mdu == "bsf" {
			metodeDaurUlang = "BSF"
		} else {
			metodeDaurUlang = "nil"
		}
	} else if kategori == "Anorganik" {
		if mdu == "p" || mdu == "pe" || mdu == "pel" || mdu == "pele" || mdu == "peleb" || mdu == "pelebu" || mdu == "pelebura" || mdu == "peleburan" {
			metodeDaurUlang = "Peleburan"
		} else if mdu == "r" || mdu == "re" || mdu == "reu" || mdu == "reus" || mdu == "reuse" {
			metodeDaurUlang = "Reuse"
		} else if mdu == "d" || mdu == "do" || mdu == "dow" || mdu == "down" || mdu == "downc" || mdu == "downcy" || mdu == "downcycl" || mdu == "downcycli" || mdu == "downcycling" {
			metodeDaurUlang = "Downcycling"
		} else if mdu == "u" || mdu == "up" || mdu == "upc" || mdu == "upcy" || mdu == "upcycl" || mdu == "upcycli" || mdu == "upcycling" {
			metodeDaurUlang = "Upcycling"
		} else if mdu == "e" || mdu == "ec" || mdu == "eco" || mdu == "ecob" || mdu == "ecobr" || mdu == "ecobri" || mdu == "ecobric" || mdu == "ecobricks" {
			metodeDaurUlang = "Ecobricks"
		} else {
			metodeDaurUlang = "nil"
		}
	} else if kategori == "B3" {
		if mdu == "i" || mdu == "in" || mdu == "ins" || mdu == "insi" || mdu == "insin" || mdu == "insine" || mdu == "insiner" || mdu == "insinera" || mdu == "insineras" || mdu == "insinerasi" {
			metodeDaurUlang = "Insinerasi"
		} else if mdu == "s" || mdu == "st" || mdu == "sta" || mdu == "stab" || mdu == "stabi" || mdu == "stabil" || mdu == "stabilis" || mdu == "stabilisa" || mdu == "stabilisasi" {
			metodeDaurUlang = "Stabilisasi"
		} else if mdu == "e" || mdu == "el" || mdu == "ele" || mdu == "elek" || mdu == "elekt" || mdu == "elektr" || mdu == "elektron" || mdu == "elektroni" || mdu == "elektronik" {
			metodeDaurUlang = "Elektronik"
		} else if mdu == "k" || mdu == "ki" || mdu == "kim" || mdu == "kimi" || mdu == "kimia" {
			metodeDaurUlang = "Kimia"
		} else {
			metodeDaurUlang = "nil"
		}
	} else {
		metodeDaurUlang = "nil"
	}

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
	clearConsole() // Clear console

	var (
		data          Sampah
		index         int
		confirmDelete string
	)

	if len(DataSampah) == 0 { // Panjang DataSampah == 0
		fmt.Println("Tidak ada data yang tersedia untuk dihapus.")
		fmt.Println() // Spacing
		return
	}

	fmt.Println("========================================================================================")
	fmt.Printf("|| %-3s || %-15s || %-6s || %-20s || %-10s || %-8s ||\n", "No", "Kategori", "Jumlah", "Metode Daur Ulang", "Lokasi", "Status")
	fmt.Println("||====================================================================================||")
	for i := 0; i < len(DataSampah); i++ { // i kurang dari panjang DataSampah -> i++
		data = DataSampah[i]
		fmt.Printf("|| %-3d || %-15s || %-6d || %-20s || %-10s || %-8s ||\n", i+1, data.Kategori, data.Jumlah, data.MetodeDaurUlang, data.Lokasi, data.Status)
	}
	fmt.Println("========================================================================================")

	fmt.Println() // Spacing

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
		DataSampah = append(DataSampah[:index], DataSampah[index+1:]...) // Hapus data dengan menggeser elemen setelahnya

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
	fmt.Printf("|| %-3s || %-15s || %-6s || %-20s || %-10s || %-8s ||\n", "No", "Kategori", "Jumlah", "Metode Daur Ulang", "Lokasi", "Status")
	fmt.Println("||====================================================================================||")
	for i := 0; i < len(DataSampah); i++ { // i kurang dari panjang DataSampah -> i++
		data = DataSampah[i]
		fmt.Printf("|| %-3d || %-15s || %-6d || %-20s || %-10s || %-8s ||\n", i+1, data.Kategori, data.Jumlah, data.MetodeDaurUlang, data.Lokasi, data.Status)
	}
	fmt.Println("========================================================================================")

	fmt.Println() // Spacing
}
