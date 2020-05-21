package actionhandlers

import (
	"image/color"

	streamdeck "github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/buttons"
)

type ColourChangeAction struct {
	NewColour color.Color
	btn       *buttons.ColourButton
}

func (action *ColourChangeAction) SetButton(b streamdeck.Button) {
	action.btn = b.(*buttons.ColourButton)
}

func (action *ColourChangeAction) Pressed() {
	action.btn.SetColour(action.NewColour)
}
