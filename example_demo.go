package streamdeck

import "image/color"

func ExampleDevice_WriteColorToButton() {
	sd := Open()
	sd.ClearButtons()
	sd.SetBrightness(50)
	for i := 0; i < 8; i++ {
		sd.WriteColorToButton(i, color.RGBA{0, 0, 0, 0})
	}
	for i := 8; i < 16; i++ {
		sd.WriteColorToButton(i, color.RGBA{255, 0, 255, 0})
	}
	for i := 16; i < 24; i++ {
		sd.WriteColorToButton(i, color.RGBA{0, 0, 255, 0})
	}
	for i := 24; i < 32; i++ {
		sd.WriteColorToButton(i, color.RGBA{0, 255, 0, 0})
	}
}

func ExampleDevice_ButtonPress() {
	sd := Open()
	sd.ClearButtons()
	sd.ButtonPress(func(btnIndex int, sd *Device, err error) {
		sd.ClearButtons()
		sd.WriteImageToButton(btnIndex, "play.jpg")
	})

}
