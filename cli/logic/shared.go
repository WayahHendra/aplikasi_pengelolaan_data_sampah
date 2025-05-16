package logic

import (
	"fmt"
	"sampah-app/cli/utils"
)

func ShowRecycleTypeTable() { // showRecycleTypeTable() untuk menampilkan tabel jenis-jenis sampah daur ulang
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

func GarbageTypes(input string) string {
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

func RecyclingMethods(input, wasteType string) string {
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
		case "1", "p", "pe", "pel", "pele", "peleb", "pelebu", "pelebura", "peleburan":
			return "Peleburan"
		case "2", "r", "re", "reu", "reus", "reuse":
			return "Reuse"
		case "3", "d", "do", "dow", "down", "downc", "downcy", "downcycl", "downcycli", "downcycling":
			return "Downcycling"
		case "4", "u", "up", "upc", "upcy", "upcycl", "upcycli", "upcycling":
			return "Upcycling"
		case "5", "e", "ec", "eco", "ecob", "ecobr", "ecobri", "ecobric", "ecobricks":
			return "Ecobricks"
		default:
			return "nil"
		}
	case "B3":
		switch utils.StrToLower(input) {
		case "1", "i", "in", "ins", "insi", "insin", "insine", "insiner", "insinera", "insineras", "insinerasi":
			return "Insinerasi"
		case "2", "s", "st", "sta", "stab", "stabi", "stabil", "stabilis", "stabilisa", "stabilisasi":
			return "Stabilisasi"
		case "3", "e", "el", "ele", "elek", "elekt", "elektr", "elektron", "elektroni", "elektronik":
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
