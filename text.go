package streamdeck

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomonobold"
	"golang.org/x/image/math/fixed"
)

func (d *Device) WriteTextToButton(btnIndex int, text string, textColour color.Color, backgroundColour color.Color) {
	img := GetImageWithText(text, textColour, backgroundColour)
	d.writeToButton(btnIndex, img)
}

// TODO make private
func GetImageWithText(text string, textColour color.Color, backgroundColour color.Color) image.Image {
	x := 0
	y := 50

	fmt.Println(textColour)
	fmt.Println(backgroundColour)
	src_img := getSolidColourImage(textColour)
	dst_img := getSolidColourImage(backgroundColour)

	textWidth := 7 * len(text)
	fmt.Println(textWidth)

	f := &font.Drawer{
		Dst:  dst_img,
		Src:  src_img,
		Face: basicfont.Face7x13,
		Dot:  fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)},
	}
	f.DrawString(text)
	return dst_img
}
