package streamdeck

import "image/color"

type ButtonDisplay interface {
	GetButtonType() string
	GetImageFile() string
	GetText() string
	GetButtonIndex() int
	SetButtonIndex(int)
	RegisterUpdateHandler(func(Button))
	Pressed()
}

type ButtonActionHandler interface {
	SetButton(Button)
	Pressed()
}

type Button interface {
	ButtonDisplay
}

type StreamDeck struct {
	dev     *Device
	buttons map[int]Button
}

func New() (*StreamDeck, error) {
	sd := &StreamDeck{}
	d, err := Open()
	if err != nil {
		return nil, err
	}
	sd.dev = d
	sd.buttons = make(map[int]Button)
	sd.dev.ButtonPress(sd.pressHandler)
	return sd, nil
}

func (sd *StreamDeck) AddButton(btnIndex int, b Button) {
	b.RegisterUpdateHandler(sd.ButtonUpdateHandler)
	b.SetButtonIndex(btnIndex)
	sd.buttons[btnIndex] = b
	sd.updateButton(b)
}

func (sd *StreamDeck) ButtonUpdateHandler(b Button) {
	sd.buttons[b.GetButtonIndex()] = b
	sd.updateButton(b)
}

func (sd *StreamDeck) pressHandler(btnIndex int, d *Device, err error) {
	if err != nil {
		panic(err)
	}
	b := sd.buttons[btnIndex]
	if b != nil {
		sd.buttons[btnIndex].Pressed()
	}
}

func (sd *StreamDeck) updateButton(b Button) {
	switch b.GetButtonType() {
	case "text":
		sd.dev.WriteTextToButton(b.GetButtonIndex(), b.GetText(), color.White, color.Black)
	case "imagefile":
		sd.dev.WriteImageToButton(b.GetButtonIndex(), b.GetImageFile())
	}
}
