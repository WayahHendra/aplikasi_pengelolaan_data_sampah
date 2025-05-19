package cli_menus

import (
	"fmt"
	"trash-app/core"
)

// ShowTableMenu menampilkan menu utama untuk aplikasi.
func ShowTableMenu() {
	if core.SwitchLanguage {
		fmt.Println("=============================================================")
		fmt.Println("||        WELCOME TO THE TRASH RECYCLING APP", core.Version, "       ||")
		fmt.Println("||=========================================================||")
		fmt.Println("||  1. Add data                        C                   ||")
		fmt.Println("||  2. Show all data                   R                   ||")
		fmt.Println("||  3. Edit data                       U                   ||")
		fmt.Println("||  4. Delete data                     D                   ||")
		fmt.Println("||  5. Search data                                         ||")
		fmt.Println("||  6. Sort data                                           ||")
		fmt.Println("||  7. Record recycling process                            ||")
		fmt.Println("||  8. Show statistics                                     ||")
		fmt.Println("||  9. Save data                                           ||")
		fmt.Println("||  10. Load data                                          ||")
		fmt.Println("||  11. Switch language                                    ||")
		fmt.Println("||  12. Exit program                                       ||")
		fmt.Println("=============================================================")
		fmt.Print("Select menu (1-11): ")
	} else {
		fmt.Println("=============================================================")
		fmt.Println("||   SELAMAT DATANG DI APLIKASI DAUR ULANG SAMPAH", core.Version, "  ||")
		fmt.Println("||=========================================================||")
		fmt.Println("||  1. Tambah data                     C                   ||")
		fmt.Println("||  2. Tampilkan semua data            R                   ||")
		fmt.Println("||  3. Ubah data                       U                   ||")
		fmt.Println("||  4. Hapus data                      D                   ||")
		fmt.Println("||  5. Cari data                                           ||")
		fmt.Println("||  6. Urutkan data                                        ||")
		fmt.Println("||  7. Catat proses daur ulang                             ||")
		fmt.Println("||  8. Tampilkan statistik                                 ||")
		fmt.Println("||  9. Simpan data                                         ||")
		fmt.Println("||  10. Muat data                                          ||")
		fmt.Println("||  11. Ganti bahasa                                       ||")
		fmt.Println("||  12. Keluar program                                     ||")
		fmt.Println("=============================================================")
		fmt.Print("Pilih menu (1-11): ")
	}
}
