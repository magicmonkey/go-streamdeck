A Go interface to an Elgato StreamdeckXL (specifically the 32-button version)

[![GoDoc](https://godoc.org/github.com/magicmonkey/go-streamdeck?status.svg)](https://godoc.org/github.com/magicmonkey/go-streamdeck)

There are 2 ways to use this: the low-level "comms-oriented" interface (using `streamdeck.Open`) which wraps the USB HID protocol, or the higher-level "button-oriented" interface (using `streamdeck.New`) which represents buttons and actions.

Example high-level usage:
```
import (
	streamdeck "github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/actionhandlers"
	"github.com/magicmonkey/go-streamdeck/buttons"
)

func main() {
	sd, err := streamdeck.New()
	if err != nil {
		panic(err)
	}

	// A simple yellow button in position 26
	cButton := buttons.NewColourButton(color.RGBA{255, 255, 0, 255})
	sd.AddButton(26, cButton)

	// A button with text on it in position 2, which echoes to the console when presesd
	myButton := buttons.NewTextButton("Hi world")
	myButton.SetActionHandler(&actionhandlers.TextPrintAction{Label: "You pressed me"})
	sd.AddButton(2, myButton)

	// A button with text on it which changes when pressed
	myNextButton := buttons.NewTextButton("7")
	myNextButton.SetActionHandler(&actionhandlers.TextLabelChange{NewLabel: "8"})
	sd.AddButton(7, myNextButton)

	// A button which performs multiple actions when pressed
	multiActionButton := buttons.NewColourButton(color.RGBA{255, 0, 255, 255})
	thisActionHandler := &actionhandlers.ChainedAction{}
	thisActionHandler.AddAction(&actionhandlers.TextPrintAction{Label: "Purple press"})
	thisActionHandler.AddAction(&actionhandlers.ColourChangeAction{NewColour: color.RGBA{255, 0, 0, 255}})
	multiActionButton.SetActionHandler(thisActionHandler)
	sd.AddButton(27, multiActionButton)

	time.Sleep(20 * time.Second)
}
```

Example low-level usage:
```
import streamdeck "github.com/magicmonkey/go-streamdeck"

func main() {
	sd, err := streamdeck.Open()
	if err != nil {
		panic(err)
	}
	sd.ClearButtons()

	sd.SetBrightness(50)

	sd.ButtonPress(func(btnIndex int, sd *streamdeck.Device, err error) {
		if err != nil {
			panic(err)
		}
		sd.ClearButtons()
		sd.WriteImageToButton("play.jpg", btnIndex)
	})
}
```
