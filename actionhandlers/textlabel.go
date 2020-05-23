package actionhandlers

import (
	streamdeck "github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/buttons"
)

type TextLabelChange struct {
	NewLabel string
}

func (action *TextLabelChange) Pressed(btn streamdeck.Button) {
	mybtn := btn.(*buttons.TextButton)
	mybtn.SetText(action.NewLabel)
}
