package main

import (
	"image/color"
	"strconv"
	"sync"

	streamdeck "github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/actionhandlers"
	"github.com/magicmonkey/go-streamdeck/buttons"
	"github.com/magicmonkey/go-streamdeck/decorators"
)

func main() {
	var current int

	sd, err := streamdeck.New()
	if err != nil {
		panic(err)
	}

	btns := make([]*buttons.TextButton, 32)
	for i := 0; i < 32; i++ {
		btns[i] = buttons.NewTextButton(strconv.Itoa(i))
		sd.AddButton(i, btns[i])
	}

	greenBorder := decorators.NewBorder(10, color.RGBA{0, 255, 0, 255})

	redBorder := decorators.NewBorder(5, color.RGBA{255, 0, 0, 255})
	for i := 0; i < 32; i++ {
		sd.SetDecorator(i, redBorder)
	}

	for i := 0; i < 32; i++ {
		h := actionhandlers.NewCustomAction(func(btn streamdeck.Button) {
			sd.SetDecorator(current, redBorder)
			sd.SetDecorator(btn.GetButtonIndex(), greenBorder)
			current = btn.GetButtonIndex()
		})
		btns[i].SetActionHandler(h)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()

}
