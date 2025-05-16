package utils

import "strings"

func StrToLower(s string) string {
	runes := []rune(s) // Ubah string ke slice rune agar mendukung karakter Unicode
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 'a' - 'A' // Ubah ke huruf kecil
		}
	}
	return string(runes) // Ubah kembali ke string
}

func StrToUpper(s string) string {
	runes := []rune(s) // Ubah string ke slice rune agar mendukung karakter Unicode
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] -= 'a' - 'A' // Ubah ke huruf kapital
		}
	}
	return string(runes) // Ubah kembali ke string
}

func EnsureJSONExtension(filename string) string {
	if !strings.HasSuffix(filename, ".json") {
		return filename + ".json"
	}
	return filename
}
