package streamdeck

import (
	"fmt"
	"image"
	"image/color"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"

	"golang.org/x/image/font/gofont/gomedium"
)

func (d *Device) WriteTextToButton(btnIndex int, text string, textColour color.Color, backgroundColour color.Color, fontsize int) {
	img := getImageWithText(text, textColour, backgroundColour, fontsize)
	d.writeToButton(btnIndex, img)
}

func getImageWithText(text string, textColour color.Color, backgroundColour color.Color, fontsize int) image.Image {

	size := float64(fontsize)

	myfont, err := truetype.Parse(gomedium.TTF)
	if err != nil {
		panic(err)
	}

	fmt.Println(textColour)
	fmt.Println(backgroundColour)
	src_img := image.NewUniform(textColour)
	dst_img := getSolidColourImage(backgroundColour)

	c := freetype.NewContext()
	c.SetFont(myfont)
	c.SetDst(dst_img)
	c.SetSrc(src_img)
	c.SetFontSize(size)
	c.SetClip(dst_img.Bounds())

	// Calculate width of string
	width := 0
	face := truetype.NewFace(myfont, &truetype.Options{Size: size})
	for _, x := range text {
		awidth, _ := face.GlyphAdvance(rune(x))
		iwidthf := int(float64(awidth) / 64)
		width += iwidthf
	}

	x := int((96 - width) / 2)
	y := 50

	pt := freetype.Pt(x, y)
	c.DrawString(text, pt)

	/*
		textWidth := 7 * len(text)
		fmt.Println(textWidth)

		f := &font.Drawer{
			Dst:  dst_img,
			Src:  src_img,
			Face: basicfont.Face7x13,
			Dot:  fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)},
		}
		f.DrawString(text)
	*/
	return dst_img
}
