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
		fmt.Println("|| No  ||                    Menu Name                     ||")
		fmt.Println("||=========================================================||")
		fmt.Println("|| 1.  || Add data                                C        ||")
		fmt.Println("|| 2.  || Show data                               R        ||")
		fmt.Println("|| 3.  || Edit data                               U        ||")
		fmt.Println("|| 4.  || Delete data                             D        ||")
		fmt.Println("|| 5.  || Search data                                      ||")
		fmt.Println("|| 6.  || Sort data                                        ||")
		fmt.Println("|| 7.  || Show statistics                                  ||")
		fmt.Println("|| 8.  || Save all data                                    ||")
		fmt.Println("|| 9.  || Load all data                                    ||")
		fmt.Println("|| 10. || Load data by date                                ||")
		fmt.Println("|| 11. || Switch language                                  ||")
		fmt.Println("|| -1. || Exit program                                     ||")
		fmt.Println("=============================================================")
		fmt.Print("Select menu [1-11 or -1]: ")
	} else {
		fmt.Println("=============================================================")
		fmt.Println("||   SELAMAT DATANG DI APLIKASI DAUR ULANG SAMPAH", core.Version, "  ||")
		fmt.Println("||=========================================================||")
		fmt.Println("|| No  ||                    Nama Menu                     ||")
		fmt.Println("||=========================================================||")
		fmt.Println("|| 1.  || Tambah data                             C        ||")
		fmt.Println("|| 2.  || Tampilkan data                          R        ||")
		fmt.Println("|| 3.  || Ubah data                               U        ||")
		fmt.Println("|| 4.  || Hapus data                              D        ||")
		fmt.Println("|| 5.  || Cari data                                        ||")
		fmt.Println("|| 6.  || Urutkan data                                     ||")
		fmt.Println("|| 7.  || Tampilkan statistik                              ||")
		fmt.Println("|| 8.  || Simpan semua data                                ||")
		fmt.Println("|| 9.  || Muat semua data                                  ||")
		fmt.Println("|| 10. || Muat data berdasarkan tanggal                    ||")
		fmt.Println("|| 11. || Ganti bahasa                                     ||")
		fmt.Println("|| -1. || Keluar program                                   ||")
		fmt.Println("=============================================================")
		fmt.Print("Pilih menu [1-11 atau -1]: ")
	}
}

func ShowTableSubMenu() {
	if core.SwitchLanguage {
		fmt.Println("=============================================================")
		fmt.Println("||                 DISPLAY WASTE DATA MENU                 ||")
		fmt.Println("||=========================================================||")
		fmt.Println("|| No  ||                    Menu Name                     ||")
		fmt.Println("||=========================================================||")
		fmt.Println("|| 1.  || Show all waste data                              ||")
		fmt.Println("|| 2.  || Show waste data by date                          ||")
		fmt.Println("|| -1. || Back to main menu                                ||")
		fmt.Println("=============================================================")
		fmt.Print("Select menu [1-2]: ")
	} else {
		fmt.Println("=============================================================")
		fmt.Println("||                MENU TAMPILKAN DATA SAMPAH               ||")
		fmt.Println("||=========================================================||")
		fmt.Println("|| No  ||                    Nama Menu                     ||")
		fmt.Println("||=========================================================||")
		fmt.Println("|| 1.  || Tampilkan semua data sampah                      ||")
		fmt.Println("|| 2.  || Tampilkan data sampah berdasarkan tanggal        ||")
		fmt.Println("|| -1. || Kembali ke menu utama                            ||")
		fmt.Println("=============================================================")
		fmt.Print("Pilih menu [1-2]: ")
	}
}
