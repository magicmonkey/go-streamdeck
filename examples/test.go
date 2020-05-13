package main

import streamdeck "github.com/magicmonkey/go-streamdeck"

func main() {
	sd := streamdeck.Open()
	sd.ClearButtons()

	sd.SetBrightness(50)

	sd.ButtonPress(func(btnIndex int, sd *streamdeck.Device) {
		sd.ClearButtons()
		sd.WriteImageToButton("play.jpg", btnIndex)
	})
}
