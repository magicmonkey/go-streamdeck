package main

import (
	"image/color"
	"image/jpeg"
	"os"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

func main() {
	/*
		sd := streamdeck.Open()
		sd.ClearButtons()

		sd.SetBrightness(50)

		sd.WriteTextToButton(2, "Hi Lorna!", color.RGBA{0, 0, 0, 255}, color.RGBA{0, 255, 255, 255})
	*/
	//sd.WriteImageToButton("test.jpg", 9)
	//sd.WriteColorToButton(color.RGBA{255, 0, 255, 0}, 1)
	/*
		sd.ButtonPress(func(btnIndex int, sd *streamdeck.Device) {
			sd.ClearButtons()
			sd.WriteImageToButton("play.jpg", btnIndex)
		})
	*/

	//streamdeck.ExampleDevice_WriteColorToButton()
	img := streamdeck.GetImageWithText("Hello!", color.RGBA{255, 255, 255, 255}, color.RGBA{255, 0, 0, 100})
	//fmt.Println(img)
	newimg := streamdeck.ResizeAndRotate(img, 96, 96)
	f, _ := createOrOpenFile("test.jpg")
	jpeg.Encode(f, newimg, nil)
}

func createOrOpenFile(fname string) (*os.File, error) {
	os.Remove(fname)
	f, err := os.Create(fname)
	if err != nil {
		return nil, err
	}
	return f, nil
}
