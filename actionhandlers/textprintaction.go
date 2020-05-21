package actionhandlers

import (
	"fmt"
	streamdeck "github.com/magicmonkey/go-streamdeck"
)

type TextPrintAction struct {
	Label string
	btn   streamdeck.Button
}

func (tpa *TextPrintAction) SetButton(b streamdeck.Button) {
	tpa.btn = b
}

func (tpa *TextPrintAction) Pressed() {
	fmt.Println(tpa.Label)
	fmt.Print("The button is: ")
	fmt.Println(tpa.btn)
}
