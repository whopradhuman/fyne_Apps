package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
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

var panContent *fyne.Container

func main() {
	myApp.Settings().SetTheme(theme.DarkTheme())
	img = canvas.NewImageFromFile("background")

	calcBtn = widget.NewButton("Calculator", func() {

	})

	picBtn = widget.NewButton("Gallery", func() {

	})

	weaBtn = widget.NewButton("Weather", func() {

	})

	ediBtn = widget.NewButton("Text Editor", func() {

	})
}