package main

import (
	"image/color"
	"strconv"
	"sync"

	streamdeck "github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/actionhandlers"
	"github.com/magicmonkey/go-streamdeck/buttons"
	"github.com/magicmonkey/go-streamdeck/decorators"
	_ "github.com/magicmonkey/go-streamdeck/devices"
)

func main() {
	// store the currently selected button
	var current int

	// initialise the streamdeck
	sd, err := streamdeck.New()
	if err != nil {
		panic(err)
	}

	// put a number onto each button
	btns := make([]*buttons.TextButton, 32)
	for i := 0; i < 32; i++ {
		btns[i] = buttons.NewTextButton(strconv.Itoa(i))
		sd.AddButton(i, btns[i])
	}

	// create some decorators for later use
	greenBorder := decorators.NewBorder(10, color.RGBA{0, 255, 0, 255})
	redBorder := decorators.NewBorder(5, color.RGBA{255, 0, 0, 255})

	// add red borders to all buttons
	for i := 0; i < 32; i++ {
		sd.SetDecorator(i, redBorder)
	}

	// add action handlers as an inline function
	for i := 0; i < 32; i++ {
		h := actionhandlers.NewCustomAction(func(btn streamdeck.Button) {
			sd.SetDecorator(current, redBorder)
			sd.SetDecorator(btn.GetButtonIndex(), greenBorder)
			current = btn.GetButtonIndex()
		})
		btns[i].SetActionHandler(h)
	}

	// don't end the program, keep running
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
