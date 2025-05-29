package routers

import (
	"log"
	"net/http"
	"trash-app/gui/api/module"
	logger "trash-app/utils"
)

func NewTrashRouters() {
	modules := modules.NewTrashModules()
	trashController := modules.TrashController

	router := http.NewServeMux()
	router.HandleFunc("/waste/add", trashController.CreateWaste)
	router.HandleFunc("/waste/get-all", trashController.GetAllWaste)
	router.HandleFunc("/waste/get-date/", trashController.GetWasteByDate)
	router.HandleFunc("/waste/update/", trashController.UpdateWasteById)
	router.HandleFunc("/waste/delete/", trashController.DeleteWasteById)

	log.Print(logger.Types["info"], "Server running on 127.0.0.1 with port 8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}
