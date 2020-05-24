package main

import (
	"image/color"
	"time"

	streamdeck "github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/actionhandlers"
	"github.com/magicmonkey/go-streamdeck/buttons"
)

func main() {
	sd, err := streamdeck.New()
	if err != nil {
		panic(err)
	}

	myButton := buttons.NewTextButton("Hi world")
	myButton.SetActionHandler(&actionhandlers.TextPrintAction{Label: "You pressed me"})
	sd.AddButton(2, myButton)

	myOtherButton := buttons.NewTextButton("4")
	myOtherButton.SetActionHandler(&actionhandlers.NumberPrintAction{Number: 5})
	sd.AddButton(3, myOtherButton)

	myNextButton := buttons.NewTextButton("7")
	myNextButton.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "8"})
	sd.AddButton(7, myNextButton)

	anotherButton, _ := buttons.NewImageFileButton("/home/kevin/streamdeck/go-streamdeck/examples/play.jpg")
	sd.AddButton(9, anotherButton)

	cButton := buttons.NewColourButton(color.RGBA{255, 255, 0, 255})
	sd.AddButton(26, cButton)

	multiActionButton := buttons.NewColourButton(color.RGBA{255, 0, 255, 255})
	thisActionHandler := &actionhandlers.ChainedAction{}
	thisActionHandler.AddAction(&actionhandlers.TextPrintAction{Label: "Purple press"})
	thisActionHandler.AddAction(&actionhandlers.ColourChangeAction{NewColour: color.RGBA{255, 0, 0, 255}})
	multiActionButton.SetActionHandler(thisActionHandler)
	sd.AddButton(27, multiActionButton)

	time.Sleep(2 * time.Second)

	myButton.SetText("Bye!")
	cButton.SetColour(color.RGBA{0, 255, 255, 255})
}
