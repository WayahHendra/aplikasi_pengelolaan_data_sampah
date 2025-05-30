package modules

import (
	controllers "trash-app/gui/api/controller"
	services "trash-app/gui/api/service"
)

type TrashModules struct {
	TrashController *controllers.TrashControllers
}

func NewTrashModules() *TrashModules {
	trashServices := services.NewTrashServices()
	trashControllers := controllers.NewTrashControllers(trashServices)

	return &TrashModules{
		TrashController: trashControllers,
	}
}
