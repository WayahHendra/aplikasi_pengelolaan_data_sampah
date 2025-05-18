package cli_menus

import (
	"fmt"
	logic "trash-app/cli/cli-logic"
	"trash-app/core"
	"trash-app/utils"
)

func HandleSelection(value int, breakLoop *bool) {
	switch value {
	case 1:
		logic.CreateWaste()
	case 2:
		logic.ReadWaste()
	case 3:
		logic.UpdateWaste()
	case 4:
		logic.DeleteWaste()
	case 5:
		logic.SearchWaste()
	case 6:
		logic.SortWaste()
	case 7:
		logic.RecordProcess()
	case 8:
		logic.ShowStatistics()
	case 9:
		logic.SaveWaste(core.WasteData)
	case 10:
		logic.LoadWaste(&core.WasteData)
	case 11:
		fmt.Println("Keluar dari program!")
		*breakLoop = true
	default:
		fmt.Println("Pilihan tidak valid, coba lagi.")
		utils.PressToContinue()
		utils.ClearConsole()
	}
}
