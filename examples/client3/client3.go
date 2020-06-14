package main

import (
	"image/color"
	"time"
	streamdeck "github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/buttons"
	"github.com/magicmonkey/go-streamdeck/decorators"
	_ "github.com/magicmonkey/go-streamdeck/devices"
)

func main() {

	sd, err := streamdeck.New()
	if err != nil {
		panic(err)
	}

	btn1 := buttons.NewTextButton("Hello")
	sd.AddButton(1, btn1)
	btn2, _ := buttons.NewImageFileButton("examples/test/play.jpg")
	sd.AddButton(2, btn2)


	greenBorder := decorators.NewBorder(10, color.RGBA{0, 255, 0, 255})
	sd.SetDecorator(1, greenBorder)
	sd.SetDecorator(2, greenBorder)

	time.Sleep(1 * time.Second)

	sd.UnsetDecorator(1)
	sd.UnsetDecorator(2)
}