package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateInput(input string) (int, error) {
	// remove leading and trailing spaces
	input = strings.TrimSpace(input)
	if input == "" {
		return 0, fmt.Errorf("input tidak boleh kosong")
	}

	// convert string input to integer
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("input harus berupa angka")
	}

	// check if input is negative
	if value < 0 {
		return 0, fmt.Errorf("input atau angka tidak boleh negatif")
	}

	return value, nil
}
