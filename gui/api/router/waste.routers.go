package routers

import (
	"log"
	"net/http"
	"strconv"
	"trash-app/core"
	"trash-app/gui/api/middleware"
	"trash-app/gui/api/module"
	"trash-app/utils"
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

	routers := middleware.CORS(middleware.APIKeyAuth(router))

	port, err := utils.FindAvailablePort(core.PRIORITY_PORTS)
	if err != nil {
		log.Fatalf(utils.Types["error"] + "Failed to find available port: %v", err)
	}

	log.Printf(utils.Types["info"] + "Server running on http://127.0.0.1:%d\n", port)
	http.ListenAndServe("127.0.0.1:" + strconv.Itoa(port), routers)
}
