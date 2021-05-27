package ui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

)

func GeneratePointTable(w fyne.Window) fyne.CanvasObject {
	splitContainer := container.NewVSplit(widget.NewLabel("測試found image"), widget.NewLabel("entrydata"))
	splitContainer.SetOffset(0.5)
	datatable := generateSHPTableCanvas(w, nil)
	splitContainer2 := container.NewHSplit(datatable, splitContainer)
	splitContainer2.SetOffset(0.5)
	//boarder := container.NewBorder(nil, nil, splitContainer, splitContainer2)
	return splitContainer2
}

func generateSHPTableCanvas(w fyne.Window, rawdata [][]string) fyne.CanvasObject {
	if rawdata == nil {
		return widget.NewButton("Open Shape File", func() {
			OpenShapeFile(w)
		})
	}
	list := widget.NewTable(
		func() (int, int) {
			return len(rawdata), len(rawdata[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(rawdata[i.Row][i.Col])
		})
	list.OnSelected = func(id widget.TableCellID) {
		fmt.Println(rawdata[id.Row][id.Col])
	}
	//vbox := container.NewVBox(list)

	return list

}
func generateSHPTableHeader(header []string) fyne.CanvasObject {
	head := widget.NewTable(func() (int, int) {
		return 1, len(header[0])
	},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(header[i.Col])
		})
	return head
}

func testSHPDATA() (data [][]string) {
	data = make([][]string, 5)
	data[0] = []string{"col1", "col2", "col3", "col4", "col5"}
	for i := 1; i < len(data); i++ {
		data[i] = []string{strconv.Itoa(i), "a", "b", "c", "d"}
	}
	return data
}
