package routers

import (
	"log"
	"net/http"
	"trash-app/gui/api/module"
)

func NewTrashRouters() {
	modules := modules.NewTrashModules()
	trashController := modules.TrashController

	router := http.NewServeMux()
	router.HandleFunc("/waste/add", trashController.CreateWaste)
	router.HandleFunc("/waste/get-all", trashController.GetAllWaste)

	log.Println("Server running on 127.0.0.1 with port 8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}
