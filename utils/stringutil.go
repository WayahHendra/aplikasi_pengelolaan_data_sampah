package utils

import "strings"

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
