package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

func SortWaste() {
	utils.ClearConsole()

	var (
		sortFieldChoice     string // Input user (1/2/3/4)
		sortOrderChoice     string // Input user (1/2)
		sortAlgorithmChoice string // Input user (1/2)
		sortField           string // Output: "jenis", "metode", "jumlah", "lokasi"
		sortAscending       bool   // Output: true (ascending), false (descending)
	)

	if len(core.WasteData) == 0 { // Panjang data sampah == 0
		fmt.Println("Belum ada data sampah yang tersedia.")
		fmt.Println() // Spacing

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	fmt.Println("Urutkan data berdasarkan kolom apa?")
	fmt.Println("1. Jenis Sampah")
	fmt.Println("2. Metode Daur Ulang")
	fmt.Println("3. Jumlah Sampah (kg)")
	fmt.Println("4. Lokasi Pengumpulan")
	fmt.Print("Masukkan pilihan (1-4): ")
	fmt.Scan(&sortFieldChoice)
	core.SetSortField(&sortField, sortFieldChoice)

	fmt.Println() // Spacing

	fmt.Println("Urutan pengurutan:")
	fmt.Println("1. Naik (Ascending)")
	fmt.Println("2. Turun (Descending)")
	fmt.Print("Masukkan pilihan (1-2): ")
	fmt.Scan(&sortOrderChoice)
	core.SetSortOrder(&sortAscending, sortOrderChoice)

	fmt.Println() // Spacing

	fmt.Println("Pilih algoritma pengurutan:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Masukkan pilihan (1-2): ")
	fmt.Scan(&sortAlgorithmChoice)
	core.SortByAlgorithm(sortAlgorithmChoice, sortField, sortAscending)

	fmt.Println() // Spacing

	switch sortField {
	case "jenis":
		if sortAscending {
			fmt.Println("Data berhasil diurutkan berdasarkan Jenis (A-Z).")
		} else {
			fmt.Println("Data berhasil diurutkan berdasarkan Jenis (Z-A).")
		}
	case "metode":
		if sortAscending {
			fmt.Println("Data berhasil diurutkan berdasarkan Metode Daur Ulang (A-Z).")
		} else {
			fmt.Println("Data berhasil diurutkan berdasarkan Metode Daur Ulang (Z-A).")
		}
	case "jumlah":
		if sortAscending {
			fmt.Println("Data berhasil diurutkan berdasarkan Jumlah (terkecil ke terbesar).")
		} else {
			fmt.Println("Data berhasil diurutkan berdasarkan Jumlah (terbesar ke terkecil).")
		}
	case "lokasi":
		if sortAscending {
			fmt.Println("Data berhasil diurutkan berdasarkan Lokasi (A-Z).")
		} else {
			fmt.Println("Data berhasil diurutkan berdasarkan Lokasi (Z-A).")
		}
	}

	fmt.Println() // Spacing

	utils.PressToContinue()
	utils.ClearConsole()
}
