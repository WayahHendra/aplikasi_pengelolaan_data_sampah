package menus

import (
	"fmt"
	"sampah-app/cli/data"
	"sampah-app/cli/logic"
	"sampah-app/cli/utils"
)

func HandleSelection(value int, breakLoop *bool) {
	switch value {
	case 1:
		logic.AddData()
	case 2:
		logic.ShowAllData()
	case 3:
		logic.UpdateData()
	case 4:
		logic.DeleteData()
	case 5:
		logic.SearchData()
	case 6:
		logic.SortData()
	case 7:
		logic.RecordProcess()
	case 8:
		logic.ShowStatistics()
	case 9:
		data.SaveWasteData(data.WasteData)
	case 10:
		data.LoadWasteData(&data.WasteData)
	case 11:
		fmt.Println("Keluar dari program!")
		*breakLoop = true
	default:
		fmt.Println("Pilihan tidak valid, coba lagi.")
		utils.PressToContinue()
		utils.ClearConsole()
	}
}
