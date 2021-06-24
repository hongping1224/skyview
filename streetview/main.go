package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/hongping1224/skyview/util"
	"github.com/webview/webview"
)

var (
	appID       = "com.ARSEM.skyview。streetview"
	windowName  = "Street View"
	entryString = binding.NewString()
	labelStirng = binding.NewString()
)

type enterEntry struct {
	widget.Entry
}

func (e *enterEntry) onEnter() {
	buttonPress()
}

func newEnterEntry() *enterEntry {
	entry := &enterEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func newEnterEntryWithData(data binding.String) *enterEntry {
	entry := newEnterEntry()
	entry.Bind(data)
	return entry
}

func (e *enterEntry) TypedKey(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyEnter:
		e.onEnter()
	default:
		e.Entry.TypedKey(key)
	}
}

func main() {
	a := app.NewWithID(appID)
	w := a.NewWindow(windowName)

	button := widget.NewButton("Open WebView", func() {
		buttonPress()
	})
	button2 := widget.NewButton("Open Native", func() {
		buttonPressNative()
	})

	label := widget.NewLabelWithData(labelStirng)
	entry := newEnterEntryWithData(entryString)
	w.SetContent(container.NewVBox(
		entry,
		label,
		button,
		button2,
	))

	w.Resize(fyne.NewSize(300, 130))
	w.ShowAndRun()

}

func buttonPress() {
	s, _ := entryString.Get()
	url, err := decodeCoordinate(s)
	labelStirng.Set("")
	if err != nil {
		labelStirng.Set("Format Error, Example:\n169,885.900  2,544,297.055")
		return
	}
	go openStreetView(url)
}

func buttonPressNative() {
	s, _ := entryString.Get()
	url, err := decodeCoordinate(s)
	labelStirng.Set("")
	if err != nil {
		labelStirng.Set("Format Error, Example:\n169,885.900  2,544,297.055")
		return
	}
	cmd := exec.Command("start", url)

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func openStreetView(input string) error {
	debug := false
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Google Map")
	w.SetSize(1500, 800, webview.HintNone)
	w.Navigate(input)
	w.Run()
	return nil
}

func decodeCoordinate(input string) (string, error) {
	input = strings.ReplaceAll(input, ",", "")
	input = strings.ReplaceAll(input, "\n", "")
	for strings.Contains(input, "  ") {
		input = strings.ReplaceAll(input, "  ", " ")
	}
	input = strings.ReplaceAll(input, " ", ",")
	s := strings.Split(input, ",")
	if len(s) != 2 {
		return "", errors.New("format error")
	}
	x, err := strconv.ParseFloat(s[0], 64)
	if err != nil {
		return "", errors.New("format error")
	}
	y, err := strconv.ParseFloat(s[1], 64)
	if err != nil {
		return "", errors.New("format error")
	}
	lat, lon := util.TWD97ToWGS84(x, y)
	fmt.Println(lat, lon)
	return fmt.Sprintf("https://www.google.com/maps/place/%f,%f", lat, lon), nil
	/*
		"https://www.google.com/maps/place/22°55'26.4\"N+120°17'15.2\""
		"https://www.google.com/maps/place/22%C2%B055'26.4%22N+120%C2%B017'15.2%22E"*/

}
