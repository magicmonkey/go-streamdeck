# Go Streamdeck

A Go interface to an Elgato Streamdeck (currently works with the 32-button XL only because that's what I have).

[![GoDoc](https://godoc.org/github.com/magicmonkey/go-streamdeck?status.svg)](https://godoc.org/github.com/magicmonkey/go-streamdeck)

_Designed for and tested with Ubuntu, Go 1.13+ and a Streamdeck XL. Images are the wrong size for other streamdecks; bug reports and patches are welcome!_

- [Installation](#installation)
- [Usage](#usage)
  * [Example high-level usage](#example-high-level-usage)
  * [Example low-level usage](#example-low-level-usage)
- [Showcase](#showcase)
- [Contributions](#contributions)

## Installation

Either include the library in your project or install it with the following command:

```
go get github.com/magicmonkey/go-streamdeck
```

## Usage

There are 2 ways to use this: the low-level "comms-oriented" interface (using `streamdeck.Open`) which wraps the USB HID protocol, or the higher-level "button-oriented" interface (using `streamdeck.New`) which represents buttons and actions.

If you want to implement your own actions, I suggest that you either instantiate a `CustomAction` or alternatively implement the `ButtonActionHandler` interface (basing your code on the `CustomAction`).

### Example high-level usage

High level usage gives some helpers to set up buttons. This example has a few things to look at:

* A button in position 2 that says "Hi world" and prints to the console when pressed

* A button in position 7 displaying the number 7 - changes to number 8 when pressed.

* A yellow button in position 26

* A purple button in position 27, it changes colour _and_ prints to the console when pressed.

```go
import (
	"image/color"
	"time"

	streamdeck "github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/actionhandlers"
	"github.com/magicmonkey/go-streamdeck/buttons"
	_ "github.com/magicmonkey/go-streamdeck/devices"
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
	myNextButton.SetActionHandler(&actionhandlers.TextLabelChangeAction{NewLabel: "8"})
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

The program runs for 20 seconds and then exits.

### Example low-level usage

The low-level usage gives more control over the operations of the streamdeck and buttons.

This example shows an image on any pressed button, updating each time another button is pressed.

```go
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

	time.Sleep(20 * time.Second)

}
```

The program runs for 20 seconds and then exits.

## Showcase

Projects using this library (pull request to add yours!)

* [Streamdeck tricks](https://github.com/lornajane/streamdeck-tricks)

## Contributions

This is a very new project but all feedback, comments, questions and patches are more than welcome. Please get in touch by opening an issue, it would be good to hear who is using the project and how things are going.

For more, see [CONTRIBUTING.md](CONTRIBUTING.md).
