package utils

import "strings"

func StrToLower(text string) string {
	return strings.ToLower(text)
}

func StrToUpper(text string) string {
	return strings.ToUpper(text)
}

func EnsureJSONExtension(filename string) string {
	if !strings.HasSuffix(filename, ".json") {
		return filename + ".json"
	}
	return filename
}
