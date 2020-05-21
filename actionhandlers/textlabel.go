package actionhandlers

import (
	streamdeck "github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/buttons"
)

type TextLabelChange struct {
	NewLabel string
	btn      *buttons.TextButton
}

func (action *TextLabelChange) SetButton(b streamdeck.Button) {
	action.btn = b.(*buttons.TextButton)
}

func (action *TextLabelChange) Pressed() {
	action.btn.SetText("8")
}
