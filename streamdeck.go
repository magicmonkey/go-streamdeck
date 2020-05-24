package streamdeck

import "image"

// ButtonDisplay is the interface to satisfy for displaying on a button
type ButtonDisplay interface {
	GetImageForButton() image.Image
	GetButtonIndex() int
	SetButtonIndex(int)
	RegisterUpdateHandler(func(Button))
	Pressed()
}

// ButtonActionHandler is the interface to satisfy for handling a button being pressed, generally via an `actionhandler`
type ButtonActionHandler interface {
	Pressed(Button)
}

// Button is the interface to satisfy for being a button; currently this is a direct proxy for the `ButtonDisplay` interface as there isn't a requirement to handle being pressed
type Button interface {
	ButtonDisplay
}

// StreamDeck is the main struct to represent a StreamDeck device, and internally contains the reference to a `Device`
type StreamDeck struct {
	dev     *Device
	buttons map[int]Button
}

// New will return a new instance of a `StreamDeck`, and is the main entry point for the higher-level interface.  It will return an error if there is no StreamDeck plugged in.
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

// AddButton adds a `Button` object to the StreamDeck at the specified index
func (sd *StreamDeck) AddButton(btnIndex int, b Button) {
	b.RegisterUpdateHandler(sd.ButtonUpdateHandler)
	b.SetButtonIndex(btnIndex)
	sd.buttons[btnIndex] = b
	sd.updateButton(b)
}

// ButtonUpdateHandler allows a user of this library to signal when something external has changed, such that this button should be update
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

func (sd *StreamDeck) updateButton(b Button) error {
	img := b.GetImageForButton()
	e := sd.dev.WriteRawImageToButton(b.GetButtonIndex(), img)
	return e
}
