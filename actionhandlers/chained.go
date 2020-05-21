package actionhandlers

import streamdeck "github.com/magicmonkey/go-streamdeck"

type ChainedAction struct {
	actions []streamdeck.ButtonActionHandler
	btn     streamdeck.Button
}

func (act *ChainedAction) SetButton(b streamdeck.Button) {
	act.btn = b
	for _, a := range act.actions {
		a.SetButton(b)
	}
}

func (act *ChainedAction) AddAction(newaction streamdeck.ButtonActionHandler) {
	act.actions = append(act.actions, newaction)
}

func (act *ChainedAction) Pressed() {
	for _, a := range act.actions {
		a.Pressed()
	}
}
