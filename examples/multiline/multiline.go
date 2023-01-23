package main

import (
	"image/color"

	streamdeck "github.com/magicmonkey/go-streamdeck"
	_ "github.com/magicmonkey/go-streamdeck/devices"
	"golang.org/x/image/colornames"
)

func main() {
	sd, err := streamdeck.Open()
	if err != nil {
		panic(err)
	}
	sd.ClearButtons()
	sd.SetBrightness(100)

	// demonstrate multiline capabilities
	sd.WriteTextToButtonExt(0, "WORD\nUP", colornames.Goldenrod, colornames.Maroon, streamdeck.ButtonOptions{})
	sd.WriteTextToButtonExt(1, "LINE 1\nLINE 2", colornames.Goldenrod, colornames.Maroon, streamdeck.ButtonOptions{MarginX: 5, MarginY: 5})
	sd.WriteTextToButtonExt(2, "Magicical\nMystery\nTour", colornames.Goldenrod, colornames.Maroon, streamdeck.ButtonOptions{MarginX: 10, MarginY: 10})
	sd.WriteTextToButtonExt(3, "this\ntoo\nshall\npass", colornames.Black, colornames.Dodgerblue, streamdeck.ButtonOptions{MaxTextSize: 20, MarginX: 0, MarginY: 0})
	sd.WriteTextToButtonExt(4, "www\njjjj\n", colornames.Yellow, colornames.Blueviolet, streamdeck.ButtonOptions{MarginX: 10, MarginY: 10})
	sd.WriteTextToButtonExt(5, "ON\nOFF", colornames.Red, colornames.White, streamdeck.ButtonOptions{MarginX: 10, MarginY: 10})
	sd.WriteTextToButtonExt(6, "AJj", color.RGBA{0, 0, 0, 255}, color.RGBA{0, 255, 255, 255}, streamdeck.ButtonOptions{MarginX: 5, MarginY: 5, MaxTextSize: 40})
	sd.WriteTextToButtonExt(7, "A\nB", color.RGBA{0, 0, 0, 150}, color.RGBA{255, 255, 255, 255}, streamdeck.ButtonOptions{MarginX: 5, MarginY: 5})

	sd.WriteTextToButtonExt(8, "READY\nLOAD", colornames.Dodgerblue, colornames.Darkblue, streamdeck.ButtonOptions{MarginX: 8, MarginY: 8})
	sd.WriteTextToButtonExt(9, "READY\nLOAD", colornames.Dodgerblue, colornames.Darkblue, streamdeck.ButtonOptions{MarginX: 10, MarginY: 10, AlignX: -1, AlignY: -1})
	sd.WriteTextToButtonExt(10, "READY\nLOAD", colornames.Dodgerblue, colornames.Darkblue, streamdeck.ButtonOptions{MarginX: 15, MarginY: 15, AlignX: 1, AlignY: 1})
	sd.WriteTextToButtonExt(11, "1\n2\n3\n4\n5", colornames.Dodgerblue, colornames.Darkblue, streamdeck.ButtonOptions{MarginX: 5, MarginY: 0, AlignX: 0, AlignY: 1})
	sd.WriteTextToButtonExt(12, "Office\nMorn", colornames.White, colornames.Darkcyan, streamdeck.ButtonOptions{MaxTextSize: 20, MarginX: 10, MarginY: 10, AlignY: -1})
	sd.WriteTextToButtonExt(13, "Office\nNoon", colornames.Black, colornames.Yellow, streamdeck.ButtonOptions{MaxTextSize: 20, MarginX: 10, MarginY: 10})
	sd.WriteTextToButtonExt(14, "Office\nNight", colornames.White, colornames.Darkblue, streamdeck.ButtonOptions{MaxTextSize: 20, MarginX: 10, MarginY: 10, AlignY: 1})
	sd.WriteTextToButtonExt(15, "OFF", colornames.Grey, colornames.Lightgrey, streamdeck.ButtonOptions{MarginX: 10, MarginY: 10})

	// demonstrate margin usage
	sd.WriteTextToButtonExt(16, "Margins", colornames.Lightsalmon, colornames.Maroon, streamdeck.ButtonOptions{})
	sd.WriteTextToButtonExt(17, "Margins", colornames.Lightsalmon, colornames.Maroon, streamdeck.ButtonOptions{MarginX: 10, MarginY: 10})
	sd.WriteTextToButtonExt(18, "Margins", colornames.Lightsalmon, colornames.Maroon, streamdeck.ButtonOptions{MarginX: 15, MarginY: 15})
	sd.WriteTextToButtonExt(19, "Margins", colornames.Lightsalmon, colornames.Maroon, streamdeck.ButtonOptions{MarginX: 20, MarginY: 20})

	// demonstrate alignment usage
	sd.WriteTextToButtonExt(20, "Justify", colornames.Sandybrown, colornames.Saddlebrown, streamdeck.ButtonOptions{MaxTextSize: 20, MarginX: 5, MarginY: 5, AlignX: -1, AlignY: -1})
	sd.WriteTextToButtonExt(21, "Justify", colornames.Sandybrown, colornames.Saddlebrown, streamdeck.ButtonOptions{MaxTextSize: 20, MarginX: 5, MarginY: 5, AlignX: 1, AlignY: 0})
	sd.WriteTextToButtonExt(22, "Justify", colornames.Sandybrown, colornames.Saddlebrown, streamdeck.ButtonOptions{MaxTextSize: 20, MarginX: 5, MarginY: 5, AlignX: 0, AlignY: 1})
	sd.WriteTextToButtonExt(23, "Justify", colornames.Sandybrown, colornames.Saddlebrown, streamdeck.ButtonOptions{MaxTextSize: 20, MarginX: 5, MarginY: 5, AlignX: 1, AlignY: 1})

	sd.WriteTextToButtonExt(24, "Chicken\nNoodle\nSoup", colornames.Black, colornames.Dodgerblue, streamdeck.ButtonOptions{MarginX: 10, MarginY: 5, AlignX: -1, AlignY: 1})
	sd.WriteTextToButtonExt(25, "Chicken\nNoodle\nSoup", colornames.Black, colornames.Dodgerblue, streamdeck.ButtonOptions{MarginX: 20, MarginY: 20, AlignX: -1, AlignY: 0})
	sd.WriteTextToButtonExt(26, "Chicken\nNoodle\nSoup", colornames.Black, colornames.Dodgerblue, streamdeck.ButtonOptions{MarginX: 10, MarginY: 5, AlignX: 1, AlignY: 1})
	sd.WriteTextToButtonExt(27, "Chicken\nNoodle\nSoup", colornames.Black, colornames.Dodgerblue, streamdeck.ButtonOptions{MarginX: 10, MarginY: 5, AlignX: 0, AlignY: -1})
	sd.WriteTextToButtonExt(28, "Mac\n&\nCheese", colornames.Black, colornames.Darkgoldenrod, streamdeck.ButtonOptions{MarginX: 10, MarginY: 10})
	sd.WriteTextToButtonExt(29, "Mac\n&\nCheese", colornames.Black, colornames.Darkgoldenrod, streamdeck.ButtonOptions{MarginX: 10, MarginY: 10, AlignX: -1, AlignY: -1})
	sd.WriteTextToButtonExt(30, "Mac & Cheese", colornames.Black, colornames.Darkgoldenrod, streamdeck.ButtonOptions{MarginX: 10, MarginY: 10, AlignX: 0, AlignY: 1})
	sd.WriteTextToButtonExt(31, "Mac\n&\nCheese", colornames.Black, colornames.Darkgoldenrod, streamdeck.ButtonOptions{MarginX: 5, MarginY: 5, AlignX: 1, AlignY: 1})
}
