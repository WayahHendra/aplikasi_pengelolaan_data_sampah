package gui

import routers "trash-app/gui/api/router"

func RunGUI() {
	routers.NewTrashRouters()
}
