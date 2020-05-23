package actionhandlers

import streamdeck "github.com/magicmonkey/go-streamdeck"

type CustomAction struct {
	handler func(streamdeck.Button)
}

func (action *CustomAction) SetHandler(f func(streamdeck.Button)) {
	action.handler = f
}

func (action *CustomAction) Pressed(btn streamdeck.Button) {
	action.handler(btn)
}
