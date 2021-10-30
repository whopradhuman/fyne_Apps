package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
	"strconv"
)

func main() {
	a := app.New()   // a is app
	w := a.NewWindow("Calculator")  //w is widget
	w.Resize(fyne.Size{500, 330})

	output := ""

	ioBox := widget.NewLabel(output)

	history := ""
	historyBox := widget.NewLabel(history)

	var historyPrinted bool
	var historyArr []string

	historyBtn := widget.NewButton("History", func() {
		for i := len(historyArr) - 1; i >= 0; i-- {
			history += historyArr[i] + "\n"
		}
		if historyPrinted == false{
			historyBox.SetText(history)
			historyPrinted = true
		}
	})

	backBtn := widget.NewButton("Back", func() {
		if len(output) == 0 {
			return
		}
		output = output[: len(output) - 1]
		ioBox.SetText(output)
	})

	clearBtn := widget.NewButton("Clear", func() {
		output = ""
		ioBox.SetText(output)
	})

	bracStartBtn := widget.NewButton("(", func() {
		output += "("
		ioBox.SetText(output)
	})

	bracEndBtn := widget.NewButton(")", func() {
		output += ")"
		ioBox.SetText(output)
	})

	divideBtn := widget.NewButton("/", func() {
		output += "/"
		ioBox.SetText(output)
	})

	sevenBtn := widget.NewButton("7", func() {
		output += "7"
		ioBox.SetText(output)
	})

	eightBtn := widget.NewButton("8", func() {
		output += "8"
		ioBox.SetText(output)
	})

	nineBtn := widget.NewButton("9", func() {
		output += "9"
		ioBox.SetText(output)
	})

	multiplyBtn := widget.NewButton("*", func() {
		output += "*"
		ioBox.SetText(output)
	})

	fourBtn := widget.NewButton("4", func() {
		output += "4"
		ioBox.SetText(output)
	})

	fiveBtn := widget.NewButton("5", func() {
		output += "5"
		ioBox.SetText(output)
	})

	sixBtn := widget.NewButton("6", func() {
		output += "6"
		ioBox.SetText(output)
	})

	subtractBtn := widget.NewButton("-", func() {
		output += "-"
		ioBox.SetText(output)
	})

	oneBtn := widget.NewButton("1", func() {
		output += "1"
		ioBox.SetText(output)
	})

	twoBtn := widget.NewButton("2", func() {
		output += "2"
		ioBox.SetText(output)
	})

	threeBtn := widget.NewButton("3", func() {
		output += "3"
		ioBox.SetText(output)
	})

	additionBtn := widget.NewButton("+", func() {
		output += "+"
		ioBox.SetText(output)
	})

	zeroBtn := widget.NewButton("0", func() {
		output += "0"
		ioBox.SetText(output)
	})

	dotBtn := widget.NewButton(".", func() {
		output += "."
		ioBox.SetText(output)
	})

	equalsBtn := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				historyPrinted = false
				history = ""
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strToAdd := output + " = " + ans
				historyArr = append(historyArr, strToAdd)
				output = ans
			} else {
				output = "error"
			}
		} else {
			output = "error"
		}
		ioBox.SetText(output)
	})

	w.SetContent(container.NewVBox(
		ioBox,
		historyBox,
		container.NewGridWithColumns(2,
			historyBtn,
			backBtn,
		),

		container.NewGridWithColumns(4,
			clearBtn,
			bracStartBtn,
			bracEndBtn,
			divideBtn,
		),

		container.NewGridWithColumns(4,
			sevenBtn,
			eightBtn,
			nineBtn,
			multiplyBtn,
		),

		container.NewGridWithColumns(4,
			fourBtn,
			fiveBtn,
			sixBtn,
			subtractBtn,
		),

		container.NewGridWithColumns(4,
			oneBtn,
			twoBtn,
			threeBtn,
			additionBtn,
		),

		container.NewGridWithColumns(2,
			container.NewGridWithColumns(2,
				zeroBtn,
				dotBtn,
				),
				equalsBtn,
		),

	))

		w.ShowAndRun()
}