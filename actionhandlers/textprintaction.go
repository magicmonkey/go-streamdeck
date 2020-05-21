package actionhandlers

import (
	"fmt"
	streamdeck "github.com/magicmonkey/go-streamdeck"
)

type TextPrintAction struct {
	btn streamdeck.Button
}

func (tpa *TextPrintAction) SetButton(b streamdeck.Button) {
	tpa.btn = b
}

func (tpa *TextPrintAction) Pressed() {
	fmt.Println("I was pressed")
	fmt.Println("I am:")
	fmt.Println(tpa.btn)
}
