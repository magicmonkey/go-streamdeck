package buttons

import (
	"image"
	"image/color"
	"image/draw"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

type ColourButton struct {
	colour        color.Color
	updateHandler func(streamdeck.Button)
	btnIndex      int
	actionHandler streamdeck.ButtonActionHandler
}

func (btn *ColourButton) GetImageForButton() image.Image {
	ButtonSize := 96
	img := image.NewRGBA(image.Rect(0, 0, ButtonSize, ButtonSize))
	//colour := color.RGBA{red, green, blue, 0}
	draw.Draw(img, img.Bounds(), image.NewUniform(btn.colour), image.Point{0, 0}, draw.Src)
	return img
}

func (btn *ColourButton) SetButtonIndex(btnIndex int) {
	btn.btnIndex = btnIndex
}

func (btn *ColourButton) GetButtonIndex() int {
	return btn.btnIndex
}

func (btn *ColourButton) SetColour(colour color.Color) {
	btn.colour = colour
	btn.updateHandler(btn)
}

func (btn *ColourButton) RegisterUpdateHandler(f func(streamdeck.Button)) {
	btn.updateHandler = f
}

func (btn *ColourButton) SetActionHandler(a streamdeck.ButtonActionHandler) {
	btn.actionHandler = a
}

func (btn *ColourButton) Pressed() {
	if btn.actionHandler != nil {
		btn.actionHandler.Pressed(btn)
	}
}

func NewColourButton(colour color.Color) *ColourButton {
	btn := &ColourButton{colour: colour}
	return btn
}
