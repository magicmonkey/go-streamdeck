package actionhandlers

import streamdeck "github.com/magicmonkey/go-streamdeck"

type ChainedAction struct {
	actions []streamdeck.ButtonActionHandler
}

func (act *ChainedAction) AddAction(newaction streamdeck.ButtonActionHandler) {
	act.actions = append(act.actions, newaction)
}

func (act *ChainedAction) Pressed(btn streamdeck.Button) {
	for _, a := range act.actions {
		a.Pressed(btn)
	}
}
