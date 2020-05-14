A Go interface to an Elgato StreamdeckXL (specifically the 32-button version)

[![GoDoc](https://godoc.org/github.com/magicmonkey/go-streamdeck?status.svg)](https://godoc.org/github.com/magicmonkey/go-streamdeck)

Example usage:
```
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
```
