package streamdeck

import (
	"image"
	"image/color"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"

	"golang.org/x/image/font/gofont/gomedium"
)

func (d *Device) WriteTextToButton(btnIndex int, text string, textColour color.Color, backgroundColour color.Color) {
	img := getImageWithText(text, textColour, backgroundColour)
	d.WriteRawImageToButton(btnIndex, img)
}

func getImageWithText(text string, textColour color.Color, backgroundColour color.Color) image.Image {

	size := float64(18)

	myfont, err := truetype.Parse(gomedium.TTF)
	if err != nil {
		panic(err)
	}

	width := 0
	for size = 1; size < 60; size++ {
		width = getTextWidth(text, size)
		if width > 90 {
			size = size - 1
			break
		}
	}

	src_img := image.NewUniform(textColour)
	dst_img := getSolidColourImage(backgroundColour)

	c := freetype.NewContext()
	c.SetFont(myfont)
	c.SetDst(dst_img)
	c.SetSrc(src_img)
	c.SetFontSize(size)
	c.SetClip(dst_img.Bounds())

	x := int((96 - width) / 2) // Horizontally centre text
	y := int(50 + (size / 3))  // Fudged vertical centre, erm, very "heuristic"

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

func getTextWidth(text string, size float64) int {

	myfont, err := truetype.Parse(gomedium.TTF)
	if err != nil {
		panic(err)
	}

	// Calculate width of string
	width := 0
	face := truetype.NewFace(myfont, &truetype.Options{Size: size})
	for _, x := range text {
		awidth, _ := face.GlyphAdvance(rune(x))
		iwidthf := int(float64(awidth) / 64)
		width += iwidthf
	}

	return width
}
