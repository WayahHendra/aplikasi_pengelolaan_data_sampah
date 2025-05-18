package cli_menus

import (
	"fmt"
	"trash-app/core"
)

func ShowTableMenu() { // showTableMenu() untuk memunculkan menu
	fmt.Println("===========================================================")
	fmt.Println("||   SELAMAT DATANG DI APLIKASI DAUR ULANG SAMPAH", core.Version, "  ||")
	fmt.Println("||=======================================================||")
	fmt.Println("||  1. Tambah data                     C                 ||")
	fmt.Println("||  2. Tampilkan semua data            R                 ||")
	fmt.Println("||  3. Ubah data                       U                 ||")
	fmt.Println("||  4. Hapus data                      D                 ||")
	fmt.Println("||  5. Cari data                                         ||")
	fmt.Println("||  6. Urutkan data                                      ||")
	fmt.Println("||  7. Catat proses daur ulang                           ||")
	fmt.Println("||  8. Tampilkan statistik                               ||")
	fmt.Println("||  9. Save data                                         ||")
	fmt.Println("||  10. Load data                                        ||")
	fmt.Println("||  11. Keluar program                                   ||")
	fmt.Println("===========================================================")
	fmt.Print("Pilih menu (1-11): ")
}
