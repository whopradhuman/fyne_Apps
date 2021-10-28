package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	a:= app.New()
	w := a.NewWindow("Image Viewer")
	w.Resize(fyne.Size{800, 450})
	//hello := widget.NewLabel("Hello")

	path := "/home/prxd/Pictures/"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	//var picsArr[] string

	tabs := container.NewAppTabs(
		//container.NewTabItem("Image 1", canvas.NewImageFromFile(picsArr[0])),
		//container.NewTabItem("Image 2", canvas.NewImageFromFile(picsArr[1])),
		//container.NewTabItem("Image 2", canvas.NewImageFromFile(picsArr[2])),
		)

	for _, file := range files {
		//fmt.Println(file.Name(), file.IsDir())

		if !file.IsDir() {
			brokenString := strings.Split(file.Name(), ".")
			imgExtension := brokenString[len(brokenString) - 1]
			if imgExtension == "png" || imgExtension == "jpeg" || imgExtension == "jpg" {
				//picsArr = append(picsArr, path + "/" + file.Name())
					image := canvas.NewImageFromFile(path + "/" + file.Name())
					image.FillMode = canvas.ImageFillContain
					tabs.Append(container.NewTabItem(file.Name(), image))

			}
		}
	}


	//image := canvas.NewImageFromFile(picsArr[0])

	//for i:= 1; i < len(picsArr); i++ {
	//	image := canvas.NewImageFromFile(picsArr[i])
	//	image.FillMode = canvas.ImageFillContain
	//	tabs.Append(container.NewTabItem("Image " + strconv.Itoa(i + 1), image))
	//}

	tabs.SetTabLocation(container.TabLocationTrailing)
	w.SetContent(
		tabs,
		)




	//w.SetContent(container.NewVBox(
	//	hello,
	//	))

	w.ShowAndRun()
}