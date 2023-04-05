package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("SleepTimer")
	myWindow.Resize(fyne.NewSize(500, 500))

	// Adiciona um widget de entrada de texto
	timeEntry := widget.NewEntry()
	timeEntry.SetPlaceHolder("Digite um tempo em minutos")

	// Adiciona um botão para selecionar o tempo
	confirmButton := widget.NewButton("Selecionar", func() {
		selectedTime := timeEntry.Text
		fmt.Println(selectedTime)
		// faça algo com o tempo selecionado...go
	})
	fifteenButton := widget.NewButton("15 minutes", func() {
		//fazer alguma acao com o botao de 15 minutos
	})
	thirtyButton := widget.NewButton("30 minutes", func() {
		//fazer alguma acao com o botao de 30 minutos
	})
	hourButton := widget.NewButton("1 hour", func() {
		//fazer alguma acao com o botao de 1 hora
	})
	myContainer := container.New(layout.NewVBoxLayout(),
		timeEntry,
		confirmButton,
		fifteenButton,
		thirtyButton,
		hourButton,
	)
	fifteenButton.Resize(fyne.NewSize(50, 50))
	// Cria um conteúdo vertical que contém o campo de entrada de tempo e o botão de seleção

	// Define o conteúdo da janela e mostra-a
	myWindow.SetContent(myContainer)
	myWindow.ShowAndRun()
}
