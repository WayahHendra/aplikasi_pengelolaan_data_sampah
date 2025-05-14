package main

import "fmt"

type Waste struct { // Struktur data sampah
	WasteType       string
	RecyclingMethod string
	Quantity        int
	Location        string
	Status          string
}

var (
	WasteData       []Waste          // Untuk menyimpan input data ke dalam array
	TriggerShowData bool    = true   // Enable dan disable -> pressToContinue() dan clearConsole() pada fungsi showAllData()
	Version         string  = "v0.8" // Versi aplikasi
)

func showTableMenu() { // showMenu function untuk memunculkan menu
	fmt.Println("===========================================================")
	fmt.Println("||   SELAMAT DATANG DI APLIKASI DAUR ULANG SAMPAH", Version, "  ||")
	fmt.Println("||=======================================================||")
	fmt.Println("||  1. Tambah data sampah              C                 ||")
	fmt.Println("||  2. Tampilkan semua data            R                 ||")
	fmt.Println("||  3. Ubah data sampah                U                 ||")
	fmt.Println("||  4. Hapus data sampah               D                 ||")
	fmt.Println("||  5. Cari data sampah                                  ||")
	fmt.Println("||  6. Urutkan data sampah                               ||")
	fmt.Println("||  7. Catat proses daur ulang                           ||")
	fmt.Println("||  8. Tampilkan statistik                               ||")
	fmt.Println("||  9. Keluar program                                    ||")
	fmt.Println("===========================================================")
	fmt.Print("Pilih menu (1-9): ")
}

func showRecycleTypeTable() { // showRecycleTypeTable fungsi untuk menampilkan tabel jenis-jenis sampah daur ulang
	fmt.Println("=============================================================================================================================")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "Jenis Sampah", "Metode Daur Ulang", "Deskripsi Singkat")
	fmt.Println("||=========================================================================================================================||")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(1) Kompos", "Sampah seperti daun dan sisa makanan diubah menjadi pupuk")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "(1) Organik", "(2) Biogas", "Limbah organik difermentasi menghasilkan gas dan pupuk")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(3) BSF", "Larva BSF mengurai sampah, hasilnya jadi pakan dan kompos")
	fmt.Println("||=========================================================================================================================||")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(1) Peleburan", "Logam & plastik dilebur")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(2) Reuse", "Botol atau kemasan dipakai kembali tanpa diproses ulang")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "(2) Anorganik", "(3) Downcycling", "Plastik diolah menjadi produk kualitas lebih rendah (mis. tekstil)")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(4) Upcycling", "Sampah dijadikan barang kerajinan atau produk kreatif")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(5) Ecobricks", "Sampah plastik dimampatkan ke dalam botol jadi bahan bangunan")
	fmt.Println("||=========================================================================================================================||")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(1) Insinerasi", "Limbah dibakar dalam suhu tinggi agar aman")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "(3) B3 (Berbahaya)", "(2) Stabilisasi", "Limbah diubah jadi bentuk padat agar tak mencemari lingkungan")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(3) Elektronik", "E-waste didaur ulang untuk ambil logam berharga (emas, tembaga)")
	fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(4) Kimia", "Limbah cair dinetralkan atau diolah dengan bahan kimia khusus")
	fmt.Println("=============================================================================================================================")
}

func garbageTypes(input string) string { // Koreksi otomatis jenis sampah
	j := strToLower(input)
	if j == "1" || j == "o" || j == "or" || j == "org" || j == "orga" || j == "organ" || j == "organi" || j == "organik" {
		return "Organik"
	} else if j == "2" || j == "a" || j == "an" || j == "ano" || j == "anor" || j == "anorg" || j == "anorga" || j == "anorgan" || j == "anorgani" || j == "anorganik" {
		return "Anorganik"
	} else if j == "3" || j == "b" || j == "b3" {
		return "B3"
	} else {
		return "nil"
	}
}

func recyclingMethods(input, jenis string) string { // Koreksi otomatis metode daur ulang
	mdu := strToLower(input)
	if jenis == "Organik" {
		if mdu == "1" || mdu == "k" || mdu == "ko" || mdu == "kom" || mdu == "komp" || mdu == "kompo" || mdu == "kompos" {
			return "Kompos"
		} else if mdu == "2" || mdu == "b" || mdu == "bi" || mdu == "bio" || mdu == "biog" || mdu == "bioga" || mdu == "biogas" {
			return "Biogas"
		} else if mdu == "3" || mdu == "b" || mdu == "bs" || mdu == "bsf" {
			return "BSF"
		} else {
			return "nil"
		}
	} else if jenis == "Anorganik" {
		if mdu == "1" || mdu == "p" || mdu == "pe" || mdu == "pel" || mdu == "pele" || mdu == "peleb" || mdu == "pelebu" || mdu == "pelebura" || mdu == "peleburan" {
			return "Peleburan"
		} else if mdu == "2" || mdu == "r" || mdu == "re" || mdu == "reu" || mdu == "reus" || mdu == "reuse" {
			return "Reuse"
		} else if mdu == "3" || mdu == "d" || mdu == "do" || mdu == "dow" || mdu == "down" || mdu == "downc" || mdu == "downcy" || mdu == "downcycl" || mdu == "downcycli" || mdu == "downcycling" {
			return "Downcycling"
		} else if mdu == "4" || mdu == "u" || mdu == "up" || mdu == "upc" || mdu == "upcy" || mdu == "upcycl" || mdu == "upcycli" || mdu == "upcycling" {
			return "Upcycling"
		} else if mdu == "5" || mdu == "e" || mdu == "ec" || mdu == "eco" || mdu == "ecob" || mdu == "ecobr" || mdu == "ecobri" || mdu == "ecobric" || mdu == "ecobricks" {
			return "Ecobricks"
		} else {
			return "nil"
		}
	} else if jenis == "B3" {
		if mdu == "1" || mdu == "i" || mdu == "in" || mdu == "ins" || mdu == "insi" || mdu == "insin" || mdu == "insine" || mdu == "insiner" || mdu == "insinera" || mdu == "insineras" || mdu == "insinerasi" {
			return "Insinerasi"
		} else if mdu == "2" || mdu == "s" || mdu == "st" || mdu == "sta" || mdu == "stab" || mdu == "stabi" || mdu == "stabil" || mdu == "stabilis" || mdu == "stabilisa" || mdu == "stabilisasi" {
			return "Stabilisasi"
		} else if mdu == "3" || mdu == "e" || mdu == "el" || mdu == "ele" || mdu == "elek" || mdu == "elekt" || mdu == "elektr" || mdu == "elektron" || mdu == "elektroni" || mdu == "elektronik" {
			return "Elektronik"
		} else if mdu == "4" || mdu == "k" || mdu == "ki" || mdu == "kim" || mdu == "kimi" || mdu == "kimia" {
			return "Kimia"
		} else {
			return "nil"
		}
	} else {
		return "nil"
	}
}
