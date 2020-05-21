package actionhandlers

import streamdeck "github.com/magicmonkey/go-streamdeck"

type CustomAction struct {
	btn     streamdeck.Button
	handler func(streamdeck.Button)
}

func (action *CustomAction) SetButton(b streamdeck.Button) {
	action.btn = b
}

func (action *CustomAction) SetHandler(f func(streamdeck.Button)) {
	action.handler = f
}

func (action *CustomAction) Pressed() {
	action.handler(action.btn)
}
