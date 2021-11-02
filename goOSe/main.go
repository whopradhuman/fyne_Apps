package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()
var myWin fyne.Window = myApp.NewWindow("goOSe")

var calcBtn fyne.Widget
var picBtn fyne.Widget
var weaBtn fyne.Widget
var ediBtn fyne.Widget

var deskBtn fyne.Widget

var img fyne.CanvasObject

var panelContent *fyne.Container


func main() {
	myApp.Settings().SetTheme(theme.DarkTheme())
	img = canvas.NewImageFromFile("background.jpg")

	calcBtn = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func() {
		showCalculatorApp()
	})

	picBtn = widget.NewButtonWithIcon("Gallery", theme.FileImageIcon(), func() {
		showGalleryApp(myWin)
	})

	weaBtn = widget.NewButtonWithIcon("Weather", theme.InfoIcon(), func() {
		showWeatherApp(myWin)
	})

	ediBtn = widget.NewButtonWithIcon("Text Editor", theme.DocumentIcon(), func() {
		showTextApp()
	})

	deskBtn = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		myWin.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})

	panelContent = container.NewVBox(
		container.NewGridWithColumns(5,
			deskBtn,
			weaBtn,
			calcBtn,
			picBtn,
			ediBtn,
		),
	)

	myWin.Resize(fyne.NewSize(1920, 1080))
	myWin.CenterOnScreen()

	myWin.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))

	myWin.ShowAndRun()
}