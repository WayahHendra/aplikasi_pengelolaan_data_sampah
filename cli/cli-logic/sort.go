package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

// SortWaste memungkinkan pengguna untuk mengurutkan data sampah berdasarkan kolom tertentu.
func SortWaste() {
	utils.ClearConsole()

	var (
		sortFieldChoice     string // Input pengguna untuk memilih kolom pengurutan (1/2/3/4)
		sortOrderChoice     string // Input pengguna untuk memilih urutan pengurutan (1/2)
		sortAlgorithmChoice string // Input pengguna untuk memilih algoritma pengurutan (1/2)
		sortField           string // Kolom pengurutan: "jenis", "metode", "jumlah", "lokasi"
		sortAscending       bool   // Urutan pengurutan: true (ascending), false (descending)
	)

	// Periksa apakah ada data sampah yang tersedia
	if len(core.WasteData) == 0 {
		if core.SwitchLanguage {
			fmt.Println("No data available to sort.")
		} else {
			fmt.Println("Tidak ada data yang tersedia untuk diurutkan.")
		}

		fmt.Println()

		utils.PressToContinue()
		utils.ClearConsole()
		return
	}

	if core.SwitchLanguage {
		fmt.Println("Sort data by which column?")
		fmt.Println("1. Waste Type")
		fmt.Println("2. Recycling Method")
		fmt.Println("3. Waste Amount (kg)")
		fmt.Println("4. Collection Location")
		fmt.Print("Enter your choice (1-4): ")
	} else {
		fmt.Println("Urutkan data berdasarkan kolom apa?")
		fmt.Println("1. Jenis Sampah")
		fmt.Println("2. Metode Daur Ulang")
		fmt.Println("3. Jumlah Sampah (kg)")
		fmt.Println("4. Lokasi Pengumpulan")
		fmt.Print("Masukkan pilihan (1-4): ")
	}
	fmt.Scan(&sortFieldChoice)
	core.SetSortField(&sortField, sortFieldChoice) // Tentukan kolom pengurutan

	fmt.Println()

	if core.SwitchLanguage {
		fmt.Println("Sort order:")
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Print("Enter your choice (1-2): ")
	} else {
		fmt.Println("Urutan pengurutan:")
		fmt.Println("1. Naik")
		fmt.Println("2. Turun")
		fmt.Print("Masukkan pilihan (1-2): ")
	}
	fmt.Scan(&sortOrderChoice)
	core.SetSortOrder(&sortAscending, sortOrderChoice) // Tentukan urutan pengurutan

	fmt.Println()

	if core.SwitchLanguage {
		fmt.Println("Choose sorting algorithm:")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		fmt.Print("Enter your choice (1-2): ")
	} else {
		fmt.Println("Pilih algoritma pengurutan:")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		fmt.Print("Masukkan pilihan (1-2): ")
	}
	fmt.Scan(&sortAlgorithmChoice)
	core.SortByAlgorithm(sortAlgorithmChoice, sortField, sortAscending) // Jalankan algoritma pengurutan

	fmt.Println()

	// Tampilkan pesan sukses berdasarkan kolom dan urutan pengurutan
	switch sortField {
	case "jenis":
		if sortAscending {
			if core.SwitchLanguage {
				fmt.Println("Data successfully sorted by Type (A-Z).")
			} else {
				fmt.Println("Data berhasil diurutkan berdasarkan Jenis (A-Z).")
			}
		} else {
			if core.SwitchLanguage {
				fmt.Println("Data successfully sorted by Type (Z-A).")
			} else {
				fmt.Println("Data berhasil diurutkan berdasarkan Jenis (Z-A).")
			}
		}
	case "metode":
		if sortAscending {
			if core.SwitchLanguage {
				fmt.Println("Data successfully sorted by Recycling Method (A-Z).")
			} else {
				fmt.Println("Data berhasil diurutkan berdasarkan Metode Daur Ulang (A-Z).")
			}
		} else {
			if core.SwitchLanguage {
				fmt.Println("Data successfully sorted by Recycling Method (Z-A).")
			} else {
				fmt.Println("Data berhasil diurutkan berdasarkan Metode Daur Ulang (Z-A).")
			}
		}
	case "jumlah":
		if sortAscending {
			if core.SwitchLanguage {
				fmt.Println("Data successfully sorted by Amount (smallest to largest).")
			} else {
				fmt.Println("Data berhasil diurutkan berdasarkan Jumlah (terkecil ke terbesar).")
			}
		} else {
			if core.SwitchLanguage {
				fmt.Println("Data successfully sorted by Amount (largest to smallest).")
			} else {
				fmt.Println("Data berhasil diurutkan berdasarkan Jumlah (terbesar ke terkecil).")
			}
		}
	case "lokasi":
		if sortAscending {
			if core.SwitchLanguage {
				fmt.Println("Data successfully sorted by Location (A-Z).")
			} else {
				fmt.Println("Data berhasil diurutkan berdasarkan Lokasi (A-Z).")
			}
		} else {
			if core.SwitchLanguage {
				fmt.Println("Data successfully sorted by Location (Z-A).")
			} else {
				fmt.Println("Data berhasil diurutkan berdasarkan Lokasi (Z-A).")
			}
		}
	}

	fmt.Println()

	utils.PressToContinue()
	utils.ClearConsole()
}
