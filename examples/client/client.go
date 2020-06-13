package main

import (
	"fmt"
	"image/color"
	"time"

	streamdeck "github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/actionhandlers"
	"github.com/magicmonkey/go-streamdeck/buttons"
	"github.com/magicmonkey/go-streamdeck/decorators"
	_ "github.com/magicmonkey/go-streamdeck/devices"
)

func main() {
	sd, err := streamdeck.New()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found device [%s]\n", sd.GetName())

	myButton := buttons.NewTextButton("Hi world")
	myButton.SetActionHandler(&actionhandlers.TextPrintAction{Label: "You pressed me"})
	sd.AddButton(2, myButton)

	myOtherButton := buttons.NewTextButton("4")
	myOtherButton.SetActionHandler(&actionhandlers.NumberPrintAction{Number: 5})
	sd.AddButton(3, myOtherButton)

	myNextButton := buttons.NewTextButton("7")
	myNextButton.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "8"})
	sd.AddButton(7, myNextButton)

	anotherButton, err := buttons.NewImageFileButton("examples/test/play.jpg")
	if err != nil {
		panic(err)
	}
	sd.AddButton(9, anotherButton)

	cButton := buttons.NewColourButton(color.RGBA{255, 255, 0, 255})
	sd.AddButton(26, cButton)

	multiActionButton := buttons.NewColourButton(color.RGBA{255, 0, 255, 255})
	thisActionHandler := &actionhandlers.ChainedAction{}
	thisActionHandler.AddAction(&actionhandlers.TextPrintAction{Label: "Purple press"})
	thisActionHandler.AddAction(&actionhandlers.ColourChangeAction{NewColour: color.RGBA{255, 0, 0, 255}})
	multiActionButton.SetActionHandler(thisActionHandler)
	sd.AddButton(27, multiActionButton)

	decoratedButton := buttons.NewTextButton("ABC")
	sd.AddButton(19, decoratedButton)

	time.Sleep(2 * time.Second)
	decorator1 := decorators.NewBorder(10, color.RGBA{0, 255, 0, 255})
	sd.SetDecorator(19, decorator1)
	time.Sleep(2 * time.Second)
	decorator2 := decorators.NewBorder(5, color.RGBA{255, 0, 0, 255})
	sd.SetDecorator(19, decorator2)

	time.Sleep(2 * time.Second)

	myButton.SetText("Bye!")
	cButton.SetColour(color.RGBA{0, 255, 255, 255})
	sd.UnsetDecorator(19)
}
