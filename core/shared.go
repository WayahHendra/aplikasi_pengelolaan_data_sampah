package core

import (
	"fmt"
	"trash-app/utils"
)

// ShowRecycleTypeTable menampilkan tabel jenis sampah, metode daur ulang, dan deskripsi singkat.
func ShowRecycleTypeTable() {
	if SwitchLanguage {
		fmt.Println("Available recycling data:")
		fmt.Println("=================================================================================================================================")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "Waste Type", "Recycling Method", "Short Description")
		fmt.Println("||=============================================================================================================================||")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[1] Compost", "Waste like leaves and food scraps turned into fertilizer")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "[1] Organic", "[2] Biogas", "Organic waste fermented to produce gas and fertilizer")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[3] BSF", "BSF larvae decompose waste, resulting in feed and compost")
		fmt.Println("||=============================================================================================================================||")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[1] Melting", "Metals & plastics melted")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[2] Reuse", "Bottles or packaging reused without reprocessing")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "[2] Inorganic", "[3] Downcycling", "Plastic processed into lower quality products (e.g. textiles)")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[4] Upcycling", "Waste turned into crafts or creative products")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[5] Ecobricks", "Plastic waste compressed into bottles for building materials")
		fmt.Println("||=============================================================================================================================||")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[1] Incineration", "Waste burned at high temperatures to ensure safety")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "[3] B3 (Hazardous)", "[2] Stabilization", "Waste turned into solid form to prevent environmental contamination")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[3] Electronics", "E-waste recycled to recover precious metals (gold, copper)")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[4] Chemical", "Liquid waste neutralized or treated with special chemicals")
		fmt.Println("=================================================================================================================================")
	} else {
		fmt.Println("Data daur ulang yang tersedia:")
		fmt.Println("=================================================================================================================================")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "Jenis Sampah", "Metode Daur Ulang", "Deskripsi Singkat")
		fmt.Println("||=============================================================================================================================||")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[1] Kompos", "Sampah seperti daun dan sisa makanan diubah menjadi pupuk")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "[1] Organik", "[2] Biogas", "Limbah organik difermentasi menghasilkan gas dan pupuk")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[3] BSF", "Larva BSF mengurai sampah, hasilnya jadi pakan dan kompos")
		fmt.Println("||=============================================================================================================================||")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[1] Peleburan", "Logam & plastik dilebur")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[2] Reuse", "Botol atau kemasan dipakai kembali tanpa diproses ulang")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "[2] Anorganik", "[3] Downcycling", "Plastik diolah menjadi produk kualitas lebih rendah (mis. tekstil)")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[4] Upcycling", "Sampah dijadikan barang kerajinan atau produk kreatif")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[5] Ecobricks", "Sampah plastik dimampatkan ke dalam botol jadi bahan bangunan")
		fmt.Println("||=============================================================================================================================||")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[1] Insinerasi", "Limbah dibakar dalam suhu tinggi agar aman")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "[3] B3 (Berbahaya)", "[2] Stabilisasi", "Limbah diubah jadi bentuk padat agar tak mencemari lingkungan")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[3] Elektronik", "E-waste didaur ulang untuk ambil logam berharga (emas, tembaga)")
		fmt.Printf("|| %-20s || %-25s || %-70s ||\n", "", "[4] Kimia", "Limbah cair dinetralkan atau diolah dengan bahan kimia khusus")
		fmt.Println("=================================================================================================================================")
	}
}

// GarbageTypes mengembalikan jenis sampah berdasarkan input pengguna.
func GarbageTypes(input string) string {
	switch utils.StrToLower(input) {
	case "1", "organic":
		return "Organic"
	case "2", "inorganic":
		return "Inorganic"
	case "3", "b3":
		return "B3"
	default:
		return "nil"
	}
}

// RecyclingMethods mengembalikan metode daur ulang berdasarkan input dan jenis sampah.
func RecyclingMethods(input, wasteType string) string {
	switch wasteType {
	case "Organic":
		switch utils.StrToLower(input) {
		case "1", "compost":
			return "Compost"
		case "2", "biogas":
			return "Biogas"
		case "3", "bsf":
			return "BSF"
		default:
			return "nil"
		}
	case "Inorganic":
		switch utils.StrToLower(input) {
		case "1", "melting":
			return "Melting"
		case "2", "reuse":
			return "Reuse"
		case "3", "downcycling":
			return "Downcycling"
		case "4", "upcycling":
			return "Upcycling"
		case "5", "ecobricks":
			return "Ecobricks"
		default:
			return "nil"
		}
	case "B3":
		switch utils.StrToLower(input) {
		case "1", "incineration":
			return "Incineration"
		case "2", "stabilization":
			return "Stabilization"
		case "3", "electronics":
			return "Electronics"
		case "4", "chemical":
			return "Chemical"
		default:
			return "nil"
		}
	default:
		return "nil"
	}
}
