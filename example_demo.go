package streamdeck

func ExampleDevice_WriteColorToButton() {
	sd := Open()
	sd.ClearButtons()
	sd.SetBrightness(50)
	for i := 0; i < 8; i++ {
		sd.WriteColorToButton(255, 255, 0, i)
	}
	for i := 8; i < 16; i++ {
		sd.WriteColorToButton(255, 0, 255, i)
	}
	for i := 16; i < 24; i++ {
		sd.WriteColorToButton(0, 0, 255, i)
	}
	for i := 24; i < 32; i++ {
		sd.WriteColorToButton(0, 255, 0, i)
	}
}

func ExampleDevice_ButtonPress() {
	sd := Open()
	sd.ClearButtons()
	sd.ButtonPress(func(btnIndex int, sd *Device) {
		sd.ClearButtons()
		sd.WriteImageToButton("play.jpg", btnIndex)
	})

}
