package core

import "trash-app/utils"

// SequentialSearch mencari data sampah berdasarkan jenis menggunakan Sequential Search.
func SequentialSearch(query string, field string) []Waste {
	var results []Waste
	query = utils.StrToLower(query)

	// Iterasi melalui semua data sampah
	for i := 0; i < len(WasteData); i++ {
		var waste Waste = WasteData[i] // Ambil data sampah dari WasteData

		// Pencarian berdasarkan kolom
		switch field {
		case "WasteType":
			if utils.Contains(utils.StrToLower(waste.WasteType), query) {
				results = append(results, waste)
			}
		case "RecyclingMethod":
			if utils.Contains(utils.StrToLower(waste.RecyclingMethod), query) {
				results = append(results, waste)
			}
		case "Location":
			if utils.Contains(utils.StrToLower(waste.Location), query) {
				results = append(results, waste)
			}
		case "Status":
			if utils.Contains(utils.StrToLower(waste.Status), query) {
				results = append(results, waste)
			}
		}
	}

	return results
}

// BinarySearch mencari data sampah berdasarkan jenis menggunakan Binary Search.
// Pastikan data sudah diurutkan berdasarkan WasteType sebelum menggunakan fungsi ini.
func BinarySearch(query string, field string) []Waste {
	var results []Waste
	query = utils.StrToLower(query)

	// Urutkan data menggunakan algoritma sorting yang sudah dibuat
	SortByAlgorithm("1", field, true) // Gunakan Selection Sort dengan urutan ascending

	// Lakukan Binary Search
	low, high := 0, len(WasteData)-1
	for low <= high {
		mid := (low + high) / 2
		var midValue string
		switch field {
		case "WasteType":
			midValue = utils.StrToLower(WasteData[mid].WasteType)
		case "RecyclingMethod":
			midValue = utils.StrToLower(WasteData[mid].RecyclingMethod)
		case "Location":
			midValue = utils.StrToLower(WasteData[mid].Location)
		case "Status":
			midValue = utils.StrToLower(WasteData[mid].Status)
		}

		if utils.Contains(midValue, query) {
			// Tambahkan elemen yang cocok pada indeks tengah
			results = append(results, WasteData[mid])

			// Cari elemen yang cocok di kiri dan kanan indeks tengah
			left, right := mid-1, mid+1
			for left >= 0 {
				var leftValue string
				switch field {
				case "WasteType":
					leftValue = utils.StrToLower(WasteData[left].WasteType)
				case "RecyclingMethod":
					leftValue = utils.StrToLower(WasteData[left].RecyclingMethod)
				case "Location":
					leftValue = utils.StrToLower(WasteData[left].Location)
				case "Status":
					leftValue = utils.StrToLower(WasteData[left].Status)
				}
				if utils.Contains(leftValue, query) {
					results = append(results, WasteData[left])
				} else {
					break
				}
				left--
			}
			for right < len(WasteData) {
				var rightValue string
				switch field {
				case "WasteType":
					rightValue = utils.StrToLower(WasteData[right].WasteType)
				case "RecyclingMethod":
					rightValue = utils.StrToLower(WasteData[right].RecyclingMethod)
				case "Location":
					rightValue = utils.StrToLower(WasteData[right].Location)
				case "Status":
					rightValue = utils.StrToLower(WasteData[right].Status)
				}
				if utils.Contains(rightValue, query) {
					results = append(results, WasteData[right])
				} else {
					break
				}
				right++
			}
			break
		} else if query < midValue {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return results
}
