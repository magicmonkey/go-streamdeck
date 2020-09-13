package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"sync"

	"github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/actionhandlers"
	"github.com/magicmonkey/go-streamdeck/buttons"
	// needed to get the device definitions
	_ "github.com/magicmonkey/go-streamdeck/devices"
)

func main() {
	// add a streamdeck
	sd, err := streamdeck.New()
	if err != nil {
		panic(err)
	}

	// create a text button labelled "Pic"
	myButton := buttons.NewTextButton("Pic")

	// create a custom action with a function as the action handler
	shotaction := &actionhandlers.CustomAction{}
	shotaction.SetHandler(func(btn streamdeck.Button) {
		// a goroutine so that we don't wait for the command to return
		go takeScreenshot()
	})

	// attach action to button
	myButton.SetActionHandler(shotaction)
	// put button in top left slot
	sd.AddButton(0, myButton)

	// now run and keep running
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

func takeScreenshot() {
	fmt.Println("Taking screenshot with delay...")
	cmd := exec.Command("/usr/bin/gnome-screenshot", "-w", "-d", "2")
	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)
	slurp2, _ := ioutil.ReadAll(stdout)
	fmt.Printf("%s\n", slurp2)

	fmt.Println("Taken screenshot")
}
