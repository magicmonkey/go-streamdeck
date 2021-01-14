package streamdeck

import (
	"image"
	"image/color"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"

	"golang.org/x/image/font/gofont/gomedium"
)

// WriteTextToButton is a low-level way to write text directly onto a button on the StreamDeck
func (d *Device) WriteTextToButton(btnIndex int, text string, textColour color.Color, backgroundColour color.Color) {
	img := getImageWithText(text, textColour, backgroundColour, d.deviceType.imageSize.X)
	d.WriteRawImageToButton(btnIndex, img)
}

func getImageWithText(text string, textColour color.Color, backgroundColour color.Color, btnSize int) image.Image {

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

	srcImg := image.NewUniform(textColour)
	dstImg := getSolidColourImage(backgroundColour, btnSize)

	c := freetype.NewContext()
	c.SetFont(myfont)
	c.SetDst(dstImg)
	c.SetSrc(srcImg)
	c.SetFontSize(size)
	c.SetClip(dstImg.Bounds())

	x := int((btnSize - width) / 2) // Horizontally centre text
	y := int(50 + (size / 3))  // Fudged vertical centre, erm, very "heuristic"

	pt := freetype.Pt(x, y)
	c.DrawString(text, pt)

	/*
		textWidth := 7 * len(text)
		fmt.Println(textWidth)

		f := &font.Drawer{
			Dst:  dstImg,
			Src:  srcImg,
			Face: basicfont.Face7x13,
			Dot:  fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)},
		}
		f.DrawString(text)
	*/
	return dstImg
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
