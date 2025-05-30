package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

// SearchWaste memungkinkan pengguna untuk mencari data sampah.
func SearchWaste() {
	utils.ClearConsole()

	var (
		query       string       // Kata kunci pencarian
		method      int          // Metode pencarian: 1 untuk Sequential Search, 2 untuk Binary Search
		results     []core.Waste // Hasil pencarian
		searchField int          // Pilihan kolom pencarian
	)

	if core.SwitchLanguage {
		fmt.Println("=========================================")
		fmt.Println("[-1] Back to main menu [Cancel operation]")
		fmt.Println("=========================================")
	} else {
		fmt.Println("=============================================")
		fmt.Println("[-1] Kembali ke menu utama [Batalkan operasi]")
		fmt.Println("=============================================")
	}

	fmt.Println()

	if core.SwitchLanguage {
		fmt.Println("Choose search field:")
		fmt.Println("[1] Waste Type")
		fmt.Println("[2] Recycling Method")
		fmt.Println("[3] Location")
		fmt.Println("[4] Status")
		fmt.Print("Enter your choice: ")
	} else {
		fmt.Println("Pilih kolom pencarian:")
		fmt.Println("[1] Jenis Sampah")
		fmt.Println("[2] Metode Daur Ulang")
		fmt.Println("[3] Lokasi")
		fmt.Println("[4] Status")
		fmt.Print("Masukkan pilihan Anda: ")
	}
	fmt.Scan(&searchField)
	if searchField == -1 {
		return // Kembali ke menu utama
	}

	if core.SwitchLanguage {
		fmt.Print("Enter search query: ")
	} else {
		fmt.Print("Masukkan kata kunci pencarian: ")
	}
	fmt.Scan(&query)
	if query == "-1" {
		return // Kembali ke menu utama
	}

	utils.ClearConsole()

	if core.SwitchLanguage {
		fmt.Println("=========================================")
		fmt.Println("[-1] Back to main menu [Cancel operation]")
		fmt.Println("=========================================")
	} else {
		fmt.Println("=============================================")
		fmt.Println("[-1] Kembali ke menu utama [Batalkan operasi]")
		fmt.Println("=============================================")
	}

	fmt.Println()

	if core.SwitchLanguage {
		fmt.Println("Choose search method:")
		fmt.Println("[1] Sequential Search")
		fmt.Println("[2] Binary Search")
		fmt.Print("Enter your choice: ")
	} else {
		fmt.Println("Pilih metode pencarian:")
		fmt.Println("[1] Sequential Search")
		fmt.Println("[2] Binary Search")
		fmt.Print("Masukkan pilihan Anda: ")
	}
	fmt.Scan(&method)
	if method == -1 {
		return // Kembali ke menu utama
	}

	// Pilih metode pencarian berdasarkan kolom
	switch searchField {
	case 1: // Jenis Sampah
		if method == 1 {
			results = core.SequentialSearch(query, "WasteType")
		} else if method == 2 {
			results = core.BinarySearch(query, "WasteType")
		}
	case 2: // Metode Daur Ulang
		if method == 1 {
			results = core.SequentialSearch(query, "RecyclingMethod")
		} else if method == 2 {
			results = core.BinarySearch(query, "RecyclingMethod")
		}
	case 3: // Lokasi
		if method == 1 {
			results = core.SequentialSearch(query, "Location")
		} else if method == 2 {
			results = core.BinarySearch(query, "Location")
		}
	case 4: // Status
		if method == 1 {
			results = core.SequentialSearch(query, "Status")
		} else if method == 2 {
			results = core.BinarySearch(query, "Status")
		}
	default:
		if core.SwitchLanguage {
			fmt.Println("Invalid choice.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
		utils.PressToContinue("")
		return
	}

	utils.ClearConsole()

	if len(results) == 0 {
		if core.SwitchLanguage {
			fmt.Println("No matching data found.")
		} else {
			fmt.Println("Tidak ada data yang cocok ditemukan.")
		}
	} else {
		if core.SwitchLanguage {
			fmt.Println("Search results:")
		} else {
			fmt.Println("Hasil pencarian:")
		}

		fmt.Println("============================================================================================================")
		fmt.Printf("|| %-3s || %-15s || %-20s || %-12s || %-20s || %-12s ||\n", "No", "Type", "Recycling Method", "Quantity", "Location", "Status")
		fmt.Println("============================================================================================================")
		for i, waste := range results {
			fmt.Printf("|| %-3d || %-15s || %-20s || %-12s || %-20s || %-12s ||\n", i+1, waste.WasteType, waste.RecyclingMethod, fmt.Sprint(waste.Quantity, " kg"), waste.Location, waste.Status)
		}
		fmt.Println("============================================================================================================")
	}

	if core.SwitchLanguage {
		utils.PressToContinue("Press Enter to return to the main menu...")
	} else {
		utils.PressToContinue("Tekan Enter untuk kembali ke menu utama...")
	}

	utils.ClearConsole()
}
