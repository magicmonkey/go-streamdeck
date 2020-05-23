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

func NewEmptyCustomAction() *CustomAction {
	return &CustomAction{}
}

func NewCustomAction(handler func(streamdeck.Button)) *CustomAction {
	return &CustomAction{handler: handler}
}
