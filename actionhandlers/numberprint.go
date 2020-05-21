package actionhandlers

import (
	"fmt"
	streamdeck "github.com/magicmonkey/go-streamdeck"
)

type NumberPrintAction struct {
	Number int
	btn    streamdeck.Button
}

func (npa *NumberPrintAction) SetButton(b streamdeck.Button) {
	npa.btn = b
}

func (npa *NumberPrintAction) Pressed() {
	fmt.Println(npa.Number)
}
