package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Contains memeriksa apakah substring ada dalam string.
func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// StrToLower mengubah string menjadi huruf kecil.
func StrToLower(text string) string {
	return strings.ToLower(text)
}

// StrToUpper mengubah string menjadi huruf besar.
func StrToUpper(text string) string {
	return strings.ToUpper(text)
}

// EnsureJSONExtension memastikan nama file memiliki ekstensi .json.
func EnsureJSONExtension(filename string) string {
	if !strings.HasSuffix(filename, ".json") {
		return filename + ".json"
	}
	return filename
}

// FormatingSaveData membuat nama file berdasarkan tanggal saat ini dan memastikan folder data ada.
func FormatingSaveData() string {
	var dataFolder, file string = "../data", ""

	if _, err := os.Stat(dataFolder); os.IsNotExist(err) {
		os.Mkdir(dataFolder, 0755)
	}

	today := time.Now()
	filename := fmt.Sprintf("%02d-%02d-%04d.json",
		today.Day(),
		today.Month(),
		today.Year(),
	)

	file = dataFolder + "/" + filename

	return file
}
