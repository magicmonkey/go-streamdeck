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

	// Button in position 2, changes to "Bye!" at the end of the program
	// When pressed, this prints "You pressed me" to the terminal
	myButton := buttons.NewTextButton("Hi world")
	myButton.SetActionHandler(&actionhandlers.TextPrintAction{Label: "You pressed me"})
	sd.AddButton(2, myButton)

	// Button in position 3, prints "5" to the terminal when pressed
	myOtherButton := buttons.NewTextButton("4")
	myOtherButton.SetActionHandler(&actionhandlers.NumberPrintAction{Number: 5})
	sd.AddButton(3, myOtherButton)

	// Button in position 7 (top right on Streamdeck XL), says 7
	// When pressed, changes to display "8"
	myNextButton := buttons.NewTextButton("7")
	myNextButton.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "8"})
	sd.AddButton(7, myNextButton)

	// Image button, no action handler
	anotherButton, err := buttons.NewImageFileButton("examples/test/play.jpg")
	if err != nil {
		panic(err)
	}
	sd.AddButton(9, anotherButton)

	// Yellow button, no action handler but it goes to blue at the end of the program
	cButton := buttons.NewColourButton(color.RGBA{255, 255, 0, 255})
	sd.AddButton(26, cButton)

	// One button, two actions (uses ChainedAction)
	// Purple button, prints to the console and turns red when pressed
	multiActionButton := buttons.NewColourButton(color.RGBA{255, 0, 255, 255})
	thisActionHandler := &actionhandlers.ChainedAction{}
	thisActionHandler.AddAction(&actionhandlers.TextPrintAction{Label: "Purple press"})
	thisActionHandler.AddAction(&actionhandlers.ColourChangeAction{NewColour: color.RGBA{255, 0, 0, 255}})
	multiActionButton.SetActionHandler(thisActionHandler)
	sd.AddButton(27, multiActionButton)

	// Text button, gets a red highlight after 2 seconds, then a green
	// highlight after another 2 seconds
	decoratedButton := buttons.NewTextButton("ABC")
	sd.AddButton(19, decoratedButton)
	time.Sleep(2 * time.Second)
	decorator1 := decorators.NewBorder(10, color.RGBA{0, 255, 0, 255})
	sd.SetDecorator(19, decorator1)
	time.Sleep(2 * time.Second)
	decorator2 := decorators.NewBorder(5, color.RGBA{255, 0, 0, 255})
	sd.SetDecorator(19, decorator2)
	time.Sleep(2 * time.Second)

	// When this button says "Bye!", the program ends
	myButton.SetText("Bye!")
	// When this button goes blue, the program ends
	cButton.SetColour(color.RGBA{0, 255, 255, 255})
	sd.UnsetDecorator(19)
}
