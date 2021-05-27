package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

func GenerateMainMenu(w fyne.Window) {
	menu := fyne.NewMenu("File",
		fyne.NewMenuItem("Open", func() { OpenShapeFile(w) }),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("About", func() { fmt.Println("About") }))
	w.SetMainMenu(fyne.NewMainMenu(menu))
}

func OpenShapeFile(parent fyne.Window) {
	fileopen := dialog.NewFileOpen(func(uc fyne.URIReadCloser, e error) {
		fmt.Println(uc.URI(), e)
	}, parent)
	fileopen.SetFilter(storage.NewExtensionFileFilter([]string{".shp"}))
	fileopen.Show()
}
