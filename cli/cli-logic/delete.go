package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

// DeleteWaste memungkinkan pengguna untuk menghapus data sampah.
func DeleteWaste() {
	utils.ClearConsole()

	var (
		index         int    // Indeks data yang akan dihapus
		confirmDelete string // Konfirmasi penghapusan
	)

	// Periksa apakah ada data yang tersedia
	if len(core.WasteData) == 0 {
		if core.SwitchLanguage {
			fmt.Println("No data available to delete.")
		} else {
			fmt.Println("Tidak ada data yang tersedia untuk dihapus.")
		}

		if core.SwitchLanguage {
			utils.PressToContinue("Press Enter to back to the main menu...")
		} else {
			utils.PressToContinue("Tekan Enter untuk kembali ke menu utama...")
		}

		utils.ClearConsole()

		return
	}

	// Tampilkan data sampah
	core.TriggerShowData = false // Nonaktifkan trigger untuk membersihkan konsol
	ReadWaste("")                // Tampilkan data
	core.TriggerShowData = true  // Aktifkan kembali trigger

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
		fmt.Print("Enter the data number you want to delete: ")
	} else {
		fmt.Print("Masukkan nomor data yang ingin dihapus: ")
	}
	fmt.Scan(&index)
	if index == -1 {
		return // Kembali ke menu utama
	}

	index-- // Konversi nomor ke indeks (dimulai dari 0)

	// Validasi indeks
	if index < 0 || index >= len(core.WasteData) {
		if core.SwitchLanguage {
			fmt.Println("Invalid data number.")
		} else {
			fmt.Println("Nomor data tidak valid.")
		}

		fmt.Println()

		if core.SwitchLanguage {
			utils.PressToContinue("Press Enter to continue...")
		} else {
			utils.PressToContinue("Tekan Enter untuk melanjutkan...")
		}
		utils.ClearConsole()

		return
	}

	if core.SwitchLanguage {
		fmt.Print("Are you sure you want to delete this data? [y/n]: ")
	} else {
		fmt.Print("Yakin ingin menghapus data? [y/n]: ")
	}
	fmt.Scan(&confirmDelete)
	if confirmDelete == "-1" {
		return // Kembali ke menu utama
	}

	// Proses penghapusan berdasarkan konfirmasi
	if utils.StrToLower(confirmDelete) == "y" {
		core.DeleteData(index) // Hapus data berdasarkan indeks

		utils.ClearConsole()

		if core.SwitchLanguage {
			fmt.Println("Data successfully deleted.")
		} else {
			fmt.Println("Data berhasil dihapus.")
		}
	} else if utils.StrToLower(confirmDelete) == "n" {
		fmt.Println()

		if core.SwitchLanguage {
			fmt.Println("Data deletion canceled.")
		} else {
			fmt.Println("Penghapusan data dibatalkan.")
		}
	} else {
		fmt.Println()

		if core.SwitchLanguage {
			fmt.Println("Invalid input!")
		} else {
			fmt.Println("Input tidak valid!")
		}
	}

	if core.SwitchLanguage {
		utils.PressToContinue("Press Enter to back to the main menu...")
	} else {
		utils.PressToContinue("Tekan Enter untuk kembali ke menu utama...")
	}

	utils.ClearConsole()
}
