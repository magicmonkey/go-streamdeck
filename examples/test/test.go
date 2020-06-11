package main

import (
	"fmt"
	"image/color"
	"sync"

	streamdeck "github.com/magicmonkey/go-streamdeck"
	_ "github.com/magicmonkey/go-streamdeck/devices"
)

func main() {
	sd, err := streamdeck.Open()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found device [%s]\n", sd.GetName())

	sd.ClearButtons()

	sd.SetBrightness(50)

	sd.WriteTextToButton(2, "Hi!", color.RGBA{0, 0, 0, 255}, color.RGBA{0, 255, 255, 255})
	sd.WriteTextToButton(3, "Hi again!", color.RGBA{0, 0, 0, 255}, color.RGBA{0, 255, 255, 255})
	sd.WriteTextToButton(4, "Hi again again!", color.RGBA{0, 0, 0, 255}, color.RGBA{0, 255, 255, 255})

	//sd.WriteImageToButton(9, "something.png")
	//sd.WriteImageToButton(10, "play.jpg")
	//sd.WriteColorToButton(color.RGBA{255, 0, 255, 0}, 1)
	sd.ButtonPress(func(btnIndex int, sd *streamdeck.Device, err error) {
		if err != nil {
			panic(err)
		}
		sd.ClearButtons()
		sd.WriteImageToButton(btnIndex, "examples/test/play.jpg")
	})

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
