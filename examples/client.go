package main

import (
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
	myNextButton.SetActionHandler(&actionhandlers.TextLabelChange{NewLabel: "8"})
	sd.AddButton(7, myNextButton)

	time.Sleep(2 * time.Second)

	myButton.SetText("Bye!")
}
