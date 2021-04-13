package main

import (
	"time"

	streamdeck "github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/buttons"
	_ "github.com/magicmonkey/go-streamdeck/devices"
)

func main() {
	// initialise the device
	sd, err := streamdeck.New()
	if err != nil {
		panic(err)
	}

	sd.SetBrightness(40)

	// create buttons
	btn1 := buttons.NewTextButton("Brightness")
	sd.AddButton(1, btn1)
	btn2 := buttons.NewTextButton("40")
	sd.AddButton(2, btn2)

	// wait for one second
	time.Sleep(1 * time.Second)

	// set brightness
	sd.SetBrightness(100)
	btn2.SetText("100")

	// program exits
}
