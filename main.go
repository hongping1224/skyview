package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/hongping1224/skyview/ui"
)

var (
	appID      = "com.ARSEM.skyview"
	windowName = "Skyview"
)

func main() {

	os.Setenv("FYNE_FONT", "./resources/font/wt014.ttf")

	a := app.NewWithID(appID)
	w := a.NewWindow(windowName)
	ui.GenerateMainMenu(w)
	splitContainer := ui.GeneratePointTable(w)
	splitContainer2 := container.NewHSplit(widget.NewLabel("left"), widget.NewLabel("right"))
	appTabs := container.NewAppTabs(container.NewTabItem("Data", splitContainer), container.NewTabItem("Wrong Entry", splitContainer2))
	w.Resize(fyne.NewSize(800, 800))
	w.SetContent(appTabs)
	w.ShowAndRun()
}
