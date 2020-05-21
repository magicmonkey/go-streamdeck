package main

import (
	"time"

	streamdeck "github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/actionhandlers"
)

func main() {
	sd, err := streamdeck.New()
	if err != nil {
		panic(err)
	}

	b1 := streamdeck.NewTextButton("Hi world")
	b1.SetActionHandler(&actionhandlers.TextPrintAction{})

	sd.AddButton(2, b1)

	time.Sleep(1 * time.Second)

	b1.SetText("Bye!")
}
