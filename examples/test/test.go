package main

import (
	"fmt"
	"image/color"
	"sync"

	streamdeck "github.com/magicmonkey/go-streamdeck"
	_ "github.com/magicmonkey/go-streamdeck/devices"
)

func main() {
	// connect to a streamdeck
	sd, err := streamdeck.Open()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found device [%s]\n", sd.GetName())

	// remove all button content
	sd.ClearButtons()

	// manage brightness level
	sd.SetBrightness(50)

	// show text in increasing lengths on three buttons
	sd.WriteTextToButton(2, "Hi!", color.RGBA{0, 0, 0, 255}, color.RGBA{0, 255, 255, 255})
	sd.WriteTextToButton(3, "Hi again!", color.RGBA{0, 0, 0, 255}, color.RGBA{0, 255, 255, 255})
	sd.WriteTextToButton(4, "Hi again again!", color.RGBA{0, 0, 0, 255}, color.RGBA{0, 255, 255, 255})

	// when any button is pressed, clear all buttons, and set an image on the pressed button
	sd.ButtonPress(func(btnIndex int, sd *streamdeck.Device, err error) {
		if err != nil {
			panic(err)
		}
		sd.ClearButtons()
		sd.WriteImageToButton(btnIndex, "play.jpg")
	})

	// keep the program running
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
