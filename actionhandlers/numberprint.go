package actionhandlers

import (
	"fmt"
	streamdeck "github.com/magicmonkey/go-streamdeck"
)

type NumberPrintAction struct {
	Number int
}

func (npa *NumberPrintAction) Pressed(btn streamdeck.Button) {
	fmt.Println(npa.Number)
}

func NewNumberPrintAction(number int) *NumberPrintAction {
	return &NumberPrintAction{Number: number}
}
