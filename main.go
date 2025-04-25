package main
import "fmt"

// banner function for showing menu
func banner() {
	fmt.Println("\nSELAMAT DATANG DI APLIKASI DAUR ULANG SAMPAH")
	fmt.Println("===========================================")
	fmt.Println("1. Tambah data sampah")
	fmt.Println("2. Ubah data sampah")
	fmt.Println("3. Hapus data sampah")
	fmt.Println("4. Catat proses daur ulang")
	fmt.Println("5. Cari data sampah")
	fmt.Println("6. Urutkan data sampah")
	fmt.Println("7. Tampilkan statistik")
	fmt.Println("8. Tampilkan semua data")
	fmt.Println("9. Keluar program")
	fmt.Println("===========================================")
	fmt.Print("Pilih menu (1-9): \n")
}

func main () {
	banner()
}