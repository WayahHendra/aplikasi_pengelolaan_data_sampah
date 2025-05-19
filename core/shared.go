package core

import (
	"fmt"
	"trash-app/utils"
)

// ShowRecycleTypeTable menampilkan tabel jenis sampah, metode daur ulang, dan deskripsi singkat.
func ShowRecycleTypeTable() {
	if SwitchLanguage {
		fmt.Println("Available recycling data:")
		fmt.Println("=============================================================================================================================")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "Waste Type", "Recycling Method", "Short Description")
		fmt.Println("||=========================================================================================================================||")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(1) Compost", "Waste like leaves and food scraps turned into fertilizer")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "(1) Organic", "(2) Biogas", "Organic waste fermented to produce gas and fertilizer")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(3) BSF", "BSF larvae decompose waste, resulting in feed and compost")
		fmt.Println("||=========================================================================================================================||")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(1) Melting", "Metals & plastics melted")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(2) Reuse", "Bottles or packaging reused without reprocessing")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "(2) Inorganic", "(3) Downcycling", "Plastic processed into lower quality products (e.g. textiles)")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(4) Upcycling", "Waste turned into crafts or creative products")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(5) Ecobricks", "Plastic waste compressed into bottles for building materials")
		fmt.Println("||=========================================================================================================================||")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(1) Incineration", "Waste burned at high temperatures to ensure safety")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "(3) B3 (Hazardous)", "(2) Stabilization", "Waste turned into solid form to prevent environmental contamination")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(3) Electronics", "E-waste recycled to recover precious metals (gold, copper)")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(4) Chemical", "Liquid waste neutralized or treated with special chemicals")
		fmt.Println("=============================================================================================================================")
	} else {
		fmt.Println("Data daur ulang yang tersedia:")
		fmt.Println("=============================================================================================================================")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "Jenis Sampah", "Metode Daur Ulang", "Deskripsi Singkat")
		fmt.Println("||=========================================================================================================================||")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(1) Kompos", "Sampah seperti daun dan sisa makanan diubah menjadi pupuk")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "(1) Organik", "(2) Biogas", "Limbah organik difermentasi menghasilkan gas dan pupuk")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(3) BSF", "Larva BSF mengurai sampah, hasilnya jadi pakan dan kompos")
		fmt.Println("||=========================================================================================================================||")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(1) Peleburan", "Logam & plastik dilebur")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(2) Reuse", "Botol atau kemasan dipakai kembali tanpa diproses ulang")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "(2) Anorganik", "(3) Downcycling", "Plastik diolah menjadi produk kualitas lebih rendah (mis. tekstil)")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(4) Upcycling", "Sampah dijadikan barang kerajinan atau produk kreatif")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(5) Ecobricks", "Sampah plastik dimampatkan ke dalam botol jadi bahan bangunan")
		fmt.Println("||=========================================================================================================================||")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(1) Insinerasi", "Limbah dibakar dalam suhu tinggi agar aman")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "(3) B3 (Berbahaya)", "(2) Stabilisasi", "Limbah diubah jadi bentuk padat agar tak mencemari lingkungan")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(3) Elektronik", "E-waste didaur ulang untuk ambil logam berharga (emas, tembaga)")
		fmt.Printf("|| %-20s || %-25s || %-66s ||\n", "", "(4) Kimia", "Limbah cair dinetralkan atau diolah dengan bahan kimia khusus")
		fmt.Println("=============================================================================================================================")
	}
}

// GarbageTypes mengembalikan jenis sampah berdasarkan input pengguna.
func GarbageTypes(input string) string {
	if SwitchLanguage {
		switch utils.StrToLower(input) {
		case "1", "o", "or", "org", "orga", "organ", "organi", "organic":
			return "Organic"
		case "2", "i", "in", "ino", "inor", "inorg", "inorga", "inorgan", "inorgani", "inorganic":
			return "Inorganic"
		case "3", "b", "b3":
			return "B3"
		default:
			return "nil"
		}
	} else {
		switch utils.StrToLower(input) {
		case "1", "o", "or", "org", "orga", "organ", "organi", "organik":
			return "Organik"
		case "2", "a", "an", "ano", "anor", "anorg", "anorga", "anorgan", "anorgani", "anorganik":
			return "Anorganik"
		case "3", "b", "b3":
			return "B3"
		default:
			return "nil"
		}
	}
}

// RecyclingMethods mengembalikan metode daur ulang berdasarkan input dan jenis sampah.
func RecyclingMethods(input, wasteType string) string {
	if SwitchLanguage {
		switch wasteType {
		case "Organic":
			switch utils.StrToLower(input) {
			case "1", "c", "co", "com", "comp", "compo", "compos", "compost":
				return "Compost"
			case "2", "b", "bi", "bio", "biog", "bioga", "biogas":
				return "Biogas"
			case "3", "bs", "bsf":
				return "BSF"
			default:
				return "nil"
			}
		case "Inorganik":
			switch utils.StrToLower(input) {
			case "1", "m", "me", "mel", "melt", "melti", "meltin", "melting":
				return "Melting"
			case "2", "r", "re", "reu", "reus", "reuse":
				return "Reuse"
			case "3", "d", "do", "dow", "down", "downc", "downcy", "downcyc", "downcycl", "downcycli", "downcyclin", "downcycling":
				return "Downcycling"
			case "4", "u", "up", "upc", "upcy", "upcyc", "upcycl", "upcycli", "upcyclin", "upcycling":
				return "Upcycling"
			case "5", "e", "ec", "eco", "ecob", "ecobr", "ecobri", "ecobric", "ecobrick", "ecobricks":
				return "Ecobricks"
			default:
				return "nil"
			}
		case "B3":
			switch utils.StrToLower(input) {
			case "1", "i", "in", "inc", "inci", "incin", "incine", "inciner", "incinera", "incinerat", "incinerati", "incineratio", "incineration":
				return "Incineration"
			case "2", "s", "st", "sta", "stab", "stabi", "stabil", "stabili", "stabiliz", "stabiliza", "stabilizat", "stabilizati", "stabilizatio", "stabilization":
				return "Stabilization"
			case "3", "e", "el", "ele", "elec", "elect", "electr", "electro", "electron", "electroni", "electronic", "electronics":
				return "Electronics"
			case "4", "c", "ch", "che", "chem", "chemi", "chemic", "chemica", "chemical":
				return "Chemical"
			default:
				return "nil"
			}
		default:
			return "nil"
		}
	} else {
		switch wasteType {
		case "Organik":
			switch utils.StrToLower(input) {
			case "1", "k", "ko", "kom", "komp", "kompo", "kompos":
				return "Kompos"
			case "2", "b", "bi", "bio", "biog", "bioga", "biogas":
				return "Biogas"
			case "3", "bs", "bsf":
				return "BSF"
			default:
				return "nil"
			}
		case "Anorganik":
			switch utils.StrToLower(input) {
			case "1", "p", "pe", "pel", "pele", "peleb", "pelebu", "pelebur", "pelebura", "peleburan":
				return "Peleburan"
			case "2", "r", "re", "reu", "reus", "reuse":
				return "Reuse"
			case "3", "d", "do", "dow", "down", "downc", "downcy", "downcyc", "downcycl", "downcycli", "downcyclin", "downcycling":
				return "Downcycling"
			case "4", "u", "up", "upc", "upcy", "upcyc", "upcycl", "upcycli", "upcyclin", "upcycling":
				return "Upcycling"
			case "5", "e", "ec", "eco", "ecob", "ecobr", "ecobri", "ecobric", "ecobrick", "ecobricks":
				return "Ecobricks"
			default:
				return "nil"
			}
		case "B3":
			switch utils.StrToLower(input) {
			case "1", "i", "in", "ins", "insi", "insin", "insine", "insiner", "insinera", "insineras", "insinerasi":
				return "Insinerasi"
			case "2", "s", "st", "sta", "stab", "stabi", "stabil", "stabili", "stabilis", "stabilisa", "stabilisas", "stabilisasi":
				return "Stabilisasi"
			case "3", "e", "el", "ele", "elek", "elekt", "elektr", "elektro", "elektron", "elektroni", "elektronik":
				return "Elektronik"
			case "4", "k", "ki", "kim", "kimi", "kimia":
				return "Kimia"
			default:
				return "nil"
			}
		default:
			return "nil"
		}
	}
}
