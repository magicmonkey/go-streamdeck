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

	myButton := streamdeck.NewTextButton("Hi world")
	myButton.SetActionHandler(&actionhandlers.TextPrintAction{Label: "You pressed me"})

	sd.AddButton(2, myButton)

	time.Sleep(1 * time.Second)

	myButton.SetText("Bye!")
}
