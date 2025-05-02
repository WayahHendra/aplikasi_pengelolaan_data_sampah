package main

import (
	"aplikasi_pengelolaan_data_sampah/internal/validation"
	"fmt"
)

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
	fmt.Println("||  9. Keluar program                               ||")
	fmt.Println("======================================================")
	fmt.Print("Pilih menu (1-9): ")
}

func main() {
	var input string

	// loop
	for {
		showMenu()

		fmt.Scan(&input)
		value, err := validation.ValidateInput(input)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		switch value {
		case 1:
			addData()
		case 2:
			updateData()
		case 3:
			deleteData()
		case 4:
			recordProcess()
		case 5:
			searchData()
		case 6:
			sortData()
		case 7:
			showStatistics()
		case 8:
			showAllData()
		case 9:
			fmt.Println("Keluar dari program!")
			return // break loop
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
