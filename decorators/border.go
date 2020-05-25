package decorators

import (
	"image"
	"image/color"
)

type Border struct {
	width  int
	colour color.Color
}

func NewBorder(width int, colour color.Color) *Border {
	b := &Border{width: width, colour: colour}
	return b
}

func (b *Border) Apply(img image.Image) image.Image {
	newimg := img.(*image.RGBA)
	// TODO base the 96 on the image bounds
	for i := 0; i < b.width; i++ {
		rect(i, i, 96-i, 96-i, newimg, b.colour)
	}
	return newimg
}

// Utility functions from https://stackoverflow.com/questions/28992396/draw-a-rectangle-in-golang

func hLine(x1, y, x2 int, img *image.RGBA, colour color.Color) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, colour)
	}
}

func vLine(x, y1, y2 int, img *image.RGBA, colour color.Color) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, colour)
	}
}

func rect(x1, y1, x2, y2 int, img *image.RGBA, colour color.Color) {
	hLine(x1, y1, x2, img, colour)
	hLine(x1, y2, x2, img, colour)
	vLine(x1, y1, y2, img, colour)
	vLine(x2, y1, y2, img, colour)
}
