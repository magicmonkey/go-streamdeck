package buttons

import (
	"image"
	"image/color"
	"image/draw"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

// ColourButton represents a button which is a solid block of a single colour
type ColourButton struct {
	colour        color.Color
	updateHandler func(streamdeck.Button)
	btnIndex      int
	actionHandler streamdeck.ButtonActionHandler
}

// GetImageForButton is the interface implemention to get the button's image as an image.Image
func (btn *ColourButton) GetImageForButton(btnSize int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, btnSize, btnSize))
	//colour := color.RGBA{red, green, blue, 0}
	draw.Draw(img, img.Bounds(), image.NewUniform(btn.colour), image.Point{0, 0}, draw.Src)
	return img
}

// SetButtonIndex is the interface implemention to set which button on the Streamdeck this is
func (btn *ColourButton) SetButtonIndex(btnIndex int) {
	btn.btnIndex = btnIndex
}

// GetButtonIndex is the interface implemention to get which button on the Streamdeck this is
func (btn *ColourButton) GetButtonIndex() int {
	return btn.btnIndex
}

// SetColour allows the colour for the button to be changed on the fly
func (btn *ColourButton) SetColour(colour color.Color) {
	btn.colour = colour
	btn.updateHandler(btn)
}

// RegisterUpdateHandler is the interface implemention to let the engine give this button a callback to
// use to request that the button image is updated on the Streamdeck.
func (btn *ColourButton) RegisterUpdateHandler(f func(streamdeck.Button)) {
	btn.updateHandler = f
}

// SetActionHandler allows a ButtonActionHandler implementation to be
// set on this button, so that something can happen when the button is pressed.
func (btn *ColourButton) SetActionHandler(a streamdeck.ButtonActionHandler) {
	btn.actionHandler = a
}

// Pressed is the interface implementation for letting the engine notify that the button has been
// pressed.  This hands-off to the specified ButtonActionHandler if it has been set.
func (btn *ColourButton) Pressed() {
	if btn.actionHandler != nil {
		btn.actionHandler.Pressed(btn)
	}
}

// NewColourButton creates a new ColourButton of the specified colour
func NewColourButton(colour color.Color) *ColourButton {
	btn := &ColourButton{colour: colour}
	return btn
}
