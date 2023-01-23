package buttons

import (
	"image"
	"image/color"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

// TextButton represents a button with text on it
type TextButton struct {
	label            string
	textColour       color.Color
	backgroundColour color.Color
	updateHandler    func(streamdeck.Button)
	btnIndex         int
	actionHandler    streamdeck.ButtonActionHandler
}

// GetImageForButton is the interface implemention to get the button's image as an image.Image
func (btn *TextButton) GetImageForButton(btnSize int) image.Image {
	return streamdeck.GetImageWithTextExt(btn.label, btn.textColour, btn.backgroundColour, btnSize, streamdeck.ButtonOptions{MarginX: 3, MarginY: 3})
}

// SetButtonIndex is the interface implemention to set which button on the Streamdeck this is
func (btn *TextButton) SetButtonIndex(btnIndex int) {
	btn.btnIndex = btnIndex
}

// GetButtonIndex is the interface implemention to get which button on the Streamdeck this is
func (btn *TextButton) GetButtonIndex() int {
	return btn.btnIndex
}

// SetText allows the text on the button to be changed on the fly
func (btn *TextButton) SetText(label string) {
	btn.label = label
	btn.updateHandler(btn)
}

// SetTextColour allows the colour of the text on the button to be changed on the fly
func (btn *TextButton) SetTextColour(textColour color.Color) {
	btn.textColour = textColour
	btn.updateHandler(btn)
}

// SetBackgroundColor allows the background colour on the button to be changed on the fly
func (btn *TextButton) SetBackgroundColor(backgroundColour color.Color) {
	btn.backgroundColour = backgroundColour
	btn.updateHandler(btn)
}

// RegisterUpdateHandler is the interface implemention to let the engine give this button a callback to
// use to request that the button image is updated on the Streamdeck.
func (btn *TextButton) RegisterUpdateHandler(f func(streamdeck.Button)) {
	btn.updateHandler = f
}

// SetActionHandler allows a ButtonActionHandler implementation to be
// set on this button, so that something can happen when the button is pressed.
func (btn *TextButton) SetActionHandler(a streamdeck.ButtonActionHandler) {
	btn.actionHandler = a
}

// Pressed is the interface implementation for letting the engine notify that the button has been
// pressed.  This hands-off to the specified ButtonActionHandler if it has been set.
func (btn *TextButton) Pressed() {
	if btn.actionHandler != nil {
		btn.actionHandler.Pressed(btn)
	}
}

// NewTextButton creates a new TextButton with the specified text on it, in white on a black
// background.  The text will be set on a single line, and auto-sized to fill the button as best
// as possible.
func NewTextButton(label string) *TextButton {
	btn := NewTextButtonWithColours(label, color.White, color.Black)
	return btn
}

// NewTextButtonWithColours creates a new TextButton with the specified text on it, in the specified
// text and background colours.  The text will be set on a single line, and auto-sized to fill the
// button as best as possible.
func NewTextButtonWithColours(label string, textColour color.Color, backgroundColour color.Color) *TextButton {
	btn := &TextButton{label: label, textColour: textColour, backgroundColour: backgroundColour}
	return btn
}
