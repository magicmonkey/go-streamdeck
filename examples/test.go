package main

import streamdeck "github.com/magicmonkey/go-streamdeck"

func main() {
	sd := streamdeck.Open()
	sd.ClearButtons()

	sd.ButtonPress(func(btnIndex int, sd *streamdeck.Device) {
		sd.WriteImageToButton("play.jpg", btnIndex)
	})
}
