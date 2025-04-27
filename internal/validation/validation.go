package validation

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateInput(input string) (int, error) {
	// remove leading and trailing spaces
	input = strings.TrimSpace(input)
	if input == "" {
		return 0, fmt.Errorf("Input tidak boleh kosong")
	}

	// convert string input to integer
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("Input harus berupa angka")
	}

	// check if input is negative
	if value < 0 {
		return 0, fmt.Errorf("Input atau angka tidak boleh negatif")
	}

	return value, nil
}
