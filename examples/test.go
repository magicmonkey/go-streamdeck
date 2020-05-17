package main

import (
	"image/color"
	"os"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

func main() {
	sd := streamdeck.Open()
	sd.ClearButtons()

	sd.SetBrightness(50)

	sd.WriteTextToButton(2, "Hi!", color.RGBA{0, 0, 0, 255}, color.RGBA{0, 255, 255, 255})
	sd.WriteTextToButton(3, "Hi again!", color.RGBA{0, 0, 0, 255}, color.RGBA{0, 255, 255, 255})
	sd.WriteTextToButton(4, "Hi again again!", color.RGBA{0, 0, 0, 255}, color.RGBA{0, 255, 255, 255})

	//sd.WriteImageToButton(9, "something.png")
	//sd.WriteImageToButton(10, "play.jpg")
	//sd.WriteColorToButton(color.RGBA{255, 0, 255, 0}, 1)
	/*
		sd.ButtonPress(func(btnIndex int, sd *streamdeck.Device) {
			sd.ClearButtons()
			sd.WriteImageToButton("play.jpg", btnIndex)
		})
	*/

	//streamdeck.ExampleDevice_WriteColorToButton()
	/*
		img := streamdeck.GetImageWithText("Hello again!", color.RGBA{255, 255, 255, 255}, color.RGBA{255, 0, 0, 100}, 18)
		newimg := streamdeck.ResizeAndRotate(img, 96, 96)
		f, _ := createOrOpenFile("test.jpg")
		jpeg.Encode(f, newimg, nil)
	*/
}

func createOrOpenFile(fname string) (*os.File, error) {
	os.Remove(fname)
	f, err := os.Create(fname)
	if err != nil {
		return nil, err
	}
	return f, nil
}
