package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
)

var count int = 1
func main() {



	fmt.Println("Hello Texty!")
	a := app.New()
	w := a.NewWindow("Texty")

	w.Resize(fyne.Size{600, 400})
	content := container.NewVBox(
		widget.NewLabel("Hey Texty"),
		)

	inputField := widget.NewMultiLineEntry()
	inputField.SetPlaceHolder("texTy, write here...")

	//content.Add(widget.NewButton("New File ", func() {
	//	content.Add(widget.NewLabel("New File " + strconv.Itoa(count)))
	//	count++
	//}))


	fileName := ""
	saveBtn := widget.NewButton("Save File", func() {
		saveDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(inputField.Text)
				uc.Write(textData)
			},
			w)

		if len(fileName) == 0 {
			fileName = "New File.txt"
		}
		saveDialog.SetFileName(fileName)
		saveDialog.Show()
	})

	openBtn := widget.NewButton("Open File", func() {
		openDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				readValue , _:= ioutil.ReadAll(r)
				out := fyne.NewStaticResource("New File", readValue)

				inputField.SetText(string(out.StaticContent))
				fileName = out.StaticName

				//viewable := widget.NewMultiLineEntry()
				//viewable.SetText(string(out.StaticContent))
				//
				//w := fyne.CurrentApp().NewWindow(
				//	string(out.StaticName))
				//w.SetContent(container.NewScroll(viewable))
				//w.Resize(fyne.NewSize(400,400))
				//
				//w.Show()
			},
		w)

		openDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		openDialog.Show()
	})

	w.SetContent(
		container.NewVBox(
			content,
			inputField,
			//container.NewHBox(
				container.NewGridWithColumns( 2,
						saveBtn,
						openBtn,
					),
				//),
			),
		)

	w.ShowAndRun()
}